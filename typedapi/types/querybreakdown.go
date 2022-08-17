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

// QueryBreakdown type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/profile.ts#L95-L114
type QueryBreakdown struct {
	Advance                     int64 `json:"advance"`
	AdvanceCount                int64 `json:"advance_count"`
	BuildScorer                 int64 `json:"build_scorer"`
	BuildScorerCount            int64 `json:"build_scorer_count"`
	ComputeMaxScore             int64 `json:"compute_max_score"`
	ComputeMaxScoreCount        int64 `json:"compute_max_score_count"`
	CreateWeight                int64 `json:"create_weight"`
	CreateWeightCount           int64 `json:"create_weight_count"`
	Match                       int64 `json:"match"`
	MatchCount                  int64 `json:"match_count"`
	NextDoc                     int64 `json:"next_doc"`
	NextDocCount                int64 `json:"next_doc_count"`
	Score                       int64 `json:"score"`
	ScoreCount                  int64 `json:"score_count"`
	SetMinCompetitiveScore      int64 `json:"set_min_competitive_score"`
	SetMinCompetitiveScoreCount int64 `json:"set_min_competitive_score_count"`
	ShallowAdvance              int64 `json:"shallow_advance"`
	ShallowAdvanceCount         int64 `json:"shallow_advance_count"`
}

// QueryBreakdownBuilder holds QueryBreakdown struct and provides a builder API.
type QueryBreakdownBuilder struct {
	v *QueryBreakdown
}

// NewQueryBreakdown provides a builder for the QueryBreakdown struct.
func NewQueryBreakdownBuilder() *QueryBreakdownBuilder {
	r := QueryBreakdownBuilder{
		&QueryBreakdown{},
	}

	return &r
}

// Build finalize the chain and returns the QueryBreakdown struct
func (rb *QueryBreakdownBuilder) Build() QueryBreakdown {
	return *rb.v
}

func (rb *QueryBreakdownBuilder) Advance(advance int64) *QueryBreakdownBuilder {
	rb.v.Advance = advance
	return rb
}

func (rb *QueryBreakdownBuilder) AdvanceCount(advancecount int64) *QueryBreakdownBuilder {
	rb.v.AdvanceCount = advancecount
	return rb
}

func (rb *QueryBreakdownBuilder) BuildScorer(buildscorer int64) *QueryBreakdownBuilder {
	rb.v.BuildScorer = buildscorer
	return rb
}

func (rb *QueryBreakdownBuilder) BuildScorerCount(buildscorercount int64) *QueryBreakdownBuilder {
	rb.v.BuildScorerCount = buildscorercount
	return rb
}

func (rb *QueryBreakdownBuilder) ComputeMaxScore(computemaxscore int64) *QueryBreakdownBuilder {
	rb.v.ComputeMaxScore = computemaxscore
	return rb
}

func (rb *QueryBreakdownBuilder) ComputeMaxScoreCount(computemaxscorecount int64) *QueryBreakdownBuilder {
	rb.v.ComputeMaxScoreCount = computemaxscorecount
	return rb
}

func (rb *QueryBreakdownBuilder) CreateWeight(createweight int64) *QueryBreakdownBuilder {
	rb.v.CreateWeight = createweight
	return rb
}

func (rb *QueryBreakdownBuilder) CreateWeightCount(createweightcount int64) *QueryBreakdownBuilder {
	rb.v.CreateWeightCount = createweightcount
	return rb
}

func (rb *QueryBreakdownBuilder) Match(match int64) *QueryBreakdownBuilder {
	rb.v.Match = match
	return rb
}

func (rb *QueryBreakdownBuilder) MatchCount(matchcount int64) *QueryBreakdownBuilder {
	rb.v.MatchCount = matchcount
	return rb
}

func (rb *QueryBreakdownBuilder) NextDoc(nextdoc int64) *QueryBreakdownBuilder {
	rb.v.NextDoc = nextdoc
	return rb
}

func (rb *QueryBreakdownBuilder) NextDocCount(nextdoccount int64) *QueryBreakdownBuilder {
	rb.v.NextDocCount = nextdoccount
	return rb
}

func (rb *QueryBreakdownBuilder) Score(score int64) *QueryBreakdownBuilder {
	rb.v.Score = score
	return rb
}

func (rb *QueryBreakdownBuilder) ScoreCount(scorecount int64) *QueryBreakdownBuilder {
	rb.v.ScoreCount = scorecount
	return rb
}

func (rb *QueryBreakdownBuilder) SetMinCompetitiveScore(setmincompetitivescore int64) *QueryBreakdownBuilder {
	rb.v.SetMinCompetitiveScore = setmincompetitivescore
	return rb
}

func (rb *QueryBreakdownBuilder) SetMinCompetitiveScoreCount(setmincompetitivescorecount int64) *QueryBreakdownBuilder {
	rb.v.SetMinCompetitiveScoreCount = setmincompetitivescorecount
	return rb
}

func (rb *QueryBreakdownBuilder) ShallowAdvance(shallowadvance int64) *QueryBreakdownBuilder {
	rb.v.ShallowAdvance = shallowadvance
	return rb
}

func (rb *QueryBreakdownBuilder) ShallowAdvanceCount(shallowadvancecount int64) *QueryBreakdownBuilder {
	rb.v.ShallowAdvanceCount = shallowadvancecount
	return rb
}
