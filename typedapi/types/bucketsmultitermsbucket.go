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

// BucketsMultiTermsBucket holds the union for the following types:
//
//	map[string]MultiTermsBucket
//	[]MultiTermsBucket
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L303-L312
type BucketsMultiTermsBucket interface{}

// BucketsMultiTermsBucketBuilder holds BucketsMultiTermsBucket struct and provides a builder API.
type BucketsMultiTermsBucketBuilder struct {
	v BucketsMultiTermsBucket
}

// NewBucketsMultiTermsBucket provides a builder for the BucketsMultiTermsBucket struct.
func NewBucketsMultiTermsBucketBuilder() *BucketsMultiTermsBucketBuilder {
	return &BucketsMultiTermsBucketBuilder{}
}

// Build finalize the chain and returns the BucketsMultiTermsBucket struct
func (u *BucketsMultiTermsBucketBuilder) Build() BucketsMultiTermsBucket {
	return u.v
}

func (u *BucketsMultiTermsBucketBuilder) Map(values map[string]*MultiTermsBucketBuilder) *BucketsMultiTermsBucketBuilder {
	tmp := make(map[string]MultiTermsBucket, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	u.v = tmp
	return u
}

func (u *BucketsMultiTermsBucketBuilder) MultiTermsBuckets(multitermsbuckets []MultiTermsBucketBuilder) *BucketsMultiTermsBucketBuilder {
	tmp := make([]MultiTermsBucket, len(multitermsbuckets))
	for _, value := range multitermsbuckets {
		tmp = append(tmp, value.Build())
	}
	u.v = tmp
	return u
}
