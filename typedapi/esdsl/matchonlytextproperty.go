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

type _matchOnlyTextProperty struct {
	v *types.MatchOnlyTextProperty
}

func NewMatchOnlyTextProperty() *_matchOnlyTextProperty {

	return &_matchOnlyTextProperty{v: types.NewMatchOnlyTextProperty()}

}

// Allows you to copy the values of multiple fields into a group
// field, which can then be queried as a single field.
func (s *_matchOnlyTextProperty) CopyTo(fields ...string) *_matchOnlyTextProperty {

	s.v.CopyTo = fields

	return s
}

// Multi-fields allow the same string value to be indexed in multiple ways for
// different purposes, such as one
// field for search and a multi-field for sorting and aggregations, or the same
// string value analyzed by different analyzers.
func (s *_matchOnlyTextProperty) Fields(fields map[string]types.Property) *_matchOnlyTextProperty {

	s.v.Fields = fields
	return s
}

func (s *_matchOnlyTextProperty) AddField(key string, value types.PropertyVariant) *_matchOnlyTextProperty {

	var tmp map[string]types.Property
	if s.v.Fields == nil {
		s.v.Fields = make(map[string]types.Property)
	} else {
		tmp = s.v.Fields
	}

	tmp[key] = *value.PropertyCaster()

	s.v.Fields = tmp
	return s
}

// Metadata about the field.
func (s *_matchOnlyTextProperty) Meta(meta map[string]string) *_matchOnlyTextProperty {

	s.v.Meta = meta
	return s
}

func (s *_matchOnlyTextProperty) AddMeta(key string, value string) *_matchOnlyTextProperty {

	var tmp map[string]string
	if s.v.Meta == nil {
		s.v.Meta = make(map[string]string)
	} else {
		tmp = s.v.Meta
	}

	tmp[key] = value

	s.v.Meta = tmp
	return s
}

func (s *_matchOnlyTextProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_matchOnlyTextProperty) MatchOnlyTextPropertyCaster() *types.MatchOnlyTextProperty {
	return s.v
}
