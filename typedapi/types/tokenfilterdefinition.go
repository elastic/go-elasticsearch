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

// TokenFilterDefinition holds the union for the following types:
//
//	AsciiFoldingTokenFilter
//	CommonGramsTokenFilter
//	ConditionTokenFilter
//	DelimitedPayloadTokenFilter
//	DictionaryDecompounderTokenFilter
//	EdgeNGramTokenFilter
//	ElisionTokenFilter
//	FingerprintTokenFilter
//	HunspellTokenFilter
//	HyphenationDecompounderTokenFilter
//	IcuCollationTokenFilter
//	IcuFoldingTokenFilter
//	IcuNormalizationTokenFilter
//	IcuTokenizer
//	IcuTransformTokenFilter
//	KStemTokenFilter
//	KeepTypesTokenFilter
//	KeepWordsTokenFilter
//	KeywordMarkerTokenFilter
//	KuromojiPartOfSpeechTokenFilter
//	KuromojiReadingFormTokenFilter
//	KuromojiStemmerTokenFilter
//	LengthTokenFilter
//	LimitTokenCountTokenFilter
//	LowercaseTokenFilter
//	MultiplexerTokenFilter
//	NGramTokenFilter
//	NoriPartOfSpeechTokenFilter
//	PatternCaptureTokenFilter
//	PatternReplaceTokenFilter
//	PhoneticTokenFilter
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
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L346-L399
type TokenFilterDefinition interface{}

// TokenFilterDefinitionBuilder holds TokenFilterDefinition struct and provides a builder API.
type TokenFilterDefinitionBuilder struct {
	v TokenFilterDefinition
}

// NewTokenFilterDefinition provides a builder for the TokenFilterDefinition struct.
func NewTokenFilterDefinitionBuilder() *TokenFilterDefinitionBuilder {
	return &TokenFilterDefinitionBuilder{}
}

// Build finalize the chain and returns the TokenFilterDefinition struct
func (u *TokenFilterDefinitionBuilder) Build() TokenFilterDefinition {
	return u.v
}

func (u *TokenFilterDefinitionBuilder) AsciiFoldingTokenFilter(asciifoldingtokenfilter *AsciiFoldingTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := asciifoldingtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) CommonGramsTokenFilter(commongramstokenfilter *CommonGramsTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := commongramstokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) ConditionTokenFilter(conditiontokenfilter *ConditionTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := conditiontokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) DelimitedPayloadTokenFilter(delimitedpayloadtokenfilter *DelimitedPayloadTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := delimitedpayloadtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) DictionaryDecompounderTokenFilter(dictionarydecompoundertokenfilter *DictionaryDecompounderTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := dictionarydecompoundertokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) EdgeNGramTokenFilter(edgengramtokenfilter *EdgeNGramTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := edgengramtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) ElisionTokenFilter(elisiontokenfilter *ElisionTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := elisiontokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) FingerprintTokenFilter(fingerprinttokenfilter *FingerprintTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := fingerprinttokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) HunspellTokenFilter(hunspelltokenfilter *HunspellTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := hunspelltokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) HyphenationDecompounderTokenFilter(hyphenationdecompoundertokenfilter *HyphenationDecompounderTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := hyphenationdecompoundertokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) IcuCollationTokenFilter(icucollationtokenfilter *IcuCollationTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := icucollationtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) IcuFoldingTokenFilter(icufoldingtokenfilter *IcuFoldingTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := icufoldingtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) IcuNormalizationTokenFilter(icunormalizationtokenfilter *IcuNormalizationTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := icunormalizationtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) IcuTokenizer(icutokenizer *IcuTokenizerBuilder) *TokenFilterDefinitionBuilder {
	v := icutokenizer.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) IcuTransformTokenFilter(icutransformtokenfilter *IcuTransformTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := icutransformtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) KStemTokenFilter(kstemtokenfilter *KStemTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := kstemtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) KeepTypesTokenFilter(keeptypestokenfilter *KeepTypesTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := keeptypestokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) KeepWordsTokenFilter(keepwordstokenfilter *KeepWordsTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := keepwordstokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) KeywordMarkerTokenFilter(keywordmarkertokenfilter *KeywordMarkerTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := keywordmarkertokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) KuromojiPartOfSpeechTokenFilter(kuromojipartofspeechtokenfilter *KuromojiPartOfSpeechTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := kuromojipartofspeechtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) KuromojiReadingFormTokenFilter(kuromojireadingformtokenfilter *KuromojiReadingFormTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := kuromojireadingformtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) KuromojiStemmerTokenFilter(kuromojistemmertokenfilter *KuromojiStemmerTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := kuromojistemmertokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) LengthTokenFilter(lengthtokenfilter *LengthTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := lengthtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) LimitTokenCountTokenFilter(limittokencounttokenfilter *LimitTokenCountTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := limittokencounttokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) LowercaseTokenFilter(lowercasetokenfilter *LowercaseTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := lowercasetokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) MultiplexerTokenFilter(multiplexertokenfilter *MultiplexerTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := multiplexertokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) NGramTokenFilter(ngramtokenfilter *NGramTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := ngramtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) NoriPartOfSpeechTokenFilter(noripartofspeechtokenfilter *NoriPartOfSpeechTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := noripartofspeechtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) PatternCaptureTokenFilter(patterncapturetokenfilter *PatternCaptureTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := patterncapturetokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) PatternReplaceTokenFilter(patternreplacetokenfilter *PatternReplaceTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := patternreplacetokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) PhoneticTokenFilter(phonetictokenfilter *PhoneticTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := phonetictokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) PorterStemTokenFilter(porterstemtokenfilter *PorterStemTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := porterstemtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) PredicateTokenFilter(predicatetokenfilter *PredicateTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := predicatetokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) RemoveDuplicatesTokenFilter(removeduplicatestokenfilter *RemoveDuplicatesTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := removeduplicatestokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) ReverseTokenFilter(reversetokenfilter *ReverseTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := reversetokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) ShingleTokenFilter(shingletokenfilter *ShingleTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := shingletokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) SnowballTokenFilter(snowballtokenfilter *SnowballTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := snowballtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) StemmerOverrideTokenFilter(stemmeroverridetokenfilter *StemmerOverrideTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := stemmeroverridetokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) StemmerTokenFilter(stemmertokenfilter *StemmerTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := stemmertokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) StopTokenFilter(stoptokenfilter *StopTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := stoptokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) SynonymGraphTokenFilter(synonymgraphtokenfilter *SynonymGraphTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := synonymgraphtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) SynonymTokenFilter(synonymtokenfilter *SynonymTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := synonymtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) TrimTokenFilter(trimtokenfilter *TrimTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := trimtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) TruncateTokenFilter(truncatetokenfilter *TruncateTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := truncatetokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) UniqueTokenFilter(uniquetokenfilter *UniqueTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := uniquetokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) UppercaseTokenFilter(uppercasetokenfilter *UppercaseTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := uppercasetokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) WordDelimiterGraphTokenFilter(worddelimitergraphtokenfilter *WordDelimiterGraphTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := worddelimitergraphtokenfilter.Build()
	u.v = &v
	return u
}

func (u *TokenFilterDefinitionBuilder) WordDelimiterTokenFilter(worddelimitertokenfilter *WordDelimiterTokenFilterBuilder) *TokenFilterDefinitionBuilder {
	v := worddelimitertokenfilter.Build()
	u.v = &v
	return u
}
