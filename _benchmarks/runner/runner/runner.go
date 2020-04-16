// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package runner

import (
	"context"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

var (
	statsIndex = "metrics-intake-" + time.Now().Format("2006-01")

	reGoVersion = regexp.MustCompile(`go(\d+\.\d+\.?.?)`)
	errs        []error

	RuntimeOS      string
	RuntimeVersion string
)

func init() {
	RuntimeOS = runtime.GOOS

	if v := reGoVersion.ReplaceAllString(runtime.Version(), "$1"); v != "" {
		RuntimeVersion = v
	} else {
		RuntimeVersion = runtime.Version()
	}
}

// NewRunner returns new benchmarking runner.
//
func NewRunner(cfg Config) (*Runner, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	indexer, _ := esutil.NewBulkIndexer(
		esutil.BulkIndexerConfig{
			Client:        cfg.ReportClient,
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
	BuildID string

	Action      string
	Category    string
	Environment string

	NumWarmups     int
	NumRepetitions int

	SetupFunc    RunnerFunc
	RunnerFunc   RunnerFunc
	RunnerClient *elasticsearch.Client

	ReportClient *elasticsearch.Client

	Target struct {
		OS      ConfigOS
		Service ConfigService
	}

	Runner struct {
		Service ConfigService
	}
}

// ConfigOS describes OS.
//
type ConfigOS struct {
	Family string
}

// ConfigService describes service.
//
type ConfigService struct {
	Type    string
	Name    string
	Version string
	Git     ConfigGit
}

// ConfigGit describes Git.
//
type ConfigGit struct {
	Branch string
	Commit string
}

// RunnerFunc represents the runner operation.
//
type RunnerFunc func(int, Config) (*esapi.Response, error)

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

// Errs returns the underlying errors.
//
func (e *Error) Errs() string {
	var errStrings []string
	for _, e := range e.errs {
		errStrings = append(errStrings, e.Error())
	}
	return strings.Join(errStrings, ";")
}

// Run executes the benchmark runs.
//
func (r *Runner) Run() error {
	if err := validateConfig(r.config); err != nil {
		return err
	}

	var errs []error

	r.stats = r.stats[:]

	if r.config.SetupFunc != nil {
		if _, err := r.config.SetupFunc(0, r.config); err != nil {
			return err
		}
	}

	for n := 1; n <= r.config.NumWarmups; n++ {
		if _, err := r.config.RunnerFunc(n, r.config); err != nil {
			errs = append(errs, err)
		}
	}

	for n := 1; n <= r.config.NumRepetitions; n++ {
		stat := Stats{Start: time.Now().UTC()}
		res, err := r.config.RunnerFunc(n, r.config)
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
			Tags:      []string{"bench", "go-elasticsearch"},
			Labels: map[string]string{
				"build_id":    r.config.BuildID,
				"environment": r.config.Environment,
			},
			Event: recordEvent{
				Action:   r.config.Action,
				Duration: s.Duration.Nanoseconds(),
			},
			HTTP: recordHTTP{
				Response: recordHTTPResponse{
					StatusCode: s.ResponseStatusCode,
				},
			},
			Benchmark: recordBenchmark{
				BuildID:     r.config.BuildID,
				Category:    r.config.Category,
				Environment: r.config.Environment,
				Repetitions: r.config.NumRepetitions,
				Runner: recordRunner{
					Service: recordService{
						Type:    "client",
						Name:    "go-elasticsearch",
						Version: elasticsearch.Version,
						Git:     recordGit{Branch: r.config.Runner.Service.Git.Branch, Commit: r.config.Runner.Service.Git.Commit},
					},
					Runtime: recordRuntime{Name: "go", Version: RuntimeVersion},
					OS:      recordOS{Family: RuntimeOS},
				},
				Target: recordTarget{
					Service: recordService{
						Type:    r.config.Target.Service.Type,
						Name:    r.config.Target.Service.Type,
						Version: r.config.Target.Service.Version,
						Git:     recordGit{Branch: r.config.Target.Service.Git.Branch, Commit: r.config.Target.Service.Git.Commit},
					},
					OS: recordOS{Family: r.config.Target.OS.Family},
				},
			},
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
type record struct {
	Timestamp time.Time         `json:"@timestamp"`
	Labels    map[string]string `json:"labels,omitempty"`
	Tags      []string          `json:"tags,omitempty"`

	Event     recordEvent     `json:"event"`
	HTTP      recordHTTP      `json:"http,omitempty"`
	Benchmark recordBenchmark `json:"benchmark"`
}

// recordEvent represents the event information for a single iteration.
// See: https://www.elastic.co/guide/en/ecs/current/ecs-event.html
type recordEvent struct {
	Action   string `json:"action"`
	Duration int64  `json:"duration"`
	Dataset  string `json:"dataset,omitempty"`
}

// recordBenchmark represents the benchmark information for a single iteration.
type recordBenchmark struct {
	BuildID     string       `json:"build_id"`
	Repetitions int          `json:"repetitions"`
	Runner      recordRunner `json:"runner"`
	Target      recordTarget `json:"target"`
	Category    string       `json:"category,omitempty"`
	Environment string       `json:"environment,omitempty"`
}

// recordRunner represents the information about the runner.
type recordRunner struct {
	Service recordService `json:"service"`
	Runtime recordRuntime `json:"runtime"`
	OS      recordOS      `json:"os"`
}

// recordTarget represents the information about the target.
type recordTarget struct {
	Service recordService `json:"service"`
	OS      recordOS      `json:"os"`
}

// recordService represents the information about the target service.
// See: https://www.elastic.co/guide/en/ecs/current/ecs-service.html
type recordService struct {
	Type    string    `json:"type"`
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Git     recordGit `json:"git,omitempty"`
}

// recordHTTP represents the HTTP information for a single iteration.
type recordHTTP struct {
	Response recordHTTPResponse `json:"response,omitempty"`
}

// recordHTTPResponse represents the HTTP response information for a single iteration.
type recordHTTPResponse struct {
	StatusCode int `json:"status_code"`
}

// recordOS represents the information about the operating system.
// See: https://www.elastic.co/guide/en/ecs/current/ecs-os.html
type recordOS struct {
	Family string `json:"family"`
}

// recordRuntime represents the information about the client runtime.
type recordRuntime struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// recordGit represents the information about the client Git repository.
type recordGit struct {
	Branch string `json:"branch,omitempty"`
	Commit string `json:"commit,omitempty"`
}

func validateConfig(cfg Config) error {
	if cfg.BuildID == "" {
		return fmt.Errorf("missing cfg.BuildID")
	}

	if cfg.Action == "" {
		return fmt.Errorf("missing cfg.Action")
	}

	if cfg.Category == "" {
		return fmt.Errorf("missing cfg.Category")
	}

	if cfg.Environment == "" {
		return fmt.Errorf("missing cfg.Environment")
	}

	if cfg.RunnerClient == nil {
		return fmt.Errorf("missing cfg.RunnerClient")
	}

	if cfg.ReportClient == nil {
		return fmt.Errorf("missing cfg.ReportClient")
	}

	if cfg.RunnerFunc == nil {
		return fmt.Errorf("missing cfg.RunnerFunc")
	}

	return nil
}
