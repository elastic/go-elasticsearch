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

// ModelPlotConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/ModelPlot.ts#L23-L41
type ModelPlotConfig struct {
	// AnnotationsEnabled If true, enables calculation and storage of the model change annotations for
	// each entity that is being analyzed.
	AnnotationsEnabled *bool `json:"annotations_enabled,omitempty"`
	// Enabled If true, enables calculation and storage of the model bounds for each entity
	// that is being analyzed.
	Enabled *bool `json:"enabled,omitempty"`
	// Terms Limits data collection to this comma separated list of partition or by field
	// values. If terms are not specified or it is an empty string, no filtering is
	// applied. Wildcards are not supported. Only the specified terms can be viewed
	// when using the Single Metric Viewer.
	Terms *Field `json:"terms,omitempty"`
}

// ModelPlotConfigBuilder holds ModelPlotConfig struct and provides a builder API.
type ModelPlotConfigBuilder struct {
	v *ModelPlotConfig
}

// NewModelPlotConfig provides a builder for the ModelPlotConfig struct.
func NewModelPlotConfigBuilder() *ModelPlotConfigBuilder {
	r := ModelPlotConfigBuilder{
		&ModelPlotConfig{},
	}

	return &r
}

// Build finalize the chain and returns the ModelPlotConfig struct
func (rb *ModelPlotConfigBuilder) Build() ModelPlotConfig {
	return *rb.v
}

// AnnotationsEnabled If true, enables calculation and storage of the model change annotations for
// each entity that is being analyzed.

func (rb *ModelPlotConfigBuilder) AnnotationsEnabled(annotationsenabled bool) *ModelPlotConfigBuilder {
	rb.v.AnnotationsEnabled = &annotationsenabled
	return rb
}

// Enabled If true, enables calculation and storage of the model bounds for each entity
// that is being analyzed.

func (rb *ModelPlotConfigBuilder) Enabled(enabled bool) *ModelPlotConfigBuilder {
	rb.v.Enabled = &enabled
	return rb
}

// Terms Limits data collection to this comma separated list of partition or by field
// values. If terms are not specified or it is an empty string, no filtering is
// applied. Wildcards are not supported. Only the specified terms can be viewed
// when using the Single Metric Viewer.

func (rb *ModelPlotConfigBuilder) Terms(terms Field) *ModelPlotConfigBuilder {
	rb.v.Terms = &terms
	return rb
}
