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

// SimpleMovingAverageAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/pipeline.ts#L195-L198
type SimpleMovingAverageAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Minimize    *bool                `json:"minimize,omitempty"`
	Model       string               `json:"model,omitempty"`
	Name        *string              `json:"name,omitempty"`
	Predict     *int                 `json:"predict,omitempty"`
	Settings    EmptyObject          `json:"settings"`
	Window      *int                 `json:"window,omitempty"`
}

// SimpleMovingAverageAggregationBuilder holds SimpleMovingAverageAggregation struct and provides a builder API.
type SimpleMovingAverageAggregationBuilder struct {
	v *SimpleMovingAverageAggregation
}

// NewSimpleMovingAverageAggregation provides a builder for the SimpleMovingAverageAggregation struct.
func NewSimpleMovingAverageAggregationBuilder() *SimpleMovingAverageAggregationBuilder {
	r := SimpleMovingAverageAggregationBuilder{
		&SimpleMovingAverageAggregation{},
	}

	r.v.Model = "simple"

	return &r
}

// Build finalize the chain and returns the SimpleMovingAverageAggregation struct
func (rb *SimpleMovingAverageAggregationBuilder) Build() SimpleMovingAverageAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.

func (rb *SimpleMovingAverageAggregationBuilder) BucketsPath(bucketspath *BucketsPathBuilder) *SimpleMovingAverageAggregationBuilder {
	v := bucketspath.Build()
	rb.v.BucketsPath = &v
	return rb
}

func (rb *SimpleMovingAverageAggregationBuilder) Format(format string) *SimpleMovingAverageAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *SimpleMovingAverageAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *SimpleMovingAverageAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

func (rb *SimpleMovingAverageAggregationBuilder) Meta(meta *MetadataBuilder) *SimpleMovingAverageAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *SimpleMovingAverageAggregationBuilder) Minimize(minimize bool) *SimpleMovingAverageAggregationBuilder {
	rb.v.Minimize = &minimize
	return rb
}

func (rb *SimpleMovingAverageAggregationBuilder) Name(name string) *SimpleMovingAverageAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *SimpleMovingAverageAggregationBuilder) Predict(predict int) *SimpleMovingAverageAggregationBuilder {
	rb.v.Predict = &predict
	return rb
}

func (rb *SimpleMovingAverageAggregationBuilder) Settings(settings *EmptyObjectBuilder) *SimpleMovingAverageAggregationBuilder {
	v := settings.Build()
	rb.v.Settings = v
	return rb
}

func (rb *SimpleMovingAverageAggregationBuilder) Window(window int) *SimpleMovingAverageAggregationBuilder {
	rb.v.Window = &window
	return rb
}
