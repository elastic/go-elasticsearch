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

package queryrole

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package queryrole
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/query_role/QueryRolesRequest.ts#L25-L85
type Request struct {

	// From The starting document offset.
	// It must not be negative.
	// By default, you cannot page through more than 10,000 hits using the `from`
	// and `size` parameters.
	// To page through more hits, use the `search_after` parameter.
	From *int `json:"from,omitempty"`
	// Query A query to filter which roles to return.
	// If the query parameter is missing, it is equivalent to a `match_all` query.
	// The query supports a subset of query types, including `match_all`, `bool`,
	// `term`, `terms`, `match`,
	// `ids`, `prefix`, `wildcard`, `exists`, `range`, and `simple_query_string`.
	// You can query the following information associated with roles: `name`,
	// `description`, `metadata`,
	// `applications.application`, `applications.privileges`, and
	// `applications.resources`.
	Query *types.RoleQueryContainer `json:"query,omitempty"`
	// SearchAfter The search after definition.
	SearchAfter []types.FieldValue `json:"search_after,omitempty"`
	// Size The number of hits to return.
	// It must not be negative.
	// By default, you cannot page through more than 10,000 hits using the `from`
	// and `size` parameters.
	// To page through more hits, use the `search_after` parameter.
	Size *int `json:"size,omitempty"`
	// Sort The sort definition.
	// You can sort on `username`, `roles`, or `enabled`.
	// In addition, sort can also be applied to the `_doc` field to sort by index
	// order.
	Sort []types.SortCombinations `json:"sort,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Queryrole request: %w", err)
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

		case "from":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "From", err)
				}
				s.From = &value
			case float64:
				f := int(v)
				s.From = &f
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}

		case "search_after":
			if err := dec.Decode(&s.SearchAfter); err != nil {
				return fmt.Errorf("%s | %w", "SearchAfter", err)
			}

		case "size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Size", err)
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		case "sort":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(types.SortCombinations)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Sort", err)
				}

				s.Sort = append(s.Sort, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Sort); err != nil {
					return fmt.Errorf("%s | %w", "Sort", err)
				}
			}

		}
	}
	return nil
}
