// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package benchmarks

import (
	"github.com/elastic/go-elasticsearch/v8/benchmarks/runner"
)

var (
	Config     map[string]string
	Operations []Operation

	DefaultRepetitions = 1000
)

// Operation represents a benchmarked operation.
//
type Operation struct {
	Action         string
	Category       string
	Environment    string
	NumWarmups     int
	NumRepetitions int
	SetupFunc      runner.RunnerFunc
	RunnerFunc     runner.RunnerFunc
}

// Register appends op to the list of operations.
//
func Register(op Operation) error {
	Operations = append(Operations, op)
	return nil
}
