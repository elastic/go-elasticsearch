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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _document struct {
	v *types.Document
}

func NewDocument(source_ json.RawMessage) *_document {

	tmp := &_document{v: types.NewDocument()}

	tmp.Source_(source_)

	return tmp

}

// Unique identifier for the document.
// This ID must be unique within the `_index`.
func (s *_document) Id_(id string) *_document {

	s.v.Id_ = &id

	return s
}

// Name of the index containing the document.
func (s *_document) Index_(indexname string) *_document {

	s.v.Index_ = &indexname

	return s
}

// JSON body for the document.
func (s *_document) Source_(source_ json.RawMessage) *_document {

	s.v.Source_ = source_

	return s
}

func (s *_document) DocumentCaster() *types.Document {
	return s.v
}
