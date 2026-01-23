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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textSimilarityReranker struct {
	v *types.TextSimilarityReranker
}

// A retriever that reranks the top documents based on a reranking model using
// the InferenceAPI
func NewTextSimilarityReranker(field string, inferencetext string, retriever types.RetrieverContainerVariant) *_textSimilarityReranker {

	tmp := &_textSimilarityReranker{v: types.NewTextSimilarityReranker()}

	tmp.Field(field)

	tmp.InferenceText(inferencetext)

	tmp.Retriever(retriever)

	return tmp

}

func (s *_textSimilarityReranker) ChunkRescorer(chunkrescorer types.ChunkRescorerVariant) *_textSimilarityReranker {

	s.v.ChunkRescorer = chunkrescorer.ChunkRescorerCaster()

	return s
}

func (s *_textSimilarityReranker) Field(field string) *_textSimilarityReranker {

	s.v.Field = field

	return s
}

func (s *_textSimilarityReranker) InferenceId(inferenceid string) *_textSimilarityReranker {

	s.v.InferenceId = &inferenceid

	return s
}

func (s *_textSimilarityReranker) InferenceText(inferencetext string) *_textSimilarityReranker {

	s.v.InferenceText = inferencetext

	return s
}

func (s *_textSimilarityReranker) RankWindowSize(rankwindowsize int) *_textSimilarityReranker {

	s.v.RankWindowSize = &rankwindowsize

	return s
}

func (s *_textSimilarityReranker) Retriever(retriever types.RetrieverContainerVariant) *_textSimilarityReranker {

	s.v.Retriever = *retriever.RetrieverContainerCaster()

	return s
}

func (s *_textSimilarityReranker) Filter(filters ...types.QueryVariant) *_textSimilarityReranker {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_textSimilarityReranker) MinScore(minscore float32) *_textSimilarityReranker {

	s.v.MinScore = &minscore

	return s
}

func (s *_textSimilarityReranker) Name_(name_ string) *_textSimilarityReranker {

	s.v.Name_ = &name_

	return s
}

func (s *_textSimilarityReranker) RetrieverContainerCaster() *types.RetrieverContainer {
	container := types.NewRetrieverContainer()

	container.TextSimilarityReranker = s.v

	return container
}

func (s *_textSimilarityReranker) TextSimilarityRerankerCaster() *types.TextSimilarityReranker {
	return s.v
}
