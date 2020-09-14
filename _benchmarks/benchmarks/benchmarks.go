// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package benchmarks

import (
	"bytes"

	"github.com/elastic/go-elasticsearch/v8/benchmarks/runner"
)

var (
	Config      map[string]string
	Actions     []Action
	DataSources = make(map[string]*bytes.Buffer)

	DefaultRepetitions = 1000
)

// Action represents a benchmarked action.
//
type Action struct {
	Name           string
	Category       string
	Environment    string
	NumWarmups     int
	NumRepetitions int
	NumOperations  int
	SetupFunc      runner.RunnerFunc
	RunnerFunc     runner.RunnerFunc
}

// Register appends op to the list of operations.
//
func Register(a Action) error {
	Actions = append(Actions, a)
	return nil
}
