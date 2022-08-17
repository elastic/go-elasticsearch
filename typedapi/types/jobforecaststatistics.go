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

// JobForecastStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Job.ts#L120-L127
type JobForecastStatistics struct {
	ForecastedJobs   int              `json:"forecasted_jobs"`
	MemoryBytes      *JobStatistics   `json:"memory_bytes,omitempty"`
	ProcessingTimeMs *JobStatistics   `json:"processing_time_ms,omitempty"`
	Records          *JobStatistics   `json:"records,omitempty"`
	Status           map[string]int64 `json:"status,omitempty"`
	Total            int64            `json:"total"`
}

// JobForecastStatisticsBuilder holds JobForecastStatistics struct and provides a builder API.
type JobForecastStatisticsBuilder struct {
	v *JobForecastStatistics
}

// NewJobForecastStatistics provides a builder for the JobForecastStatistics struct.
func NewJobForecastStatisticsBuilder() *JobForecastStatisticsBuilder {
	r := JobForecastStatisticsBuilder{
		&JobForecastStatistics{
			Status: make(map[string]int64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the JobForecastStatistics struct
func (rb *JobForecastStatisticsBuilder) Build() JobForecastStatistics {
	return *rb.v
}

func (rb *JobForecastStatisticsBuilder) ForecastedJobs(forecastedjobs int) *JobForecastStatisticsBuilder {
	rb.v.ForecastedJobs = forecastedjobs
	return rb
}

func (rb *JobForecastStatisticsBuilder) MemoryBytes(memorybytes *JobStatisticsBuilder) *JobForecastStatisticsBuilder {
	v := memorybytes.Build()
	rb.v.MemoryBytes = &v
	return rb
}

func (rb *JobForecastStatisticsBuilder) ProcessingTimeMs(processingtimems *JobStatisticsBuilder) *JobForecastStatisticsBuilder {
	v := processingtimems.Build()
	rb.v.ProcessingTimeMs = &v
	return rb
}

func (rb *JobForecastStatisticsBuilder) Records(records *JobStatisticsBuilder) *JobForecastStatisticsBuilder {
	v := records.Build()
	rb.v.Records = &v
	return rb
}

func (rb *JobForecastStatisticsBuilder) Status(value map[string]int64) *JobForecastStatisticsBuilder {
	rb.v.Status = value
	return rb
}

func (rb *JobForecastStatisticsBuilder) Total(total int64) *JobForecastStatisticsBuilder {
	rb.v.Total = total
	return rb
}
