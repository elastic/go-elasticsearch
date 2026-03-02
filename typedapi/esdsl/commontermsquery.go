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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
)

type _commonTermsQuery struct {
	k string
	v *types.CommonTermsQuery
}

func NewCommonTermsQuery(field string, query string) *_commonTermsQuery {
	tmp := &_commonTermsQuery{
		k: field,
		v: types.NewCommonTermsQuery(),
	}

	tmp.Query(query)
	return tmp
}

func (s *_commonTermsQuery) Analyzer(analyzer string) *_commonTermsQuery {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_commonTermsQuery) CutoffFrequency(cutofffrequency types.Float64) *_commonTermsQuery {

	s.v.CutoffFrequency = &cutofffrequency

	return s
}

func (s *_commonTermsQuery) HighFreqOperator(highfreqoperator operator.Operator) *_commonTermsQuery {

	s.v.HighFreqOperator = &highfreqoperator
	return s
}

func (s *_commonTermsQuery) LowFreqOperator(lowfreqoperator operator.Operator) *_commonTermsQuery {

	s.v.LowFreqOperator = &lowfreqoperator
	return s
}

func (s *_commonTermsQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_commonTermsQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

func (s *_commonTermsQuery) Query(query string) *_commonTermsQuery {

	s.v.Query = query

	return s
}

func (s *_commonTermsQuery) Boost(boost float32) *_commonTermsQuery {

	s.v.Boost = &boost

	return s
}

func (s *_commonTermsQuery) QueryName_(queryname_ string) *_commonTermsQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_commonTermsQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Common = map[string]types.CommonTermsQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleCommonTermsQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleCommonTermsQuery() *_commonTermsQuery {
	return &_commonTermsQuery{
		k: "",
		v: types.NewCommonTermsQuery(),
	}
}

func (s *_commonTermsQuery) CommonTermsQueryCaster() *types.CommonTermsQuery {
	return s.v.CommonTermsQueryCaster()
}
