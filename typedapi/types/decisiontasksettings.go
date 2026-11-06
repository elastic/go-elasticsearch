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

// The task settings of the `decision` task.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/inference/decision/types.ts#L31-L40
type DecisionTaskSettings struct {
	// Questions The questions to answer, keyed by question name. The answers in the response
	// use the same names. The questions replace the ones of the inference endpoint
	// as a whole. At least one question is required.
	Questions map[string]DecisionQuestion `json:"questions,omitempty"`
}

func (s *DecisionTaskSettings) UnmarshalJSON(data []byte) error {

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

		case "questions":
			if s.Questions == nil {
				s.Questions = make(map[string]DecisionQuestion, 0)
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
					oo := NewPredicateDecisionQuestion()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Questions | %w", err)
					}
					s.Questions[key] = oo
				case "choice":
					oo := NewChoiceDecisionQuestion()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Questions | %w", err)
					}
					s.Questions[key] = oo
				case "score":
					oo := NewScoreDecisionQuestion()
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("Questions | %w", err)
					}
					s.Questions[key] = oo
				default:
					oo := new(DecisionQuestion)
					if err := localDec.Decode(&oo); err != nil {
						return fmt.Errorf("new(DecisionQuestion) | %w", err)
					}
					s.Questions[key] = oo
				}
			}

		}
	}
	return nil
}

// NewDecisionTaskSettings returns a DecisionTaskSettings.
func NewDecisionTaskSettings() *DecisionTaskSettings {
	r := &DecisionTaskSettings{
		Questions: make(map[string]DecisionQuestion),
	}

	return r
}

type DecisionTaskSettingsVariant interface {
	DecisionTaskSettingsCaster() *DecisionTaskSettings
}

func (s *DecisionTaskSettings) DecisionTaskSettingsCaster() *DecisionTaskSettings {
	return s
}
