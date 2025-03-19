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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _semanticTextProperty struct {
	v *types.SemanticTextProperty
}

func NewSemanticTextProperty() *_semanticTextProperty {

	return &_semanticTextProperty{v: types.NewSemanticTextProperty()}

}

// Inference endpoint that will be used to generate embeddings for the field.
// This parameter cannot be updated. Use the Create inference API to create the
// endpoint.
// If `search_inference_id` is specified, the inference endpoint will only be
// used at index time.
func (s *_semanticTextProperty) InferenceId(id string) *_semanticTextProperty {

	s.v.InferenceId = &id

	return s
}

func (s *_semanticTextProperty) Meta(meta map[string]string) *_semanticTextProperty {

	s.v.Meta = meta
	return s
}

func (s *_semanticTextProperty) AddMeta(key string, value string) *_semanticTextProperty {

	var tmp map[string]string
	if s.v.Meta == nil {
		s.v.Meta = make(map[string]string)
	} else {
		tmp = s.v.Meta
	}

	tmp[key] = value

	s.v.Meta = tmp
	return s
}

// Inference endpoint that will be used to generate embeddings at query time.
// You can update this parameter by using the Update mapping API. Use the Create
// inference API to create the endpoint.
// If not specified, the inference endpoint defined by inference_id will be used
// at both index and query time.
func (s *_semanticTextProperty) SearchInferenceId(id string) *_semanticTextProperty {

	s.v.SearchInferenceId = &id

	return s
}

func (s *_semanticTextProperty) DynamicTemplateCaster() *types.DynamicTemplate {
	container := types.NewDynamicTemplate()

	container.Mapping = s.v

	return container
}

func (s *_semanticTextProperty) SemanticTextPropertyCaster() *types.SemanticTextProperty {
	return s.v
}
