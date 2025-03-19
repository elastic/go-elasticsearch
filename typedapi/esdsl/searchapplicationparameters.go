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

type _searchApplicationParameters struct {
	v *types.SearchApplicationParameters
}

func NewSearchApplicationParameters() *_searchApplicationParameters {

	return &_searchApplicationParameters{v: types.NewSearchApplicationParameters()}

}

// Analytics collection associated to the Search Application.
func (s *_searchApplicationParameters) AnalyticsCollectionName(name string) *_searchApplicationParameters {

	s.v.AnalyticsCollectionName = &name

	return s
}

// Indices that are part of the Search Application.
func (s *_searchApplicationParameters) Indices(indices ...string) *_searchApplicationParameters {

	for _, v := range indices {

		s.v.Indices = append(s.v.Indices, v)

	}
	return s
}

// Search template to use on search operations.
func (s *_searchApplicationParameters) Template(template types.SearchApplicationTemplateVariant) *_searchApplicationParameters {

	s.v.Template = template.SearchApplicationTemplateCaster()

	return s
}

func (s *_searchApplicationParameters) SearchApplicationParametersCaster() *types.SearchApplicationParameters {
	return s.v
}
