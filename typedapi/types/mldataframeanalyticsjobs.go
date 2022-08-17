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

// MlDataFrameAnalyticsJobs type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L168-L173
type MlDataFrameAnalyticsJobs struct {
	All_           MlDataFrameAnalyticsJobsCount     `json:"_all"`
	AnalysisCounts *MlDataFrameAnalyticsJobsAnalysis `json:"analysis_counts,omitempty"`
	MemoryUsage    *MlDataFrameAnalyticsJobsMemory   `json:"memory_usage,omitempty"`
	Stopped        *MlDataFrameAnalyticsJobsCount    `json:"stopped,omitempty"`
}

// MlDataFrameAnalyticsJobsBuilder holds MlDataFrameAnalyticsJobs struct and provides a builder API.
type MlDataFrameAnalyticsJobsBuilder struct {
	v *MlDataFrameAnalyticsJobs
}

// NewMlDataFrameAnalyticsJobs provides a builder for the MlDataFrameAnalyticsJobs struct.
func NewMlDataFrameAnalyticsJobsBuilder() *MlDataFrameAnalyticsJobsBuilder {
	r := MlDataFrameAnalyticsJobsBuilder{
		&MlDataFrameAnalyticsJobs{},
	}

	return &r
}

// Build finalize the chain and returns the MlDataFrameAnalyticsJobs struct
func (rb *MlDataFrameAnalyticsJobsBuilder) Build() MlDataFrameAnalyticsJobs {
	return *rb.v
}

func (rb *MlDataFrameAnalyticsJobsBuilder) All_(all_ *MlDataFrameAnalyticsJobsCountBuilder) *MlDataFrameAnalyticsJobsBuilder {
	v := all_.Build()
	rb.v.All_ = v
	return rb
}

func (rb *MlDataFrameAnalyticsJobsBuilder) AnalysisCounts(analysiscounts *MlDataFrameAnalyticsJobsAnalysisBuilder) *MlDataFrameAnalyticsJobsBuilder {
	v := analysiscounts.Build()
	rb.v.AnalysisCounts = &v
	return rb
}

func (rb *MlDataFrameAnalyticsJobsBuilder) MemoryUsage(memoryusage *MlDataFrameAnalyticsJobsMemoryBuilder) *MlDataFrameAnalyticsJobsBuilder {
	v := memoryusage.Build()
	rb.v.MemoryUsage = &v
	return rb
}

func (rb *MlDataFrameAnalyticsJobsBuilder) Stopped(stopped *MlDataFrameAnalyticsJobsCountBuilder) *MlDataFrameAnalyticsJobsBuilder {
	v := stopped.Build()
	rb.v.Stopped = &v
	return rb
}
