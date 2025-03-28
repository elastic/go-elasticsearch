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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _adaptiveAllocations struct {
	v *types.AdaptiveAllocations
}

func NewAdaptiveAllocations() *_adaptiveAllocations {

	return &_adaptiveAllocations{v: types.NewAdaptiveAllocations()}

}

// Turn on `adaptive_allocations`.
func (s *_adaptiveAllocations) Enabled(enabled bool) *_adaptiveAllocations {

	s.v.Enabled = &enabled

	return s
}

// The maximum number of allocations to scale to.
// If set, it must be greater than or equal to `min_number_of_allocations`.
func (s *_adaptiveAllocations) MaxNumberOfAllocations(maxnumberofallocations int) *_adaptiveAllocations {

	s.v.MaxNumberOfAllocations = &maxnumberofallocations

	return s
}

// The minimum number of allocations to scale to.
// If set, it must be greater than or equal to 0.
// If not defined, the deployment scales to 0.
func (s *_adaptiveAllocations) MinNumberOfAllocations(minnumberofallocations int) *_adaptiveAllocations {

	s.v.MinNumberOfAllocations = &minnumberofallocations

	return s
}

func (s *_adaptiveAllocations) AdaptiveAllocationsCaster() *types.AdaptiveAllocations {
	return s.v
}
