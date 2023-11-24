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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

// BucketCorrelationFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/aggregations/pipeline.ts#L137-L142
type BucketCorrelationFunction struct {
	// CountCorrelation The configuration to calculate a count correlation. This function is designed
	// for determining the correlation of a term value and a given metric.
	CountCorrelation BucketCorrelationFunctionCountCorrelation `json:"count_correlation"`
}

// NewBucketCorrelationFunction returns a BucketCorrelationFunction.
func NewBucketCorrelationFunction() *BucketCorrelationFunction {
	r := &BucketCorrelationFunction{}

	return r
}
