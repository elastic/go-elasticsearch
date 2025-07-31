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
)

// IndexState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/indices/_types/IndexState.ts#L27-L40
type IndexState struct {
	Aliases    map[string]Alias `json:"aliases,omitempty"`
	DataStream *string          `json:"data_stream,omitempty"`
	// Defaults Default settings, included when the request's `include_default` is `true`.
	Defaults *IndexSettings `json:"defaults,omitempty"`
	// Lifecycle Data stream lifecycle applicable if this is a data stream.
	Lifecycle *DataStreamLifecycle `json:"lifecycle,omitempty"`
	Mappings  *TypeMapping         `json:"mappings,omitempty"`
	Settings  *IndexSettings       `json:"settings,omitempty"`
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
				return fmt.Errorf("%s | %w", "Aliases", err)
			}

		case "data_stream":
			if err := dec.Decode(&s.DataStream); err != nil {
				return fmt.Errorf("%s | %w", "DataStream", err)
			}

		case "defaults":
			if err := dec.Decode(&s.Defaults); err != nil {
				return fmt.Errorf("%s | %w", "Defaults", err)
			}

		case "lifecycle":
			if err := dec.Decode(&s.Lifecycle); err != nil {
				return fmt.Errorf("%s | %w", "Lifecycle", err)
			}

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return fmt.Errorf("%s | %w", "Mappings", err)
			}

		case "settings":
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		}
	}
	return nil
}

// NewIndexState returns a IndexState.
func NewIndexState() *IndexState {
	r := &IndexState{
		Aliases: make(map[string]Alias),
	}

	return r
}
