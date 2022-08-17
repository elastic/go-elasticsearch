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

// MachineLearning type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L360-L368
type MachineLearning struct {
	Available              bool                     `json:"available"`
	DataFrameAnalyticsJobs MlDataFrameAnalyticsJobs `json:"data_frame_analytics_jobs"`
	Datafeeds              map[string]Datafeed      `json:"datafeeds"`
	Enabled                bool                     `json:"enabled"`
	Inference              MlInference              `json:"inference"`
	Jobs                   Jobs                     `json:"jobs"`
	NodeCount              int                      `json:"node_count"`
}

// MachineLearningBuilder holds MachineLearning struct and provides a builder API.
type MachineLearningBuilder struct {
	v *MachineLearning
}

// NewMachineLearning provides a builder for the MachineLearning struct.
func NewMachineLearningBuilder() *MachineLearningBuilder {
	r := MachineLearningBuilder{
		&MachineLearning{
			Datafeeds: make(map[string]Datafeed, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MachineLearning struct
func (rb *MachineLearningBuilder) Build() MachineLearning {
	return *rb.v
}

func (rb *MachineLearningBuilder) Available(available bool) *MachineLearningBuilder {
	rb.v.Available = available
	return rb
}

func (rb *MachineLearningBuilder) DataFrameAnalyticsJobs(dataframeanalyticsjobs *MlDataFrameAnalyticsJobsBuilder) *MachineLearningBuilder {
	v := dataframeanalyticsjobs.Build()
	rb.v.DataFrameAnalyticsJobs = v
	return rb
}

func (rb *MachineLearningBuilder) Datafeeds(values map[string]*DatafeedBuilder) *MachineLearningBuilder {
	tmp := make(map[string]Datafeed, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Datafeeds = tmp
	return rb
}

func (rb *MachineLearningBuilder) Enabled(enabled bool) *MachineLearningBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *MachineLearningBuilder) Inference(inference *MlInferenceBuilder) *MachineLearningBuilder {
	v := inference.Build()
	rb.v.Inference = v
	return rb
}

func (rb *MachineLearningBuilder) Jobs(jobs *JobsBuilder) *MachineLearningBuilder {
	v := jobs.Build()
	rb.v.Jobs = v
	return rb
}

func (rb *MachineLearningBuilder) NodeCount(nodecount int) *MachineLearningBuilder {
	rb.v.NodeCount = nodecount
	return rb
}
