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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// The answer to a `score` question.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/inference/_types/Results.ts#L218-L235
type ScoreDecisionAnswer struct {
	// Confidence The confidence in the score.
	Confidence Float64 `json:"confidence"`
	// Probabilities The probability of each level, ordered from the lowest to the highest level.
	Probabilities []ScoreDecisionProbability `json:"probabilities"`
	// Score The probability-weighted position on the scale, from `0` for the lowest
	// level.
	Score Float64 `json:"score"`
	Type  string  `json:"type,omitempty"`
}

func (s *ScoreDecisionAnswer) UnmarshalJSON(data []byte) error {

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

		case "confidence":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Confidence", err)
				}
				f := Float64(value)
				s.Confidence = f
			case float64:
				f := Float64(v)
				s.Confidence = f
			}

		case "probabilities":
			if err := dec.Decode(&s.Probabilities); err != nil {
				return fmt.Errorf("%s | %w", "Probabilities", err)
			}

		case "score":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Score", err)
				}
				f := Float64(value)
				s.Score = f
			case float64:
				f := Float64(v)
				s.Score = f
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s ScoreDecisionAnswer) MarshalJSON() ([]byte, error) {
	type innerScoreDecisionAnswer ScoreDecisionAnswer
	tmp := innerScoreDecisionAnswer{
		Confidence:    s.Confidence,
		Probabilities: s.Probabilities,
		Score:         s.Score,
		Type:          s.Type,
	}

	tmp.Type = "score"

	return json.Marshal(tmp)
}

// NewScoreDecisionAnswer returns a ScoreDecisionAnswer.
func NewScoreDecisionAnswer() *ScoreDecisionAnswer {
	r := &ScoreDecisionAnswer{}

	return r
}
