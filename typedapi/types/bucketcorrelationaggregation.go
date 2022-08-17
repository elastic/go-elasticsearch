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

// BucketCorrelationAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/pipeline.ts#L102-L108
type BucketCorrelationAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath `json:"buckets_path,omitempty"`
	// Function The correlation function to execute.
	Function BucketCorrelationFunction `json:"function"`
	Meta     *Metadata                 `json:"meta,omitempty"`
	Name     *string                   `json:"name,omitempty"`
}

// BucketCorrelationAggregationBuilder holds BucketCorrelationAggregation struct and provides a builder API.
type BucketCorrelationAggregationBuilder struct {
	v *BucketCorrelationAggregation
}

// NewBucketCorrelationAggregation provides a builder for the BucketCorrelationAggregation struct.
func NewBucketCorrelationAggregationBuilder() *BucketCorrelationAggregationBuilder {
	r := BucketCorrelationAggregationBuilder{
		&BucketCorrelationAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the BucketCorrelationAggregation struct
func (rb *BucketCorrelationAggregationBuilder) Build() BucketCorrelationAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.

func (rb *BucketCorrelationAggregationBuilder) BucketsPath(bucketspath *BucketsPathBuilder) *BucketCorrelationAggregationBuilder {
	v := bucketspath.Build()
	rb.v.BucketsPath = &v
	return rb
}

// Function The correlation function to execute.

func (rb *BucketCorrelationAggregationBuilder) Function(function *BucketCorrelationFunctionBuilder) *BucketCorrelationAggregationBuilder {
	v := function.Build()
	rb.v.Function = v
	return rb
}

func (rb *BucketCorrelationAggregationBuilder) Meta(meta *MetadataBuilder) *BucketCorrelationAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *BucketCorrelationAggregationBuilder) Name(name string) *BucketCorrelationAggregationBuilder {
	rb.v.Name = &name
	return rb
}
