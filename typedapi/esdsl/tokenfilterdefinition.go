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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// This is provide all the types that are part of the union.
type _tokenFilterDefinition struct {
	v types.TokenFilterDefinition
}

func NewTokenFilterDefinition() *_tokenFilterDefinition {
	return &_tokenFilterDefinition{v: nil}
}

// UnknownTokenFilterDefinition is used to set the unknown value of the union.
// Highlited as @non_exhaustive in the specification.
func (u *_tokenFilterDefinition) UnknownTokenFilterDefinition(unknown json.RawMessage) *_tokenFilterDefinition {
	u.v = unknown
	return u
}

func (u *_tokenFilterDefinition) AsciiFoldingTokenFilter(asciifoldingtokenfilter types.AsciiFoldingTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &asciifoldingtokenfilter

	return u
}

// Interface implementation for AsciiFoldingTokenFilter in TokenFilterDefinition union
func (u *_asciiFoldingTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) CommonGramsTokenFilter(commongramstokenfilter types.CommonGramsTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &commongramstokenfilter

	return u
}

// Interface implementation for CommonGramsTokenFilter in TokenFilterDefinition union
func (u *_commonGramsTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) ConditionTokenFilter(conditiontokenfilter types.ConditionTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &conditiontokenfilter

	return u
}

// Interface implementation for ConditionTokenFilter in TokenFilterDefinition union
func (u *_conditionTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) DelimitedPayloadTokenFilter(delimitedpayloadtokenfilter types.DelimitedPayloadTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &delimitedpayloadtokenfilter

	return u
}

// Interface implementation for DelimitedPayloadTokenFilter in TokenFilterDefinition union
func (u *_delimitedPayloadTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) EdgeNGramTokenFilter(edgengramtokenfilter types.EdgeNGramTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &edgengramtokenfilter

	return u
}

// Interface implementation for EdgeNGramTokenFilter in TokenFilterDefinition union
func (u *_edgeNGramTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) ElisionTokenFilter(elisiontokenfilter types.ElisionTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &elisiontokenfilter

	return u
}

// Interface implementation for ElisionTokenFilter in TokenFilterDefinition union
func (u *_elisionTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) FingerprintTokenFilter(fingerprinttokenfilter types.FingerprintTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &fingerprinttokenfilter

	return u
}

// Interface implementation for FingerprintTokenFilter in TokenFilterDefinition union
func (u *_fingerprintTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) HunspellTokenFilter(hunspelltokenfilter types.HunspellTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &hunspelltokenfilter

	return u
}

// Interface implementation for HunspellTokenFilter in TokenFilterDefinition union
func (u *_hunspellTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) HyphenationDecompounderTokenFilter(hyphenationdecompoundertokenfilter types.HyphenationDecompounderTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &hyphenationdecompoundertokenfilter

	return u
}

// Interface implementation for HyphenationDecompounderTokenFilter in TokenFilterDefinition union
func (u *_hyphenationDecompounderTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) KeepTypesTokenFilter(keeptypestokenfilter types.KeepTypesTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &keeptypestokenfilter

	return u
}

// Interface implementation for KeepTypesTokenFilter in TokenFilterDefinition union
func (u *_keepTypesTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) KeepWordsTokenFilter(keepwordstokenfilter types.KeepWordsTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &keepwordstokenfilter

	return u
}

// Interface implementation for KeepWordsTokenFilter in TokenFilterDefinition union
func (u *_keepWordsTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) KeywordMarkerTokenFilter(keywordmarkertokenfilter types.KeywordMarkerTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &keywordmarkertokenfilter

	return u
}

// Interface implementation for KeywordMarkerTokenFilter in TokenFilterDefinition union
func (u *_keywordMarkerTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) KStemTokenFilter(kstemtokenfilter types.KStemTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &kstemtokenfilter

	return u
}

// Interface implementation for KStemTokenFilter in TokenFilterDefinition union
func (u *_kStemTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) LengthTokenFilter(lengthtokenfilter types.LengthTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &lengthtokenfilter

	return u
}

// Interface implementation for LengthTokenFilter in TokenFilterDefinition union
func (u *_lengthTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) LimitTokenCountTokenFilter(limittokencounttokenfilter types.LimitTokenCountTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &limittokencounttokenfilter

	return u
}

// Interface implementation for LimitTokenCountTokenFilter in TokenFilterDefinition union
func (u *_limitTokenCountTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) LowercaseTokenFilter(lowercasetokenfilter types.LowercaseTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &lowercasetokenfilter

	return u
}

// Interface implementation for LowercaseTokenFilter in TokenFilterDefinition union
func (u *_lowercaseTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) MultiplexerTokenFilter(multiplexertokenfilter types.MultiplexerTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &multiplexertokenfilter

	return u
}

// Interface implementation for MultiplexerTokenFilter in TokenFilterDefinition union
func (u *_multiplexerTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) NGramTokenFilter(ngramtokenfilter types.NGramTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &ngramtokenfilter

	return u
}

// Interface implementation for NGramTokenFilter in TokenFilterDefinition union
func (u *_nGramTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) NoriPartOfSpeechTokenFilter(noripartofspeechtokenfilter types.NoriPartOfSpeechTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &noripartofspeechtokenfilter

	return u
}

// Interface implementation for NoriPartOfSpeechTokenFilter in TokenFilterDefinition union
func (u *_noriPartOfSpeechTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) PatternCaptureTokenFilter(patterncapturetokenfilter types.PatternCaptureTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &patterncapturetokenfilter

	return u
}

// Interface implementation for PatternCaptureTokenFilter in TokenFilterDefinition union
func (u *_patternCaptureTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) PatternReplaceTokenFilter(patternreplacetokenfilter types.PatternReplaceTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &patternreplacetokenfilter

	return u
}

// Interface implementation for PatternReplaceTokenFilter in TokenFilterDefinition union
func (u *_patternReplaceTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) PorterStemTokenFilter(porterstemtokenfilter types.PorterStemTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &porterstemtokenfilter

	return u
}

// Interface implementation for PorterStemTokenFilter in TokenFilterDefinition union
func (u *_porterStemTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) PredicateTokenFilter(predicatetokenfilter types.PredicateTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &predicatetokenfilter

	return u
}

// Interface implementation for PredicateTokenFilter in TokenFilterDefinition union
func (u *_predicateTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) RemoveDuplicatesTokenFilter(removeduplicatestokenfilter types.RemoveDuplicatesTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &removeduplicatestokenfilter

	return u
}

// Interface implementation for RemoveDuplicatesTokenFilter in TokenFilterDefinition union
func (u *_removeDuplicatesTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) ReverseTokenFilter(reversetokenfilter types.ReverseTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &reversetokenfilter

	return u
}

// Interface implementation for ReverseTokenFilter in TokenFilterDefinition union
func (u *_reverseTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) ShingleTokenFilter(shingletokenfilter types.ShingleTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &shingletokenfilter

	return u
}

// Interface implementation for ShingleTokenFilter in TokenFilterDefinition union
func (u *_shingleTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) SnowballTokenFilter(snowballtokenfilter types.SnowballTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &snowballtokenfilter

	return u
}

// Interface implementation for SnowballTokenFilter in TokenFilterDefinition union
func (u *_snowballTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) StemmerOverrideTokenFilter(stemmeroverridetokenfilter types.StemmerOverrideTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &stemmeroverridetokenfilter

	return u
}

// Interface implementation for StemmerOverrideTokenFilter in TokenFilterDefinition union
func (u *_stemmerOverrideTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) StemmerTokenFilter(stemmertokenfilter types.StemmerTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &stemmertokenfilter

	return u
}

// Interface implementation for StemmerTokenFilter in TokenFilterDefinition union
func (u *_stemmerTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) StopTokenFilter(stoptokenfilter types.StopTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &stoptokenfilter

	return u
}

// Interface implementation for StopTokenFilter in TokenFilterDefinition union
func (u *_stopTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) SynonymGraphTokenFilter(synonymgraphtokenfilter types.SynonymGraphTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &synonymgraphtokenfilter

	return u
}

// Interface implementation for SynonymGraphTokenFilter in TokenFilterDefinition union
func (u *_synonymGraphTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) SynonymTokenFilter(synonymtokenfilter types.SynonymTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &synonymtokenfilter

	return u
}

// Interface implementation for SynonymTokenFilter in TokenFilterDefinition union
func (u *_synonymTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) TrimTokenFilter(trimtokenfilter types.TrimTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &trimtokenfilter

	return u
}

// Interface implementation for TrimTokenFilter in TokenFilterDefinition union
func (u *_trimTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) TruncateTokenFilter(truncatetokenfilter types.TruncateTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &truncatetokenfilter

	return u
}

// Interface implementation for TruncateTokenFilter in TokenFilterDefinition union
func (u *_truncateTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) UniqueTokenFilter(uniquetokenfilter types.UniqueTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &uniquetokenfilter

	return u
}

// Interface implementation for UniqueTokenFilter in TokenFilterDefinition union
func (u *_uniqueTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) UppercaseTokenFilter(uppercasetokenfilter types.UppercaseTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &uppercasetokenfilter

	return u
}

// Interface implementation for UppercaseTokenFilter in TokenFilterDefinition union
func (u *_uppercaseTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) WordDelimiterGraphTokenFilter(worddelimitergraphtokenfilter types.WordDelimiterGraphTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &worddelimitergraphtokenfilter

	return u
}

// Interface implementation for WordDelimiterGraphTokenFilter in TokenFilterDefinition union
func (u *_wordDelimiterGraphTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) WordDelimiterTokenFilter(worddelimitertokenfilter types.WordDelimiterTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &worddelimitertokenfilter

	return u
}

// Interface implementation for WordDelimiterTokenFilter in TokenFilterDefinition union
func (u *_wordDelimiterTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) KuromojiStemmerTokenFilter(kuromojistemmertokenfilter types.KuromojiStemmerTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &kuromojistemmertokenfilter

	return u
}

// Interface implementation for KuromojiStemmerTokenFilter in TokenFilterDefinition union
func (u *_kuromojiStemmerTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) KuromojiReadingFormTokenFilter(kuromojireadingformtokenfilter types.KuromojiReadingFormTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &kuromojireadingformtokenfilter

	return u
}

// Interface implementation for KuromojiReadingFormTokenFilter in TokenFilterDefinition union
func (u *_kuromojiReadingFormTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) KuromojiPartOfSpeechTokenFilter(kuromojipartofspeechtokenfilter types.KuromojiPartOfSpeechTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &kuromojipartofspeechtokenfilter

	return u
}

// Interface implementation for KuromojiPartOfSpeechTokenFilter in TokenFilterDefinition union
func (u *_kuromojiPartOfSpeechTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) IcuCollationTokenFilter(icucollationtokenfilter types.IcuCollationTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &icucollationtokenfilter

	return u
}

// Interface implementation for IcuCollationTokenFilter in TokenFilterDefinition union
func (u *_icuCollationTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) IcuFoldingTokenFilter(icufoldingtokenfilter types.IcuFoldingTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &icufoldingtokenfilter

	return u
}

// Interface implementation for IcuFoldingTokenFilter in TokenFilterDefinition union
func (u *_icuFoldingTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) IcuNormalizationTokenFilter(icunormalizationtokenfilter types.IcuNormalizationTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &icunormalizationtokenfilter

	return u
}

// Interface implementation for IcuNormalizationTokenFilter in TokenFilterDefinition union
func (u *_icuNormalizationTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) IcuTransformTokenFilter(icutransformtokenfilter types.IcuTransformTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &icutransformtokenfilter

	return u
}

// Interface implementation for IcuTransformTokenFilter in TokenFilterDefinition union
func (u *_icuTransformTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) PhoneticTokenFilter(phonetictokenfilter types.PhoneticTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &phonetictokenfilter

	return u
}

// Interface implementation for PhoneticTokenFilter in TokenFilterDefinition union
func (u *_phoneticTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) DictionaryDecompounderTokenFilter(dictionarydecompoundertokenfilter types.DictionaryDecompounderTokenFilterVariant) *_tokenFilterDefinition {

	u.v = &dictionarydecompoundertokenfilter

	return u
}

// Interface implementation for DictionaryDecompounderTokenFilter in TokenFilterDefinition union
func (u *_dictionaryDecompounderTokenFilter) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	t := types.TokenFilterDefinition(u.v)
	return &t
}

func (u *_tokenFilterDefinition) TokenFilterDefinitionCaster() *types.TokenFilterDefinition {
	return &u.v
}
