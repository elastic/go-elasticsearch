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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rescorerRetriever struct {
	v *types.RescorerRetriever
}

// A retriever that re-scores only the results produced by its child retriever.
func NewRescorerRetriever(retriever types.RetrieverContainerVariant) *_rescorerRetriever {

	tmp := &_rescorerRetriever{v: types.NewRescorerRetriever()}

	tmp.Retriever(retriever)

	return tmp

}

func (s *_rescorerRetriever) Filter(filters ...types.QueryVariant) *_rescorerRetriever {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_rescorerRetriever) MinScore(minscore float32) *_rescorerRetriever {

	s.v.MinScore = &minscore

	return s
}

func (s *_rescorerRetriever) Name_(name_ string) *_rescorerRetriever {

	s.v.Name_ = &name_

	return s
}

func (s *_rescorerRetriever) Rescore(rescores ...types.RescoreVariant) *_rescorerRetriever {

	s.v.Rescore = make([]types.Rescore, len(rescores))
	for i, v := range rescores {
		s.v.Rescore[i] = *v.RescoreCaster()
	}

	return s
}

func (s *_rescorerRetriever) Retriever(retriever types.RetrieverContainerVariant) *_rescorerRetriever {

	s.v.Retriever = *retriever.RetrieverContainerCaster()

	return s
}

func (s *_rescorerRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	container := types.NewRetrieverContainer()

	container.Rescorer = s.v

	return container
}

func (s *_rescorerRetriever) RescorerRetrieverCaster() *types.RescorerRetriever {
	return s.v
}
