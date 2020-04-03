// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package runner

import (
	"context"
	"fmt"
	"regexp"
	"runtime"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

var (
	statsIndex = "metrics-intake-" + time.Now().Format("2006-01")

	reGoVersion = regexp.MustCompile(`go(\d+\.\d+\.?.?)`)
	errs        []error

	OSFamily       string
	RuntimeVersion string
)

func init() {
	OSFamily = runtime.GOOS

	if v := reGoVersion.ReplaceAllString(runtime.Version(), "$1"); v != "" {
		RuntimeVersion = v
	} else {
		RuntimeVersion = runtime.Version()
	}
}

// NewRunner returns new benchmarking runner.
//
func NewRunner(cfg Config) (*Runner, error) {
	if cfg.RunnerClient == nil {
		return nil, fmt.Errorf("missing cfg.RunnerClient")
	}

	if cfg.StatsClient == nil {
		return nil, fmt.Errorf("missing cfg.StatsClient")
	}

	if cfg.RunnerFunc == nil {
		return nil, fmt.Errorf("missing cfg.RunnerFunc")
	}

	if cfg.Action == "" {
		return nil, fmt.Errorf("missing cfg.Action")
	}

	indexer, _ := esutil.NewBulkIndexer(
		esutil.BulkIndexerConfig{
			Client:        cfg.StatsClient,
			Index:         statsIndex,
			FlushBytes:    2e+6,
			FlushInterval: 15 * time.Second,
			OnError:       func(ctx context.Context, err error) { errs = append(errs, err) },
		},
	)

	return &Runner{
		config:  cfg,
		indexer: indexer,
	}, nil
}

// Runner represents the benchmarking runner.
//
type Runner struct {
	config  Config
	stats   []Stats
	indexer esutil.BulkIndexer
}

// Config represents configuration for Runner.
//
type Config struct {
	Action string

	NumWarmups     int
	NumRepetitions int

	SetupFunc    RunnerFunc
	RunnerFunc   RunnerFunc
	RunnerClient *elasticsearch.Client

	StatsClient *elasticsearch.Client
}

// RunnerFunc represents the runner operation.
//
type RunnerFunc func(Config) (*esapi.Response, error)

// Stats represents statistics about a single run.
//
type Stats struct {
	Start              time.Time
	Duration           time.Duration
	Outcome            string
	ResponseStatusCode int
}

// Error describes an error occurring in during run.
//
type Error struct {
	err  string
	errs []error
}

// Error returns string message for error.
//
func (e *Error) Error() string { return e.err }

// Run executes the benchmark runs.
//
func (r *Runner) Run() error {
	var errs []error

	r.stats = r.stats[:]

	if r.config.SetupFunc != nil {
		if _, err := r.config.SetupFunc(r.config); err != nil {
			return err
		}
	}

	for n := 1; n <= r.config.NumWarmups; n++ {
		if _, err := r.config.RunnerFunc(r.config); err != nil {
			errs = append(errs, err)
		}
	}

	for n := 1; n <= r.config.NumRepetitions; n++ {
		stat := Stats{Start: time.Now().UTC()}
		res, err := r.config.RunnerFunc(r.config)
		if err != nil {
			errs = append(errs, err)
			stat.Outcome = "failure"
		} else {
			stat.Duration = time.Since(stat.Start)
			stat.ResponseStatusCode = res.StatusCode
			if res.IsError() {
				errs = append(errs, fmt.Errorf("HTTP error: %s", res.String()))
				stat.Outcome = "failure"
			} else {
				stat.Outcome = "success"
			}
			r.stats = append(r.stats, stat)
		}
	}

	if err := r.SaveStats(); err != nil {
		return err
	}

	if len(errs) > 0 {
		return &Error{err: fmt.Sprintf("encountered %d errors during the run", len(errs)), errs: errs}
	}
	return nil
}

// Stats returns statistics about the run.
//
func (r *Runner) Stats() []Stats {
	return r.stats
}

// Errs returns the runner errors.
//
func (r *Runner) Errs() []error {
	return errs
}

// SaveStats stores runner statistics in Elasticsearch.
//
func (r *Runner) SaveStats() error {
	var errs []error

	for _, s := range r.stats {
		record := record{
			Timestamp: s.Start,
			Tags:      []string{"go-elasticsearch"},
			Labels: map[string]string{
				"client": "go-elasticsearch",
			},
			Event: recordEvent{
				Action:   r.config.Action,
				Duration: s.Duration.Nanoseconds(),
			},
			Benchmark: recordBenchmark{
				Warmups:     r.config.NumWarmups,
				Repetitions: r.config.NumRepetitions,
			},
			HTTP: recordHTTP{
				Response: recordHTTPResponse{
					StatusCode: s.ResponseStatusCode,
				},
			},
			OS:      recordOS{Family: OSFamily},
			Git:     recordGit{Branch: "master"},
			Runtime: recordRuntime{Name: "go", Version: RuntimeVersion},
		}
		if err := r.indexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Action: "index",
				Body:   esutil.NewJSONReader(record),
				OnFailure: func(
					ctx context.Context,
					item esutil.BulkIndexerItem,
					res esutil.BulkIndexerResponseItem, err error,
				) {
					if err != nil {
						errs = append(errs, err)
					} else {
						errs = append(errs, fmt.Errorf("HTTP error: %s: %s", res.Error.Type, res.Error.Reason))
					}
				},
			},
		); err != nil {
			return err
		}
	}

	if err := r.indexer.Close(context.Background()); err != nil {
		return err
	}

	if len(errs) > 0 {
		return fmt.Errorf("%d errors when saving stats: %s", len(errs), errs)
	}
	return nil
}

// record represents statistics about a single iteration.
//
type record struct {
	Timestamp time.Time         `json:"@timestamp"`
	Labels    map[string]string `json:"labels,omitempty"`
	Tags      []string          `json:"tags,omitempty"`

	Event     recordEvent     `json:"event"`
	Benchmark recordBenchmark `json:"benchmark,omitempty"`
	HTTP      recordHTTP      `json:"http,omitempty"`
	OS        recordOS        `json:"os,omitempty"`
	Runtime   recordRuntime   `json:"runtime,omitempty"`
	Git       recordGit       `json:"git,omitempty"`
}

// recordEvent represents the event information for a single iteration.
//
type recordEvent struct {
	Action   string `json:"action,omitempty"`
	Duration int64  `json:"duration,omitempty"`
}

// recordBenchmark represents the benchmark information for a single iteration.
//
type recordBenchmark struct {
	Warmups     int `json:"warmups,omitempty"`
	Repetitions int `json:"repetitions,omitempty"`
}

// recordHTTP represents the HTTP information for a single iteration.
//
type recordHTTP struct {
	Response recordHTTPResponse `json:"response,omitempty"`
}

// recordHTTPResponse represents the HTTP response information for a single iteration.
//
type recordHTTPResponse struct {
	StatusCode int `json:"status_code,omitempty"`
}

// recordOS represents the information about the operating system.
//
type recordOS struct {
	Family string `json:"family,omitempty"`
}

// recordRuntime represents the information about the runtime.
//
type recordRuntime struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// recordGit represents the information about the Git repository.
//
type recordGit struct {
	Branch string `json:"branch,omitempty"`
}
