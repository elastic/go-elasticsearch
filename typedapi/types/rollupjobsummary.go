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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// RollupJobSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/rollup/get_rollup_index_caps/types.ts#L28-L33
type RollupJobSummary struct {
	Fields       map[Field][]RollupJobSummaryField `json:"fields"`
	IndexPattern string                            `json:"index_pattern"`
	JobId        Id                                `json:"job_id"`
	RollupIndex  IndexName                         `json:"rollup_index"`
}

// RollupJobSummaryBuilder holds RollupJobSummary struct and provides a builder API.
type RollupJobSummaryBuilder struct {
	v *RollupJobSummary
}

// NewRollupJobSummary provides a builder for the RollupJobSummary struct.
func NewRollupJobSummaryBuilder() *RollupJobSummaryBuilder {
	r := RollupJobSummaryBuilder{
		&RollupJobSummary{
			Fields: make(map[Field][]RollupJobSummaryField, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the RollupJobSummary struct
func (rb *RollupJobSummaryBuilder) Build() RollupJobSummary {
	return *rb.v
}

func (rb *RollupJobSummaryBuilder) Fields(value map[Field][]RollupJobSummaryField) *RollupJobSummaryBuilder {
	rb.v.Fields = value
	return rb
}

func (rb *RollupJobSummaryBuilder) IndexPattern(indexpattern string) *RollupJobSummaryBuilder {
	rb.v.IndexPattern = indexpattern
	return rb
}

func (rb *RollupJobSummaryBuilder) JobId(jobid Id) *RollupJobSummaryBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *RollupJobSummaryBuilder) RollupIndex(rollupindex IndexName) *RollupJobSummaryBuilder {
	rb.v.RollupIndex = rollupindex
	return rb
}
