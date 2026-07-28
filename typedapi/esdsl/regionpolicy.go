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
// https://github.com/elastic/elasticsearch-specification/tree/7560c979602e6941815872bdaec801200bc7ec4e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _regionPolicy struct {
	v *types.RegionPolicy
}

func NewRegionPolicy() *_regionPolicy {
	return &_regionPolicy{v: types.NewRegionPolicy()}
}

func (s *_regionPolicy) AllowedGeos(allowedgeos ...string) *_regionPolicy {

	for _, v := range allowedgeos {

		s.v.AllowedGeos = append(s.v.AllowedGeos, v)

	}
	return s
}

func (s *_regionPolicy) AllowedRegions(allowedregions ...types.CspRegionVariant) *_regionPolicy {

	for _, v := range allowedregions {

		s.v.AllowedRegions = append(s.v.AllowedRegions, *v.CspRegionCaster())

	}
	return s
}

func (s *_regionPolicy) AllowedRegionsValues(allowedregionsvalues []types.CspRegion) *_regionPolicy {

	s.v.AllowedRegions = allowedregionsvalues
	return s
}

func (s *_regionPolicy) RegionPolicyCaster() *types.RegionPolicy {
	return s.v
}
