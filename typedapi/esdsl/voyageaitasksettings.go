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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _voyageAITaskSettings struct {
	v *types.VoyageAITaskSettings
}

func NewVoyageAITaskSettings() *_voyageAITaskSettings {

	return &_voyageAITaskSettings{v: types.NewVoyageAITaskSettings()}

}

// Type of the input text.
// Permitted values: `ingest` (maps to `document` in the VoyageAI
// documentation), `search` (maps to `query` in the VoyageAI documentation).
// Only for the `text_embedding` task type.
func (s *_voyageAITaskSettings) InputType(inputtype string) *_voyageAITaskSettings {

	s.v.InputType = &inputtype

	return s
}

// Whether to return the source documents in the response.
// Only for the `rerank` task type.
func (s *_voyageAITaskSettings) ReturnDocuments(returndocuments bool) *_voyageAITaskSettings {

	s.v.ReturnDocuments = &returndocuments

	return s
}

// The number of most relevant documents to return.
// If not specified, the reranking results of all documents will be returned.
// Only for the `rerank` task type.
func (s *_voyageAITaskSettings) TopK(topk int) *_voyageAITaskSettings {

	s.v.TopK = &topk

	return s
}

// Whether to truncate the input texts to fit within the context length.
func (s *_voyageAITaskSettings) Truncation(truncation bool) *_voyageAITaskSettings {

	s.v.Truncation = &truncation

	return s
}

func (s *_voyageAITaskSettings) VoyageAITaskSettingsCaster() *types.VoyageAITaskSettings {
	return s.v
}
