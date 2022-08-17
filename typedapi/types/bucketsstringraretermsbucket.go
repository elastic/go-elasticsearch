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

// BucketsStringRareTermsBucket holds the union for the following types:
//
//	map[string]StringRareTermsBucket
//	[]StringRareTermsBucket
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L303-L312
type BucketsStringRareTermsBucket interface{}

// BucketsStringRareTermsBucketBuilder holds BucketsStringRareTermsBucket struct and provides a builder API.
type BucketsStringRareTermsBucketBuilder struct {
	v BucketsStringRareTermsBucket
}

// NewBucketsStringRareTermsBucket provides a builder for the BucketsStringRareTermsBucket struct.
func NewBucketsStringRareTermsBucketBuilder() *BucketsStringRareTermsBucketBuilder {
	return &BucketsStringRareTermsBucketBuilder{}
}

// Build finalize the chain and returns the BucketsStringRareTermsBucket struct
func (u *BucketsStringRareTermsBucketBuilder) Build() BucketsStringRareTermsBucket {
	return u.v
}

func (u *BucketsStringRareTermsBucketBuilder) Map(values map[string]*StringRareTermsBucketBuilder) *BucketsStringRareTermsBucketBuilder {
	tmp := make(map[string]StringRareTermsBucket, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	u.v = tmp
	return u
}

func (u *BucketsStringRareTermsBucketBuilder) StringRareTermsBuckets(stringraretermsbuckets []StringRareTermsBucketBuilder) *BucketsStringRareTermsBucketBuilder {
	tmp := make([]StringRareTermsBucket, len(stringraretermsbuckets))
	for _, value := range stringraretermsbuckets {
		tmp = append(tmp, value.Build())
	}
	u.v = tmp
	return u
}
