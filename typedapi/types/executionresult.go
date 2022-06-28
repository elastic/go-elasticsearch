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


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// ExecutionResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/watcher/_types/Execution.ts#L60-L66
type ExecutionResult struct {
	Actions           []ExecutionResultAction  `json:"actions"`
	Condition         ExecutionResultCondition `json:"condition"`
	ExecutionDuration int                      `json:"execution_duration"`
	ExecutionTime     DateString               `json:"execution_time"`
	Input             ExecutionResultInput     `json:"input"`
}

// ExecutionResultBuilder holds ExecutionResult struct and provides a builder API.
type ExecutionResultBuilder struct {
	v *ExecutionResult
}

// NewExecutionResult provides a builder for the ExecutionResult struct.
func NewExecutionResultBuilder() *ExecutionResultBuilder {
	r := ExecutionResultBuilder{
		&ExecutionResult{},
	}

	return &r
}

// Build finalize the chain and returns the ExecutionResult struct
func (rb *ExecutionResultBuilder) Build() ExecutionResult {
	return *rb.v
}

func (rb *ExecutionResultBuilder) Actions(actions []ExecutionResultActionBuilder) *ExecutionResultBuilder {
	tmp := make([]ExecutionResultAction, len(actions))
	for _, value := range actions {
		tmp = append(tmp, value.Build())
	}
	rb.v.Actions = tmp
	return rb
}

func (rb *ExecutionResultBuilder) Condition(condition *ExecutionResultConditionBuilder) *ExecutionResultBuilder {
	v := condition.Build()
	rb.v.Condition = v
	return rb
}

func (rb *ExecutionResultBuilder) ExecutionDuration(executionduration int) *ExecutionResultBuilder {
	rb.v.ExecutionDuration = executionduration
	return rb
}

func (rb *ExecutionResultBuilder) ExecutionTime(executiontime DateString) *ExecutionResultBuilder {
	rb.v.ExecutionTime = executiontime
	return rb
}

func (rb *ExecutionResultBuilder) Input(input *ExecutionResultInputBuilder) *ExecutionResultBuilder {
	v := input.Build()
	rb.v.Input = v
	return rb
}
