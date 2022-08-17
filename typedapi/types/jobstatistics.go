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

// JobStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Job.ts#L44-L49
type JobStatistics struct {
	Avg   float64 `json:"avg"`
	Max   float64 `json:"max"`
	Min   float64 `json:"min"`
	Total float64 `json:"total"`
}

// JobStatisticsBuilder holds JobStatistics struct and provides a builder API.
type JobStatisticsBuilder struct {
	v *JobStatistics
}

// NewJobStatistics provides a builder for the JobStatistics struct.
func NewJobStatisticsBuilder() *JobStatisticsBuilder {
	r := JobStatisticsBuilder{
		&JobStatistics{},
	}

	return &r
}

// Build finalize the chain and returns the JobStatistics struct
func (rb *JobStatisticsBuilder) Build() JobStatistics {
	return *rb.v
}

func (rb *JobStatisticsBuilder) Avg(avg float64) *JobStatisticsBuilder {
	rb.v.Avg = avg
	return rb
}

func (rb *JobStatisticsBuilder) Max(max float64) *JobStatisticsBuilder {
	rb.v.Max = max
	return rb
}

func (rb *JobStatisticsBuilder) Min(min float64) *JobStatisticsBuilder {
	rb.v.Min = min
	return rb
}

func (rb *JobStatisticsBuilder) Total(total float64) *JobStatisticsBuilder {
	rb.v.Total = total
	return rb
}
