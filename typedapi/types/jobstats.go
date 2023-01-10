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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jobstate"
)

// JobStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/Job.ts#L96-L107
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

// NewJobStats returns a JobStats.
func NewJobStats() *JobStats {
	r := &JobStats{}

	return r
}
