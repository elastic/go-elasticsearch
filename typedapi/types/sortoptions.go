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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"fmt"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// SortOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/sort.ts#L82-L91
type SortOptions struct {
	Doc_         *ScoreSort           `json:"_doc,omitempty"`
	GeoDistance_ *GeoDistanceSort     `json:"_geo_distance,omitempty"`
	Score_       *ScoreSort           `json:"_score,omitempty"`
	Script_      *ScriptSort          `json:"_script,omitempty"`
	SortOptions  map[string]FieldSort `json:"SortOptions,omitempty"`
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
				return err
			}

		case "_geo_distance":
			if err := dec.Decode(&s.GeoDistance_); err != nil {
				return err
			}

		case "_score":
			if err := dec.Decode(&s.Score_); err != nil {
				return err
			}

		case "_script":
			if err := dec.Decode(&s.Script_); err != nil {
				return err
			}

		case "SortOptions":
			if s.SortOptions == nil {
				s.SortOptions = make(map[string]FieldSort, 0)
			}
			if err := dec.Decode(&s.SortOptions); err != nil {
				return err
			}

		default:

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s SortOptions) MarshalJSON() ([]byte, error) {
	type opt SortOptions
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

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
		tmp[fmt.Sprintf("%s", key)] = value
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
		SortOptions: make(map[string]FieldSort, 0),
	}

	return r
}
