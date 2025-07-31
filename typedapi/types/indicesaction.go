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
	"encoding/json"
	"fmt"
)

// IndicesAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/update_aliases/types.ts#L23-L39
type IndicesAction struct {
	// Add Adds a data stream or index to an alias.
	// If the alias doesn’t exist, the `add` action creates it.
	Add                             *AddAction                 `json:"add,omitempty"`
	AdditionalIndicesActionProperty map[string]json.RawMessage `json:"-"`
	// Remove Removes a data stream or index from an alias.
	Remove *RemoveAction `json:"remove,omitempty"`
	// RemoveIndex Deletes an index.
	// You cannot use this action on aliases or data streams.
	RemoveIndex *RemoveIndexAction `json:"remove_index,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s IndicesAction) MarshalJSON() ([]byte, error) {
	type opt IndicesAction
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
	for key, value := range s.AdditionalIndicesActionProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalIndicesActionProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewIndicesAction returns a IndicesAction.
func NewIndicesAction() *IndicesAction {
	r := &IndicesAction{
		AdditionalIndicesActionProperty: make(map[string]json.RawMessage),
	}

	return r
}
