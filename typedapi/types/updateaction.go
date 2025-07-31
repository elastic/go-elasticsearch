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
)

// UpdateAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/bulk/types.ts#L182-L217
type UpdateAction struct {
	// DetectNoop If true, the `result` in the response is set to 'noop' when no changes to the
	// document occur.
	DetectNoop *bool `json:"detect_noop,omitempty"`
	// Doc A partial update to an existing document.
	Doc json.RawMessage `json:"doc,omitempty"`
	// DocAsUpsert Set to `true` to use the contents of `doc` as the value of `upsert`.
	DocAsUpsert *bool `json:"doc_as_upsert,omitempty"`
	// Script The script to run to update the document.
	Script *Script `json:"script,omitempty"`
	// ScriptedUpsert Set to `true` to run the script whether or not the document exists.
	ScriptedUpsert *bool `json:"scripted_upsert,omitempty"`
	// Source_ If `false`, source retrieval is turned off.
	// You can also specify a comma-separated list of the fields you want to
	// retrieve.
	Source_ SourceConfig `json:"_source,omitempty"`
	// Upsert If the document does not already exist, the contents of `upsert` are inserted
	// as a new document.
	// If the document exists, the `script` is run.
	Upsert json.RawMessage `json:"upsert,omitempty"`
}

func (s *UpdateAction) UnmarshalJSON(data []byte) error {

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

		case "detect_noop":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DetectNoop", err)
				}
				s.DetectNoop = &value
			case bool:
				s.DetectNoop = &v
			}

		case "doc":
			if err := dec.Decode(&s.Doc); err != nil {
				return fmt.Errorf("%s | %w", "Doc", err)
			}

		case "doc_as_upsert":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DocAsUpsert", err)
				}
				s.DocAsUpsert = &value
			case bool:
				s.DocAsUpsert = &v
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		case "scripted_upsert":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ScriptedUpsert", err)
				}
				s.ScriptedUpsert = &value
			case bool:
				s.ScriptedUpsert = &v
			}

		case "_source":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Source_", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		source__field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Source_", err)
				}

				switch t {

				case "exclude_vectors", "excludes", "includes":
					o := NewSourceFilter()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Source_", err)
					}
					s.Source_ = o
					break source__field

				}
			}
			if s.Source_ == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Source_); err != nil {
					return fmt.Errorf("%s | %w", "Source_", err)
				}
			}

		case "upsert":
			if err := dec.Decode(&s.Upsert); err != nil {
				return fmt.Errorf("%s | %w", "Upsert", err)
			}

		}
	}
	return nil
}

// NewUpdateAction returns a UpdateAction.
func NewUpdateAction() *UpdateAction {
	r := &UpdateAction{}

	return r
}
