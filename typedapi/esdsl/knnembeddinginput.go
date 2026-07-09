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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _knnEmbeddingInput struct {
	v types.KnnEmbeddingInput
}

func NewKnnEmbeddingInput() *_knnEmbeddingInput {
	return &_knnEmbeddingInput{v: nil}
}

func (u *_knnEmbeddingInput) String(string string) *_knnEmbeddingInput {

	u.v = &string

	return u
}

func (u *_knnEmbeddingInput) InferenceStringGroup(inferencestringgroups ...types.InferenceStringVariant) *_knnEmbeddingInput {

	convertedItems := make([]types.InferenceString, 0, len(inferencestringgroups))
	for _, v := range inferencestringgroups {
		convertedItems = append(convertedItems, *v.InferenceStringCaster())
	}
	u.v = convertedItems

	return u
}

func (u *_knnEmbeddingInput) InferenceStringGroupValues(inferencestringgroupvalues []types.InferenceString) *_knnEmbeddingInput {

	u.v = inferencestringgroupvalues
	return u
}

// Interface implementation for InferenceStringGroup in KnnEmbeddingInput union
func (u *_inferenceStringGroup) KnnEmbeddingInputCaster() *types.KnnEmbeddingInput {
	t := types.KnnEmbeddingInput(u.v)
	return &t
}

func (u *_knnEmbeddingInput) KnnEmbeddingInputCaster() *types.KnnEmbeddingInput {
	return &u.v
}
