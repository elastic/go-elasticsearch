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

type _fieldLookup struct {
	v *types.FieldLookup
}

func NewFieldLookup() *_fieldLookup {

	return &_fieldLookup{v: types.NewFieldLookup()}

}

// `id` of the document.
func (s *_fieldLookup) Id(id string) *_fieldLookup {

	s.v.Id = id

	return s
}

// Index from which to retrieve the document.
func (s *_fieldLookup) Index(indexname string) *_fieldLookup {

	s.v.Index = &indexname

	return s
}

// Name of the field.
func (s *_fieldLookup) Path(field string) *_fieldLookup {

	s.v.Path = &field

	return s
}

// Custom routing value.
func (s *_fieldLookup) Routing(routing string) *_fieldLookup {

	s.v.Routing = &routing

	return s
}

func (s *_fieldLookup) FieldLookupCaster() *types.FieldLookup {
	return s.v
}
