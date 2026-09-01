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
// https://github.com/elastic/elasticsearch-specification/tree/9665eef0d78c41f20c4c83e69b7c8155efd58f24

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _rerankInput struct {
	v types.RerankInput
}

func NewRerankInput() *_rerankInput {
	return &_rerankInput{v: nil}
}

func (u *_rerankInput) RerankStringInput(rerankstringinputs ...string) *_rerankInput {

	u.v = rerankstringinputs

	return u
}

// Interface implementation for RerankStringInput in RerankInput union
func (u *_rerankStringInput) RerankInputCaster() *types.RerankInput {
	t := types.RerankInput(u.v)
	return &t
}

func (u *_rerankInput) RerankObjectInput(rerankobjectinputs ...types.RerankInputObjectVariant) *_rerankInput {

	convertedItems := make([]types.RerankInputObject, 0, len(rerankobjectinputs))
	for _, v := range rerankobjectinputs {
		convertedItems = append(convertedItems, *v.RerankInputObjectCaster())
	}
	u.v = convertedItems

	return u
}

func (u *_rerankInput) RerankObjectInputValues(rerankobjectinputvalues []types.RerankInputObject) *_rerankInput {

	u.v = rerankobjectinputvalues
	return u
}

// Interface implementation for RerankObjectInput in RerankInput union
func (u *_rerankObjectInput) RerankInputCaster() *types.RerankInput {
	t := types.RerankInput(u.v)
	return &t
}

func (u *_rerankInput) RerankInputCaster() *types.RerankInput {
	return &u.v
}
