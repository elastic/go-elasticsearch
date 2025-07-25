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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _completionToolType struct {
	v types.CompletionToolType
}

func NewCompletionToolType() *_completionToolType {
	return &_completionToolType{v: nil}
}

func (u *_completionToolType) String(string string) *_completionToolType {

	u.v = &string

	return u
}

func (u *_completionToolType) CompletionToolChoice(completiontoolchoice types.CompletionToolChoiceVariant) *_completionToolType {

	u.v = completiontoolchoice.CompletionToolChoiceCaster()

	return u
}

// Interface implementation for CompletionToolChoice in CompletionToolType union
func (u *_completionToolChoice) CompletionToolTypeCaster() *types.CompletionToolType {
	t := types.CompletionToolType(u.v)
	return &t
}

func (u *_completionToolType) CompletionToolTypeCaster() *types.CompletionToolType {
	return &u.v
}
