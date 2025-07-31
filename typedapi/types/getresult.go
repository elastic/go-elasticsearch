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

// GetResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/get/types.ts#L25-L67
type GetResult struct {
	// Fields If the `stored_fields` parameter is set to `true` and `found` is `true`, it
	// contains the document fields stored in the index.
	Fields map[string]json.RawMessage `json:"fields,omitempty"`
	// Found Indicates whether the document exists.
	Found bool `json:"found"`
	// Id_ The unique identifier for the document.
	Id_      string   `json:"_id"`
	Ignored_ []string `json:"_ignored,omitempty"`
	// Index_ The name of the index the document belongs to.
	Index_ string `json:"_index"`
	// PrimaryTerm_ The primary term assigned to the document for the indexing operation.
	PrimaryTerm_ *int64 `json:"_primary_term,omitempty"`
	// Routing_ The explicit routing, if set.
	Routing_ *string `json:"_routing,omitempty"`
	// SeqNo_ The sequence number assigned to the document for the indexing operation.
	// Sequence numbers are used to ensure an older version of a document doesn't
	// overwrite a newer version.
	SeqNo_ *int64 `json:"_seq_no,omitempty"`
	// Source_ If `found` is `true`, it contains the document data formatted in JSON.
	// If the `_source` parameter is set to `false` or the `stored_fields` parameter
	// is set to `true`, it is excluded.
	Source_ json.RawMessage `json:"_source,omitempty"`
	// Version_ The document version, which is ncremented each time the document is updated.
	Version_ *int64 `json:"_version,omitempty"`
}

func (s *GetResult) UnmarshalJSON(data []byte) error {

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
				s.Fields = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
			}

		case "found":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Found", err)
				}
				s.Found = value
			case bool:
				s.Found = v
			}

		case "_id":
			if err := dec.Decode(&s.Id_); err != nil {
				return fmt.Errorf("%s | %w", "Id_", err)
			}

		case "_ignored":
			if err := dec.Decode(&s.Ignored_); err != nil {
				return fmt.Errorf("%s | %w", "Ignored_", err)
			}

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return fmt.Errorf("%s | %w", "Index_", err)
			}

		case "_primary_term":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryTerm_", err)
				}
				s.PrimaryTerm_ = &value
			case float64:
				f := int64(v)
				s.PrimaryTerm_ = &f
			}

		case "_routing":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Routing_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Routing_ = &o

		case "_seq_no":
			if err := dec.Decode(&s.SeqNo_); err != nil {
				return fmt.Errorf("%s | %w", "SeqNo_", err)
			}

		case "_source":
			if err := dec.Decode(&s.Source_); err != nil {
				return fmt.Errorf("%s | %w", "Source_", err)
			}

		case "_version":
			if err := dec.Decode(&s.Version_); err != nil {
				return fmt.Errorf("%s | %w", "Version_", err)
			}

		}
	}
	return nil
}

// NewGetResult returns a GetResult.
func NewGetResult() *GetResult {
	r := &GetResult{
		Fields: make(map[string]json.RawMessage),
	}

	return r
}
