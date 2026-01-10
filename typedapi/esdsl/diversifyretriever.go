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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/diversifyretrievertypes"
)

type _diversifyRetriever struct {
	v *types.DiversifyRetriever
}

// A retriever that diversifies the results from its child retriever.
func NewDiversifyRetriever(field string, retriever types.RetrieverContainerVariant, type_ diversifyretrievertypes.DiversifyRetrieverTypes) *_diversifyRetriever {

	tmp := &_diversifyRetriever{v: types.NewDiversifyRetriever()}

	tmp.Field(field)

	tmp.Retriever(retriever)

	tmp.Type(type_)

	return tmp

}

func (s *_diversifyRetriever) Field(field string) *_diversifyRetriever {

	s.v.Field = field

	return s
}

func (s *_diversifyRetriever) Lambda(lambda float32) *_diversifyRetriever {

	s.v.Lambda = &lambda

	return s
}

func (s *_diversifyRetriever) QueryVector(queryvectors ...float32) *_diversifyRetriever {

	s.v.QueryVector = queryvectors

	return s
}

func (s *_diversifyRetriever) QueryVectorBuilder(queryvectorbuilder types.QueryVectorBuilderVariant) *_diversifyRetriever {

	s.v.QueryVectorBuilder = queryvectorbuilder.QueryVectorBuilderCaster()

	return s
}

func (s *_diversifyRetriever) RankWindowSize(rankwindowsize int) *_diversifyRetriever {

	s.v.RankWindowSize = &rankwindowsize

	return s
}

func (s *_diversifyRetriever) Retriever(retriever types.RetrieverContainerVariant) *_diversifyRetriever {

	s.v.Retriever = *retriever.RetrieverContainerCaster()

	return s
}

func (s *_diversifyRetriever) Size(size int) *_diversifyRetriever {

	s.v.Size = &size

	return s
}

func (s *_diversifyRetriever) Type(type_ diversifyretrievertypes.DiversifyRetrieverTypes) *_diversifyRetriever {

	s.v.Type = type_
	return s
}

func (s *_diversifyRetriever) Filter(filters ...types.QueryVariant) *_diversifyRetriever {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_diversifyRetriever) MinScore(minscore float32) *_diversifyRetriever {

	s.v.MinScore = &minscore

	return s
}

func (s *_diversifyRetriever) Name_(name_ string) *_diversifyRetriever {

	s.v.Name_ = &name_

	return s
}

func (s *_diversifyRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	container := types.NewRetrieverContainer()

	container.Diversify = s.v

	return container
}

func (s *_diversifyRetriever) DiversifyRetrieverCaster() *types.DiversifyRetriever {
	return s.v
}
