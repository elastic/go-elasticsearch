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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"encoding/json"
	"fmt"
)

// SortOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/sort.ts#L82-L91
type SortOptions struct {
	Doc_         *ScoreSort           `json:"_doc,omitempty"`
	GeoDistance_ *GeoDistanceSort     `json:"_geo_distance,omitempty"`
	Score_       *ScoreSort           `json:"_score,omitempty"`
	Script_      *ScriptSort          `json:"_script,omitempty"`
	SortOptions  map[string]FieldSort `json:"SortOptions,omitempty"`
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
