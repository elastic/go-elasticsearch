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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// BucketKsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/aggregations/pipeline.ts#L79-L112
type BucketKsAggregation struct {
	// Alternative A list of string values indicating which K-S test alternative to calculate.
	// The valid values
	// are: "greater", "less", "two_sided". This parameter is key for determining
	// the K-S statistic used
	// when calculating the K-S test. Default value is all possible alternative
	// hypotheses.
	Alternative []string `json:"alternative,omitempty"`
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath `json:"buckets_path,omitempty"`
	// Fractions A list of doubles indicating the distribution of the samples with which to
	// compare to the `buckets_path` results.
	// In typical usage this is the overall proportion of documents in each bucket,
	// which is compared with the actual
	// document proportions in each bucket from the sibling aggregation counts. The
	// default is to assume that overall
	// documents are uniformly distributed on these buckets, which they would be if
	// one used equal percentiles of a
	// metric to define the bucket end points.
	Fractions []float64              `json:"fractions,omitempty"`
	Meta      map[string]interface{} `json:"meta,omitempty"`
	Name      *string                `json:"name,omitempty"`
	// SamplingMethod Indicates the sampling methodology when calculating the K-S test. Note, this
	// is sampling of the returned values.
	// This determines the cumulative distribution function (CDF) points used
	// comparing the two samples. Default is
	// `upper_tail`, which emphasizes the upper end of the CDF points. Valid options
	// are: `upper_tail`, `uniform`,
	// and `lower_tail`.
	SamplingMethod *string `json:"sampling_method,omitempty"`
}

// NewBucketKsAggregation returns a BucketKsAggregation.
func NewBucketKsAggregation() *BucketKsAggregation {
	r := &BucketKsAggregation{}

	return r
}
