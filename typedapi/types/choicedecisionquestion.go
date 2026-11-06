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

// A question answered with one option out of a fixed set of choices.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/inference/decision/types.ts#L66-L75
type ChoiceDecisionQuestion struct {
	// Choices The options to choose from, each mapped to an optional description. At least
	// two choices are required.
	Choices map[string]string `json:"choices"`
	// Instructions A description of what to decide. It must not be empty.
	Instructions string `json:"instructions"`
	Type         string `json:"type,omitempty"`
}

func (s *ChoiceDecisionQuestion) UnmarshalJSON(data []byte) error {

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

		case "choices":
			if s.Choices == nil {
				s.Choices = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Choices); err != nil {
				return fmt.Errorf("%s | %w", "Choices", err)
			}

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
func (s ChoiceDecisionQuestion) MarshalJSON() ([]byte, error) {
	type innerChoiceDecisionQuestion ChoiceDecisionQuestion
	tmp := innerChoiceDecisionQuestion{
		Choices:      s.Choices,
		Instructions: s.Instructions,
		Type:         s.Type,
	}

	tmp.Type = "choice"

	return json.Marshal(tmp)
}

// NewChoiceDecisionQuestion returns a ChoiceDecisionQuestion.
func NewChoiceDecisionQuestion() *ChoiceDecisionQuestion {
	r := &ChoiceDecisionQuestion{
		Choices: make(map[string]string),
	}

	return r
}

type ChoiceDecisionQuestionVariant interface {
	ChoiceDecisionQuestionCaster() *ChoiceDecisionQuestion
}

func (s *ChoiceDecisionQuestion) ChoiceDecisionQuestionCaster() *ChoiceDecisionQuestion {
	return s
}

func (s *ChoiceDecisionQuestion) DecisionQuestionCaster() *DecisionQuestion {
	if s == nil {
		return nil
	}
	o := DecisionQuestion(s)
	return &o
}
