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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _configuration struct {
	v *types.Configuration
}

func NewConfiguration() *_configuration {

	return &_configuration{v: types.NewConfiguration()}

}

func (s *_configuration) FeatureStates(featurestates ...string) *_configuration {

	for _, v := range featurestates {

		s.v.FeatureStates = append(s.v.FeatureStates, v)

	}
	return s
}

func (s *_configuration) IgnoreUnavailable(ignoreunavailable bool) *_configuration {

	s.v.IgnoreUnavailable = &ignoreunavailable

	return s
}

func (s *_configuration) IncludeGlobalState(includeglobalstate bool) *_configuration {

	s.v.IncludeGlobalState = &includeglobalstate

	return s
}

func (s *_configuration) Indices(indices ...string) *_configuration {

	s.v.Indices = indices

	return s
}

func (s *_configuration) Metadata(metadata types.MetadataVariant) *_configuration {

	s.v.Metadata = *metadata.MetadataCaster()

	return s
}

func (s *_configuration) Partial(partial bool) *_configuration {

	s.v.Partial = &partial

	return s
}

func (s *_configuration) ConfigurationCaster() *types.Configuration {
	return s.v
}
