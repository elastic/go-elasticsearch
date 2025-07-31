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

// CategorizationAnalyzerDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/Analysis.ts#L184-L198
type CategorizationAnalyzerDefinition struct {
	// CharFilter One or more character filters. In addition to the built-in character filters,
	// other plugins can provide more character filters. If this property is not
	// specified, no character filters are applied prior to categorization. If you
	// are customizing some other aspect of the analyzer and you need to achieve the
	// equivalent of `categorization_filters` (which are not permitted when some
	// other aspect of the analyzer is customized), add them here as pattern replace
	// character filters.
	CharFilter []CharFilter `json:"char_filter,omitempty"`
	// Filter One or more token filters. In addition to the built-in token filters, other
	// plugins can provide more token filters. If this property is not specified, no
	// token filters are applied prior to categorization.
	Filter []TokenFilter `json:"filter,omitempty"`
	// Tokenizer The name or definition of the tokenizer to use after character filters are
	// applied. This property is compulsory if `categorization_analyzer` is
	// specified as an object. Machine learning provides a tokenizer called
	// `ml_standard` that tokenizes in a way that has been determined to produce
	// good categorization results on a variety of log file formats for logs in
	// English. If you want to use that tokenizer but change the character or token
	// filters, specify "tokenizer": "ml_standard" in your
	// `categorization_analyzer`. Additionally, the `ml_classic` tokenizer is
	// available, which tokenizes in the same way as the non-customizable tokenizer
	// in old versions of the product (before 6.2). `ml_classic` was the default
	// categorization tokenizer in versions 6.2 to 7.13, so if you need
	// categorization identical to the default for jobs created in these versions,
	// specify "tokenizer": "ml_classic" in your `categorization_analyzer`.
	Tokenizer Tokenizer `json:"tokenizer,omitempty"`
}

func (s *CategorizationAnalyzerDefinition) UnmarshalJSON(data []byte) error {

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

		case "char_filter":

			buf := []json.RawMessage{}
			dec.Decode(&buf)
			for _, rawMsg := range buf {

				source := bytes.NewReader(rawMsg)
				localDec := json.NewDecoder(source)
				kind := make(map[string]string, 0)
				localDec.Decode(&kind)
				source.Seek(0, io.SeekStart)

				switch kind["type"] {

				case "html_strip":
					o := NewHtmlStripCharFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "html_strip", err)
					}
					s.CharFilter = append(s.CharFilter, *o)
				case "mapping":
					o := NewMappingCharFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "mapping", err)
					}
					s.CharFilter = append(s.CharFilter, *o)
				case "pattern_replace":
					o := NewPatternReplaceCharFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "pattern_replace", err)
					}
					s.CharFilter = append(s.CharFilter, *o)
				case "icu_normalizer":
					o := NewIcuNormalizationCharFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "icu_normalizer", err)
					}
					s.CharFilter = append(s.CharFilter, *o)
				case "kuromoji_iteration_mark":
					o := NewKuromojiIterationMarkCharFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "kuromoji_iteration_mark", err)
					}
					s.CharFilter = append(s.CharFilter, *o)
				default:
					o := new(any)
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("CharFilter | %w", err)
					}
					s.CharFilter = append(s.CharFilter, *o)
				}
			}

		case "filter":

			buf := []json.RawMessage{}
			dec.Decode(&buf)
			for _, rawMsg := range buf {

				source := bytes.NewReader(rawMsg)
				localDec := json.NewDecoder(source)
				kind := make(map[string]string, 0)
				localDec.Decode(&kind)
				source.Seek(0, io.SeekStart)

				switch kind["type"] {

				case "apostrophe":
					o := NewApostropheTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "apostrophe", err)
					}
					s.Filter = append(s.Filter, *o)
				case "arabic_stem":
					o := NewArabicStemTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "arabic_stem", err)
					}
					s.Filter = append(s.Filter, *o)
				case "arabic_normalization":
					o := NewArabicNormalizationTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "arabic_normalization", err)
					}
					s.Filter = append(s.Filter, *o)
				case "asciifolding":
					o := NewAsciiFoldingTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "asciifolding", err)
					}
					s.Filter = append(s.Filter, *o)
				case "bengali_normalization":
					o := NewBengaliNormalizationTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "bengali_normalization", err)
					}
					s.Filter = append(s.Filter, *o)
				case "brazilian_stem":
					o := NewBrazilianStemTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "brazilian_stem", err)
					}
					s.Filter = append(s.Filter, *o)
				case "cjk_bigram":
					o := NewCjkBigramTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "cjk_bigram", err)
					}
					s.Filter = append(s.Filter, *o)
				case "cjk_width":
					o := NewCjkWidthTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "cjk_width", err)
					}
					s.Filter = append(s.Filter, *o)
				case "classic":
					o := NewClassicTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "classic", err)
					}
					s.Filter = append(s.Filter, *o)
				case "common_grams":
					o := NewCommonGramsTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "common_grams", err)
					}
					s.Filter = append(s.Filter, *o)
				case "condition":
					o := NewConditionTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "condition", err)
					}
					s.Filter = append(s.Filter, *o)
				case "czech_stem":
					o := NewCzechStemTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "czech_stem", err)
					}
					s.Filter = append(s.Filter, *o)
				case "decimal_digit":
					o := NewDecimalDigitTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "decimal_digit", err)
					}
					s.Filter = append(s.Filter, *o)
				case "delimited_payload":
					o := NewDelimitedPayloadTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "delimited_payload", err)
					}
					s.Filter = append(s.Filter, *o)
				case "dutch_stem":
					o := NewDutchStemTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "dutch_stem", err)
					}
					s.Filter = append(s.Filter, *o)
				case "edge_ngram":
					o := NewEdgeNGramTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "edge_ngram", err)
					}
					s.Filter = append(s.Filter, *o)
				case "elision":
					o := NewElisionTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "elision", err)
					}
					s.Filter = append(s.Filter, *o)
				case "fingerprint":
					o := NewFingerprintTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "fingerprint", err)
					}
					s.Filter = append(s.Filter, *o)
				case "flatten_graph":
					o := NewFlattenGraphTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "flatten_graph", err)
					}
					s.Filter = append(s.Filter, *o)
				case "french_stem":
					o := NewFrenchStemTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "french_stem", err)
					}
					s.Filter = append(s.Filter, *o)
				case "german_normalization":
					o := NewGermanNormalizationTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "german_normalization", err)
					}
					s.Filter = append(s.Filter, *o)
				case "german_stem":
					o := NewGermanStemTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "german_stem", err)
					}
					s.Filter = append(s.Filter, *o)
				case "hindi_normalization":
					o := NewHindiNormalizationTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "hindi_normalization", err)
					}
					s.Filter = append(s.Filter, *o)
				case "hunspell":
					o := NewHunspellTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "hunspell", err)
					}
					s.Filter = append(s.Filter, *o)
				case "hyphenation_decompounder":
					o := NewHyphenationDecompounderTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "hyphenation_decompounder", err)
					}
					s.Filter = append(s.Filter, *o)
				case "indic_normalization":
					o := NewIndicNormalizationTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "indic_normalization", err)
					}
					s.Filter = append(s.Filter, *o)
				case "keep_types":
					o := NewKeepTypesTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "keep_types", err)
					}
					s.Filter = append(s.Filter, *o)
				case "keep":
					o := NewKeepWordsTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "keep", err)
					}
					s.Filter = append(s.Filter, *o)
				case "keyword_marker":
					o := NewKeywordMarkerTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "keyword_marker", err)
					}
					s.Filter = append(s.Filter, *o)
				case "keyword_repeat":
					o := NewKeywordRepeatTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "keyword_repeat", err)
					}
					s.Filter = append(s.Filter, *o)
				case "kstem":
					o := NewKStemTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "kstem", err)
					}
					s.Filter = append(s.Filter, *o)
				case "length":
					o := NewLengthTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "length", err)
					}
					s.Filter = append(s.Filter, *o)
				case "limit":
					o := NewLimitTokenCountTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "limit", err)
					}
					s.Filter = append(s.Filter, *o)
				case "lowercase":
					o := NewLowercaseTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "lowercase", err)
					}
					s.Filter = append(s.Filter, *o)
				case "min_hash":
					o := NewMinHashTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "min_hash", err)
					}
					s.Filter = append(s.Filter, *o)
				case "multiplexer":
					o := NewMultiplexerTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "multiplexer", err)
					}
					s.Filter = append(s.Filter, *o)
				case "ngram":
					o := NewNGramTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "ngram", err)
					}
					s.Filter = append(s.Filter, *o)
				case "nori_part_of_speech":
					o := NewNoriPartOfSpeechTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "nori_part_of_speech", err)
					}
					s.Filter = append(s.Filter, *o)
				case "pattern_capture":
					o := NewPatternCaptureTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "pattern_capture", err)
					}
					s.Filter = append(s.Filter, *o)
				case "pattern_replace":
					o := NewPatternReplaceTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "pattern_replace", err)
					}
					s.Filter = append(s.Filter, *o)
				case "persian_normalization":
					o := NewPersianNormalizationTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "persian_normalization", err)
					}
					s.Filter = append(s.Filter, *o)
				case "persian_stem":
					o := NewPersianStemTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "persian_stem", err)
					}
					s.Filter = append(s.Filter, *o)
				case "porter_stem":
					o := NewPorterStemTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "porter_stem", err)
					}
					s.Filter = append(s.Filter, *o)
				case "predicate_token_filter":
					o := NewPredicateTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "predicate_token_filter", err)
					}
					s.Filter = append(s.Filter, *o)
				case "remove_duplicates":
					o := NewRemoveDuplicatesTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "remove_duplicates", err)
					}
					s.Filter = append(s.Filter, *o)
				case "reverse":
					o := NewReverseTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "reverse", err)
					}
					s.Filter = append(s.Filter, *o)
				case "russian_stem":
					o := NewRussianStemTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "russian_stem", err)
					}
					s.Filter = append(s.Filter, *o)
				case "scandinavian_folding":
					o := NewScandinavianFoldingTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "scandinavian_folding", err)
					}
					s.Filter = append(s.Filter, *o)
				case "scandinavian_normalization":
					o := NewScandinavianNormalizationTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "scandinavian_normalization", err)
					}
					s.Filter = append(s.Filter, *o)
				case "serbian_normalization":
					o := NewSerbianNormalizationTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "serbian_normalization", err)
					}
					s.Filter = append(s.Filter, *o)
				case "shingle":
					o := NewShingleTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "shingle", err)
					}
					s.Filter = append(s.Filter, *o)
				case "snowball":
					o := NewSnowballTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "snowball", err)
					}
					s.Filter = append(s.Filter, *o)
				case "sorani_normalization":
					o := NewSoraniNormalizationTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "sorani_normalization", err)
					}
					s.Filter = append(s.Filter, *o)
				case "stemmer_override":
					o := NewStemmerOverrideTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "stemmer_override", err)
					}
					s.Filter = append(s.Filter, *o)
				case "stemmer":
					o := NewStemmerTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "stemmer", err)
					}
					s.Filter = append(s.Filter, *o)
				case "stop":
					o := NewStopTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "stop", err)
					}
					s.Filter = append(s.Filter, *o)
				case "synonym_graph":
					o := NewSynonymGraphTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "synonym_graph", err)
					}
					s.Filter = append(s.Filter, *o)
				case "synonym":
					o := NewSynonymTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "synonym", err)
					}
					s.Filter = append(s.Filter, *o)
				case "trim":
					o := NewTrimTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "trim", err)
					}
					s.Filter = append(s.Filter, *o)
				case "truncate":
					o := NewTruncateTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "truncate", err)
					}
					s.Filter = append(s.Filter, *o)
				case "unique":
					o := NewUniqueTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "unique", err)
					}
					s.Filter = append(s.Filter, *o)
				case "uppercase":
					o := NewUppercaseTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "uppercase", err)
					}
					s.Filter = append(s.Filter, *o)
				case "word_delimiter_graph":
					o := NewWordDelimiterGraphTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "word_delimiter_graph", err)
					}
					s.Filter = append(s.Filter, *o)
				case "word_delimiter":
					o := NewWordDelimiterTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "word_delimiter", err)
					}
					s.Filter = append(s.Filter, *o)
				case "ja_stop":
					o := NewJaStopTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "ja_stop", err)
					}
					s.Filter = append(s.Filter, *o)
				case "kuromoji_stemmer":
					o := NewKuromojiStemmerTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "kuromoji_stemmer", err)
					}
					s.Filter = append(s.Filter, *o)
				case "kuromoji_readingform":
					o := NewKuromojiReadingFormTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "kuromoji_readingform", err)
					}
					s.Filter = append(s.Filter, *o)
				case "kuromoji_part_of_speech":
					o := NewKuromojiPartOfSpeechTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "kuromoji_part_of_speech", err)
					}
					s.Filter = append(s.Filter, *o)
				case "icu_collation":
					o := NewIcuCollationTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "icu_collation", err)
					}
					s.Filter = append(s.Filter, *o)
				case "icu_folding":
					o := NewIcuFoldingTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "icu_folding", err)
					}
					s.Filter = append(s.Filter, *o)
				case "icu_normalizer":
					o := NewIcuNormalizationTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "icu_normalizer", err)
					}
					s.Filter = append(s.Filter, *o)
				case "icu_transform":
					o := NewIcuTransformTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "icu_transform", err)
					}
					s.Filter = append(s.Filter, *o)
				case "phonetic":
					o := NewPhoneticTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "phonetic", err)
					}
					s.Filter = append(s.Filter, *o)
				case "dictionary_decompounder":
					o := NewDictionaryDecompounderTokenFilter()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "dictionary_decompounder", err)
					}
					s.Filter = append(s.Filter, *o)
				default:
					o := new(any)
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("Filter | %w", err)
					}
					s.Filter = append(s.Filter, *o)
				}
			}

		case "tokenizer":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			kind := make(map[string]string, 0)
			localDec := json.NewDecoder(source)
			localDec.Decode(&kind)
			source.Seek(0, io.SeekStart)

			switch kind["type"] {

			case "char_group":
				o := NewCharGroupTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "char_group", err)
				}
				s.Tokenizer = *o
			case "classic":
				o := NewClassicTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "classic", err)
				}
				s.Tokenizer = *o
			case "edge_ngram":
				o := NewEdgeNGramTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "edge_ngram", err)
				}
				s.Tokenizer = *o
			case "keyword":
				o := NewKeywordTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "keyword", err)
				}
				s.Tokenizer = *o
			case "letter":
				o := NewLetterTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "letter", err)
				}
				s.Tokenizer = *o
			case "lowercase":
				o := NewLowercaseTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "lowercase", err)
				}
				s.Tokenizer = *o
			case "ngram":
				o := NewNGramTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "ngram", err)
				}
				s.Tokenizer = *o
			case "path_hierarchy":
				o := NewPathHierarchyTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "path_hierarchy", err)
				}
				s.Tokenizer = *o
			case "pattern":
				o := NewPatternTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "pattern", err)
				}
				s.Tokenizer = *o
			case "simple_pattern":
				o := NewSimplePatternTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "simple_pattern", err)
				}
				s.Tokenizer = *o
			case "simple_pattern_split":
				o := NewSimplePatternSplitTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "simple_pattern_split", err)
				}
				s.Tokenizer = *o
			case "standard":
				o := NewStandardTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "standard", err)
				}
				s.Tokenizer = *o
			case "thai":
				o := NewThaiTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "thai", err)
				}
				s.Tokenizer = *o
			case "uax_url_email":
				o := NewUaxEmailUrlTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "uax_url_email", err)
				}
				s.Tokenizer = *o
			case "whitespace":
				o := NewWhitespaceTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "whitespace", err)
				}
				s.Tokenizer = *o
			case "icu_tokenizer":
				o := NewIcuTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "icu_tokenizer", err)
				}
				s.Tokenizer = *o
			case "kuromoji_tokenizer":
				o := NewKuromojiTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "kuromoji_tokenizer", err)
				}
				s.Tokenizer = *o
			case "nori_tokenizer":
				o := NewNoriTokenizer()
				if err := localDec.Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "nori_tokenizer", err)
				}
				s.Tokenizer = *o
			default:
				if err := localDec.Decode(&s.Tokenizer); err != nil {
					return fmt.Errorf("Tokenizer | %w", err)
				}
			}

		}
	}
	return nil
}

// NewCategorizationAnalyzerDefinition returns a CategorizationAnalyzerDefinition.
func NewCategorizationAnalyzerDefinition() *CategorizationAnalyzerDefinition {
	r := &CategorizationAnalyzerDefinition{}

	return r
}
