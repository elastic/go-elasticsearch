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

// CategorizeTextAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L433-L497
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
	Field Field `json:"field"`
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
	MaxUniqueTokens *int      `json:"max_unique_tokens,omitempty"`
	Meta            *Metadata `json:"meta,omitempty"`
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

// CategorizeTextAggregationBuilder holds CategorizeTextAggregation struct and provides a builder API.
type CategorizeTextAggregationBuilder struct {
	v *CategorizeTextAggregation
}

// NewCategorizeTextAggregation provides a builder for the CategorizeTextAggregation struct.
func NewCategorizeTextAggregationBuilder() *CategorizeTextAggregationBuilder {
	r := CategorizeTextAggregationBuilder{
		&CategorizeTextAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the CategorizeTextAggregation struct
func (rb *CategorizeTextAggregationBuilder) Build() CategorizeTextAggregation {
	return *rb.v
}

// CategorizationAnalyzer The categorization analyzer specifies how the text is analyzed and tokenized
// before being categorized.
// The syntax is very similar to that used to define the analyzer in the
// [Analyze
// endpoint](https://www.elastic.co/guide/en/elasticsearch/reference/8.0/indices-analyze.html).
// This property
// cannot be used at the same time as categorization_filters.

func (rb *CategorizeTextAggregationBuilder) CategorizationAnalyzer(categorizationanalyzer *CategorizeTextAnalyzerBuilder) *CategorizeTextAggregationBuilder {
	v := categorizationanalyzer.Build()
	rb.v.CategorizationAnalyzer = &v
	return rb
}

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

func (rb *CategorizeTextAggregationBuilder) CategorizationFilters(categorization_filters ...string) *CategorizeTextAggregationBuilder {
	rb.v.CategorizationFilters = categorization_filters
	return rb
}

// Field The semi-structured text field to categorize.

func (rb *CategorizeTextAggregationBuilder) Field(field Field) *CategorizeTextAggregationBuilder {
	rb.v.Field = field
	return rb
}

// MaxMatchedTokens The maximum number of token positions to match on before attempting to merge
// categories. Larger
// values will use more memory and create narrower categories. Max allowed value
// is 100.

func (rb *CategorizeTextAggregationBuilder) MaxMatchedTokens(maxmatchedtokens int) *CategorizeTextAggregationBuilder {
	rb.v.MaxMatchedTokens = &maxmatchedtokens
	return rb
}

// MaxUniqueTokens The maximum number of unique tokens at any position up to max_matched_tokens.
// Must be larger than 1.
// Smaller values use less memory and create fewer categories. Larger values
// will use more memory and
// create narrower categories. Max allowed value is 100.

func (rb *CategorizeTextAggregationBuilder) MaxUniqueTokens(maxuniquetokens int) *CategorizeTextAggregationBuilder {
	rb.v.MaxUniqueTokens = &maxuniquetokens
	return rb
}

func (rb *CategorizeTextAggregationBuilder) Meta(meta *MetadataBuilder) *CategorizeTextAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

// MinDocCount The minimum number of documents for a bucket to be returned to the results.

func (rb *CategorizeTextAggregationBuilder) MinDocCount(mindoccount int) *CategorizeTextAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

func (rb *CategorizeTextAggregationBuilder) Name(name string) *CategorizeTextAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// ShardMinDocCount The minimum number of documents for a bucket to be returned from the shard
// before merging.

func (rb *CategorizeTextAggregationBuilder) ShardMinDocCount(shardmindoccount int) *CategorizeTextAggregationBuilder {
	rb.v.ShardMinDocCount = &shardmindoccount
	return rb
}

// ShardSize The number of categorization buckets to return from each shard before merging
// all the results.

func (rb *CategorizeTextAggregationBuilder) ShardSize(shardsize int) *CategorizeTextAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

// SimilarityThreshold The minimum percentage of tokens that must match for text to be added to the
// category bucket. Must
// be between 1 and 100. The larger the value the narrower the categories.
// Larger values will increase memory
// usage and create narrower categories.

func (rb *CategorizeTextAggregationBuilder) SimilarityThreshold(similaritythreshold int) *CategorizeTextAggregationBuilder {
	rb.v.SimilarityThreshold = &similaritythreshold
	return rb
}

// Size The number of buckets to return.

func (rb *CategorizeTextAggregationBuilder) Size(size int) *CategorizeTextAggregationBuilder {
	rb.v.Size = &size
	return rb
}
