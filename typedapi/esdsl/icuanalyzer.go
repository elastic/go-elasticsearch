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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/icunormalizationtype"
)

type _icuAnalyzer struct {
	v *types.IcuAnalyzer
}

func NewIcuAnalyzer(method icunormalizationtype.IcuNormalizationType, mode icunormalizationmode.IcuNormalizationMode) *_icuAnalyzer {

	tmp := &_icuAnalyzer{v: types.NewIcuAnalyzer()}

	tmp.Method(method)

	tmp.Mode(mode)

	return tmp

}

func (s *_icuAnalyzer) Method(method icunormalizationtype.IcuNormalizationType) *_icuAnalyzer {

	s.v.Method = method
	return s
}

func (s *_icuAnalyzer) Mode(mode icunormalizationmode.IcuNormalizationMode) *_icuAnalyzer {

	s.v.Mode = mode
	return s
}

func (s *_icuAnalyzer) IcuAnalyzerCaster() *types.IcuAnalyzer {
	return s.v
}
