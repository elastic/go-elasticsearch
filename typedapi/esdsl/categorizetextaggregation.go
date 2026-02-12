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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _categorizeTextAggregation struct {
	v *types.CategorizeTextAggregation
}

// A multi-bucket aggregation that groups semi-structured text into buckets.
func NewCategorizeTextAggregation() *_categorizeTextAggregation {

	return &_categorizeTextAggregation{v: types.NewCategorizeTextAggregation()}

}

func (s *_categorizeTextAggregation) CategorizationAnalyzer(categorizetextanalyzer types.CategorizeTextAnalyzerVariant) *_categorizeTextAggregation {

	s.v.CategorizationAnalyzer = *categorizetextanalyzer.CategorizeTextAnalyzerCaster()

	return s
}

func (s *_categorizeTextAggregation) CategorizationFilters(categorizationfilters ...string) *_categorizeTextAggregation {

	for _, v := range categorizationfilters {

		s.v.CategorizationFilters = append(s.v.CategorizationFilters, v)

	}
	return s
}

func (s *_categorizeTextAggregation) Field(field string) *_categorizeTextAggregation {

	s.v.Field = field

	return s
}

func (s *_categorizeTextAggregation) MaxMatchedTokens(maxmatchedtokens int) *_categorizeTextAggregation {

	s.v.MaxMatchedTokens = &maxmatchedtokens

	return s
}

func (s *_categorizeTextAggregation) MaxUniqueTokens(maxuniquetokens int) *_categorizeTextAggregation {

	s.v.MaxUniqueTokens = &maxuniquetokens

	return s
}

func (s *_categorizeTextAggregation) MinDocCount(mindoccount int) *_categorizeTextAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

func (s *_categorizeTextAggregation) ShardMinDocCount(shardmindoccount int) *_categorizeTextAggregation {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

func (s *_categorizeTextAggregation) ShardSize(shardsize int) *_categorizeTextAggregation {

	s.v.ShardSize = &shardsize

	return s
}

func (s *_categorizeTextAggregation) SimilarityThreshold(similaritythreshold int) *_categorizeTextAggregation {

	s.v.SimilarityThreshold = &similaritythreshold

	return s
}

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
