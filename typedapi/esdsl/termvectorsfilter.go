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

type _termVectorsFilter struct {
	v *types.TermVectorsFilter
}

func NewTermVectorsFilter() *_termVectorsFilter {

	return &_termVectorsFilter{v: types.NewTermVectorsFilter()}

}

func (s *_termVectorsFilter) MaxDocFreq(maxdocfreq int) *_termVectorsFilter {

	s.v.MaxDocFreq = &maxdocfreq

	return s
}

func (s *_termVectorsFilter) MaxNumTerms(maxnumterms int) *_termVectorsFilter {

	s.v.MaxNumTerms = &maxnumterms

	return s
}

func (s *_termVectorsFilter) MaxTermFreq(maxtermfreq int) *_termVectorsFilter {

	s.v.MaxTermFreq = &maxtermfreq

	return s
}

func (s *_termVectorsFilter) MaxWordLength(maxwordlength int) *_termVectorsFilter {

	s.v.MaxWordLength = &maxwordlength

	return s
}

func (s *_termVectorsFilter) MinDocFreq(mindocfreq int) *_termVectorsFilter {

	s.v.MinDocFreq = &mindocfreq

	return s
}

func (s *_termVectorsFilter) MinTermFreq(mintermfreq int) *_termVectorsFilter {

	s.v.MinTermFreq = &mintermfreq

	return s
}

func (s *_termVectorsFilter) MinWordLength(minwordlength int) *_termVectorsFilter {

	s.v.MinWordLength = &minwordlength

	return s
}

func (s *_termVectorsFilter) TermVectorsFilterCaster() *types.TermVectorsFilter {
	return s.v
}
