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
// https://github.com/elastic/elasticsearch-specification/tree/ec3159eb31c62611202a4fb157ea88fa6ff78e1a


package types

// JobForecastStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/ml/_types/Job.ts#L120-L127
type JobForecastStatistics struct {
	ForecastedJobs   int              `json:"forecasted_jobs"`
	MemoryBytes      *JobStatistics   `json:"memory_bytes,omitempty"`
	ProcessingTimeMs *JobStatistics   `json:"processing_time_ms,omitempty"`
	Records          *JobStatistics   `json:"records,omitempty"`
	Status           map[string]int64 `json:"status,omitempty"`
	Total            int64            `json:"total"`
}

// NewJobForecastStatistics returns a JobForecastStatistics.
func NewJobForecastStatistics() *JobForecastStatistics {
	r := &JobForecastStatistics{
		Status: make(map[string]int64, 0),
	}

	return r
}
