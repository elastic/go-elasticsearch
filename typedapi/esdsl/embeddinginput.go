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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _embeddingInput struct {
	v types.EmbeddingInput
}

func NewEmbeddingInput() *_embeddingInput {
	return &_embeddingInput{v: nil}
}

func (u *_embeddingInput) EmbeddingStringInput(embeddingstringinputs ...string) *_embeddingInput {

	u.v = embeddingstringinputs

	return u
}

// Interface implementation for EmbeddingStringInput in EmbeddingInput union
func (u *_embeddingStringInput) EmbeddingInputCaster() *types.EmbeddingInput {
	t := types.EmbeddingInput(u.v)
	return &t
}

func (u *_embeddingInput) EmbeddingContentInput(embeddingcontentinputs ...types.EmbeddingContentObjectVariant) *_embeddingInput {

	convertedItems := make([]types.EmbeddingContentObject, 0, len(embeddingcontentinputs))
	for _, v := range embeddingcontentinputs {
		convertedItems = append(convertedItems, *v.EmbeddingContentObjectCaster())
	}
	u.v = convertedItems

	return u
}

// Interface implementation for EmbeddingContentInput in EmbeddingInput union
func (u *_embeddingContentInput) EmbeddingInputCaster() *types.EmbeddingInput {
	t := types.EmbeddingInput(u.v)
	return &t
}

func (u *_embeddingInput) EmbeddingInputCaster() *types.EmbeddingInput {
	return &u.v
}
