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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textSimilarityReranker struct {
	v *types.TextSimilarityReranker
}

// A retriever that reranks the top documents based on a reranking model using
// the InferenceAPI
func NewTextSimilarityReranker(retriever types.RetrieverContainerVariant) *_textSimilarityReranker {

	tmp := &_textSimilarityReranker{v: types.NewTextSimilarityReranker()}

	tmp.Retriever(retriever)

	return tmp

}

func (s *_textSimilarityReranker) Field(field string) *_textSimilarityReranker {

	s.v.Field = &field

	return s
}

func (s *_textSimilarityReranker) Filter(filters ...types.QueryVariant) *_textSimilarityReranker {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_textSimilarityReranker) InferenceId(inferenceid string) *_textSimilarityReranker {

	s.v.InferenceId = &inferenceid

	return s
}

func (s *_textSimilarityReranker) InferenceText(inferencetext string) *_textSimilarityReranker {

	s.v.InferenceText = &inferencetext

	return s
}

func (s *_textSimilarityReranker) MinScore(minscore float32) *_textSimilarityReranker {

	s.v.MinScore = &minscore

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

func (s *_textSimilarityReranker) RetrieverContainerCaster() *types.RetrieverContainer {
	container := types.NewRetrieverContainer()

	container.TextSimilarityReranker = s.v

	return container
}

func (s *_textSimilarityReranker) TextSimilarityRerankerCaster() *types.TextSimilarityReranker {
	return s.v
}
