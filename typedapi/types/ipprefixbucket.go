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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

import (
	"encoding/json"
	"fmt"
)

// IpPrefixBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/aggregations/Aggregate.ts#L616-L621
type IpPrefixBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"-"`
	DocCount     int64                       `json:"doc_count"`
	IsIpv6       bool                        `json:"is_ipv6"`
	Key          string                      `json:"key"`
	Netmask      *string                     `json:"netmask,omitempty"`
	PrefixLength int                         `json:"prefix_length"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s IpPrefixBucket) MarshalJSON() ([]byte, error) {
	type opt IpPrefixBucket
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.Aggregations {
		tmp[fmt.Sprintf("%s", key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// IpPrefixBucketBuilder holds IpPrefixBucket struct and provides a builder API.
type IpPrefixBucketBuilder struct {
	v *IpPrefixBucket
}

// NewIpPrefixBucket provides a builder for the IpPrefixBucket struct.
func NewIpPrefixBucketBuilder() *IpPrefixBucketBuilder {
	r := IpPrefixBucketBuilder{
		&IpPrefixBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IpPrefixBucket struct
func (rb *IpPrefixBucketBuilder) Build() IpPrefixBucket {
	return *rb.v
}

func (rb *IpPrefixBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *IpPrefixBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *IpPrefixBucketBuilder) DocCount(doccount int64) *IpPrefixBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *IpPrefixBucketBuilder) IsIpv6(isipv6 bool) *IpPrefixBucketBuilder {
	rb.v.IsIpv6 = isipv6
	return rb
}

func (rb *IpPrefixBucketBuilder) Key(key string) *IpPrefixBucketBuilder {
	rb.v.Key = key
	return rb
}

func (rb *IpPrefixBucketBuilder) Netmask(netmask string) *IpPrefixBucketBuilder {
	rb.v.Netmask = &netmask
	return rb
}

func (rb *IpPrefixBucketBuilder) PrefixLength(prefixlength int) *IpPrefixBucketBuilder {
	rb.v.PrefixLength = prefixlength
	return rb
}
