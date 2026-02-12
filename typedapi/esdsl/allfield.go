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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _allField struct {
	v *types.AllField
}

func NewAllField(analyzer string, enabled bool, omitnorms bool, searchanalyzer string, similarity string, store bool, storetermvectoroffsets bool, storetermvectorpayloads bool, storetermvectorpositions bool, storetermvectors bool) *_allField {

	tmp := &_allField{v: types.NewAllField()}

	tmp.Analyzer(analyzer)

	tmp.Enabled(enabled)

	tmp.OmitNorms(omitnorms)

	tmp.SearchAnalyzer(searchanalyzer)

	tmp.Similarity(similarity)

	tmp.Store(store)

	tmp.StoreTermVectorOffsets(storetermvectoroffsets)

	tmp.StoreTermVectorPayloads(storetermvectorpayloads)

	tmp.StoreTermVectorPositions(storetermvectorpositions)

	tmp.StoreTermVectors(storetermvectors)

	return tmp

}

func (s *_allField) Analyzer(analyzer string) *_allField {

	s.v.Analyzer = analyzer

	return s
}

func (s *_allField) Enabled(enabled bool) *_allField {

	s.v.Enabled = enabled

	return s
}

func (s *_allField) OmitNorms(omitnorms bool) *_allField {

	s.v.OmitNorms = omitnorms

	return s
}

func (s *_allField) SearchAnalyzer(searchanalyzer string) *_allField {

	s.v.SearchAnalyzer = searchanalyzer

	return s
}

func (s *_allField) Similarity(similarity string) *_allField {

	s.v.Similarity = similarity

	return s
}

func (s *_allField) Store(store bool) *_allField {

	s.v.Store = store

	return s
}

func (s *_allField) StoreTermVectorOffsets(storetermvectoroffsets bool) *_allField {

	s.v.StoreTermVectorOffsets = storetermvectoroffsets

	return s
}

func (s *_allField) StoreTermVectorPayloads(storetermvectorpayloads bool) *_allField {

	s.v.StoreTermVectorPayloads = storetermvectorpayloads

	return s
}

func (s *_allField) StoreTermVectorPositions(storetermvectorpositions bool) *_allField {

	s.v.StoreTermVectorPositions = storetermvectorpositions

	return s
}

func (s *_allField) StoreTermVectors(storetermvectors bool) *_allField {

	s.v.StoreTermVectors = storetermvectors

	return s
}

func (s *_allField) AllFieldCaster() *types.AllField {
	return s.v
}
