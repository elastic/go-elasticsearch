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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ipPrefixAggregation struct {
	v *types.IpPrefixAggregation
}

// A bucket aggregation that groups documents based on the network or
// sub-network of an IP address.
func NewIpPrefixAggregation(prefixlength int) *_ipPrefixAggregation {

	tmp := &_ipPrefixAggregation{v: types.NewIpPrefixAggregation()}

	tmp.PrefixLength(prefixlength)

	return tmp

}

func (s *_ipPrefixAggregation) AppendPrefixLength(appendprefixlength bool) *_ipPrefixAggregation {

	s.v.AppendPrefixLength = &appendprefixlength

	return s
}

func (s *_ipPrefixAggregation) Field(field string) *_ipPrefixAggregation {

	s.v.Field = field

	return s
}

func (s *_ipPrefixAggregation) IsIpv6(isipv6 bool) *_ipPrefixAggregation {

	s.v.IsIpv6 = &isipv6

	return s
}

func (s *_ipPrefixAggregation) Keyed(keyed bool) *_ipPrefixAggregation {

	s.v.Keyed = &keyed

	return s
}

func (s *_ipPrefixAggregation) MinDocCount(mindoccount int64) *_ipPrefixAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

func (s *_ipPrefixAggregation) PrefixLength(prefixlength int) *_ipPrefixAggregation {

	s.v.PrefixLength = prefixlength

	return s
}

func (s *_ipPrefixAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.IpPrefix = s.v

	return container
}

func (s *_ipPrefixAggregation) IpPrefixAggregationCaster() *types.IpPrefixAggregation {
	return s.v
}
