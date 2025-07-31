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

// Rescore type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/search/_types/rescoring.ts#L25-L38
type Rescore struct {
	AdditionalRescoreProperty map[string]json.RawMessage `json:"-"`
	LearningToRank            *LearningToRank            `json:"learning_to_rank,omitempty"`
	Query                     *RescoreQuery              `json:"query,omitempty"`
	WindowSize                *int                       `json:"window_size,omitempty"`
}

func (s *Rescore) UnmarshalJSON(data []byte) error {

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

		case "learning_to_rank":
			if err := dec.Decode(&s.LearningToRank); err != nil {
				return fmt.Errorf("%s | %w", "LearningToRank", err)
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}

		case "window_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "WindowSize", err)
				}
				s.WindowSize = &value
			case float64:
				f := int(v)
				s.WindowSize = &f
			}

		default:

			if key, ok := t.(string); ok {
				if s.AdditionalRescoreProperty == nil {
					s.AdditionalRescoreProperty = make(map[string]json.RawMessage, 0)
				}
				raw := new(json.RawMessage)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "AdditionalRescoreProperty", err)
				}
				s.AdditionalRescoreProperty[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s Rescore) MarshalJSON() ([]byte, error) {
	type opt Rescore
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalRescoreProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalRescoreProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewRescore returns a Rescore.
func NewRescore() *Rescore {
	r := &Rescore{
		AdditionalRescoreProperty: make(map[string]json.RawMessage),
	}

	return r
}
