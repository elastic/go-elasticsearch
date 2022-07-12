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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// BucketCorrelationFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/pipeline.ts#L110-L115
type BucketCorrelationFunction struct {
	// CountCorrelation The configuration to calculate a count correlation. This function is designed
	// for determining the correlation of a term value and a given metric.
	CountCorrelation BucketCorrelationFunctionCountCorrelation `json:"count_correlation"`
}

// BucketCorrelationFunctionBuilder holds BucketCorrelationFunction struct and provides a builder API.
type BucketCorrelationFunctionBuilder struct {
	v *BucketCorrelationFunction
}

// NewBucketCorrelationFunction provides a builder for the BucketCorrelationFunction struct.
func NewBucketCorrelationFunctionBuilder() *BucketCorrelationFunctionBuilder {
	r := BucketCorrelationFunctionBuilder{
		&BucketCorrelationFunction{},
	}

	return &r
}

// Build finalize the chain and returns the BucketCorrelationFunction struct
func (rb *BucketCorrelationFunctionBuilder) Build() BucketCorrelationFunction {
	return *rb.v
}

// CountCorrelation The configuration to calculate a count correlation. This function is designed
// for determining the correlation of a term value and a given metric.

func (rb *BucketCorrelationFunctionBuilder) CountCorrelation(countcorrelation *BucketCorrelationFunctionCountCorrelationBuilder) *BucketCorrelationFunctionBuilder {
	v := countcorrelation.Build()
	rb.v.CountCorrelation = v
	return rb
}
