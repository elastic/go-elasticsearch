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

// This is provide all the types that are part of the union.
type _tokenFilter struct {
	v types.TokenFilter
}

func NewTokenFilter() *_tokenFilter {
	return &_tokenFilter{v: nil}
}

func (u *_tokenFilter) String(string string) *_tokenFilter {

	u.v = &string

	return u
}

func (u *_tokenFilter) TokenFilterDefinition(tokenfilterdefinition types.TokenFilterDefinitionVariant) *_tokenFilter {

	u.v = *tokenfilterdefinition.TokenFilterDefinitionCaster()

	return u
}

// Interface implementation for TokenFilterDefinition in TokenFilter union
func (u *_tokenFilterDefinition) TokenFilterCaster() *types.TokenFilter {
	t := types.TokenFilter(u.v)
	return &t
}

func (u *_tokenFilter) TokenFilterCaster() *types.TokenFilter {
	return &u.v
}
