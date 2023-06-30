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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// IndexState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/indices/_types/IndexState.ts#L27-L40
type IndexState struct {
	Aliases    map[string]Alias `json:"aliases,omitempty"`
	DataStream *string          `json:"data_stream,omitempty"`
	// Defaults Default settings, included when the request's `include_default` is `true`.
	Defaults *IndexSettings `json:"defaults,omitempty"`
	// Lifecycle Data lifecycle applicable iff this is a data stream.
	Lifecycle *DataLifecycle `json:"lifecycle,omitempty"`
	Mappings  *TypeMapping   `json:"mappings,omitempty"`
	Settings  *IndexSettings `json:"settings,omitempty"`
}

func (s *IndexState) UnmarshalJSON(data []byte) error {

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

		case "aliases":
			if s.Aliases == nil {
				s.Aliases = make(map[string]Alias, 0)
			}
			if err := dec.Decode(&s.Aliases); err != nil {
				return err
			}

		case "data_stream":
			if err := dec.Decode(&s.DataStream); err != nil {
				return err
			}

		case "defaults":
			if err := dec.Decode(&s.Defaults); err != nil {
				return err
			}

		case "lifecycle":
			if err := dec.Decode(&s.Lifecycle); err != nil {
				return err
			}

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return err
			}

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIndexState returns a IndexState.
func NewIndexState() *IndexState {
	r := &IndexState{
		Aliases: make(map[string]Alias, 0),
	}

	return r
}
