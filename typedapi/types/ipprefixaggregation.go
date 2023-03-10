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

import (
	"encoding/json"
)

// IpPrefixAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/aggregations/bucket.ts#L514-L543
type IpPrefixAggregation struct {
	// AppendPrefixLength Defines whether the prefix length is appended to IP address keys in the
	// response.
	AppendPrefixLength *bool `json:"append_prefix_length,omitempty"`
	// Field The document IP address field to aggregation on. The field mapping type must
	// be `ip`
	Field string `json:"field"`
	// IsIpv6 Defines whether the prefix applies to IPv6 addresses.
	IsIpv6 *bool `json:"is_ipv6,omitempty"`
	// Keyed Defines whether buckets are returned as a hash rather than an array in the
	// response.
	Keyed *bool                      `json:"keyed,omitempty"`
	Meta  map[string]json.RawMessage `json:"meta,omitempty"`
	// MinDocCount Minimum number of documents for buckets to be included in the response.
	MinDocCount *int64  `json:"min_doc_count,omitempty"`
	Name        *string `json:"name,omitempty"`
	// PrefixLength Length of the network prefix. For IPv4 addresses the accepted range is [0,
	// 32].
	// For IPv6 addresses the accepted range is [0, 128].
	PrefixLength int `json:"prefix_length"`
}

// NewIpPrefixAggregation returns a IpPrefixAggregation.
func NewIpPrefixAggregation() *IpPrefixAggregation {
	r := &IpPrefixAggregation{}

	return r
}
