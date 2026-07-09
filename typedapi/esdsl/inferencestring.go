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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontentformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontenttype"
)

type _inferenceString struct {
	v *types.InferenceString
}

func NewInferenceString(type_ embeddingcontenttype.EmbeddingContentType, value string) *_inferenceString {

	tmp := &_inferenceString{v: types.NewInferenceString()}

	tmp.Type(type_)

	tmp.Value(value)

	return tmp

}

func (s *_inferenceString) Format(format embeddingcontentformat.EmbeddingContentFormat) *_inferenceString {

	s.v.Format = &format

	return s
}

func (s *_inferenceString) Type(type_ embeddingcontenttype.EmbeddingContentType) *_inferenceString {

	s.v.Type = type_
	return s
}

func (s *_inferenceString) Value(value string) *_inferenceString {

	s.v.Value = value

	return s
}

func (s *_inferenceString) InferenceStringCaster() *types.InferenceString {
	return s.v
}
