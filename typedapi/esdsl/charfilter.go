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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _charFilter struct {
	v types.CharFilter
}

func NewCharFilter() *_charFilter {
	return &_charFilter{v: nil}
}

func (u *_charFilter) String(string string) *_charFilter {

	u.v = &string

	return u
}

func (u *_charFilter) CharFilterDefinition(charfilterdefinition types.CharFilterDefinitionVariant) *_charFilter {

	u.v = *charfilterdefinition.CharFilterDefinitionCaster()

	return u
}

// Interface implementation for CharFilterDefinition in CharFilter union
func (u *_charFilterDefinition) CharFilterCaster() *types.CharFilter {
	t := types.CharFilter(u.v)
	return &t
}

func (u *_charFilter) CharFilterCaster() *types.CharFilter {
	return &u.v
}
