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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// TokenizerDefinition holds the union for the following types:
//
//	CharGroupTokenizer
//	EdgeNGramTokenizer
//	IcuTokenizer
//	KeywordTokenizer
//	KuromojiTokenizer
//	LetterTokenizer
//	LowercaseTokenizer
//	NGramTokenizer
//	NoriTokenizer
//	PathHierarchyTokenizer
//	PatternTokenizer
//	StandardTokenizer
//	UaxEmailUrlTokenizer
//	WhitespaceTokenizer
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/tokenizers.ts#L123-L141
type TokenizerDefinition interface{}

// TokenizerDefinitionBuilder holds TokenizerDefinition struct and provides a builder API.
type TokenizerDefinitionBuilder struct {
	v TokenizerDefinition
}

// NewTokenizerDefinition provides a builder for the TokenizerDefinition struct.
func NewTokenizerDefinitionBuilder() *TokenizerDefinitionBuilder {
	return &TokenizerDefinitionBuilder{}
}

// Build finalize the chain and returns the TokenizerDefinition struct
func (u *TokenizerDefinitionBuilder) Build() TokenizerDefinition {
	return u.v
}

func (u *TokenizerDefinitionBuilder) CharGroupTokenizer(chargrouptokenizer *CharGroupTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := chargrouptokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) EdgeNGramTokenizer(edgengramtokenizer *EdgeNGramTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := edgengramtokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) IcuTokenizer(icutokenizer *IcuTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := icutokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) KeywordTokenizer(keywordtokenizer *KeywordTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := keywordtokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) KuromojiTokenizer(kuromojitokenizer *KuromojiTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := kuromojitokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) LetterTokenizer(lettertokenizer *LetterTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := lettertokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) LowercaseTokenizer(lowercasetokenizer *LowercaseTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := lowercasetokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) NGramTokenizer(ngramtokenizer *NGramTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := ngramtokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) NoriTokenizer(noritokenizer *NoriTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := noritokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) PathHierarchyTokenizer(pathhierarchytokenizer *PathHierarchyTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := pathhierarchytokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) PatternTokenizer(patterntokenizer *PatternTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := patterntokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) StandardTokenizer(standardtokenizer *StandardTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := standardtokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) UaxEmailUrlTokenizer(uaxemailurltokenizer *UaxEmailUrlTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := uaxemailurltokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenizerDefinitionBuilder) WhitespaceTokenizer(whitespacetokenizer *WhitespaceTokenizerBuilder) *TokenizerDefinitionBuilder {
	v := whitespacetokenizer.Build()
	u.v = &v
	return u
}
