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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"
)

// HoltMovingAverageAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/pipeline.ts#L205-L208
type HoltMovingAverageAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath            `json:"buckets_path,omitempty"`
	Format      *string                 `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy    `json:"gap_policy,omitempty"`
	Meta        *Metadata               `json:"meta,omitempty"`
	Minimize    *bool                   `json:"minimize,omitempty"`
	Model       string                  `json:"model,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	Predict     *int                    `json:"predict,omitempty"`
	Settings    HoltLinearModelSettings `json:"settings"`
	Window      *int                    `json:"window,omitempty"`
}

// HoltMovingAverageAggregationBuilder holds HoltMovingAverageAggregation struct and provides a builder API.
type HoltMovingAverageAggregationBuilder struct {
	v *HoltMovingAverageAggregation
}

// NewHoltMovingAverageAggregation provides a builder for the HoltMovingAverageAggregation struct.
func NewHoltMovingAverageAggregationBuilder() *HoltMovingAverageAggregationBuilder {
	r := HoltMovingAverageAggregationBuilder{
		&HoltMovingAverageAggregation{},
	}

	r.v.Model = "holt"

	return &r
}

// Build finalize the chain and returns the HoltMovingAverageAggregation struct
func (rb *HoltMovingAverageAggregationBuilder) Build() HoltMovingAverageAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.

func (rb *HoltMovingAverageAggregationBuilder) BucketsPath(bucketspath *BucketsPathBuilder) *HoltMovingAverageAggregationBuilder {
	v := bucketspath.Build()
	rb.v.BucketsPath = &v
	return rb
}

func (rb *HoltMovingAverageAggregationBuilder) Format(format string) *HoltMovingAverageAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *HoltMovingAverageAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *HoltMovingAverageAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

func (rb *HoltMovingAverageAggregationBuilder) Meta(meta *MetadataBuilder) *HoltMovingAverageAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *HoltMovingAverageAggregationBuilder) Minimize(minimize bool) *HoltMovingAverageAggregationBuilder {
	rb.v.Minimize = &minimize
	return rb
}

func (rb *HoltMovingAverageAggregationBuilder) Name(name string) *HoltMovingAverageAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *HoltMovingAverageAggregationBuilder) Predict(predict int) *HoltMovingAverageAggregationBuilder {
	rb.v.Predict = &predict
	return rb
}

func (rb *HoltMovingAverageAggregationBuilder) Settings(settings *HoltLinearModelSettingsBuilder) *HoltMovingAverageAggregationBuilder {
	v := settings.Build()
	rb.v.Settings = v
	return rb
}

func (rb *HoltMovingAverageAggregationBuilder) Window(window int) *HoltMovingAverageAggregationBuilder {
	rb.v.Window = &window
	return rb
}
