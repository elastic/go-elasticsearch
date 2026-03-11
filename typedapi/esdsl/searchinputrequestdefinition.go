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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
)

type _searchInputRequestDefinition struct {
	v *types.SearchInputRequestDefinition
}

func NewSearchInputRequestDefinition() *_searchInputRequestDefinition {

	return &_searchInputRequestDefinition{v: types.NewSearchInputRequestDefinition()}

}

func (s *_searchInputRequestDefinition) Body(body types.SearchInputRequestBodyVariant) *_searchInputRequestDefinition {

	s.v.Body = body.SearchInputRequestBodyCaster()

	return s
}

func (s *_searchInputRequestDefinition) Indices(indices ...string) *_searchInputRequestDefinition {

	for _, v := range indices {

		s.v.Indices = append(s.v.Indices, v)

	}
	return s
}

func (s *_searchInputRequestDefinition) IndicesOptions(indicesoptions types.IndicesOptionsVariant) *_searchInputRequestDefinition {

	s.v.IndicesOptions = indicesoptions.IndicesOptionsCaster()

	return s
}

func (s *_searchInputRequestDefinition) RestTotalHitsAsInt(resttotalhitsasint bool) *_searchInputRequestDefinition {

	s.v.RestTotalHitsAsInt = &resttotalhitsasint

	return s
}

func (s *_searchInputRequestDefinition) SearchType(searchtype searchtype.SearchType) *_searchInputRequestDefinition {

	s.v.SearchType = &searchtype
	return s
}

func (s *_searchInputRequestDefinition) Template(template types.SearchTemplateRequestBodyVariant) *_searchInputRequestDefinition {

	s.v.Template = template.SearchTemplateRequestBodyCaster()

	return s
}

func (s *_searchInputRequestDefinition) SearchInputRequestDefinitionCaster() *types.SearchInputRequestDefinition {
	return s.v
}
