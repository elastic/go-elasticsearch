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

// SerialDifferencingAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/pipeline.ts#L268-L270
type SerialDifferencingAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Lag         *int                 `json:"lag,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
}

// SerialDifferencingAggregationBuilder holds SerialDifferencingAggregation struct and provides a builder API.
type SerialDifferencingAggregationBuilder struct {
	v *SerialDifferencingAggregation
}

// NewSerialDifferencingAggregation provides a builder for the SerialDifferencingAggregation struct.
func NewSerialDifferencingAggregationBuilder() *SerialDifferencingAggregationBuilder {
	r := SerialDifferencingAggregationBuilder{
		&SerialDifferencingAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the SerialDifferencingAggregation struct
func (rb *SerialDifferencingAggregationBuilder) Build() SerialDifferencingAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.

func (rb *SerialDifferencingAggregationBuilder) BucketsPath(bucketspath *BucketsPathBuilder) *SerialDifferencingAggregationBuilder {
	v := bucketspath.Build()
	rb.v.BucketsPath = &v
	return rb
}

func (rb *SerialDifferencingAggregationBuilder) Format(format string) *SerialDifferencingAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *SerialDifferencingAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *SerialDifferencingAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

func (rb *SerialDifferencingAggregationBuilder) Lag(lag int) *SerialDifferencingAggregationBuilder {
	rb.v.Lag = &lag
	return rb
}

func (rb *SerialDifferencingAggregationBuilder) Meta(meta *MetadataBuilder) *SerialDifferencingAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *SerialDifferencingAggregationBuilder) Name(name string) *SerialDifferencingAggregationBuilder {
	rb.v.Name = &name
	return rb
}
