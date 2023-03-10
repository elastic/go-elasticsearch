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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"encoding/json"
	"fmt"
)

// InlineGet type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/common.ts#L286-L295
type InlineGet struct {
	Fields       map[string]json.RawMessage `json:"fields,omitempty"`
	Found        bool                       `json:"found"`
	Metadata     map[string]json.RawMessage `json:"-"`
	PrimaryTerm_ *int64                     `json:"_primary_term,omitempty"`
	Routing_     *string                    `json:"_routing,omitempty"`
	SeqNo_       *int64                     `json:"_seq_no,omitempty"`
	Source_      json.RawMessage            `json:"_source,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s InlineGet) MarshalJSON() ([]byte, error) {
	type opt InlineGet
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
	for key, value := range s.Metadata {
		tmp[fmt.Sprintf("%s", key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewInlineGet returns a InlineGet.
func NewInlineGet() *InlineGet {
	r := &InlineGet{
		Fields:   make(map[string]json.RawMessage, 0),
		Metadata: make(map[string]json.RawMessage, 0),
	}

	return r
}
