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

// BucketsHistogramBucket holds the union for the following types:
//
//	[]HistogramBucket
//	map[string]HistogramBucket
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L303-L312
type BucketsHistogramBucket interface{}

// BucketsHistogramBucketBuilder holds BucketsHistogramBucket struct and provides a builder API.
type BucketsHistogramBucketBuilder struct {
	v BucketsHistogramBucket
}

// NewBucketsHistogramBucket provides a builder for the BucketsHistogramBucket struct.
func NewBucketsHistogramBucketBuilder() *BucketsHistogramBucketBuilder {
	return &BucketsHistogramBucketBuilder{}
}

// Build finalize the chain and returns the BucketsHistogramBucket struct
func (u *BucketsHistogramBucketBuilder) Build() BucketsHistogramBucket {
	return u.v
}

func (u *BucketsHistogramBucketBuilder) HistogramBuckets(histogrambuckets []HistogramBucketBuilder) *BucketsHistogramBucketBuilder {
	tmp := make([]HistogramBucket, len(histogrambuckets))
	for _, value := range histogrambuckets {
		tmp = append(tmp, value.Build())
	}
	u.v = tmp
	return u
}

func (u *BucketsHistogramBucketBuilder) Map(values map[string]*HistogramBucketBuilder) *BucketsHistogramBucketBuilder {
	tmp := make(map[string]HistogramBucket, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	u.v = tmp
	return u
}
