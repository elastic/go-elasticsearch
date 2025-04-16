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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _indicesModifyAction struct {
	v *types.IndicesModifyAction
}

func NewIndicesModifyAction() *_indicesModifyAction {
	return &_indicesModifyAction{v: types.NewIndicesModifyAction()}
}

func (s *_indicesModifyAction) AddBackingIndex(addbackingindex types.IndexAndDataStreamActionVariant) *_indicesModifyAction {

	s.v.AddBackingIndex = addbackingindex.IndexAndDataStreamActionCaster()

	return s
}

// AdditionalIndicesModifyActionProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_indicesModifyAction) AdditionalIndicesModifyActionProperty(key string, value json.RawMessage) *_indicesModifyAction {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalIndicesModifyActionProperty = tmp
	return s
}

func (s *_indicesModifyAction) RemoveBackingIndex(removebackingindex types.IndexAndDataStreamActionVariant) *_indicesModifyAction {

	s.v.RemoveBackingIndex = removebackingindex.IndexAndDataStreamActionCaster()

	return s
}

func (s *_indicesModifyAction) IndicesModifyActionCaster() *types.IndicesModifyAction {
	return s.v
}
