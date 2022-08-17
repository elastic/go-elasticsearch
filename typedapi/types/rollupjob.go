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

// RollupJob type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/rollup/get_jobs/types.ts#L28-L32
type RollupJob struct {
	Config RollupJobConfiguration `json:"config"`
	Stats  RollupJobStats         `json:"stats"`
	Status RollupJobStatus        `json:"status"`
}

// RollupJobBuilder holds RollupJob struct and provides a builder API.
type RollupJobBuilder struct {
	v *RollupJob
}

// NewRollupJob provides a builder for the RollupJob struct.
func NewRollupJobBuilder() *RollupJobBuilder {
	r := RollupJobBuilder{
		&RollupJob{},
	}

	return &r
}

// Build finalize the chain and returns the RollupJob struct
func (rb *RollupJobBuilder) Build() RollupJob {
	return *rb.v
}

func (rb *RollupJobBuilder) Config(config *RollupJobConfigurationBuilder) *RollupJobBuilder {
	v := config.Build()
	rb.v.Config = v
	return rb
}

func (rb *RollupJobBuilder) Stats(stats *RollupJobStatsBuilder) *RollupJobBuilder {
	v := stats.Build()
	rb.v.Stats = v
	return rb
}

func (rb *RollupJobBuilder) Status(status *RollupJobStatusBuilder) *RollupJobBuilder {
	v := status.Build()
	rb.v.Status = v
	return rb
}
