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

// BucketsSignificantLongTermsBucket holds the union for the following types:
//
//	map[string]SignificantLongTermsBucket
//	[]SignificantLongTermsBucket
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L303-L312
type BucketsSignificantLongTermsBucket interface{}

// BucketsSignificantLongTermsBucketBuilder holds BucketsSignificantLongTermsBucket struct and provides a builder API.
type BucketsSignificantLongTermsBucketBuilder struct {
	v BucketsSignificantLongTermsBucket
}

// NewBucketsSignificantLongTermsBucket provides a builder for the BucketsSignificantLongTermsBucket struct.
func NewBucketsSignificantLongTermsBucketBuilder() *BucketsSignificantLongTermsBucketBuilder {
	return &BucketsSignificantLongTermsBucketBuilder{}
}

// Build finalize the chain and returns the BucketsSignificantLongTermsBucket struct
func (u *BucketsSignificantLongTermsBucketBuilder) Build() BucketsSignificantLongTermsBucket {
	return u.v
}

func (u *BucketsSignificantLongTermsBucketBuilder) Map(values map[string]*SignificantLongTermsBucketBuilder) *BucketsSignificantLongTermsBucketBuilder {
	tmp := make(map[string]SignificantLongTermsBucket, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	u.v = tmp
	return u
}

func (u *BucketsSignificantLongTermsBucketBuilder) SignificantLongTermsBuckets(significantlongtermsbuckets []SignificantLongTermsBucketBuilder) *BucketsSignificantLongTermsBucketBuilder {
	tmp := make([]SignificantLongTermsBucket, len(significantlongtermsbuckets))
	for _, value := range significantlongtermsbuckets {
		tmp = append(tmp, value.Build())
	}
	u.v = tmp
	return u
}
