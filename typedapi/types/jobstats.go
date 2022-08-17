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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jobstate"
)

// JobStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Job.ts#L96-L107
type JobStats struct {
	AssignmentExplanation *string               `json:"assignment_explanation,omitempty"`
	DataCounts            DataCounts            `json:"data_counts"`
	Deleting              *bool                 `json:"deleting,omitempty"`
	ForecastsStats        JobForecastStatistics `json:"forecasts_stats"`
	JobId                 string                `json:"job_id"`
	ModelSizeStats        ModelSizeStats        `json:"model_size_stats"`
	Node                  *DiscoveryNode        `json:"node,omitempty"`
	OpenTime              *DateTime             `json:"open_time,omitempty"`
	State                 jobstate.JobState     `json:"state"`
	TimingStats           JobTimingStats        `json:"timing_stats"`
}

// JobStatsBuilder holds JobStats struct and provides a builder API.
type JobStatsBuilder struct {
	v *JobStats
}

// NewJobStats provides a builder for the JobStats struct.
func NewJobStatsBuilder() *JobStatsBuilder {
	r := JobStatsBuilder{
		&JobStats{},
	}

	return &r
}

// Build finalize the chain and returns the JobStats struct
func (rb *JobStatsBuilder) Build() JobStats {
	return *rb.v
}

func (rb *JobStatsBuilder) AssignmentExplanation(assignmentexplanation string) *JobStatsBuilder {
	rb.v.AssignmentExplanation = &assignmentexplanation
	return rb
}

func (rb *JobStatsBuilder) DataCounts(datacounts *DataCountsBuilder) *JobStatsBuilder {
	v := datacounts.Build()
	rb.v.DataCounts = v
	return rb
}

func (rb *JobStatsBuilder) Deleting(deleting bool) *JobStatsBuilder {
	rb.v.Deleting = &deleting
	return rb
}

func (rb *JobStatsBuilder) ForecastsStats(forecastsstats *JobForecastStatisticsBuilder) *JobStatsBuilder {
	v := forecastsstats.Build()
	rb.v.ForecastsStats = v
	return rb
}

func (rb *JobStatsBuilder) JobId(jobid string) *JobStatsBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *JobStatsBuilder) ModelSizeStats(modelsizestats *ModelSizeStatsBuilder) *JobStatsBuilder {
	v := modelsizestats.Build()
	rb.v.ModelSizeStats = v
	return rb
}

func (rb *JobStatsBuilder) Node(node *DiscoveryNodeBuilder) *JobStatsBuilder {
	v := node.Build()
	rb.v.Node = &v
	return rb
}

func (rb *JobStatsBuilder) OpenTime(opentime *DateTimeBuilder) *JobStatsBuilder {
	v := opentime.Build()
	rb.v.OpenTime = &v
	return rb
}

func (rb *JobStatsBuilder) State(state jobstate.JobState) *JobStatsBuilder {
	rb.v.State = state
	return rb
}

func (rb *JobStatsBuilder) TimingStats(timingstats *JobTimingStatsBuilder) *JobStatsBuilder {
	v := timingstats.Build()
	rb.v.TimingStats = v
	return rb
}
