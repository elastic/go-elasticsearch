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
)

// DecisionResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/inference/_types/Results.ts#L159-L164
type DecisionResult struct {
	// Answers The answers, keyed by the name of the question they answer.
	Answers map[string]DecisionAnswer `json:"answers"`
}

func (s *DecisionResult) UnmarshalJSON(data []byte) error {

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

		case "answers":
			if s.Answers == nil {
				s.Answers = make(map[string]DecisionAnswer, 0)
			}
			refs := make(map[string]json.RawMessage, 0)
			dec.Decode(&refs)
			for key, message := range refs {
				kind := make(map[string]any)
				buf := bytes.NewReader(message)
				localDec := json.NewDecoder(buf)
				localDec.Decode(&kind)
				buf.Seek(0, io.SeekStart)

				switch kind["type"] {
				case "predicate":
					oo := NewPredicateDecisionAnswer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Answers | %w", err)
					}
					s.Answers[key] = oo
				case "choice":
					oo := NewChoiceDecisionAnswer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Answers | %w", err)
					}
					s.Answers[key] = oo
				case "score":
					oo := NewScoreDecisionAnswer()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Answers | %w", err)
					}
					s.Answers[key] = oo
				default:
					oo := new(DecisionAnswer)
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("new(DecisionAnswer) | %w", err)
					}
					s.Answers[key] = oo
				}
			}

		}
	}
	return nil
}

// NewDecisionResult returns a DecisionResult.
func NewDecisionResult() *DecisionResult {
	r := &DecisionResult{
		Answers: make(map[string]DecisionAnswer),
	}

	return r
}
