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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// UpdateAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/bulk/types.ts#L169-L205
type UpdateAction struct {
	// DetectNoop Set to false to disable setting 'result' in the response
	// to 'noop' if no change to the document occurred.
	DetectNoop *bool `json:"detect_noop,omitempty"`
	// Doc A partial update to an existing document.
	Doc json.RawMessage `json:"doc,omitempty"`
	// DocAsUpsert Set to true to use the contents of 'doc' as the value of 'upsert'
	DocAsUpsert *bool `json:"doc_as_upsert,omitempty"`
	// Script Script to execute to update the document.
	Script Script `json:"script,omitempty"`
	// ScriptedUpsert Set to true to execute the script whether or not the document exists.
	ScriptedUpsert *bool `json:"scripted_upsert,omitempty"`
	// Source_ Set to false to disable source retrieval. You can also specify a
	// comma-separated
	// list of the fields you want to retrieve.
	Source_ SourceConfig `json:"_source,omitempty"`
	// Upsert If the document does not already exist, the contents of 'upsert' are inserted
	// as a
	// new document. If the document exists, the 'script' is executed.
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
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.DetectNoop = &value
			case bool:
				s.DetectNoop = &v
			}

		case "doc":
			if err := dec.Decode(&s.Doc); err != nil {
				return err
			}

		case "doc_as_upsert":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.DocAsUpsert = &value
			case bool:
				s.DocAsUpsert = &v
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		case "scripted_upsert":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.ScriptedUpsert = &value
			case bool:
				s.ScriptedUpsert = &v
			}

		case "_source":
			if err := dec.Decode(&s.Source_); err != nil {
				return err
			}

		case "upsert":
			if err := dec.Decode(&s.Upsert); err != nil {
				return err
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
