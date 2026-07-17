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
// https://github.com/elastic/elasticsearch-specification/tree/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c

package types

// The region policy configuration. Specify exactly one of `allowed_geos` or
// `allowed_regions`.
//
// https://github.com/elastic/elasticsearch-specification/blob/8076b1c4ff3b8bd4eb5372bc75372577a21d1b0c/specification/inference/_types/RegionPolicy.ts#L36-L52
type RegionPolicy struct {
	// AllowedGeos The list of allowed geographic areas. Mutually exclusive with
	// `allowed_regions`.
	AllowedGeos []string `json:"allowed_geos,omitempty"`
	// AllowedRegions The list of allowed cloud service provider regions. Mutually exclusive with
	// `allowed_geos`.
	AllowedRegions []CspRegion `json:"allowed_regions,omitempty"`
}

// NewRegionPolicy returns a RegionPolicy.
func NewRegionPolicy() *RegionPolicy {
	r := &RegionPolicy{}

	return r
}

type RegionPolicyVariant interface {
	RegionPolicyCaster() *RegionPolicy
}

func (s *RegionPolicy) RegionPolicyCaster() *RegionPolicy {
	return s
}
