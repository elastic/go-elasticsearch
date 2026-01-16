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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationexecutionhint"
)

type _significantTermsAggregation struct {
	v *types.SignificantTermsAggregation
}

// Returns interesting or unusual occurrences of terms in a set.
func NewSignificantTermsAggregation() *_significantTermsAggregation {

	return &_significantTermsAggregation{v: types.NewSignificantTermsAggregation()}

}

func (s *_significantTermsAggregation) BackgroundFilter(backgroundfilter types.QueryVariant) *_significantTermsAggregation {

	s.v.BackgroundFilter = backgroundfilter.QueryCaster()

	return s
}

func (s *_significantTermsAggregation) ChiSquare(chisquare types.ChiSquareHeuristicVariant) *_significantTermsAggregation {

	s.v.ChiSquare = chisquare.ChiSquareHeuristicCaster()

	return s
}

func (s *_significantTermsAggregation) Exclude(termsexcludes ...string) *_significantTermsAggregation {

	s.v.Exclude = termsexcludes

	return s
}

func (s *_significantTermsAggregation) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *_significantTermsAggregation {

	s.v.ExecutionHint = &executionhint
	return s
}

func (s *_significantTermsAggregation) Field(field string) *_significantTermsAggregation {

	s.v.Field = &field

	return s
}

func (s *_significantTermsAggregation) Gnd(gnd types.GoogleNormalizedDistanceHeuristicVariant) *_significantTermsAggregation {

	s.v.Gnd = gnd.GoogleNormalizedDistanceHeuristicCaster()

	return s
}

func (s *_significantTermsAggregation) Include(termsinclude types.TermsIncludeVariant) *_significantTermsAggregation {

	s.v.Include = *termsinclude.TermsIncludeCaster()

	return s
}

func (s *_significantTermsAggregation) Jlh(jlh types.EmptyObjectVariant) *_significantTermsAggregation {

	s.v.Jlh = jlh.EmptyObjectCaster()

	return s
}

func (s *_significantTermsAggregation) MinDocCount(mindoccount int64) *_significantTermsAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

func (s *_significantTermsAggregation) MutualInformation(mutualinformation types.MutualInformationHeuristicVariant) *_significantTermsAggregation {

	s.v.MutualInformation = mutualinformation.MutualInformationHeuristicCaster()

	return s
}

func (s *_significantTermsAggregation) PValue(pvalue types.PValueHeuristicVariant) *_significantTermsAggregation {

	s.v.PValue = pvalue.PValueHeuristicCaster()

	return s
}

func (s *_significantTermsAggregation) Percentage(percentage types.PercentageScoreHeuristicVariant) *_significantTermsAggregation {

	s.v.Percentage = percentage.PercentageScoreHeuristicCaster()

	return s
}

func (s *_significantTermsAggregation) ScriptHeuristic(scriptheuristic types.ScriptedHeuristicVariant) *_significantTermsAggregation {

	s.v.ScriptHeuristic = scriptheuristic.ScriptedHeuristicCaster()

	return s
}

func (s *_significantTermsAggregation) ShardMinDocCount(shardmindoccount int64) *_significantTermsAggregation {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

func (s *_significantTermsAggregation) ShardSize(shardsize int) *_significantTermsAggregation {

	s.v.ShardSize = &shardsize

	return s
}

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
