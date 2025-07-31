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

package fieldcaps

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package fieldcaps
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/field_caps/FieldCapabilitiesRequest.ts#L25-L130
type Request struct {

	// Fields A list of fields to retrieve capabilities for. Wildcard (`*`) expressions are
	// supported.
	Fields []string `json:"fields,omitempty"`
	// IndexFilter Filter indices if the provided query rewrites to `match_none` on every shard.
	//
	// IMPORTANT: The filtering is done on a best-effort basis, it uses index
	// statistics and mappings to rewrite queries to `match_none` instead of fully
	// running the request.
	// For instance a range query over a date field can rewrite to `match_none` if
	// all documents within a shard (including deleted documents) are outside of the
	// provided range.
	// However, not all queries can rewrite to `match_none` so this API may return
	// an index even if the provided filter matches no document.
	IndexFilter *types.Query `json:"index_filter,omitempty"`
	// RuntimeMappings Define ad-hoc runtime fields in the request similar to the way it is done in
	// search requests.
	// These fields exist only as part of the query and take precedence over fields
	// defined with the same name in the index mappings.
	RuntimeMappings types.RuntimeFields `json:"runtime_mappings,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Fieldcaps request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Fields", err)
				}

				s.Fields = append(s.Fields, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Fields); err != nil {
					return fmt.Errorf("%s | %w", "Fields", err)
				}
			}

		case "index_filter":
			if err := dec.Decode(&s.IndexFilter); err != nil {
				return fmt.Errorf("%s | %w", "IndexFilter", err)
			}

		case "runtime_mappings":
			if err := dec.Decode(&s.RuntimeMappings); err != nil {
				return fmt.Errorf("%s | %w", "RuntimeMappings", err)
			}

		}
	}
	return nil
}
