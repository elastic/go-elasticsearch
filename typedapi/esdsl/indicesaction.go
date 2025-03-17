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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _indicesAction struct {
	v *types.IndicesAction
}

func NewIndicesAction() *_indicesAction {
	return &_indicesAction{v: types.NewIndicesAction()}
}

// Adds a data stream or index to an alias.
// If the alias doesn’t exist, the `add` action creates it.
func (s *_indicesAction) Add(add types.AddActionVariant) *_indicesAction {

	s.v.Add = add.AddActionCaster()

	return s
}

// AdditionalIndicesActionProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_indicesAction) AdditionalIndicesActionProperty(key string, value json.RawMessage) *_indicesAction {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalIndicesActionProperty = tmp
	return s
}

// Removes a data stream or index from an alias.
func (s *_indicesAction) Remove(remove types.RemoveActionVariant) *_indicesAction {

	s.v.Remove = remove.RemoveActionCaster()

	return s
}

// Deletes an index.
// You cannot use this action on aliases or data streams.
func (s *_indicesAction) RemoveIndex(removeindex types.RemoveIndexActionVariant) *_indicesAction {

	s.v.RemoveIndex = removeindex.RemoveIndexActionCaster()

	return s
}

func (s *_indicesAction) IndicesActionCaster() *types.IndicesAction {
	return s.v
}
