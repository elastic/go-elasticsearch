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
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package types

// BucketsIpPrefixBucket holds the union for the following types:
//
//	[]IpPrefixBucket
//	map[string]IpPrefixBucket
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/Aggregate.ts#L307-L316
type BucketsIpPrefixBucket interface{}

// BucketsIpPrefixBucketBuilder holds BucketsIpPrefixBucket struct and provides a builder API.
type BucketsIpPrefixBucketBuilder struct {
	v BucketsIpPrefixBucket
}

// NewBucketsIpPrefixBucket provides a builder for the BucketsIpPrefixBucket struct.
func NewBucketsIpPrefixBucketBuilder() *BucketsIpPrefixBucketBuilder {
	return &BucketsIpPrefixBucketBuilder{}
}

// Build finalize the chain and returns the BucketsIpPrefixBucket struct
func (u *BucketsIpPrefixBucketBuilder) Build() BucketsIpPrefixBucket {
	return u.v
}

func (u *BucketsIpPrefixBucketBuilder) IpPrefixBuckets(ipprefixbuckets []IpPrefixBucketBuilder) *BucketsIpPrefixBucketBuilder {
	tmp := make([]IpPrefixBucket, len(ipprefixbuckets))
	for _, value := range ipprefixbuckets {
		tmp = append(tmp, value.Build())
	}
	u.v = tmp
	return u
}

func (u *BucketsIpPrefixBucketBuilder) Map(values map[string]*IpPrefixBucketBuilder) *BucketsIpPrefixBucketBuilder {
	tmp := make(map[string]IpPrefixBucket, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	u.v = tmp
	return u
}
