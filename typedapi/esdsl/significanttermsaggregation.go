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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"
)

type _significantTermsAggregation struct {
	v *types.SignificantTermsAggregation
}

// Returns interesting or unusual occurrences of terms in a set.
func NewSignificantTermsAggregation() *_significantTermsAggregation {

	return &_significantTermsAggregation{v: types.NewSignificantTermsAggregation()}

}

// A background filter that can be used to focus in on significant terms within
// a narrower context, instead of the entire index.
func (s *_significantTermsAggregation) BackgroundFilter(backgroundfilter types.QueryVariant) *_significantTermsAggregation {

	s.v.BackgroundFilter = backgroundfilter.QueryCaster()

	return s
}

// Use Chi square, as described in "Information Retrieval", Manning et al.,
// Chapter 13.5.2, as the significance score.
func (s *_significantTermsAggregation) ChiSquare(chisquare types.ChiSquareHeuristicVariant) *_significantTermsAggregation {

	s.v.ChiSquare = chisquare.ChiSquareHeuristicCaster()

	return s
}

// Terms to exclude.
func (s *_significantTermsAggregation) Exclude(termsexcludes ...string) *_significantTermsAggregation {

	s.v.Exclude = termsexcludes

	return s
}

// Mechanism by which the aggregation should be executed: using field values
// directly or using global ordinals.
func (s *_significantTermsAggregation) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *_significantTermsAggregation {

	s.v.ExecutionHint = &executionhint
	return s
}

// The field from which to return significant terms.
func (s *_significantTermsAggregation) Field(field string) *_significantTermsAggregation {

	s.v.Field = &field

	return s
}

// Use Google normalized distance as described in "The Google Similarity
// Distance", Cilibrasi and Vitanyi, 2007, as the significance score.
func (s *_significantTermsAggregation) Gnd(gnd types.GoogleNormalizedDistanceHeuristicVariant) *_significantTermsAggregation {

	s.v.Gnd = gnd.GoogleNormalizedDistanceHeuristicCaster()

	return s
}

// Terms to include.
func (s *_significantTermsAggregation) Include(termsinclude types.TermsIncludeVariant) *_significantTermsAggregation {

	s.v.Include = *termsinclude.TermsIncludeCaster()

	return s
}

// Use JLH score as the significance score.
func (s *_significantTermsAggregation) Jlh(jlh types.EmptyObjectVariant) *_significantTermsAggregation {

	s.v.Jlh = jlh.EmptyObjectCaster()

	return s
}

// Only return terms that are found in more than `min_doc_count` hits.
func (s *_significantTermsAggregation) MinDocCount(mindoccount int64) *_significantTermsAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

// Use mutual information as described in "Information Retrieval", Manning et
// al., Chapter 13.5.1, as the significance score.
func (s *_significantTermsAggregation) MutualInformation(mutualinformation types.MutualInformationHeuristicVariant) *_significantTermsAggregation {

	s.v.MutualInformation = mutualinformation.MutualInformationHeuristicCaster()

	return s
}

// A simple calculation of the number of documents in the foreground sample with
// a term divided by the number of documents in the background with the term.
func (s *_significantTermsAggregation) Percentage(percentage types.PercentageScoreHeuristicVariant) *_significantTermsAggregation {

	s.v.Percentage = percentage.PercentageScoreHeuristicCaster()

	return s
}

// Customized score, implemented via a script.
func (s *_significantTermsAggregation) ScriptHeuristic(scriptheuristic types.ScriptedHeuristicVariant) *_significantTermsAggregation {

	s.v.ScriptHeuristic = scriptheuristic.ScriptedHeuristicCaster()

	return s
}

// Regulates the certainty a shard has if the term should actually be added to
// the candidate list or not with respect to the `min_doc_count`.
// Terms will only be considered if their local shard frequency within the set
// is higher than the `shard_min_doc_count`.
func (s *_significantTermsAggregation) ShardMinDocCount(shardmindoccount int64) *_significantTermsAggregation {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

// Can be used to control the volumes of candidate terms produced by each shard.
// By default, `shard_size` will be automatically estimated based on the number
// of shards and the `size` parameter.
func (s *_significantTermsAggregation) ShardSize(shardsize int) *_significantTermsAggregation {

	s.v.ShardSize = &shardsize

	return s
}

// The number of buckets returned out of the overall terms list.
func (s *_significantTermsAggregation) Size(size int) *_significantTermsAggregation {

	s.v.Size = &size

	return s
}

func (s *_significantTermsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.SignificantTerms = s.v

	return container
}

func (s *_significantTermsAggregation) SignificantTermsAggregationCaster() *types.SignificantTermsAggregation {
	return s.v
}
