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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontentformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontenttype"
)

type _embeddingContentObjectContents struct {
	v *types.EmbeddingContentObjectContents
}

func NewEmbeddingContentObjectContents(type_ embeddingcontenttype.EmbeddingContentType, value string) *_embeddingContentObjectContents {

	tmp := &_embeddingContentObjectContents{v: types.NewEmbeddingContentObjectContents()}

	tmp.Type(type_)

	tmp.Value(value)

	return tmp

}

func (s *_embeddingContentObjectContents) Format(format embeddingcontentformat.EmbeddingContentFormat) *_embeddingContentObjectContents {

	s.v.Format = &format
	return s
}

func (s *_embeddingContentObjectContents) Type(type_ embeddingcontenttype.EmbeddingContentType) *_embeddingContentObjectContents {

	s.v.Type = type_
	return s
}

func (s *_embeddingContentObjectContents) Value(value string) *_embeddingContentObjectContents {

	s.v.Value = value

	return s
}

func (s *_embeddingContentObjectContents) EmbeddingContentObjectContentsCaster() *types.EmbeddingContentObjectContents {
	return s.v
}
