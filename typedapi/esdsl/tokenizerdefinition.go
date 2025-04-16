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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// This is provide all the types that are part of the union.
type _tokenizerDefinition struct {
	v types.TokenizerDefinition
}

func NewTokenizerDefinition() *_tokenizerDefinition {
	return &_tokenizerDefinition{v: nil}
}

// UnknownTokenizerDefinition is used to set the unknown value of the union.
// Highlited as @non_exhaustive in the specification.
func (u *_tokenizerDefinition) UnknownTokenizerDefinition(unknown json.RawMessage) *_tokenizerDefinition {
	u.v = unknown
	return u
}

func (u *_tokenizerDefinition) CharGroupTokenizer(chargrouptokenizer types.CharGroupTokenizerVariant) *_tokenizerDefinition {

	u.v = &chargrouptokenizer

	return u
}

// Interface implementation for CharGroupTokenizer in TokenizerDefinition union
func (u *_charGroupTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) ClassicTokenizer(classictokenizer types.ClassicTokenizerVariant) *_tokenizerDefinition {

	u.v = &classictokenizer

	return u
}

// Interface implementation for ClassicTokenizer in TokenizerDefinition union
func (u *_classicTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) EdgeNGramTokenizer(edgengramtokenizer types.EdgeNGramTokenizerVariant) *_tokenizerDefinition {

	u.v = &edgengramtokenizer

	return u
}

// Interface implementation for EdgeNGramTokenizer in TokenizerDefinition union
func (u *_edgeNGramTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) KeywordTokenizer(keywordtokenizer types.KeywordTokenizerVariant) *_tokenizerDefinition {

	u.v = &keywordtokenizer

	return u
}

// Interface implementation for KeywordTokenizer in TokenizerDefinition union
func (u *_keywordTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) LetterTokenizer(lettertokenizer types.LetterTokenizerVariant) *_tokenizerDefinition {

	u.v = &lettertokenizer

	return u
}

// Interface implementation for LetterTokenizer in TokenizerDefinition union
func (u *_letterTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) LowercaseTokenizer(lowercasetokenizer types.LowercaseTokenizerVariant) *_tokenizerDefinition {

	u.v = &lowercasetokenizer

	return u
}

// Interface implementation for LowercaseTokenizer in TokenizerDefinition union
func (u *_lowercaseTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) NGramTokenizer(ngramtokenizer types.NGramTokenizerVariant) *_tokenizerDefinition {

	u.v = &ngramtokenizer

	return u
}

// Interface implementation for NGramTokenizer in TokenizerDefinition union
func (u *_nGramTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) PathHierarchyTokenizer(pathhierarchytokenizer types.PathHierarchyTokenizerVariant) *_tokenizerDefinition {

	u.v = &pathhierarchytokenizer

	return u
}

// Interface implementation for PathHierarchyTokenizer in TokenizerDefinition union
func (u *_pathHierarchyTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) PatternTokenizer(patterntokenizer types.PatternTokenizerVariant) *_tokenizerDefinition {

	u.v = &patterntokenizer

	return u
}

// Interface implementation for PatternTokenizer in TokenizerDefinition union
func (u *_patternTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) SimplePatternTokenizer(simplepatterntokenizer types.SimplePatternTokenizerVariant) *_tokenizerDefinition {

	u.v = &simplepatterntokenizer

	return u
}

// Interface implementation for SimplePatternTokenizer in TokenizerDefinition union
func (u *_simplePatternTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) SimplePatternSplitTokenizer(simplepatternsplittokenizer types.SimplePatternSplitTokenizerVariant) *_tokenizerDefinition {

	u.v = &simplepatternsplittokenizer

	return u
}

// Interface implementation for SimplePatternSplitTokenizer in TokenizerDefinition union
func (u *_simplePatternSplitTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) StandardTokenizer(standardtokenizer types.StandardTokenizerVariant) *_tokenizerDefinition {

	u.v = &standardtokenizer

	return u
}

// Interface implementation for StandardTokenizer in TokenizerDefinition union
func (u *_standardTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) ThaiTokenizer(thaitokenizer types.ThaiTokenizerVariant) *_tokenizerDefinition {

	u.v = &thaitokenizer

	return u
}

// Interface implementation for ThaiTokenizer in TokenizerDefinition union
func (u *_thaiTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) UaxEmailUrlTokenizer(uaxemailurltokenizer types.UaxEmailUrlTokenizerVariant) *_tokenizerDefinition {

	u.v = &uaxemailurltokenizer

	return u
}

// Interface implementation for UaxEmailUrlTokenizer in TokenizerDefinition union
func (u *_uaxEmailUrlTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) WhitespaceTokenizer(whitespacetokenizer types.WhitespaceTokenizerVariant) *_tokenizerDefinition {

	u.v = &whitespacetokenizer

	return u
}

// Interface implementation for WhitespaceTokenizer in TokenizerDefinition union
func (u *_whitespaceTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) IcuTokenizer(icutokenizer types.IcuTokenizerVariant) *_tokenizerDefinition {

	u.v = &icutokenizer

	return u
}

// Interface implementation for IcuTokenizer in TokenizerDefinition union
func (u *_icuTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) KuromojiTokenizer(kuromojitokenizer types.KuromojiTokenizerVariant) *_tokenizerDefinition {

	u.v = &kuromojitokenizer

	return u
}

// Interface implementation for KuromojiTokenizer in TokenizerDefinition union
func (u *_kuromojiTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) NoriTokenizer(noritokenizer types.NoriTokenizerVariant) *_tokenizerDefinition {

	u.v = &noritokenizer

	return u
}

// Interface implementation for NoriTokenizer in TokenizerDefinition union
func (u *_noriTokenizer) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	t := types.TokenizerDefinition(u.v)
	return &t
}

func (u *_tokenizerDefinition) TokenizerDefinitionCaster() *types.TokenizerDefinition {
	return &u.v
}
