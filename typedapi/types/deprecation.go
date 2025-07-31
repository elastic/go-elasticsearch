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
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deprecationlevel"
)

// Deprecation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/migration/deprecations/types.ts#L32-L47
type Deprecation struct {
	// Details Optional details about the deprecation warning.
	Details *string `json:"details,omitempty"`
	// Level The level property describes the significance of the issue.
	Level deprecationlevel.DeprecationLevel `json:"level"`
	// Message Descriptive information about the deprecation warning.
	Message                     string                     `json:"message"`
	Meta_                       map[string]json.RawMessage `json:"_meta,omitempty"`
	ResolveDuringRollingUpgrade bool                       `json:"resolve_during_rolling_upgrade"`
	// Url A link to the breaking change documentation, where you can find more
	// information about this change.
	Url string `json:"url"`
}

func (s *Deprecation) UnmarshalJSON(data []byte) error {

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

		case "details":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Details", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Details = &o

		case "level":
			if err := dec.Decode(&s.Level); err != nil {
				return fmt.Errorf("%s | %w", "Level", err)
			}

		case "message":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Message", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Message = o

		case "_meta":
			if s.Meta_ == nil {
				s.Meta_ = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Meta_); err != nil {
				return fmt.Errorf("%s | %w", "Meta_", err)
			}

		case "resolve_during_rolling_upgrade":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ResolveDuringRollingUpgrade", err)
				}
				s.ResolveDuringRollingUpgrade = value
			case bool:
				s.ResolveDuringRollingUpgrade = v
			}

		case "url":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Url", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Url = o

		}
	}
	return nil
}

// NewDeprecation returns a Deprecation.
func NewDeprecation() *Deprecation {
	r := &Deprecation{
		Meta_: make(map[string]json.RawMessage),
	}

	return r
}
