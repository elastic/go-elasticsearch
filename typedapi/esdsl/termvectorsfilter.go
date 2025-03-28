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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _termVectorsFilter struct {
	v *types.TermVectorsFilter
}

func NewTermVectorsFilter() *_termVectorsFilter {

	return &_termVectorsFilter{v: types.NewTermVectorsFilter()}

}

// Ignore words which occur in more than this many docs.
// Defaults to unbounded.
func (s *_termVectorsFilter) MaxDocFreq(maxdocfreq int) *_termVectorsFilter {

	s.v.MaxDocFreq = &maxdocfreq

	return s
}

// The maximum number of terms that must be returned per field.
func (s *_termVectorsFilter) MaxNumTerms(maxnumterms int) *_termVectorsFilter {

	s.v.MaxNumTerms = &maxnumterms

	return s
}

// Ignore words with more than this frequency in the source doc.
// It defaults to unbounded.
func (s *_termVectorsFilter) MaxTermFreq(maxtermfreq int) *_termVectorsFilter {

	s.v.MaxTermFreq = &maxtermfreq

	return s
}

// The maximum word length above which words will be ignored.
// Defaults to unbounded.
func (s *_termVectorsFilter) MaxWordLength(maxwordlength int) *_termVectorsFilter {

	s.v.MaxWordLength = &maxwordlength

	return s
}

// Ignore terms which do not occur in at least this many docs.
func (s *_termVectorsFilter) MinDocFreq(mindocfreq int) *_termVectorsFilter {

	s.v.MinDocFreq = &mindocfreq

	return s
}

// Ignore words with less than this frequency in the source doc.
func (s *_termVectorsFilter) MinTermFreq(mintermfreq int) *_termVectorsFilter {

	s.v.MinTermFreq = &mintermfreq

	return s
}

// The minimum word length below which words will be ignored.
func (s *_termVectorsFilter) MinWordLength(minwordlength int) *_termVectorsFilter {

	s.v.MinWordLength = &minwordlength

	return s
}

func (s *_termVectorsFilter) TermVectorsFilterCaster() *types.TermVectorsFilter {
	return s.v
}
