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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _decisionInput struct {
	v types.DecisionInput
}

func NewDecisionInput() *_decisionInput {
	return &_decisionInput{v: nil}
}

func (u *_decisionInput) String(string string) *_decisionInput {

	u.v = &string

	return u
}

func (u *_decisionInput) EmbeddingContentObjectItem(embeddingcontentobjectitem types.EmbeddingContentObjectItemVariant) *_decisionInput {

	u.v = embeddingcontentobjectitem.EmbeddingContentObjectItemCaster()

	return u
}

// Interface implementation for EmbeddingContentObjectItem in DecisionInput union
func (u *_embeddingContentObjectItem) DecisionInputCaster() *types.DecisionInput {
	t := types.DecisionInput(u.v)
	return &t
}

func (u *_decisionInput) DecisionInputCaster() *types.DecisionInput {
	return &u.v
}
