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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// QueryBreakdown type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/search/_types/profile.ts#L105-L126
type QueryBreakdown struct {
	Advance                     int64 `json:"advance"`
	AdvanceCount                int64 `json:"advance_count"`
	BuildScorer                 int64 `json:"build_scorer"`
	BuildScorerCount            int64 `json:"build_scorer_count"`
	ComputeMaxScore             int64 `json:"compute_max_score"`
	ComputeMaxScoreCount        int64 `json:"compute_max_score_count"`
	CountWeight                 int64 `json:"count_weight"`
	CountWeightCount            int64 `json:"count_weight_count"`
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

func (s *QueryBreakdown) UnmarshalJSON(data []byte) error {

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

		case "advance":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Advance", err)
				}
				s.Advance = value
			case float64:
				f := int64(v)
				s.Advance = f
			}

		case "advance_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "AdvanceCount", err)
				}
				s.AdvanceCount = value
			case float64:
				f := int64(v)
				s.AdvanceCount = f
			}

		case "build_scorer":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "BuildScorer", err)
				}
				s.BuildScorer = value
			case float64:
				f := int64(v)
				s.BuildScorer = f
			}

		case "build_scorer_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "BuildScorerCount", err)
				}
				s.BuildScorerCount = value
			case float64:
				f := int64(v)
				s.BuildScorerCount = f
			}

		case "compute_max_score":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ComputeMaxScore", err)
				}
				s.ComputeMaxScore = value
			case float64:
				f := int64(v)
				s.ComputeMaxScore = f
			}

		case "compute_max_score_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ComputeMaxScoreCount", err)
				}
				s.ComputeMaxScoreCount = value
			case float64:
				f := int64(v)
				s.ComputeMaxScoreCount = f
			}

		case "count_weight":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CountWeight", err)
				}
				s.CountWeight = value
			case float64:
				f := int64(v)
				s.CountWeight = f
			}

		case "count_weight_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CountWeightCount", err)
				}
				s.CountWeightCount = value
			case float64:
				f := int64(v)
				s.CountWeightCount = f
			}

		case "create_weight":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CreateWeight", err)
				}
				s.CreateWeight = value
			case float64:
				f := int64(v)
				s.CreateWeight = f
			}

		case "create_weight_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CreateWeightCount", err)
				}
				s.CreateWeightCount = value
			case float64:
				f := int64(v)
				s.CreateWeightCount = f
			}

		case "match":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Match", err)
				}
				s.Match = value
			case float64:
				f := int64(v)
				s.Match = f
			}

		case "match_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MatchCount", err)
				}
				s.MatchCount = value
			case float64:
				f := int64(v)
				s.MatchCount = f
			}

		case "next_doc":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NextDoc", err)
				}
				s.NextDoc = value
			case float64:
				f := int64(v)
				s.NextDoc = f
			}

		case "next_doc_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NextDocCount", err)
				}
				s.NextDocCount = value
			case float64:
				f := int64(v)
				s.NextDocCount = f
			}

		case "score":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Score", err)
				}
				s.Score = value
			case float64:
				f := int64(v)
				s.Score = f
			}

		case "score_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ScoreCount", err)
				}
				s.ScoreCount = value
			case float64:
				f := int64(v)
				s.ScoreCount = f
			}

		case "set_min_competitive_score":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SetMinCompetitiveScore", err)
				}
				s.SetMinCompetitiveScore = value
			case float64:
				f := int64(v)
				s.SetMinCompetitiveScore = f
			}

		case "set_min_competitive_score_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SetMinCompetitiveScoreCount", err)
				}
				s.SetMinCompetitiveScoreCount = value
			case float64:
				f := int64(v)
				s.SetMinCompetitiveScoreCount = f
			}

		case "shallow_advance":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShallowAdvance", err)
				}
				s.ShallowAdvance = value
			case float64:
				f := int64(v)
				s.ShallowAdvance = f
			}

		case "shallow_advance_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShallowAdvanceCount", err)
				}
				s.ShallowAdvanceCount = value
			case float64:
				f := int64(v)
				s.ShallowAdvanceCount = f
			}

		}
	}
	return nil
}

// NewQueryBreakdown returns a QueryBreakdown.
func NewQueryBreakdown() *QueryBreakdown {
	r := &QueryBreakdown{}

	return r
}
