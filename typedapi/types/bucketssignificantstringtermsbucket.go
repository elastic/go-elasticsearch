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

// BucketsSignificantStringTermsBucket holds the union for the following types:
//
//	map[string]SignificantStringTermsBucket
//	[]SignificantStringTermsBucket
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L303-L312
type BucketsSignificantStringTermsBucket interface{}

// BucketsSignificantStringTermsBucketBuilder holds BucketsSignificantStringTermsBucket struct and provides a builder API.
type BucketsSignificantStringTermsBucketBuilder struct {
	v BucketsSignificantStringTermsBucket
}

// NewBucketsSignificantStringTermsBucket provides a builder for the BucketsSignificantStringTermsBucket struct.
func NewBucketsSignificantStringTermsBucketBuilder() *BucketsSignificantStringTermsBucketBuilder {
	return &BucketsSignificantStringTermsBucketBuilder{}
}

// Build finalize the chain and returns the BucketsSignificantStringTermsBucket struct
func (u *BucketsSignificantStringTermsBucketBuilder) Build() BucketsSignificantStringTermsBucket {
	return u.v
}

func (u *BucketsSignificantStringTermsBucketBuilder) Map(values map[string]*SignificantStringTermsBucketBuilder) *BucketsSignificantStringTermsBucketBuilder {
	tmp := make(map[string]SignificantStringTermsBucket, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	u.v = tmp
	return u
}

func (u *BucketsSignificantStringTermsBucketBuilder) SignificantStringTermsBuckets(significantstringtermsbuckets []SignificantStringTermsBucketBuilder) *BucketsSignificantStringTermsBucketBuilder {
	tmp := make([]SignificantStringTermsBucket, len(significantstringtermsbuckets))
	for _, value := range significantstringtermsbuckets {
		tmp = append(tmp, value.Build())
	}
	u.v = tmp
	return u
}
