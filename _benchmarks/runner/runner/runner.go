// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package runner

import (
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

var (
	statsIndex = "metrics"
)

// NewRunner returns new benchmarking runner.
//
func NewRunner(cfg Config) (*Runner, error) {
	if cfg.RunnerClient == nil {
		return nil, fmt.Errorf("missing cfg.RunnerClient")
	}

	if cfg.StatsClient == nil {
		return nil, fmt.Errorf("missing cfg.StatsClient")
	}

	if cfg.SetupFunc == nil {
		return nil, fmt.Errorf("missing cfg.SetupFunc")
	}

	if cfg.RunnerFunc == nil {
		return nil, fmt.Errorf("missing cfg.RunnerFunc")
	}

	if cfg.Action == "" {
		return nil, fmt.Errorf("missing cfg.Action")
	}

	return &Runner{config: cfg}, nil
}

// Runner represents the benchmarking runner.
//
type Runner struct {
	config Config
	stats  []Stats
}

// Config represents configuration for Runner.
//
type Config struct {
	Action string

	NumWarmups     int
	NumRepetitions int
	NumIterations  int

	SetupFunc    func(Config) error
	RunnerFunc   func(Config) error
	RunnerClient *elasticsearch.Client

	StatsClient *elasticsearch.Client
}

// Stats represents statistics about a single run.
//
type Stats struct {
	Iterations int           `json:"iterations"`
	Duration   time.Duration `json:"duration"`
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
	var (
		errs     []error
		start    time.Time
		duration time.Duration
		stats    Stats
	)

	if err := r.config.SetupFunc(r.config); err != nil {
		return err
	}

	for n := 1; n <= r.config.NumWarmups; n++ {
		for i := 1; i <= r.config.NumIterations; i++ {
			if err := r.config.RunnerFunc(r.config); err != nil {
				errs = append(errs, err)
			}
		}
	}

	for n := 1; n <= r.config.NumRepetitions; n++ {
		start = time.Now().UTC()
		for i := 1; i <= r.config.NumIterations; i++ {
			if err := r.config.RunnerFunc(r.config); err != nil {
				errs = append(errs, err)
			}
		}
		duration = time.Since(start)

		stats = Stats{
			Duration:   duration,
			Iterations: r.config.NumIterations,
		}
		r.stats = append(r.stats, stats)
	}

	r.SaveStats()

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

// SaveStats stores runner statistics in Elasticsearch.
//
func (r *Runner) SaveStats() error {
	var errs []error

	for _, s := range r.stats {
		record := record{
			Timestamp: time.Now().UTC(),
			Labels:    map[string]string{"client": "go-elasticsearch"},
			Tags:      []string{"go-elasticsearch"},
			Event: recordEvent{
				Action:   r.config.Action,
				Duration: s.Duration.Nanoseconds() / int64(s.Iterations),
			},
		}
		payload := esutil.NewJSONReader(record)
		res, err := r.config.StatsClient.Index(statsIndex, payload)
		if err != nil {
			errs = append(errs, err)
		} else {
			if res.IsError() {
				errs = append(errs, fmt.Errorf("error: %s", res.String()))
			}
			res.Body.Close()
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("%d errors when saving stats: %s", len(errs), errs)
	}
	return nil
}

// record represents statistics about a single run.
//
type record struct {
	Timestamp time.Time         `json:"@timestamp"`
	Labels    map[string]string `json:"labels,omitempty"`
	Tags      []string          `json:"tags,omitempty"`

	Event recordEvent `json:"event"`
}

// recordEvent represents the event information for a single run.
//
type recordEvent struct {
	Action   string `json:"action,omitempty"`
	Duration int64  `json:"duration,omitempty"`
}
