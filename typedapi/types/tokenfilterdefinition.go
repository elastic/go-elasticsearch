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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// TokenFilterDefinition holds the union for the following types:
//
//	AsciiFoldingTokenFilter
//	CommonGramsTokenFilter
//	ConditionTokenFilter
//	DelimitedPayloadTokenFilter
//	EdgeNGramTokenFilter
//	ElisionTokenFilter
//	FingerprintTokenFilter
//	HunspellTokenFilter
//	HyphenationDecompounderTokenFilter
//	KeepTypesTokenFilter
//	KeepWordsTokenFilter
//	KeywordMarkerTokenFilter
//	KStemTokenFilter
//	LengthTokenFilter
//	LimitTokenCountTokenFilter
//	LowercaseTokenFilter
//	MultiplexerTokenFilter
//	NGramTokenFilter
//	NoriPartOfSpeechTokenFilter
//	PatternCaptureTokenFilter
//	PatternReplaceTokenFilter
//	PorterStemTokenFilter
//	PredicateTokenFilter
//	RemoveDuplicatesTokenFilter
//	ReverseTokenFilter
//	ShingleTokenFilter
//	SnowballTokenFilter
//	StemmerOverrideTokenFilter
//	StemmerTokenFilter
//	StopTokenFilter
//	SynonymGraphTokenFilter
//	SynonymTokenFilter
//	TrimTokenFilter
//	TruncateTokenFilter
//	UniqueTokenFilter
//	UppercaseTokenFilter
//	WordDelimiterGraphTokenFilter
//	WordDelimiterTokenFilter
//	KuromojiStemmerTokenFilter
//	KuromojiReadingFormTokenFilter
//	KuromojiPartOfSpeechTokenFilter
//	IcuTokenizer
//	IcuCollationTokenFilter
//	IcuFoldingTokenFilter
//	IcuNormalizationTokenFilter
//	IcuTransformTokenFilter
//	PhoneticTokenFilter
//	DictionaryDecompounderTokenFilter
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/analysis/token_filters.ts#L346-L399
type TokenFilterDefinition interface {
	isTokenFilterDefinition()
}

func (i AsciiFoldingTokenFilter) isTokenFilterDefinition() {}

func (i CommonGramsTokenFilter) isTokenFilterDefinition() {}

func (i ConditionTokenFilter) isTokenFilterDefinition() {}

func (i DelimitedPayloadTokenFilter) isTokenFilterDefinition() {}

func (i EdgeNGramTokenFilter) isTokenFilterDefinition() {}

func (i ElisionTokenFilter) isTokenFilterDefinition() {}

func (i FingerprintTokenFilter) isTokenFilterDefinition() {}

func (i HunspellTokenFilter) isTokenFilterDefinition() {}

func (i HyphenationDecompounderTokenFilter) isTokenFilterDefinition() {}

func (i KeepTypesTokenFilter) isTokenFilterDefinition() {}

func (i KeepWordsTokenFilter) isTokenFilterDefinition() {}

func (i KeywordMarkerTokenFilter) isTokenFilterDefinition() {}

func (i KStemTokenFilter) isTokenFilterDefinition() {}

func (i LengthTokenFilter) isTokenFilterDefinition() {}

func (i LimitTokenCountTokenFilter) isTokenFilterDefinition() {}

func (i LowercaseTokenFilter) isTokenFilterDefinition() {}

func (i MultiplexerTokenFilter) isTokenFilterDefinition() {}

func (i NGramTokenFilter) isTokenFilterDefinition() {}

func (i NoriPartOfSpeechTokenFilter) isTokenFilterDefinition() {}

func (i PatternCaptureTokenFilter) isTokenFilterDefinition() {}

func (i PatternReplaceTokenFilter) isTokenFilterDefinition() {}

func (i PorterStemTokenFilter) isTokenFilterDefinition() {}

func (i PredicateTokenFilter) isTokenFilterDefinition() {}

func (i RemoveDuplicatesTokenFilter) isTokenFilterDefinition() {}

func (i ReverseTokenFilter) isTokenFilterDefinition() {}

func (i ShingleTokenFilter) isTokenFilterDefinition() {}

func (i SnowballTokenFilter) isTokenFilterDefinition() {}

func (i StemmerOverrideTokenFilter) isTokenFilterDefinition() {}

func (i StemmerTokenFilter) isTokenFilterDefinition() {}

func (i StopTokenFilter) isTokenFilterDefinition() {}

func (i SynonymGraphTokenFilter) isTokenFilterDefinition() {}

func (i SynonymTokenFilter) isTokenFilterDefinition() {}

func (i TrimTokenFilter) isTokenFilterDefinition() {}

func (i TruncateTokenFilter) isTokenFilterDefinition() {}

func (i UniqueTokenFilter) isTokenFilterDefinition() {}

func (i UppercaseTokenFilter) isTokenFilterDefinition() {}

func (i WordDelimiterGraphTokenFilter) isTokenFilterDefinition() {}

func (i WordDelimiterTokenFilter) isTokenFilterDefinition() {}

func (i KuromojiStemmerTokenFilter) isTokenFilterDefinition() {}

func (i KuromojiReadingFormTokenFilter) isTokenFilterDefinition() {}

func (i KuromojiPartOfSpeechTokenFilter) isTokenFilterDefinition() {}

func (i IcuTokenizer) isTokenFilterDefinition() {}

func (i IcuCollationTokenFilter) isTokenFilterDefinition() {}

func (i IcuFoldingTokenFilter) isTokenFilterDefinition() {}

func (i IcuNormalizationTokenFilter) isTokenFilterDefinition() {}

func (i IcuTransformTokenFilter) isTokenFilterDefinition() {}

func (i PhoneticTokenFilter) isTokenFilterDefinition() {}

func (i DictionaryDecompounderTokenFilter) isTokenFilterDefinition() {}
