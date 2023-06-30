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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// IndexSettingsAnalysis type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/indices/_types/IndexSettings.ts#L310-L316
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
				kind := make(map[string]interface{})
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["type"] {
				case "custom":
					oo := NewCustomAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "fingerprint":
					oo := NewFingerprintAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "keyword":
					oo := NewKeywordAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "language":
					oo := NewLanguageAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "nori":
					oo := NewNoriAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "pattern":
					oo := NewPatternAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "simple":
					oo := NewSimpleAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "standard":
					oo := NewStandardAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "stop":
					oo := NewStopAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "whitespace":
					oo := NewWhitespaceAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "icu_analyzer":
					oo := NewIcuAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "kuromoji":
					oo := NewKuromojiAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "snowball":
					oo := NewSnowballAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				case "dutch":
					oo := NewDutchAnalyzer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Analyzer[key] = oo
				default:
					if err := localDec.Decode(&s.Analyzer); err != nil {
						return err
					}
				}
			}

		case "char_filter":
			if s.CharFilter == nil {
				s.CharFilter = make(map[string]CharFilter, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]interface{})
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["type"] {
				case "html_strip":
					oo := NewHtmlStripCharFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.CharFilter[key] = oo
				case "mapping":
					oo := NewMappingCharFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.CharFilter[key] = oo
				case "pattern_replace":
					oo := NewPatternReplaceCharFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.CharFilter[key] = oo
				case "icu_normalizer":
					oo := NewIcuNormalizationCharFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.CharFilter[key] = oo
				case "kuromoji_iteration_mark":
					oo := NewKuromojiIterationMarkCharFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.CharFilter[key] = oo
				default:
					if err := localDec.Decode(&s.CharFilter); err != nil {
						return err
					}
				}
			}

		case "filter":
			if s.Filter == nil {
				s.Filter = make(map[string]TokenFilter, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]interface{})
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["type"] {
				case "asciifolding":
					oo := NewAsciiFoldingTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "common_grams":
					oo := NewCommonGramsTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "condition":
					oo := NewConditionTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "delimited_payload":
					oo := NewDelimitedPayloadTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "edge_ngram":
					oo := NewEdgeNGramTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "elision":
					oo := NewElisionTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "fingerprint":
					oo := NewFingerprintTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "hunspell":
					oo := NewHunspellTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "hyphenation_decompounder":
					oo := NewHyphenationDecompounderTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "keep_types":
					oo := NewKeepTypesTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "keep":
					oo := NewKeepWordsTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "keyword_marker":
					oo := NewKeywordMarkerTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "kstem":
					oo := NewKStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "length":
					oo := NewLengthTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "limit":
					oo := NewLimitTokenCountTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "lowercase":
					oo := NewLowercaseTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "multiplexer":
					oo := NewMultiplexerTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "ngram":
					oo := NewNGramTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "nori_part_of_speech":
					oo := NewNoriPartOfSpeechTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "pattern_capture":
					oo := NewPatternCaptureTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "pattern_replace":
					oo := NewPatternReplaceTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "porter_stem":
					oo := NewPorterStemTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "predicate_token_filter":
					oo := NewPredicateTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "remove_duplicates":
					oo := NewRemoveDuplicatesTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "reverse":
					oo := NewReverseTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "shingle":
					oo := NewShingleTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "snowball":
					oo := NewSnowballTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "stemmer_override":
					oo := NewStemmerOverrideTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "stemmer":
					oo := NewStemmerTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "stop":
					oo := NewStopTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "synonym_graph":
					oo := NewSynonymGraphTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "synonym":
					oo := NewSynonymTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "trim":
					oo := NewTrimTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "truncate":
					oo := NewTruncateTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "unique":
					oo := NewUniqueTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "uppercase":
					oo := NewUppercaseTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "word_delimiter_graph":
					oo := NewWordDelimiterGraphTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "word_delimiter":
					oo := NewWordDelimiterTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "kuromoji_stemmer":
					oo := NewKuromojiStemmerTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "kuromoji_readingform":
					oo := NewKuromojiReadingFormTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "kuromoji_part_of_speech":
					oo := NewKuromojiPartOfSpeechTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "icu_tokenizer":
					oo := NewIcuTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "icu_collation":
					oo := NewIcuCollationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "icu_folding":
					oo := NewIcuFoldingTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "icu_normalizer":
					oo := NewIcuNormalizationTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "icu_transform":
					oo := NewIcuTransformTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "phonetic":
					oo := NewPhoneticTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				case "dictionary_decompounder":
					oo := NewDictionaryDecompounderTokenFilter()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Filter[key] = oo
				default:
					if err := localDec.Decode(&s.Filter); err != nil {
						return err
					}
				}
			}

		case "normalizer":
			if s.Normalizer == nil {
				s.Normalizer = make(map[string]Normalizer, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]interface{})
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["type"] {
				case "lowercase":
					oo := NewLowercaseNormalizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Normalizer[key] = oo
				case "custom":
					oo := NewCustomNormalizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Normalizer[key] = oo
				default:
					if err := localDec.Decode(&s.Normalizer); err != nil {
						return err
					}
				}
			}

		case "tokenizer":
			if s.Tokenizer == nil {
				s.Tokenizer = make(map[string]Tokenizer, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]interface{})
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["type"] {
				case "char_group":
					oo := NewCharGroupTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "edge_ngram":
					oo := NewEdgeNGramTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "keyword":
					oo := NewKeywordTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "letter":
					oo := NewLetterTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "lowercase":
					oo := NewLowercaseTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "ngram":
					oo := NewNGramTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "nori_tokenizer":
					oo := NewNoriTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "path_hierarchy":
					oo := NewPathHierarchyTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "standard":
					oo := NewStandardTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "uax_url_email":
					oo := NewUaxEmailUrlTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "whitespace":
					oo := NewWhitespaceTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "kuromoji_tokenizer":
					oo := NewKuromojiTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "pattern":
					oo := NewPatternTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				case "icu_tokenizer":
					oo := NewIcuTokenizer()
					if err := localDec.Decode(&oo); err != nil {
						return err
					}
					s.Tokenizer[key] = oo
				default:
					if err := localDec.Decode(&s.Tokenizer); err != nil {
						return err
					}
				}
			}

		}
	}
	return nil
}

// NewIndexSettingsAnalysis returns a IndexSettingsAnalysis.
func NewIndexSettingsAnalysis() *IndexSettingsAnalysis {
	r := &IndexSettingsAnalysis{
		Analyzer:   make(map[string]Analyzer, 0),
		CharFilter: make(map[string]CharFilter, 0),
		Filter:     make(map[string]TokenFilter, 0),
		Normalizer: make(map[string]Normalizer, 0),
		Tokenizer:  make(map[string]Tokenizer, 0),
	}

	return r
}
