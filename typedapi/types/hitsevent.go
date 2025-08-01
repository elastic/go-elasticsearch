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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// HitsEvent type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/eql/_types/EqlHits.ts#L41-L54
type HitsEvent struct {
	Fields map[string][]json.RawMessage `json:"fields,omitempty"`
	// Id_ Unique identifier for the event. This ID is only unique within the index.
	Id_ string `json:"_id"`
	// Index_ Name of the index containing the event.
	Index_ string `json:"_index"`
	// Missing Set to `true` for events in a timespan-constrained sequence that do not meet
	// a given condition.
	Missing *bool `json:"missing,omitempty"`
	// Source_ Original JSON body passed for the event at index time.
	Source_ json.RawMessage `json:"_source,omitempty"`
}

func (s *HitsEvent) UnmarshalJSON(data []byte) error {

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

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string][]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
			}

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

		case "missing":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Missing", err)
				}
				s.Missing = &value
			case bool:
				s.Missing = &v
			}

		case "_source":
			if err := dec.Decode(&s.Source_); err != nil {
				return fmt.Errorf("%s | %w", "Source_", err)
			}

		}
	}
	return nil
}

// NewHitsEvent returns a HitsEvent.
func NewHitsEvent() *HitsEvent {
	r := &HitsEvent{
		Fields: make(map[string][]json.RawMessage),
	}

	return r
}
