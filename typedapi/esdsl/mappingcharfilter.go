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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _mappingCharFilter struct {
	v *types.MappingCharFilter
}

func NewMappingCharFilter() *_mappingCharFilter {

	return &_mappingCharFilter{v: types.NewMappingCharFilter()}

}

func (s *_mappingCharFilter) Mappings(mappings ...string) *_mappingCharFilter {

	for _, v := range mappings {

		s.v.Mappings = append(s.v.Mappings, v)

	}
	return s
}

func (s *_mappingCharFilter) MappingsPath(mappingspath string) *_mappingCharFilter {

	s.v.MappingsPath = &mappingspath

	return s
}

func (s *_mappingCharFilter) Version(versionstring string) *_mappingCharFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_mappingCharFilter) MappingCharFilterCaster() *types.MappingCharFilter {
	return s.v
}
