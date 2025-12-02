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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _standardRetriever struct {
	v *types.StandardRetriever
}

// A retriever that replaces the functionality of a traditional query.
func NewStandardRetriever() *_standardRetriever {

	return &_standardRetriever{v: types.NewStandardRetriever()}

}

func (s *_standardRetriever) Collapse(collapse types.FieldCollapseVariant) *_standardRetriever {

	s.v.Collapse = collapse.FieldCollapseCaster()

	return s
}

func (s *_standardRetriever) Filter(filters ...types.QueryVariant) *_standardRetriever {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_standardRetriever) MinScore(minscore float32) *_standardRetriever {

	s.v.MinScore = &minscore

	return s
}

func (s *_standardRetriever) Name_(name_ string) *_standardRetriever {

	s.v.Name_ = &name_

	return s
}

func (s *_standardRetriever) Query(query types.QueryVariant) *_standardRetriever {

	s.v.Query = query.QueryCaster()

	return s
}

func (s *_standardRetriever) SearchAfter(sortresults ...types.FieldValueVariant) *_standardRetriever {

	for _, v := range sortresults {
		s.v.SearchAfter = append(s.v.SearchAfter, *v.FieldValueCaster())
	}

	return s
}

func (s *_standardRetriever) Sort(sorts ...types.SortCombinationsVariant) *_standardRetriever {

	for _, v := range sorts {
		s.v.Sort = append(s.v.Sort, *v.SortCombinationsCaster())
	}

	return s
}

func (s *_standardRetriever) TerminateAfter(terminateafter int) *_standardRetriever {

	s.v.TerminateAfter = &terminateafter

	return s
}

func (s *_standardRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	container := types.NewRetrieverContainer()

	container.Standard = s.v

	return container
}

func (s *_standardRetriever) StandardRetrieverCaster() *types.StandardRetriever {
	return s.v
}
