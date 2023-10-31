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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// CategorizationAnalyzerDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Analysis.ts#L184-L197
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

				switch rawMsg[0] {
				case '{':

					source := bytes.NewReader(rawMsg)
					localDec := json.NewDecoder(source)
					kind := make(map[string]string, 0)
					localDec.Decode(&kind)
					source.Seek(0, io.SeekStart)

					switch kind["type"] {

					case "html_strip":
						o := NewHtmlStripCharFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.CharFilter = append(s.CharFilter, *o)
					case "mapping":
						o := NewMappingCharFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.CharFilter = append(s.CharFilter, *o)
					case "pattern_replace":
						o := NewPatternReplaceCharFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.CharFilter = append(s.CharFilter, *o)
					case "icu_normalizer":
						o := NewIcuNormalizationCharFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.CharFilter = append(s.CharFilter, *o)
					case "kuromoji_iteration_mark":
						o := NewKuromojiIterationMarkCharFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.CharFilter = append(s.CharFilter, *o)
					default:
						o := new(interface{})
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.CharFilter = append(s.CharFilter, *o)
					}
				default:
					source := bytes.NewReader(rawMsg)
					o := new(interface{})
					if err := json.NewDecoder(source).Decode(&o); err != nil {
						return err
					}
					s.CharFilter = append(s.CharFilter, *o)
				}
			}

		case "filter":

			buf := []json.RawMessage{}
			dec.Decode(&buf)
			for _, rawMsg := range buf {

				switch rawMsg[0] {
				case '{':

					source := bytes.NewReader(rawMsg)
					localDec := json.NewDecoder(source)
					kind := make(map[string]string, 0)
					localDec.Decode(&kind)
					source.Seek(0, io.SeekStart)

					switch kind["type"] {

					case "asciifolding":
						o := NewAsciiFoldingTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "common_grams":
						o := NewCommonGramsTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "condition":
						o := NewConditionTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "delimited_payload":
						o := NewDelimitedPayloadTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "edge_ngram":
						o := NewEdgeNGramTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "elision":
						o := NewElisionTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "fingerprint":
						o := NewFingerprintTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "hunspell":
						o := NewHunspellTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "hyphenation_decompounder":
						o := NewHyphenationDecompounderTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "keep_types":
						o := NewKeepTypesTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "keep":
						o := NewKeepWordsTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "keyword_marker":
						o := NewKeywordMarkerTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "kstem":
						o := NewKStemTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "length":
						o := NewLengthTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "limit":
						o := NewLimitTokenCountTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "lowercase":
						o := NewLowercaseTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "multiplexer":
						o := NewMultiplexerTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "ngram":
						o := NewNGramTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "nori_part_of_speech":
						o := NewNoriPartOfSpeechTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "pattern_capture":
						o := NewPatternCaptureTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "pattern_replace":
						o := NewPatternReplaceTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "porter_stem":
						o := NewPorterStemTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "predicate_token_filter":
						o := NewPredicateTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "remove_duplicates":
						o := NewRemoveDuplicatesTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "reverse":
						o := NewReverseTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "shingle":
						o := NewShingleTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "snowball":
						o := NewSnowballTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "stemmer_override":
						o := NewStemmerOverrideTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "stemmer":
						o := NewStemmerTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "stop":
						o := NewStopTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "synonym_graph":
						o := NewSynonymGraphTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "synonym":
						o := NewSynonymTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "trim":
						o := NewTrimTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "truncate":
						o := NewTruncateTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "unique":
						o := NewUniqueTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "uppercase":
						o := NewUppercaseTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "word_delimiter_graph":
						o := NewWordDelimiterGraphTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "word_delimiter":
						o := NewWordDelimiterTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "kuromoji_stemmer":
						o := NewKuromojiStemmerTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "kuromoji_readingform":
						o := NewKuromojiReadingFormTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "kuromoji_part_of_speech":
						o := NewKuromojiPartOfSpeechTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "icu_tokenizer":
						o := NewIcuTokenizer()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "icu_collation":
						o := NewIcuCollationTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "icu_folding":
						o := NewIcuFoldingTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "icu_normalizer":
						o := NewIcuNormalizationTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "icu_transform":
						o := NewIcuTransformTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "phonetic":
						o := NewPhoneticTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					case "dictionary_decompounder":
						o := NewDictionaryDecompounderTokenFilter()
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					default:
						o := new(interface{})
						if err := localDec.Decode(&o); err != nil {
							return err
						}
						s.Filter = append(s.Filter, *o)
					}
				default:
					source := bytes.NewReader(rawMsg)
					o := new(interface{})
					if err := json.NewDecoder(source).Decode(&o); err != nil {
						return err
					}
					s.Filter = append(s.Filter, *o)
				}
			}

		case "tokenizer":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':

				kind := make(map[string]string, 0)
				localDec.Decode(&kind)
				source.Seek(0, io.SeekStart)

				switch kind["type"] {

				case "char_group":
					o := NewCharGroupTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "edge_ngram":
					o := NewEdgeNGramTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "keyword":
					o := NewKeywordTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "letter":
					o := NewLetterTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "lowercase":
					o := NewLowercaseTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "ngram":
					o := NewNGramTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "nori_tokenizer":
					o := NewNoriTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "path_hierarchy":
					o := NewPathHierarchyTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "standard":
					o := NewStandardTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "uax_url_email":
					o := NewUaxEmailUrlTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "whitespace":
					o := NewWhitespaceTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "kuromoji_tokenizer":
					o := NewKuromojiTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "pattern":
					o := NewPatternTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				case "icu_tokenizer":
					o := NewIcuTokenizer()
					if err := localDec.Decode(&o); err != nil {
						return err
					}
					s.Tokenizer = *o
				default:
					if err := localDec.Decode(&s.Tokenizer); err != nil {
						return err
					}
				}
			default:
				if err := localDec.Decode(&s.Tokenizer); err != nil {
					return err
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
