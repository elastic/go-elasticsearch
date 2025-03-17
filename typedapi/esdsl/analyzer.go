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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// This is provide all the types that are part of the union.
type _analyzer struct {
	v types.Analyzer
}

func NewAnalyzer() *_analyzer {
	return &_analyzer{v: nil}
}

func (u *_analyzer) CustomAnalyzer(customanalyzer types.CustomAnalyzerVariant) *_analyzer {

	u.v = &customanalyzer

	return u
}

// Interface implementation for CustomAnalyzer in Analyzer union
func (u *_customAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) FingerprintAnalyzer(fingerprintanalyzer types.FingerprintAnalyzerVariant) *_analyzer {

	u.v = &fingerprintanalyzer

	return u
}

// Interface implementation for FingerprintAnalyzer in Analyzer union
func (u *_fingerprintAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) KeywordAnalyzer(keywordanalyzer types.KeywordAnalyzerVariant) *_analyzer {

	u.v = &keywordanalyzer

	return u
}

// Interface implementation for KeywordAnalyzer in Analyzer union
func (u *_keywordAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) NoriAnalyzer(norianalyzer types.NoriAnalyzerVariant) *_analyzer {

	u.v = &norianalyzer

	return u
}

// Interface implementation for NoriAnalyzer in Analyzer union
func (u *_noriAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) PatternAnalyzer(patternanalyzer types.PatternAnalyzerVariant) *_analyzer {

	u.v = &patternanalyzer

	return u
}

// Interface implementation for PatternAnalyzer in Analyzer union
func (u *_patternAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) SimpleAnalyzer(simpleanalyzer types.SimpleAnalyzerVariant) *_analyzer {

	u.v = &simpleanalyzer

	return u
}

// Interface implementation for SimpleAnalyzer in Analyzer union
func (u *_simpleAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) StandardAnalyzer(standardanalyzer types.StandardAnalyzerVariant) *_analyzer {

	u.v = &standardanalyzer

	return u
}

// Interface implementation for StandardAnalyzer in Analyzer union
func (u *_standardAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) StopAnalyzer(stopanalyzer types.StopAnalyzerVariant) *_analyzer {

	u.v = &stopanalyzer

	return u
}

// Interface implementation for StopAnalyzer in Analyzer union
func (u *_stopAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) WhitespaceAnalyzer(whitespaceanalyzer types.WhitespaceAnalyzerVariant) *_analyzer {

	u.v = &whitespaceanalyzer

	return u
}

// Interface implementation for WhitespaceAnalyzer in Analyzer union
func (u *_whitespaceAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) IcuAnalyzer(icuanalyzer types.IcuAnalyzerVariant) *_analyzer {

	u.v = &icuanalyzer

	return u
}

// Interface implementation for IcuAnalyzer in Analyzer union
func (u *_icuAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) KuromojiAnalyzer(kuromojianalyzer types.KuromojiAnalyzerVariant) *_analyzer {

	u.v = &kuromojianalyzer

	return u
}

// Interface implementation for KuromojiAnalyzer in Analyzer union
func (u *_kuromojiAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) SnowballAnalyzer(snowballanalyzer types.SnowballAnalyzerVariant) *_analyzer {

	u.v = &snowballanalyzer

	return u
}

// Interface implementation for SnowballAnalyzer in Analyzer union
func (u *_snowballAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) ArabicAnalyzer(arabicanalyzer types.ArabicAnalyzerVariant) *_analyzer {

	u.v = &arabicanalyzer

	return u
}

// Interface implementation for ArabicAnalyzer in Analyzer union
func (u *_arabicAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) ArmenianAnalyzer(armeniananalyzer types.ArmenianAnalyzerVariant) *_analyzer {

	u.v = &armeniananalyzer

	return u
}

// Interface implementation for ArmenianAnalyzer in Analyzer union
func (u *_armenianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) BasqueAnalyzer(basqueanalyzer types.BasqueAnalyzerVariant) *_analyzer {

	u.v = &basqueanalyzer

	return u
}

// Interface implementation for BasqueAnalyzer in Analyzer union
func (u *_basqueAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) BengaliAnalyzer(bengalianalyzer types.BengaliAnalyzerVariant) *_analyzer {

	u.v = &bengalianalyzer

	return u
}

// Interface implementation for BengaliAnalyzer in Analyzer union
func (u *_bengaliAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) BrazilianAnalyzer(braziliananalyzer types.BrazilianAnalyzerVariant) *_analyzer {

	u.v = &braziliananalyzer

	return u
}

// Interface implementation for BrazilianAnalyzer in Analyzer union
func (u *_brazilianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) BulgarianAnalyzer(bulgariananalyzer types.BulgarianAnalyzerVariant) *_analyzer {

	u.v = &bulgariananalyzer

	return u
}

// Interface implementation for BulgarianAnalyzer in Analyzer union
func (u *_bulgarianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) CatalanAnalyzer(catalananalyzer types.CatalanAnalyzerVariant) *_analyzer {

	u.v = &catalananalyzer

	return u
}

// Interface implementation for CatalanAnalyzer in Analyzer union
func (u *_catalanAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) ChineseAnalyzer(chineseanalyzer types.ChineseAnalyzerVariant) *_analyzer {

	u.v = &chineseanalyzer

	return u
}

// Interface implementation for ChineseAnalyzer in Analyzer union
func (u *_chineseAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) CjkAnalyzer(cjkanalyzer types.CjkAnalyzerVariant) *_analyzer {

	u.v = &cjkanalyzer

	return u
}

// Interface implementation for CjkAnalyzer in Analyzer union
func (u *_cjkAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) CzechAnalyzer(czechanalyzer types.CzechAnalyzerVariant) *_analyzer {

	u.v = &czechanalyzer

	return u
}

// Interface implementation for CzechAnalyzer in Analyzer union
func (u *_czechAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) DanishAnalyzer(danishanalyzer types.DanishAnalyzerVariant) *_analyzer {

	u.v = &danishanalyzer

	return u
}

// Interface implementation for DanishAnalyzer in Analyzer union
func (u *_danishAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) DutchAnalyzer(dutchanalyzer types.DutchAnalyzerVariant) *_analyzer {

	u.v = &dutchanalyzer

	return u
}

// Interface implementation for DutchAnalyzer in Analyzer union
func (u *_dutchAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) EnglishAnalyzer(englishanalyzer types.EnglishAnalyzerVariant) *_analyzer {

	u.v = &englishanalyzer

	return u
}

// Interface implementation for EnglishAnalyzer in Analyzer union
func (u *_englishAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) EstonianAnalyzer(estoniananalyzer types.EstonianAnalyzerVariant) *_analyzer {

	u.v = &estoniananalyzer

	return u
}

// Interface implementation for EstonianAnalyzer in Analyzer union
func (u *_estonianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) FinnishAnalyzer(finnishanalyzer types.FinnishAnalyzerVariant) *_analyzer {

	u.v = &finnishanalyzer

	return u
}

// Interface implementation for FinnishAnalyzer in Analyzer union
func (u *_finnishAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) FrenchAnalyzer(frenchanalyzer types.FrenchAnalyzerVariant) *_analyzer {

	u.v = &frenchanalyzer

	return u
}

// Interface implementation for FrenchAnalyzer in Analyzer union
func (u *_frenchAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) GalicianAnalyzer(galiciananalyzer types.GalicianAnalyzerVariant) *_analyzer {

	u.v = &galiciananalyzer

	return u
}

// Interface implementation for GalicianAnalyzer in Analyzer union
func (u *_galicianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) GermanAnalyzer(germananalyzer types.GermanAnalyzerVariant) *_analyzer {

	u.v = &germananalyzer

	return u
}

// Interface implementation for GermanAnalyzer in Analyzer union
func (u *_germanAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) GreekAnalyzer(greekanalyzer types.GreekAnalyzerVariant) *_analyzer {

	u.v = &greekanalyzer

	return u
}

// Interface implementation for GreekAnalyzer in Analyzer union
func (u *_greekAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) HindiAnalyzer(hindianalyzer types.HindiAnalyzerVariant) *_analyzer {

	u.v = &hindianalyzer

	return u
}

// Interface implementation for HindiAnalyzer in Analyzer union
func (u *_hindiAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) HungarianAnalyzer(hungariananalyzer types.HungarianAnalyzerVariant) *_analyzer {

	u.v = &hungariananalyzer

	return u
}

// Interface implementation for HungarianAnalyzer in Analyzer union
func (u *_hungarianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) IndonesianAnalyzer(indonesiananalyzer types.IndonesianAnalyzerVariant) *_analyzer {

	u.v = &indonesiananalyzer

	return u
}

// Interface implementation for IndonesianAnalyzer in Analyzer union
func (u *_indonesianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) IrishAnalyzer(irishanalyzer types.IrishAnalyzerVariant) *_analyzer {

	u.v = &irishanalyzer

	return u
}

// Interface implementation for IrishAnalyzer in Analyzer union
func (u *_irishAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) ItalianAnalyzer(italiananalyzer types.ItalianAnalyzerVariant) *_analyzer {

	u.v = &italiananalyzer

	return u
}

// Interface implementation for ItalianAnalyzer in Analyzer union
func (u *_italianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) LatvianAnalyzer(latviananalyzer types.LatvianAnalyzerVariant) *_analyzer {

	u.v = &latviananalyzer

	return u
}

// Interface implementation for LatvianAnalyzer in Analyzer union
func (u *_latvianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) LithuanianAnalyzer(lithuaniananalyzer types.LithuanianAnalyzerVariant) *_analyzer {

	u.v = &lithuaniananalyzer

	return u
}

// Interface implementation for LithuanianAnalyzer in Analyzer union
func (u *_lithuanianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) NorwegianAnalyzer(norwegiananalyzer types.NorwegianAnalyzerVariant) *_analyzer {

	u.v = &norwegiananalyzer

	return u
}

// Interface implementation for NorwegianAnalyzer in Analyzer union
func (u *_norwegianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) PersianAnalyzer(persiananalyzer types.PersianAnalyzerVariant) *_analyzer {

	u.v = &persiananalyzer

	return u
}

// Interface implementation for PersianAnalyzer in Analyzer union
func (u *_persianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) PortugueseAnalyzer(portugueseanalyzer types.PortugueseAnalyzerVariant) *_analyzer {

	u.v = &portugueseanalyzer

	return u
}

// Interface implementation for PortugueseAnalyzer in Analyzer union
func (u *_portugueseAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) RomanianAnalyzer(romaniananalyzer types.RomanianAnalyzerVariant) *_analyzer {

	u.v = &romaniananalyzer

	return u
}

// Interface implementation for RomanianAnalyzer in Analyzer union
func (u *_romanianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) RussianAnalyzer(russiananalyzer types.RussianAnalyzerVariant) *_analyzer {

	u.v = &russiananalyzer

	return u
}

// Interface implementation for RussianAnalyzer in Analyzer union
func (u *_russianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) SerbianAnalyzer(serbiananalyzer types.SerbianAnalyzerVariant) *_analyzer {

	u.v = &serbiananalyzer

	return u
}

// Interface implementation for SerbianAnalyzer in Analyzer union
func (u *_serbianAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) SoraniAnalyzer(soranianalyzer types.SoraniAnalyzerVariant) *_analyzer {

	u.v = &soranianalyzer

	return u
}

// Interface implementation for SoraniAnalyzer in Analyzer union
func (u *_soraniAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) SpanishAnalyzer(spanishanalyzer types.SpanishAnalyzerVariant) *_analyzer {

	u.v = &spanishanalyzer

	return u
}

// Interface implementation for SpanishAnalyzer in Analyzer union
func (u *_spanishAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) SwedishAnalyzer(swedishanalyzer types.SwedishAnalyzerVariant) *_analyzer {

	u.v = &swedishanalyzer

	return u
}

// Interface implementation for SwedishAnalyzer in Analyzer union
func (u *_swedishAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) TurkishAnalyzer(turkishanalyzer types.TurkishAnalyzerVariant) *_analyzer {

	u.v = &turkishanalyzer

	return u
}

// Interface implementation for TurkishAnalyzer in Analyzer union
func (u *_turkishAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) ThaiAnalyzer(thaianalyzer types.ThaiAnalyzerVariant) *_analyzer {

	u.v = &thaianalyzer

	return u
}

// Interface implementation for ThaiAnalyzer in Analyzer union
func (u *_thaiAnalyzer) AnalyzerCaster() *types.Analyzer {
	t := types.Analyzer(u.v)
	return &t
}

func (u *_analyzer) AnalyzerCaster() *types.Analyzer {
	return &u.v
}
