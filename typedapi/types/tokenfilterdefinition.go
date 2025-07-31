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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

// TokenFilterDefinition holds the union for the following types:
//
//	ApostropheTokenFilter
//	ArabicStemTokenFilter
//	ArabicNormalizationTokenFilter
//	AsciiFoldingTokenFilter
//	BengaliNormalizationTokenFilter
//	BrazilianStemTokenFilter
//	CjkBigramTokenFilter
//	CjkWidthTokenFilter
//	ClassicTokenFilter
//	CommonGramsTokenFilter
//	ConditionTokenFilter
//	CzechStemTokenFilter
//	DecimalDigitTokenFilter
//	DelimitedPayloadTokenFilter
//	DutchStemTokenFilter
//	EdgeNGramTokenFilter
//	ElisionTokenFilter
//	FingerprintTokenFilter
//	FlattenGraphTokenFilter
//	FrenchStemTokenFilter
//	GermanNormalizationTokenFilter
//	GermanStemTokenFilter
//	HindiNormalizationTokenFilter
//	HunspellTokenFilter
//	HyphenationDecompounderTokenFilter
//	IndicNormalizationTokenFilter
//	KeepTypesTokenFilter
//	KeepWordsTokenFilter
//	KeywordMarkerTokenFilter
//	KeywordRepeatTokenFilter
//	KStemTokenFilter
//	LengthTokenFilter
//	LimitTokenCountTokenFilter
//	LowercaseTokenFilter
//	MinHashTokenFilter
//	MultiplexerTokenFilter
//	NGramTokenFilter
//	NoriPartOfSpeechTokenFilter
//	PatternCaptureTokenFilter
//	PatternReplaceTokenFilter
//	PersianNormalizationTokenFilter
//	PersianStemTokenFilter
//	PorterStemTokenFilter
//	PredicateTokenFilter
//	RemoveDuplicatesTokenFilter
//	ReverseTokenFilter
//	RussianStemTokenFilter
//	ScandinavianFoldingTokenFilter
//	ScandinavianNormalizationTokenFilter
//	SerbianNormalizationTokenFilter
//	ShingleTokenFilter
//	SnowballTokenFilter
//	SoraniNormalizationTokenFilter
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
//	JaStopTokenFilter
//	KuromojiStemmerTokenFilter
//	KuromojiReadingFormTokenFilter
//	KuromojiPartOfSpeechTokenFilter
//	IcuCollationTokenFilter
//	IcuFoldingTokenFilter
//	IcuNormalizationTokenFilter
//	IcuTransformTokenFilter
//	PhoneticTokenFilter
//	DictionaryDecompounderTokenFilter
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/token_filters.ts#L582-L660
type TokenFilterDefinition any

type TokenFilterDefinitionVariant interface {
	TokenFilterDefinitionCaster() *TokenFilterDefinition
}
