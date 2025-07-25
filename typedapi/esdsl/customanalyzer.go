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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _customAnalyzer struct {
	v *types.CustomAnalyzer
}

func NewCustomAnalyzer(tokenizer string) *_customAnalyzer {

	tmp := &_customAnalyzer{v: types.NewCustomAnalyzer()}

	tmp.Tokenizer(tokenizer)

	return tmp

}

func (s *_customAnalyzer) CharFilter(charfilters ...string) *_customAnalyzer {

	s.v.CharFilter = make([]string, len(charfilters))
	s.v.CharFilter = charfilters

	return s
}

func (s *_customAnalyzer) Filter(filters ...string) *_customAnalyzer {

	s.v.Filter = make([]string, len(filters))
	s.v.Filter = filters

	return s
}

func (s *_customAnalyzer) PositionIncrementGap(positionincrementgap int) *_customAnalyzer {

	s.v.PositionIncrementGap = &positionincrementgap

	return s
}

func (s *_customAnalyzer) PositionOffsetGap(positionoffsetgap int) *_customAnalyzer {

	s.v.PositionOffsetGap = &positionoffsetgap

	return s
}

func (s *_customAnalyzer) Tokenizer(tokenizer string) *_customAnalyzer {

	s.v.Tokenizer = tokenizer

	return s
}

func (s *_customAnalyzer) CustomAnalyzerCaster() *types.CustomAnalyzer {
	return s.v
}
