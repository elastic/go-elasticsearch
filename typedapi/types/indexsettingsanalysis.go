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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// IndexSettingsAnalysis type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/_types/IndexSettings.ts#L333-L339
type IndexSettingsAnalysis struct {
	Analyzer   map[string]Analyzer    `json:"analyzer,omitempty"`
	CharFilter map[string]CharFilter  `json:"char_filter,omitempty"`
	Filter     map[string]TokenFilter `json:"filter,omitempty"`
	Normalizer map[string]Normalizer  `json:"normalizer,omitempty"`
	Tokenizer  map[string]Tokenizer   `json:"tokenizer,omitempty"`
}

func (s *IndexSettingsAnalysis) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "analyzer":
			if s.Analyzer == nil {
				s.Analyzer = make(map[string]Analyzer, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]any)
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)
				if _, ok := kind["type"]; !ok {
					kind["type"] = "custom"
				}
				switch kind["type"] {
				case "custom":
					oo := NewCustomAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "fingerprint":
					oo := NewFingerprintAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "keyword":
					oo := NewKeywordAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "nori":
					oo := NewNoriAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "pattern":
					oo := NewPatternAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "simple":
					oo := NewSimpleAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "standard":
					oo := NewStandardAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "stop":
					oo := NewStopAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "whitespace":
					oo := NewWhitespaceAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "icu_analyzer":
					oo := NewIcuAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "kuromoji":
					oo := NewKuromojiAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "snowball":
					oo := NewSnowballAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "arabic":
					oo := NewArabicAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "armenian":
					oo := NewArmenianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "basque":
					oo := NewBasqueAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "bengali":
					oo := NewBengaliAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "brazilian":
					oo := NewBrazilianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "bulgarian":
					oo := NewBulgarianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "catalan":
					oo := NewCatalanAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "chinese":
					oo := NewChineseAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "cjk":
					oo := NewCjkAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "czech":
					oo := NewCzechAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "danish":
					oo := NewDanishAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "dutch":
					oo := NewDutchAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "english":
					oo := NewEnglishAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "estonian":
					oo := NewEstonianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "finnish":
					oo := NewFinnishAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "french":
					oo := NewFrenchAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "galician":
					oo := NewGalicianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "german":
					oo := NewGermanAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "greek":
					oo := NewGreekAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "hindi":
					oo := NewHindiAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "hungarian":
					oo := NewHungarianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "indonesian":
					oo := NewIndonesianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "irish":
					oo := NewIrishAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "italian":
					oo := NewItalianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "latvian":
					oo := NewLatvianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "lithuanian":
					oo := NewLithuanianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "norwegian":
					oo := NewNorwegianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "persian":
					oo := NewPersianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "portuguese":
					oo := NewPortugueseAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "romanian":
					oo := NewRomanianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "russian":
					oo := NewRussianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "serbian":
					oo := NewSerbianAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "sorani":
					oo := NewSoraniAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "spanish":
					oo := NewSpanishAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "swedish":
					oo := NewSwedishAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "turkish":
					oo := NewTurkishAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				case "thai":
					oo := NewThaiAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Analyzer | %w", err)
					}
					s.Analyzer[key] = oo
				default:
					oo := new(Analyzer)
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("new(Analyzer) | %w", err)
					}
					s.Analyzer[key] = oo
				}
			}

		case "char_filter":
			if s.CharFilter == nil {
				s.CharFilter = make(map[string]CharFilter, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]any)
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["type"] {
				case "html_strip":
					oo := NewHtmlStripCharFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("CharFilter | %w", err)
					}
					s.CharFilter[key] = oo
				case "mapping":
					oo := NewMappingCharFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("CharFilter | %w", err)
					}
					s.CharFilter[key] = oo
				case "pattern_replace":
					oo := NewPatternReplaceCharFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("CharFilter | %w", err)
					}
					s.CharFilter[key] = oo
				case "icu_normalizer":
					oo := NewIcuNormalizationCharFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("CharFilter | %w", err)
					}
					s.CharFilter[key] = oo
				case "kuromoji_iteration_mark":
					oo := NewKuromojiIterationMarkCharFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("CharFilter | %w", err)
					}
					s.CharFilter[key] = oo
				default:
					oo := new(CharFilter)
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("new(CharFilter) | %w", err)
					}
					s.CharFilter[key] = oo
				}
			}

		case "filter":
			if s.Filter == nil {
				s.Filter = make(map[string]TokenFilter, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]any)
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["type"] {
				case "apostrophe":
					oo := NewApostropheTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "arabic_stem":
					oo := NewArabicStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "arabic_normalization":
					oo := NewArabicNormalizationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "asciifolding":
					oo := NewAsciiFoldingTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "bengali_normalization":
					oo := NewBengaliNormalizationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "brazilian_stem":
					oo := NewBrazilianStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "cjk_bigram":
					oo := NewCjkBigramTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "cjk_width":
					oo := NewCjkWidthTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "classic":
					oo := NewClassicTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "common_grams":
					oo := NewCommonGramsTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "condition":
					oo := NewConditionTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "czech_stem":
					oo := NewCzechStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "decimal_digit":
					oo := NewDecimalDigitTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "delimited_payload":
					oo := NewDelimitedPayloadTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "dutch_stem":
					oo := NewDutchStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "edge_ngram":
					oo := NewEdgeNGramTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "elision":
					oo := NewElisionTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "fingerprint":
					oo := NewFingerprintTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "flatten_graph":
					oo := NewFlattenGraphTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "french_stem":
					oo := NewFrenchStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "german_normalization":
					oo := NewGermanNormalizationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "german_stem":
					oo := NewGermanStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "hindi_normalization":
					oo := NewHindiNormalizationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "hunspell":
					oo := NewHunspellTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "hyphenation_decompounder":
					oo := NewHyphenationDecompounderTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "indic_normalization":
					oo := NewIndicNormalizationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "keep_types":
					oo := NewKeepTypesTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "keep":
					oo := NewKeepWordsTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "keyword_marker":
					oo := NewKeywordMarkerTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "keyword_repeat":
					oo := NewKeywordRepeatTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "kstem":
					oo := NewKStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "length":
					oo := NewLengthTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "limit":
					oo := NewLimitTokenCountTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "lowercase":
					oo := NewLowercaseTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "min_hash":
					oo := NewMinHashTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "multiplexer":
					oo := NewMultiplexerTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "ngram":
					oo := NewNGramTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "nori_part_of_speech":
					oo := NewNoriPartOfSpeechTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "pattern_capture":
					oo := NewPatternCaptureTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "pattern_replace":
					oo := NewPatternReplaceTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "persian_normalization":
					oo := NewPersianNormalizationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "persian_stem":
					oo := NewPersianStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "porter_stem":
					oo := NewPorterStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "predicate_token_filter":
					oo := NewPredicateTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "remove_duplicates":
					oo := NewRemoveDuplicatesTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "reverse":
					oo := NewReverseTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "russian_stem":
					oo := NewRussianStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "scandinavian_folding":
					oo := NewScandinavianFoldingTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "scandinavian_normalization":
					oo := NewScandinavianNormalizationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "serbian_normalization":
					oo := NewSerbianNormalizationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "shingle":
					oo := NewShingleTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "snowball":
					oo := NewSnowballTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "sorani_normalization":
					oo := NewSoraniNormalizationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "stemmer_override":
					oo := NewStemmerOverrideTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "stemmer":
					oo := NewStemmerTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "stop":
					oo := NewStopTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "synonym_graph":
					oo := NewSynonymGraphTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "synonym":
					oo := NewSynonymTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "trim":
					oo := NewTrimTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "truncate":
					oo := NewTruncateTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "unique":
					oo := NewUniqueTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "uppercase":
					oo := NewUppercaseTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "word_delimiter_graph":
					oo := NewWordDelimiterGraphTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "word_delimiter":
					oo := NewWordDelimiterTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "ja_stop":
					oo := NewJaStopTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "kuromoji_stemmer":
					oo := NewKuromojiStemmerTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "kuromoji_readingform":
					oo := NewKuromojiReadingFormTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "kuromoji_part_of_speech":
					oo := NewKuromojiPartOfSpeechTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "icu_collation":
					oo := NewIcuCollationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "icu_folding":
					oo := NewIcuFoldingTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "icu_normalizer":
					oo := NewIcuNormalizationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "icu_transform":
					oo := NewIcuTransformTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "phonetic":
					oo := NewPhoneticTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				case "dictionary_decompounder":
					oo := NewDictionaryDecompounderTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter[key] = oo
				default:
					oo := new(TokenFilter)
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("new(TokenFilter) | %w", err)
					}
					s.Filter[key] = oo
				}
			}

		case "normalizer":
			if s.Normalizer == nil {
				s.Normalizer = make(map[string]Normalizer, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]any)
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)
				if _, ok := kind["type"]; !ok {
					kind["type"] = "custom"
				}
				switch kind["type"] {
				case "lowercase":
					oo := NewLowercaseNormalizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Normalizer | %w", err)
					}
					s.Normalizer[key] = oo
				case "custom":
					oo := NewCustomNormalizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Normalizer | %w", err)
					}
					s.Normalizer[key] = oo
				default:
					oo := new(Normalizer)
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("new(Normalizer) | %w", err)
					}
					s.Normalizer[key] = oo
				}
			}

		case "tokenizer":
			if s.Tokenizer == nil {
				s.Tokenizer = make(map[string]Tokenizer, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]any)
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["type"] {
				case "char_group":
					oo := NewCharGroupTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "classic":
					oo := NewClassicTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "edge_ngram":
					oo := NewEdgeNGramTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "keyword":
					oo := NewKeywordTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "letter":
					oo := NewLetterTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "lowercase":
					oo := NewLowercaseTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "ngram":
					oo := NewNGramTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "path_hierarchy":
					oo := NewPathHierarchyTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "pattern":
					oo := NewPatternTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "simple_pattern":
					oo := NewSimplePatternTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "simple_pattern_split":
					oo := NewSimplePatternSplitTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "standard":
					oo := NewStandardTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "thai":
					oo := NewThaiTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "uax_url_email":
					oo := NewUaxEmailUrlTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "whitespace":
					oo := NewWhitespaceTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "icu_tokenizer":
					oo := NewIcuTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "kuromoji_tokenizer":
					oo := NewKuromojiTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				case "nori_tokenizer":
					oo := NewNoriTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Tokenizer | %w", err)
					}
					s.Tokenizer[key] = oo
				default:
					oo := new(Tokenizer)
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("new(Tokenizer) | %w", err)
					}
					s.Tokenizer[key] = oo
				}
			}

		}
	}
	return nil
}

// NewIndexSettingsAnalysis returns a IndexSettingsAnalysis.
func NewIndexSettingsAnalysis() *IndexSettingsAnalysis {
	r := &IndexSettingsAnalysis{
		Analyzer:   make(map[string]Analyzer),
		CharFilter: make(map[string]CharFilter),
		Filter:     make(map[string]TokenFilter),
		Normalizer: make(map[string]Normalizer),
		Tokenizer:  make(map[string]Tokenizer),
	}

	return r
}
