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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/cohereinputtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/coheretruncatetype"
)

type _cohereTaskSettings struct {
	v *types.CohereTaskSettings
}

func NewCohereTaskSettings() *_cohereTaskSettings {

	return &_cohereTaskSettings{v: types.NewCohereTaskSettings()}

}

// For a `text_embedding` task, the type of input passed to the model.
// Valid values are:
//
// * `classification`: Use it for embeddings passed through a text classifier.
// * `clustering`: Use it for the embeddings run through a clustering algorithm.
// * `ingest`: Use it for storing document embeddings in a vector database.
// * `search`: Use it for storing embeddings of search queries run against a
// vector database to find relevant documents.
//
// IMPORTANT: The `input_type` field is required when using embedding models
// `v3` and higher.
func (s *_cohereTaskSettings) InputType(inputtype cohereinputtype.CohereInputType) *_cohereTaskSettings {

	s.v.InputType = &inputtype
	return s
}

// For a `rerank` task, return doc text within the results.
func (s *_cohereTaskSettings) ReturnDocuments(returndocuments bool) *_cohereTaskSettings {

	s.v.ReturnDocuments = &returndocuments

	return s
}

// For a `rerank` task, the number of most relevant documents to return.
// It defaults to the number of the documents.
// If this inference endpoint is used in a `text_similarity_reranker` retriever
// query and `top_n` is set, it must be greater than or equal to
// `rank_window_size` in the query.
func (s *_cohereTaskSettings) TopN(topn int) *_cohereTaskSettings {

	s.v.TopN = &topn

	return s
}

// For a `text_embedding` task, the method to handle inputs longer than the
// maximum token length.
// Valid values are:
//
// * `END`: When the input exceeds the maximum input token length, the end of
// the input is discarded.
// * `NONE`: When the input exceeds the maximum input token length, an error is
// returned.
// * `START`: When the input exceeds the maximum input token length, the start
// of the input is discarded.
func (s *_cohereTaskSettings) Truncate(truncate coheretruncatetype.CohereTruncateType) *_cohereTaskSettings {

	s.v.Truncate = &truncate
	return s
}

func (s *_cohereTaskSettings) CohereTaskSettingsCaster() *types.CohereTaskSettings {
	return s.v
}
