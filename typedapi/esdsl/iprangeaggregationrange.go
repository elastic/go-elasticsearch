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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _ipRangeAggregationRange struct {
	v *types.IpRangeAggregationRange
}

func NewIpRangeAggregationRange() *_ipRangeAggregationRange {

	return &_ipRangeAggregationRange{v: types.NewIpRangeAggregationRange()}

}

// Start of the range.
func (s *_ipRangeAggregationRange) From(from string) *_ipRangeAggregationRange {

	s.v.From = &from

	return s
}

// IP range defined as a CIDR mask.
func (s *_ipRangeAggregationRange) Mask(mask string) *_ipRangeAggregationRange {

	s.v.Mask = &mask

	return s
}

// End of the range.
func (s *_ipRangeAggregationRange) To(to string) *_ipRangeAggregationRange {

	s.v.To = &to

	return s
}

func (s *_ipRangeAggregationRange) IpRangeAggregationRangeCaster() *types.IpRangeAggregationRange {
	return s.v
}
