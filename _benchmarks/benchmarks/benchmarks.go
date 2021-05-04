// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

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
