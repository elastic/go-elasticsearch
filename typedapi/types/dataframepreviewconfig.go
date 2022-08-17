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

// DataframePreviewConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/preview_data_frame_analytics/types.ts#L27-L33
type DataframePreviewConfig struct {
	Analysis         DataframeAnalysisContainer       `json:"analysis"`
	AnalyzedFields   *DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`
	MaxNumThreads    *int                             `json:"max_num_threads,omitempty"`
	ModelMemoryLimit *string                          `json:"model_memory_limit,omitempty"`
	Source           DataframeAnalyticsSource         `json:"source"`
}

// DataframePreviewConfigBuilder holds DataframePreviewConfig struct and provides a builder API.
type DataframePreviewConfigBuilder struct {
	v *DataframePreviewConfig
}

// NewDataframePreviewConfig provides a builder for the DataframePreviewConfig struct.
func NewDataframePreviewConfigBuilder() *DataframePreviewConfigBuilder {
	r := DataframePreviewConfigBuilder{
		&DataframePreviewConfig{},
	}

	return &r
}

// Build finalize the chain and returns the DataframePreviewConfig struct
func (rb *DataframePreviewConfigBuilder) Build() DataframePreviewConfig {
	return *rb.v
}

func (rb *DataframePreviewConfigBuilder) Analysis(analysis *DataframeAnalysisContainerBuilder) *DataframePreviewConfigBuilder {
	v := analysis.Build()
	rb.v.Analysis = v
	return rb
}

func (rb *DataframePreviewConfigBuilder) AnalyzedFields(analyzedfields *DataframeAnalysisAnalyzedFieldsBuilder) *DataframePreviewConfigBuilder {
	v := analyzedfields.Build()
	rb.v.AnalyzedFields = &v
	return rb
}

func (rb *DataframePreviewConfigBuilder) MaxNumThreads(maxnumthreads int) *DataframePreviewConfigBuilder {
	rb.v.MaxNumThreads = &maxnumthreads
	return rb
}

func (rb *DataframePreviewConfigBuilder) ModelMemoryLimit(modelmemorylimit string) *DataframePreviewConfigBuilder {
	rb.v.ModelMemoryLimit = &modelmemorylimit
	return rb
}

func (rb *DataframePreviewConfigBuilder) Source(source *DataframeAnalyticsSourceBuilder) *DataframePreviewConfigBuilder {
	v := source.Build()
	rb.v.Source = v
	return rb
}
