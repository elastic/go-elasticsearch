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

// LinearMovingAverageAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/pipeline.ts#L190-L193
type LinearMovingAverageAggregation struct {
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

// LinearMovingAverageAggregationBuilder holds LinearMovingAverageAggregation struct and provides a builder API.
type LinearMovingAverageAggregationBuilder struct {
	v *LinearMovingAverageAggregation
}

// NewLinearMovingAverageAggregation provides a builder for the LinearMovingAverageAggregation struct.
func NewLinearMovingAverageAggregationBuilder() *LinearMovingAverageAggregationBuilder {
	r := LinearMovingAverageAggregationBuilder{
		&LinearMovingAverageAggregation{},
	}

	r.v.Model = "linear"

	return &r
}

// Build finalize the chain and returns the LinearMovingAverageAggregation struct
func (rb *LinearMovingAverageAggregationBuilder) Build() LinearMovingAverageAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.

func (rb *LinearMovingAverageAggregationBuilder) BucketsPath(bucketspath *BucketsPathBuilder) *LinearMovingAverageAggregationBuilder {
	v := bucketspath.Build()
	rb.v.BucketsPath = &v
	return rb
}

func (rb *LinearMovingAverageAggregationBuilder) Format(format string) *LinearMovingAverageAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *LinearMovingAverageAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *LinearMovingAverageAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

func (rb *LinearMovingAverageAggregationBuilder) Meta(meta *MetadataBuilder) *LinearMovingAverageAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *LinearMovingAverageAggregationBuilder) Minimize(minimize bool) *LinearMovingAverageAggregationBuilder {
	rb.v.Minimize = &minimize
	return rb
}

func (rb *LinearMovingAverageAggregationBuilder) Name(name string) *LinearMovingAverageAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *LinearMovingAverageAggregationBuilder) Predict(predict int) *LinearMovingAverageAggregationBuilder {
	rb.v.Predict = &predict
	return rb
}

func (rb *LinearMovingAverageAggregationBuilder) Settings(settings *EmptyObjectBuilder) *LinearMovingAverageAggregationBuilder {
	v := settings.Build()
	rb.v.Settings = v
	return rb
}

func (rb *LinearMovingAverageAggregationBuilder) Window(window int) *LinearMovingAverageAggregationBuilder {
	rb.v.Window = &window
	return rb
}
