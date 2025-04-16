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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _categorizeTextAnalyzer struct {
	v types.CategorizeTextAnalyzer
}

func NewCategorizeTextAnalyzer() *_categorizeTextAnalyzer {
	return &_categorizeTextAnalyzer{v: nil}
}

func (u *_categorizeTextAnalyzer) String(string string) *_categorizeTextAnalyzer {

	u.v = &string

	return u
}

func (u *_categorizeTextAnalyzer) CustomCategorizeTextAnalyzer(customcategorizetextanalyzer types.CustomCategorizeTextAnalyzerVariant) *_categorizeTextAnalyzer {

	u.v = &customcategorizetextanalyzer

	return u
}

// Interface implementation for CustomCategorizeTextAnalyzer in CategorizeTextAnalyzer union
func (u *_customCategorizeTextAnalyzer) CategorizeTextAnalyzerCaster() *types.CategorizeTextAnalyzer {
	t := types.CategorizeTextAnalyzer(u.v)
	return &t
}

func (u *_categorizeTextAnalyzer) CategorizeTextAnalyzerCaster() *types.CategorizeTextAnalyzer {
	return &u.v
}
