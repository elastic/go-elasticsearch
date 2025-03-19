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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/searchtype"
)

type _multisearchHeader struct {
	v *types.MultisearchHeader
}

func NewMultisearchHeader() *_multisearchHeader {

	return &_multisearchHeader{v: types.NewMultisearchHeader()}

}

func (s *_multisearchHeader) AllowNoIndices(allownoindices bool) *_multisearchHeader {

	s.v.AllowNoIndices = &allownoindices

	return s
}

func (s *_multisearchHeader) AllowPartialSearchResults(allowpartialsearchresults bool) *_multisearchHeader {

	s.v.AllowPartialSearchResults = &allowpartialsearchresults

	return s
}

func (s *_multisearchHeader) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *_multisearchHeader {

	s.v.CcsMinimizeRoundtrips = &ccsminimizeroundtrips

	return s
}

func (s *_multisearchHeader) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *_multisearchHeader {

	s.v.ExpandWildcards = expandwildcards

	return s
}

func (s *_multisearchHeader) IgnoreThrottled(ignorethrottled bool) *_multisearchHeader {

	s.v.IgnoreThrottled = &ignorethrottled

	return s
}

func (s *_multisearchHeader) IgnoreUnavailable(ignoreunavailable bool) *_multisearchHeader {

	s.v.IgnoreUnavailable = &ignoreunavailable

	return s
}

func (s *_multisearchHeader) Index(indices ...string) *_multisearchHeader {

	s.v.Index = indices

	return s
}

func (s *_multisearchHeader) Preference(preference string) *_multisearchHeader {

	s.v.Preference = &preference

	return s
}

func (s *_multisearchHeader) RequestCache(requestcache bool) *_multisearchHeader {

	s.v.RequestCache = &requestcache

	return s
}

func (s *_multisearchHeader) Routing(routing string) *_multisearchHeader {

	s.v.Routing = &routing

	return s
}

func (s *_multisearchHeader) SearchType(searchtype searchtype.SearchType) *_multisearchHeader {

	s.v.SearchType = &searchtype
	return s
}

func (s *_multisearchHeader) MultisearchHeaderCaster() *types.MultisearchHeader {
	return s.v
}
