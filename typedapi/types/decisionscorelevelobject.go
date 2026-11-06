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

// DecisionScoreLevelObject type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/inference/decision/types.ts#L94-L103
type DecisionScoreLevelObject struct {
	// Description A description of the level.
	Description *string `json:"description,omitempty"`
	// Label The label of the level.
	Label string `json:"label"`
}

func (s *DecisionScoreLevelObject) UnmarshalJSON(data []byte) error {

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

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "label":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Label", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Label = o

		}
	}
	return nil
}

// NewDecisionScoreLevelObject returns a DecisionScoreLevelObject.
func NewDecisionScoreLevelObject() *DecisionScoreLevelObject {
	r := &DecisionScoreLevelObject{}

	return r
}

type DecisionScoreLevelObjectVariant interface {
	DecisionScoreLevelObjectCaster() *DecisionScoreLevelObject
}

func (s *DecisionScoreLevelObject) DecisionScoreLevelObjectCaster() *DecisionScoreLevelObject {
	return s
}

func (s *DecisionScoreLevelObject) DecisionScoreLevelCaster() *DecisionScoreLevel {
	if s == nil {
		return nil
	}
	o := DecisionScoreLevel(s)
	return &o
}
