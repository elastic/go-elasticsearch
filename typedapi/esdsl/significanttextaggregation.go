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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationexecutionhint"
)

type _significantTextAggregation struct {
	v *types.SignificantTextAggregation
}

// Returns interesting or unusual occurrences of free-text terms in a set.
func NewSignificantTextAggregation() *_significantTextAggregation {

	return &_significantTextAggregation{v: types.NewSignificantTextAggregation()}

}

func (s *_significantTextAggregation) BackgroundFilter(backgroundfilter types.QueryVariant) *_significantTextAggregation {

	s.v.BackgroundFilter = backgroundfilter.QueryCaster()

	return s
}

func (s *_significantTextAggregation) ChiSquare(chisquare types.ChiSquareHeuristicVariant) *_significantTextAggregation {

	s.v.ChiSquare = chisquare.ChiSquareHeuristicCaster()

	return s
}

func (s *_significantTextAggregation) Exclude(termsexcludes ...string) *_significantTextAggregation {

	s.v.Exclude = termsexcludes

	return s
}

func (s *_significantTextAggregation) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *_significantTextAggregation {

	s.v.ExecutionHint = &executionhint
	return s
}

func (s *_significantTextAggregation) Field(field string) *_significantTextAggregation {

	s.v.Field = &field

	return s
}

func (s *_significantTextAggregation) FilterDuplicateText(filterduplicatetext bool) *_significantTextAggregation {

	s.v.FilterDuplicateText = &filterduplicatetext

	return s
}

func (s *_significantTextAggregation) Gnd(gnd types.GoogleNormalizedDistanceHeuristicVariant) *_significantTextAggregation {

	s.v.Gnd = gnd.GoogleNormalizedDistanceHeuristicCaster()

	return s
}

func (s *_significantTextAggregation) Include(termsinclude types.TermsIncludeVariant) *_significantTextAggregation {

	s.v.Include = *termsinclude.TermsIncludeCaster()

	return s
}

func (s *_significantTextAggregation) Jlh(jlh types.EmptyObjectVariant) *_significantTextAggregation {

	s.v.Jlh = jlh.EmptyObjectCaster()

	return s
}

func (s *_significantTextAggregation) MinDocCount(mindoccount int64) *_significantTextAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

func (s *_significantTextAggregation) MutualInformation(mutualinformation types.MutualInformationHeuristicVariant) *_significantTextAggregation {

	s.v.MutualInformation = mutualinformation.MutualInformationHeuristicCaster()

	return s
}

func (s *_significantTextAggregation) Percentage(percentage types.PercentageScoreHeuristicVariant) *_significantTextAggregation {

	s.v.Percentage = percentage.PercentageScoreHeuristicCaster()

	return s
}

func (s *_significantTextAggregation) ScriptHeuristic(scriptheuristic types.ScriptedHeuristicVariant) *_significantTextAggregation {

	s.v.ScriptHeuristic = scriptheuristic.ScriptedHeuristicCaster()

	return s
}

func (s *_significantTextAggregation) ShardMinDocCount(shardmindoccount int64) *_significantTextAggregation {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

func (s *_significantTextAggregation) ShardSize(shardsize int) *_significantTextAggregation {

	s.v.ShardSize = &shardsize

	return s
}

func (s *_significantTextAggregation) Size(size int) *_significantTextAggregation {

	s.v.Size = &size

	return s
}

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
