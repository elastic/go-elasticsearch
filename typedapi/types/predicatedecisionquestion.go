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

// A question answered with the probability that the statement in `instructions`
// is true.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/inference/decision/types.ts#L59-L64
type PredicateDecisionQuestion struct {
	// Instructions A description of what to decide. It must not be empty.
	Instructions string `json:"instructions"`
	Type         string `json:"type,omitempty"`
}

func (s *PredicateDecisionQuestion) UnmarshalJSON(data []byte) error {

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

		case "instructions":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Instructions", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Instructions = o

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s PredicateDecisionQuestion) MarshalJSON() ([]byte, error) {
	type innerPredicateDecisionQuestion PredicateDecisionQuestion
	tmp := innerPredicateDecisionQuestion{
		Instructions: s.Instructions,
		Type:         s.Type,
	}

	tmp.Type = "predicate"

	return json.Marshal(tmp)
}

// NewPredicateDecisionQuestion returns a PredicateDecisionQuestion.
func NewPredicateDecisionQuestion() *PredicateDecisionQuestion {
	r := &PredicateDecisionQuestion{}

	return r
}

type PredicateDecisionQuestionVariant interface {
	PredicateDecisionQuestionCaster() *PredicateDecisionQuestion
}

func (s *PredicateDecisionQuestion) PredicateDecisionQuestionCaster() *PredicateDecisionQuestion {
	return s
}

func (s *PredicateDecisionQuestion) DecisionQuestionCaster() *DecisionQuestion {
	if s == nil {
		return nil
	}
	o := DecisionQuestion(s)
	return &o
}
