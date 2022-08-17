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

// InferenceAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/pipeline.ts#L159-L162
type InferenceAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath     *BucketsPath              `json:"buckets_path,omitempty"`
	Format          *string                   `json:"format,omitempty"`
	GapPolicy       *gappolicy.GapPolicy      `json:"gap_policy,omitempty"`
	InferenceConfig *InferenceConfigContainer `json:"inference_config,omitempty"`
	Meta            *Metadata                 `json:"meta,omitempty"`
	ModelId         Name                      `json:"model_id"`
	Name            *string                   `json:"name,omitempty"`
}

// InferenceAggregationBuilder holds InferenceAggregation struct and provides a builder API.
type InferenceAggregationBuilder struct {
	v *InferenceAggregation
}

// NewInferenceAggregation provides a builder for the InferenceAggregation struct.
func NewInferenceAggregationBuilder() *InferenceAggregationBuilder {
	r := InferenceAggregationBuilder{
		&InferenceAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceAggregation struct
func (rb *InferenceAggregationBuilder) Build() InferenceAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.

func (rb *InferenceAggregationBuilder) BucketsPath(bucketspath *BucketsPathBuilder) *InferenceAggregationBuilder {
	v := bucketspath.Build()
	rb.v.BucketsPath = &v
	return rb
}

func (rb *InferenceAggregationBuilder) Format(format string) *InferenceAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *InferenceAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *InferenceAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

func (rb *InferenceAggregationBuilder) InferenceConfig(inferenceconfig *InferenceConfigContainerBuilder) *InferenceAggregationBuilder {
	v := inferenceconfig.Build()
	rb.v.InferenceConfig = &v
	return rb
}

func (rb *InferenceAggregationBuilder) Meta(meta *MetadataBuilder) *InferenceAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *InferenceAggregationBuilder) ModelId(modelid Name) *InferenceAggregationBuilder {
	rb.v.ModelId = modelid
	return rb
}

func (rb *InferenceAggregationBuilder) Name(name string) *InferenceAggregationBuilder {
	rb.v.Name = &name
	return rb
}
