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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/datafeedstate"
)

// DatafeedStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Datafeed.ts#L140-L147
type DatafeedStats struct {
	AssignmentExplanation *string                     `json:"assignment_explanation,omitempty"`
	DatafeedId            Id                          `json:"datafeed_id"`
	Node                  *DiscoveryNode              `json:"node,omitempty"`
	RunningState          *DatafeedRunningState       `json:"running_state,omitempty"`
	State                 datafeedstate.DatafeedState `json:"state"`
	TimingStats           DatafeedTimingStats         `json:"timing_stats"`
}

// DatafeedStatsBuilder holds DatafeedStats struct and provides a builder API.
type DatafeedStatsBuilder struct {
	v *DatafeedStats
}

// NewDatafeedStats provides a builder for the DatafeedStats struct.
func NewDatafeedStatsBuilder() *DatafeedStatsBuilder {
	r := DatafeedStatsBuilder{
		&DatafeedStats{},
	}

	return &r
}

// Build finalize the chain and returns the DatafeedStats struct
func (rb *DatafeedStatsBuilder) Build() DatafeedStats {
	return *rb.v
}

func (rb *DatafeedStatsBuilder) AssignmentExplanation(assignmentexplanation string) *DatafeedStatsBuilder {
	rb.v.AssignmentExplanation = &assignmentexplanation
	return rb
}

func (rb *DatafeedStatsBuilder) DatafeedId(datafeedid Id) *DatafeedStatsBuilder {
	rb.v.DatafeedId = datafeedid
	return rb
}

func (rb *DatafeedStatsBuilder) Node(node *DiscoveryNodeBuilder) *DatafeedStatsBuilder {
	v := node.Build()
	rb.v.Node = &v
	return rb
}

func (rb *DatafeedStatsBuilder) RunningState(runningstate *DatafeedRunningStateBuilder) *DatafeedStatsBuilder {
	v := runningstate.Build()
	rb.v.RunningState = &v
	return rb
}

func (rb *DatafeedStatsBuilder) State(state datafeedstate.DatafeedState) *DatafeedStatsBuilder {
	rb.v.State = state
	return rb
}

func (rb *DatafeedStatsBuilder) TimingStats(timingstats *DatafeedTimingStatsBuilder) *DatafeedStatsBuilder {
	v := timingstats.Build()
	rb.v.TimingStats = v
	return rb
}
