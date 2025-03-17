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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"
)

type _significantTextAggregation struct {
	v *types.SignificantTextAggregation
}

// Returns interesting or unusual occurrences of free-text terms in a set.
func NewSignificantTextAggregation() *_significantTextAggregation {

	return &_significantTextAggregation{v: types.NewSignificantTextAggregation()}

}

// A background filter that can be used to focus in on significant terms within
// a narrower context, instead of the entire index.
func (s *_significantTextAggregation) BackgroundFilter(backgroundfilter types.QueryVariant) *_significantTextAggregation {

	s.v.BackgroundFilter = backgroundfilter.QueryCaster()

	return s
}

// Use Chi square, as described in "Information Retrieval", Manning et al.,
// Chapter 13.5.2, as the significance score.
func (s *_significantTextAggregation) ChiSquare(chisquare types.ChiSquareHeuristicVariant) *_significantTextAggregation {

	s.v.ChiSquare = chisquare.ChiSquareHeuristicCaster()

	return s
}

// Values to exclude.
func (s *_significantTextAggregation) Exclude(termsexcludes ...string) *_significantTextAggregation {

	s.v.Exclude = termsexcludes

	return s
}

// Determines whether the aggregation will use field values directly or global
// ordinals.
func (s *_significantTextAggregation) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *_significantTextAggregation {

	s.v.ExecutionHint = &executionhint
	return s
}

// The field from which to return significant text.
func (s *_significantTextAggregation) Field(field string) *_significantTextAggregation {

	s.v.Field = &field

	return s
}

// Whether to out duplicate text to deal with noisy data.
func (s *_significantTextAggregation) FilterDuplicateText(filterduplicatetext bool) *_significantTextAggregation {

	s.v.FilterDuplicateText = &filterduplicatetext

	return s
}

// Use Google normalized distance as described in "The Google Similarity
// Distance", Cilibrasi and Vitanyi, 2007, as the significance score.
func (s *_significantTextAggregation) Gnd(gnd types.GoogleNormalizedDistanceHeuristicVariant) *_significantTextAggregation {

	s.v.Gnd = gnd.GoogleNormalizedDistanceHeuristicCaster()

	return s
}

// Values to include.
func (s *_significantTextAggregation) Include(termsinclude types.TermsIncludeVariant) *_significantTextAggregation {

	s.v.Include = *termsinclude.TermsIncludeCaster()

	return s
}

// Use JLH score as the significance score.
func (s *_significantTextAggregation) Jlh(jlh types.EmptyObjectVariant) *_significantTextAggregation {

	s.v.Jlh = jlh.EmptyObjectCaster()

	return s
}

// Only return values that are found in more than `min_doc_count` hits.
func (s *_significantTextAggregation) MinDocCount(mindoccount int64) *_significantTextAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

// Use mutual information as described in "Information Retrieval", Manning et
// al., Chapter 13.5.1, as the significance score.
func (s *_significantTextAggregation) MutualInformation(mutualinformation types.MutualInformationHeuristicVariant) *_significantTextAggregation {

	s.v.MutualInformation = mutualinformation.MutualInformationHeuristicCaster()

	return s
}

// A simple calculation of the number of documents in the foreground sample with
// a term divided by the number of documents in the background with the term.
func (s *_significantTextAggregation) Percentage(percentage types.PercentageScoreHeuristicVariant) *_significantTextAggregation {

	s.v.Percentage = percentage.PercentageScoreHeuristicCaster()

	return s
}

// Customized score, implemented via a script.
func (s *_significantTextAggregation) ScriptHeuristic(scriptheuristic types.ScriptedHeuristicVariant) *_significantTextAggregation {

	s.v.ScriptHeuristic = scriptheuristic.ScriptedHeuristicCaster()

	return s
}

// Regulates the certainty a shard has if the values should actually be added to
// the candidate list or not with respect to the min_doc_count.
// Values will only be considered if their local shard frequency within the set
// is higher than the `shard_min_doc_count`.
func (s *_significantTextAggregation) ShardMinDocCount(shardmindoccount int64) *_significantTextAggregation {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

// The number of candidate terms produced by each shard.
// By default, `shard_size` will be automatically estimated based on the number
// of shards and the `size` parameter.
func (s *_significantTextAggregation) ShardSize(shardsize int) *_significantTextAggregation {

	s.v.ShardSize = &shardsize

	return s
}

// The number of buckets returned out of the overall terms list.
func (s *_significantTextAggregation) Size(size int) *_significantTextAggregation {

	s.v.Size = &size

	return s
}

// Overrides the JSON `_source` fields from which text will be analyzed.
func (s *_significantTextAggregation) SourceFields(fields ...string) *_significantTextAggregation {

	s.v.SourceFields = fields

	return s
}

func (s *_significantTextAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.SignificantText = s.v

	return container
}

func (s *_significantTextAggregation) SignificantTextAggregationCaster() *types.SignificantTextAggregation {
	return s.v
}
