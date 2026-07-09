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

type _embeddingContentObject struct {
	v *types.EmbeddingContentObject
}

func NewEmbeddingContentObject() *_embeddingContentObject {

	return &_embeddingContentObject{v: types.NewEmbeddingContentObject()}

}

func (s *_embeddingContentObject) Content(embeddingcontentobjectgroups ...types.EmbeddingContentObjectItemVariant) *_embeddingContentObject {

	convertedItems := make([]types.EmbeddingContentObjectItem, 0, len(embeddingcontentobjectgroups))
	for _, v := range embeddingcontentobjectgroups {
		convertedItems = append(convertedItems, *v.EmbeddingContentObjectItemCaster())
	}
	s.v.Content = convertedItems

	return s
}

func (s *_embeddingContentObject) ContentValues(embeddingcontentobjectgroupvalues []types.EmbeddingContentObjectItem) *_embeddingContentObject {

	s.v.Content = embeddingcontentobjectgroupvalues
	return s
}

func (s *_embeddingContentObject) EmbeddingContentObjectCaster() *types.EmbeddingContentObject {
	return s.v
}
