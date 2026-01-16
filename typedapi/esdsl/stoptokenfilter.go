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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _stopTokenFilter struct {
	v *types.StopTokenFilter
}

func NewStopTokenFilter() *_stopTokenFilter {

	return &_stopTokenFilter{v: types.NewStopTokenFilter()}

}

func (s *_stopTokenFilter) IgnoreCase(ignorecase bool) *_stopTokenFilter {

	s.v.IgnoreCase = &ignorecase

	return s
}

func (s *_stopTokenFilter) RemoveTrailing(removetrailing bool) *_stopTokenFilter {

	s.v.RemoveTrailing = &removetrailing

	return s
}

func (s *_stopTokenFilter) Stopwords(stopwords types.StopWordsVariant) *_stopTokenFilter {

	s.v.Stopwords = *stopwords.StopWordsCaster()

	return s
}

func (s *_stopTokenFilter) StopwordsPath(stopwordspath string) *_stopTokenFilter {

	s.v.StopwordsPath = &stopwordspath

	return s
}

func (s *_stopTokenFilter) Version(versionstring string) *_stopTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_stopTokenFilter) StopTokenFilterCaster() *types.StopTokenFilter {
	return s.v
}
