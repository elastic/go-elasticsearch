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
// https://github.com/elastic/elasticsearch-specification/tree/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _cspRegion struct {
	v *types.CspRegion
}

// The list of allowed cloud service provider regions. Mutually exclusive with
// `allowed_geos`.
func NewCspRegion(csp string, region string) *_cspRegion {

	tmp := &_cspRegion{v: types.NewCspRegion()}

	tmp.Csp(csp)

	tmp.Region(region)

	return tmp

}

func (s *_cspRegion) Csp(csp string) *_cspRegion {

	s.v.Csp = csp

	return s
}

func (s *_cspRegion) Region(region string) *_cspRegion {

	s.v.Region = region

	return s
}

func (s *_cspRegion) RegionPolicyCaster() *types.RegionPolicy {
	container := types.NewRegionPolicy()

	container.AllowedRegions = append(container.AllowedRegions, *s.v)

	return container
}

func (s *_cspRegion) CspRegionCaster() *types.CspRegion {
	return s.v
}
