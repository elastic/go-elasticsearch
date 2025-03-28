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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _categorizeTextAggregation struct {
	v *types.CategorizeTextAggregation
}

// A multi-bucket aggregation that groups semi-structured text into buckets.
func NewCategorizeTextAggregation() *_categorizeTextAggregation {

	return &_categorizeTextAggregation{v: types.NewCategorizeTextAggregation()}

}

// The categorization analyzer specifies how the text is analyzed and tokenized
// before being categorized.
// The syntax is very similar to that used to define the analyzer in the
// [Analyze
// endpoint](https://www.elastic.co/guide/en/elasticsearch/reference/8.0/indices-analyze.html).
// This property
// cannot be used at the same time as categorization_filters.
func (s *_categorizeTextAggregation) CategorizationAnalyzer(categorizetextanalyzer types.CategorizeTextAnalyzerVariant) *_categorizeTextAggregation {

	s.v.CategorizationAnalyzer = *categorizetextanalyzer.CategorizeTextAnalyzerCaster()

	return s
}

// This property expects an array of regular expressions. The expressions are
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
func (s *_categorizeTextAggregation) CategorizationFilters(categorizationfilters ...string) *_categorizeTextAggregation {

	for _, v := range categorizationfilters {

		s.v.CategorizationFilters = append(s.v.CategorizationFilters, v)

	}
	return s
}

// The semi-structured text field to categorize.
func (s *_categorizeTextAggregation) Field(field string) *_categorizeTextAggregation {

	s.v.Field = field

	return s
}

// The maximum number of token positions to match on before attempting to merge
// categories. Larger
// values will use more memory and create narrower categories. Max allowed value
// is 100.
func (s *_categorizeTextAggregation) MaxMatchedTokens(maxmatchedtokens int) *_categorizeTextAggregation {

	s.v.MaxMatchedTokens = &maxmatchedtokens

	return s
}

// The maximum number of unique tokens at any position up to max_matched_tokens.
// Must be larger than 1.
// Smaller values use less memory and create fewer categories. Larger values
// will use more memory and
// create narrower categories. Max allowed value is 100.
func (s *_categorizeTextAggregation) MaxUniqueTokens(maxuniquetokens int) *_categorizeTextAggregation {

	s.v.MaxUniqueTokens = &maxuniquetokens

	return s
}

// The minimum number of documents in a bucket to be returned to the results.
func (s *_categorizeTextAggregation) MinDocCount(mindoccount int) *_categorizeTextAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

// The minimum number of documents in a bucket to be returned from the shard
// before merging.
func (s *_categorizeTextAggregation) ShardMinDocCount(shardmindoccount int) *_categorizeTextAggregation {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

// The number of categorization buckets to return from each shard before merging
// all the results.
func (s *_categorizeTextAggregation) ShardSize(shardsize int) *_categorizeTextAggregation {

	s.v.ShardSize = &shardsize

	return s
}

// The minimum percentage of tokens that must match for text to be added to the
// category bucket. Must
// be between 1 and 100. The larger the value the narrower the categories.
// Larger values will increase memory
// usage and create narrower categories.
func (s *_categorizeTextAggregation) SimilarityThreshold(similaritythreshold int) *_categorizeTextAggregation {

	s.v.SimilarityThreshold = &similaritythreshold

	return s
}

// The number of buckets to return.
func (s *_categorizeTextAggregation) Size(size int) *_categorizeTextAggregation {

	s.v.Size = &size

	return s
}

func (s *_categorizeTextAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.CategorizeText = s.v

	return container
}

func (s *_categorizeTextAggregation) CategorizeTextAggregationCaster() *types.CategorizeTextAggregation {
	return s.v
}
