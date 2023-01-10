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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// CategorizeTextAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/aggregations/bucket.ts#L437-L501
type CategorizeTextAggregation struct {
	// CategorizationAnalyzer The categorization analyzer specifies how the text is analyzed and tokenized
	// before being categorized.
	// The syntax is very similar to that used to define the analyzer in the
	// [Analyze
	// endpoint](https://www.elastic.co/guide/en/elasticsearch/reference/8.0/indices-analyze.html).
	// This property
	// cannot be used at the same time as categorization_filters.
	CategorizationAnalyzer *CategorizeTextAnalyzer `json:"categorization_analyzer,omitempty"`
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
	MaxUniqueTokens *int                   `json:"max_unique_tokens,omitempty"`
	Meta            map[string]interface{} `json:"meta,omitempty"`
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

// NewCategorizeTextAggregation returns a CategorizeTextAggregation.
func NewCategorizeTextAggregation() *CategorizeTextAggregation {
	r := &CategorizeTextAggregation{}

	return r
}
