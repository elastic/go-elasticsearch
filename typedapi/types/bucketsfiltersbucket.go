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

// BucketsFiltersBucket holds the union for the following types:
//
//	[]FiltersBucket
//	map[string]FiltersBucket
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L303-L312
type BucketsFiltersBucket interface{}

// BucketsFiltersBucketBuilder holds BucketsFiltersBucket struct and provides a builder API.
type BucketsFiltersBucketBuilder struct {
	v BucketsFiltersBucket
}

// NewBucketsFiltersBucket provides a builder for the BucketsFiltersBucket struct.
func NewBucketsFiltersBucketBuilder() *BucketsFiltersBucketBuilder {
	return &BucketsFiltersBucketBuilder{}
}

// Build finalize the chain and returns the BucketsFiltersBucket struct
func (u *BucketsFiltersBucketBuilder) Build() BucketsFiltersBucket {
	return u.v
}

func (u *BucketsFiltersBucketBuilder) FiltersBuckets(filtersbuckets []FiltersBucketBuilder) *BucketsFiltersBucketBuilder {
	tmp := make([]FiltersBucket, len(filtersbuckets))
	for _, value := range filtersbuckets {
		tmp = append(tmp, value.Build())
	}
	u.v = tmp
	return u
}

func (u *BucketsFiltersBucketBuilder) Map(values map[string]*FiltersBucketBuilder) *BucketsFiltersBucketBuilder {
	tmp := make(map[string]FiltersBucket, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	u.v = tmp
	return u
}
