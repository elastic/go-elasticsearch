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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// CategorizeTextAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/aggregations/bucket.ts#L437-L501
type CategorizeTextAggregation struct {
	// CategorizationAnalyzer The categorization analyzer specifies how the text is analyzed and tokenized
	// before being categorized.
	// The syntax is very similar to that used to define the analyzer in the
	// [Analyze
	// endpoint](https://www.elastic.co/guide/en/elasticsearch/reference/8.0/indices-analyze.html).
	// This property
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
	MaxUniqueTokens *int                       `json:"max_unique_tokens,omitempty"`
	Meta            map[string]json.RawMessage `json:"meta,omitempty"`
	// MinDocCount The minimum number of documents for a bucket to be returned to the results.
	MinDocCount *int    `json:"min_doc_count,omitempty"`
	Name        *string `json:"name,omitempty"`
	// ShardMinDocCount The minimum number of documents for a bucket to be returned from the shard
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
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := NewCustomCategorizeTextAnalyzer()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.CategorizationAnalyzer = *o

			default:
				if err := localDec.Decode(&s.CategorizationAnalyzer); err != nil {
					return err
				}
			}

		case "categorization_filters":
			if err := dec.Decode(&s.CategorizationFilters); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "max_matched_tokens":
			if err := dec.Decode(&s.MaxMatchedTokens); err != nil {
				return err
			}

		case "max_unique_tokens":
			if err := dec.Decode(&s.MaxUniqueTokens); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "min_doc_count":
			if err := dec.Decode(&s.MinDocCount); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "shard_min_doc_count":
			if err := dec.Decode(&s.ShardMinDocCount); err != nil {
				return err
			}

		case "shard_size":
			if err := dec.Decode(&s.ShardSize); err != nil {
				return err
			}

		case "similarity_threshold":
			if err := dec.Decode(&s.SimilarityThreshold); err != nil {
				return err
			}

		case "size":
			if err := dec.Decode(&s.Size); err != nil {
				return err
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
