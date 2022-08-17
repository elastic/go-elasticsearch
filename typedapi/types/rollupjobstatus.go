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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexingjobstate"
)

// RollupJobStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/rollup/get_jobs/types.ts#L60-L64
type RollupJobStatus struct {
	CurrentPosition map[string]interface{}            `json:"current_position,omitempty"`
	JobState        indexingjobstate.IndexingJobState `json:"job_state"`
	UpgradedDocId   *bool                             `json:"upgraded_doc_id,omitempty"`
}

// RollupJobStatusBuilder holds RollupJobStatus struct and provides a builder API.
type RollupJobStatusBuilder struct {
	v *RollupJobStatus
}

// NewRollupJobStatus provides a builder for the RollupJobStatus struct.
func NewRollupJobStatusBuilder() *RollupJobStatusBuilder {
	r := RollupJobStatusBuilder{
		&RollupJobStatus{
			CurrentPosition: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the RollupJobStatus struct
func (rb *RollupJobStatusBuilder) Build() RollupJobStatus {
	return *rb.v
}

func (rb *RollupJobStatusBuilder) CurrentPosition(value map[string]interface{}) *RollupJobStatusBuilder {
	rb.v.CurrentPosition = value
	return rb
}

func (rb *RollupJobStatusBuilder) JobState(jobstate indexingjobstate.IndexingJobState) *RollupJobStatusBuilder {
	rb.v.JobState = jobstate
	return rb
}

func (rb *RollupJobStatusBuilder) UpgradedDocId(upgradeddocid bool) *RollupJobStatusBuilder {
	rb.v.UpgradedDocId = &upgradeddocid
	return rb
}
