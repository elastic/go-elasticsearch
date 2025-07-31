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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pinnedRetriever struct {
	v *types.PinnedRetriever
}

// A pinned retriever applies pinned documents to the underlying retriever.
// This retriever will rewrite to a PinnedQueryBuilder.
func NewPinnedRetriever(retriever types.RetrieverContainerVariant) *_pinnedRetriever {

	tmp := &_pinnedRetriever{v: types.NewPinnedRetriever()}

	tmp.Retriever(retriever)

	return tmp

}

func (s *_pinnedRetriever) Docs(docs ...types.SpecifiedDocumentVariant) *_pinnedRetriever {

	for _, v := range docs {

		s.v.Docs = append(s.v.Docs, *v.SpecifiedDocumentCaster())

	}
	return s
}

func (s *_pinnedRetriever) Filter(filters ...types.QueryVariant) *_pinnedRetriever {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_pinnedRetriever) Ids(ids ...string) *_pinnedRetriever {

	for _, v := range ids {

		s.v.Ids = append(s.v.Ids, v)

	}
	return s
}

func (s *_pinnedRetriever) MinScore(minscore float32) *_pinnedRetriever {

	s.v.MinScore = &minscore

	return s
}

func (s *_pinnedRetriever) Name_(name_ string) *_pinnedRetriever {

	s.v.Name_ = &name_

	return s
}

func (s *_pinnedRetriever) RankWindowSize(rankwindowsize int) *_pinnedRetriever {

	s.v.RankWindowSize = &rankwindowsize

	return s
}

func (s *_pinnedRetriever) Retriever(retriever types.RetrieverContainerVariant) *_pinnedRetriever {

	s.v.Retriever = *retriever.RetrieverContainerCaster()

	return s
}

func (s *_pinnedRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	container := types.NewRetrieverContainer()

	container.Pinned = s.v

	return container
}

func (s *_pinnedRetriever) PinnedRetrieverCaster() *types.PinnedRetriever {
	return s.v
}
