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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// BucketCorrelationFunctionCountCorrelationIndicator type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/aggregations/pipeline.ts#L134-L152
type BucketCorrelationFunctionCountCorrelationIndicator struct {
	// DocCount The total number of documents that initially created the expectations. It’s
	// required to be greater
	// than or equal to the sum of all values in the buckets_path as this is the
	// originating superset of data
	// to which the term values are correlated.
	DocCount int `json:"doc_count"`
	// Expectations An array of numbers with which to correlate the configured `bucket_path`
	// values.
	// The length of this value must always equal the number of buckets returned by
	// the `bucket_path`.
	Expectations []Float64 `json:"expectations"`
	// Fractions An array of fractions to use when averaging and calculating variance. This
	// should be used if
	// the pre-calculated data and the buckets_path have known gaps. The length of
	// fractions, if provided,
	// must equal expectations.
	Fractions []Float64 `json:"fractions,omitempty"`
}

// NewBucketCorrelationFunctionCountCorrelationIndicator returns a BucketCorrelationFunctionCountCorrelationIndicator.
func NewBucketCorrelationFunctionCountCorrelationIndicator() *BucketCorrelationFunctionCountCorrelationIndicator {
	r := &BucketCorrelationFunctionCountCorrelationIndicator{}

	return r
}
