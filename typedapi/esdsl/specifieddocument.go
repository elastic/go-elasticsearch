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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _specifiedDocument struct {
	v *types.SpecifiedDocument
}

func NewSpecifiedDocument() *_specifiedDocument {

	return &_specifiedDocument{v: types.NewSpecifiedDocument()}

}

func (s *_specifiedDocument) Id(id string) *_specifiedDocument {

	s.v.Id = id

	return s
}

func (s *_specifiedDocument) Index(indexname string) *_specifiedDocument {

	s.v.Index = &indexname

	return s
}

func (s *_specifiedDocument) SpecifiedDocumentCaster() *types.SpecifiedDocument {
	return s.v
}
