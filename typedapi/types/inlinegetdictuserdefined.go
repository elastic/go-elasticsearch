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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

import (
	"encoding/json"
	"fmt"
)

// InlineGetDictUserDefined type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_types/common.ts#L286-L295
type InlineGetDictUserDefined struct {
	Fields                   map[string]interface{} `json:"fields,omitempty"`
	Found                    bool                   `json:"found"`
	InlineGetDictUserDefined map[string]interface{} `json:"-"`
	PrimaryTerm_             *int64                 `json:"_primary_term,omitempty"`
	Routing_                 *string                `json:"_routing,omitempty"`
	SeqNo_                   *int64                 `json:"_seq_no,omitempty"`
	Source_                  map[string]interface{} `json:"_source"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s InlineGetDictUserDefined) MarshalJSON() ([]byte, error) {
	type opt InlineGetDictUserDefined
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
	for key, value := range s.InlineGetDictUserDefined {
		tmp[fmt.Sprintf("%s", key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewInlineGetDictUserDefined returns a InlineGetDictUserDefined.
func NewInlineGetDictUserDefined() *InlineGetDictUserDefined {
	r := &InlineGetDictUserDefined{
		Fields:                   make(map[string]interface{}, 0),
		InlineGetDictUserDefined: make(map[string]interface{}, 0),
		Source_:                  make(map[string]interface{}, 0),
	}

	return r
}
