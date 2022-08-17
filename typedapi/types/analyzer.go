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

// Analyzer holds the union for the following types:
//
//	CustomAnalyzer
//	DutchAnalyzer
//	FingerprintAnalyzer
//	IcuAnalyzer
//	KeywordAnalyzer
//	KuromojiAnalyzer
//	LanguageAnalyzer
//	NoriAnalyzer
//	PatternAnalyzer
//	SimpleAnalyzer
//	SnowballAnalyzer
//	StandardAnalyzer
//	StopAnalyzer
//	WhitespaceAnalyzer
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/analyzers.ts#L113-L131
type Analyzer interface{}

// AnalyzerBuilder holds Analyzer struct and provides a builder API.
type AnalyzerBuilder struct {
	v Analyzer
}

// NewAnalyzer provides a builder for the Analyzer struct.
func NewAnalyzerBuilder() *AnalyzerBuilder {
	return &AnalyzerBuilder{}
}

// Build finalize the chain and returns the Analyzer struct
func (u *AnalyzerBuilder) Build() Analyzer {
	return u.v
}

func (u *AnalyzerBuilder) CustomAnalyzer(customanalyzer *CustomAnalyzerBuilder) *AnalyzerBuilder {
	v := customanalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) DutchAnalyzer(dutchanalyzer *DutchAnalyzerBuilder) *AnalyzerBuilder {
	v := dutchanalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) FingerprintAnalyzer(fingerprintanalyzer *FingerprintAnalyzerBuilder) *AnalyzerBuilder {
	v := fingerprintanalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) IcuAnalyzer(icuanalyzer *IcuAnalyzerBuilder) *AnalyzerBuilder {
	v := icuanalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) KeywordAnalyzer(keywordanalyzer *KeywordAnalyzerBuilder) *AnalyzerBuilder {
	v := keywordanalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) KuromojiAnalyzer(kuromojianalyzer *KuromojiAnalyzerBuilder) *AnalyzerBuilder {
	v := kuromojianalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) LanguageAnalyzer(languageanalyzer *LanguageAnalyzerBuilder) *AnalyzerBuilder {
	v := languageanalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) NoriAnalyzer(norianalyzer *NoriAnalyzerBuilder) *AnalyzerBuilder {
	v := norianalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) PatternAnalyzer(patternanalyzer *PatternAnalyzerBuilder) *AnalyzerBuilder {
	v := patternanalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) SimpleAnalyzer(simpleanalyzer *SimpleAnalyzerBuilder) *AnalyzerBuilder {
	v := simpleanalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) SnowballAnalyzer(snowballanalyzer *SnowballAnalyzerBuilder) *AnalyzerBuilder {
	v := snowballanalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) StandardAnalyzer(standardanalyzer *StandardAnalyzerBuilder) *AnalyzerBuilder {
	v := standardanalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) StopAnalyzer(stopanalyzer *StopAnalyzerBuilder) *AnalyzerBuilder {
	v := stopanalyzer.Build()
	u.v = &v
	return u
}

func (u *AnalyzerBuilder) WhitespaceAnalyzer(whitespaceanalyzer *WhitespaceAnalyzerBuilder) *AnalyzerBuilder {
	v := whitespaceanalyzer.Build()
	u.v = &v
	return u
}
