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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// SortOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b1811e10a0722431d79d1c234dd412ff47d8656f/specification/_types/sort.ts#L86-L96
type SortOptions struct {
	Doc_         *ScoreSort           `json:"_doc,omitempty"`
	GeoDistance_ *GeoDistanceSort     `json:"_geo_distance,omitempty"`
	Score_       *ScoreSort           `json:"_score,omitempty"`
	Script_      *ScriptSort          `json:"_script,omitempty"`
	SortOptions  map[string]FieldSort `json:"-"`
}

func (s *SortOptions) UnmarshalJSON(data []byte) error {

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

		case "_doc":
			if err := dec.Decode(&s.Doc_); err != nil {
				return fmt.Errorf("%s | %w", "Doc_", err)
			}

		case "_geo_distance":
			if err := dec.Decode(&s.GeoDistance_); err != nil {
				return fmt.Errorf("%s | %w", "GeoDistance_", err)
			}

		case "_score":
			if err := dec.Decode(&s.Score_); err != nil {
				return fmt.Errorf("%s | %w", "Score_", err)
			}

		case "_script":
			if err := dec.Decode(&s.Script_); err != nil {
				return fmt.Errorf("%s | %w", "Script_", err)
			}

		default:

			if key, ok := t.(string); ok {
				if s.SortOptions == nil {
					s.SortOptions = make(map[string]FieldSort, 0)
				}
				raw := NewFieldSort()
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "SortOptions", err)
				}
				if raw != nil {
					s.SortOptions[key] = *raw
				}
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s SortOptions) MarshalJSON() ([]byte, error) {
	type opt SortOptions
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]json.RawMessage, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.SortOptions {
		marshaled, err := json.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal additional property %q: %w", key, err)
		}
		tmp[fmt.Sprintf("%s", key)] = marshaled
	}
	delete(tmp, "SortOptions")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewSortOptions returns a SortOptions.
func NewSortOptions() *SortOptions {
	r := &SortOptions{
		SortOptions: make(map[string]FieldSort),
	}

	return r
}

type SortOptionsVariant interface {
	SortOptionsCaster() *SortOptions
}

func (s *SortOptions) SortOptionsCaster() *SortOptions {
	return s
}

func (s *SortOptions) SortCombinationsCaster() *SortCombinations {
	o := SortCombinations(s)
	return &o
}
