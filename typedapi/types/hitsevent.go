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
)

// HitsEvent type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/eql/_types/EqlHits.ts#L41-L49
type HitsEvent struct {
	Fields map[string][]json.RawMessage `json:"fields,omitempty"`
	// Id_ Unique identifier for the event. This ID is only unique within the index.
	Id_ string `json:"_id"`
	// Index_ Name of the index containing the event.
	Index_ string `json:"_index"`
	// Source_ Original JSON body passed for the event at index time.
	Source_ json.RawMessage `json:"_source,omitempty"`
}

// NewHitsEvent returns a HitsEvent.
func NewHitsEvent() *HitsEvent {
	r := &HitsEvent{
		Fields: make(map[string][]json.RawMessage, 0),
	}

	return r
}
