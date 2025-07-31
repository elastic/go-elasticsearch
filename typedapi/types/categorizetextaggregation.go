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
	"strconv"
)

// CategorizeTextAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/bucket.ts#L1117-L1183
type CategorizeTextAggregation struct {
	// CategorizationAnalyzer The categorization analyzer specifies how the text is analyzed and tokenized
	// before being categorized.
	// The syntax is very similar to that used to define the analyzer in the
	// `_analyze` endpoint. This property
	// cannot be used at the same time as categorization_filters.
	CategorizationAnalyzer CategorizeTextAnalyzer `json:"categorization_analyzer,omitempty"`
	// CategorizationFilters This property expects an array of regular expressions. The expressions are
	// used to filter out matching
	// sequences from the categorization field values. You can use this
	// functionality to fine tune the categorization
	// by excluding sequences from consideration when categories are defined. For
	// example, you can exclude SQL
	// statements that appear in your log files. This property cannot be used at the
	// same time as categorization_analyzer.
	// If you only want to define simple regular expression filters that are applied
	// prior to tokenization, setting
	// this property is the easiest method. If you also want to customize the
	// tokenizer or post-tokenization filtering,
	// use the categorization_analyzer property instead and include the filters as
	// pattern_replace character filters.
	CategorizationFilters []string `json:"categorization_filters,omitempty"`
	// Field The semi-structured text field to categorize.
	Field string `json:"field"`
	// MaxMatchedTokens The maximum number of token positions to match on before attempting to merge
	// categories. Larger
	// values will use more memory and create narrower categories. Max allowed value
	// is 100.
	MaxMatchedTokens *int `json:"max_matched_tokens,omitempty"`
	// MaxUniqueTokens The maximum number of unique tokens at any position up to max_matched_tokens.
	// Must be larger than 1.
	// Smaller values use less memory and create fewer categories. Larger values
	// will use more memory and
	// create narrower categories. Max allowed value is 100.
	MaxUniqueTokens *int `json:"max_unique_tokens,omitempty"`
	// MinDocCount The minimum number of documents in a bucket to be returned to the results.
	MinDocCount *int `json:"min_doc_count,omitempty"`
	// ShardMinDocCount The minimum number of documents in a bucket to be returned from the shard
	// before merging.
	ShardMinDocCount *int `json:"shard_min_doc_count,omitempty"`
	// ShardSize The number of categorization buckets to return from each shard before merging
	// all the results.
	ShardSize *int `json:"shard_size,omitempty"`
	// SimilarityThreshold The minimum percentage of tokens that must match for text to be added to the
	// category bucket. Must
	// be between 1 and 100. The larger the value the narrower the categories.
	// Larger values will increase memory
	// usage and create narrower categories.
	SimilarityThreshold *int `json:"similarity_threshold,omitempty"`
	// Size The number of buckets to return.
	Size *int `json:"size,omitempty"`
}

func (s *CategorizeTextAggregation) UnmarshalJSON(data []byte) error {

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

		case "categorization_analyzer":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "CategorizationAnalyzer", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		categorizationanalyzer_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "CategorizationAnalyzer", err)
				}

				switch t {

				case "char_filter", "filter", "tokenizer":
					o := NewCustomCategorizeTextAnalyzer()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "CategorizationAnalyzer", err)
					}
					s.CategorizationAnalyzer = o
					break categorizationanalyzer_field

				}
			}
			if s.CategorizationAnalyzer == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.CategorizationAnalyzer); err != nil {
					return fmt.Errorf("%s | %w", "CategorizationAnalyzer", err)
				}
			}

		case "categorization_filters":
			if err := dec.Decode(&s.CategorizationFilters); err != nil {
				return fmt.Errorf("%s | %w", "CategorizationFilters", err)
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "max_matched_tokens":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxMatchedTokens", err)
				}
				s.MaxMatchedTokens = &value
			case float64:
				f := int(v)
				s.MaxMatchedTokens = &f
			}

		case "max_unique_tokens":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxUniqueTokens", err)
				}
				s.MaxUniqueTokens = &value
			case float64:
				f := int(v)
				s.MaxUniqueTokens = &f
			}

		case "min_doc_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinDocCount", err)
				}
				s.MinDocCount = &value
			case float64:
				f := int(v)
				s.MinDocCount = &f
			}

		case "shard_min_doc_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardMinDocCount", err)
				}
				s.ShardMinDocCount = &value
			case float64:
				f := int(v)
				s.ShardMinDocCount = &f
			}

		case "shard_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardSize", err)
				}
				s.ShardSize = &value
			case float64:
				f := int(v)
				s.ShardSize = &f
			}

		case "similarity_threshold":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SimilarityThreshold", err)
				}
				s.SimilarityThreshold = &value
			case float64:
				f := int(v)
				s.SimilarityThreshold = &f
			}

		case "size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Size", err)
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		}
	}
	return nil
}

// NewCategorizeTextAggregation returns a CategorizeTextAggregation.
func NewCategorizeTextAggregation() *CategorizeTextAggregation {
	r := &CategorizeTextAggregation{}

	return r
}
