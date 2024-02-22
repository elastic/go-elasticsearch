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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package queryapikeys

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package queryapikeys
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/security/query_api_keys/QueryApiKeysRequest.ts#L25-L74
type Request struct {

	// From Starting document offset.
	// By default, you cannot page through more than 10,000 hits using the from and
	// size parameters.
	// To page through more hits, use the `search_after` parameter.
	From *int `json:"from,omitempty"`
	// Query A query to filter which API keys to return.
	// The query supports a subset of query types, including `match_all`, `bool`,
	// `term`, `terms`, `ids`, `prefix`, `wildcard`, and `range`.
	// You can query all public information associated with an API key.
	Query *types.Query `json:"query,omitempty"`
	// SearchAfter Search after definition
	SearchAfter []types.FieldValue `json:"search_after,omitempty"`
	// Size The number of hits to return.
	// By default, you cannot page through more than 10,000 hits using the `from`
	// and `size` parameters.
	// To page through more hits, use the `search_after` parameter.
	Size *int `json:"size,omitempty"`
	// Sort Other than `id`, all public fields of an API key are eligible for sorting.
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
		return nil, fmt.Errorf("could not deserialise json into Queryapikeys request: %w", err)
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

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.From = &value
			case float64:
				f := int(v)
				s.From = &f
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return err
			}

		case "search_after":
			if err := dec.Decode(&s.SearchAfter); err != nil {
				return err
			}

		case "size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
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
					return err
				}

				s.Sort = append(s.Sort, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Sort); err != nil {
					return err
				}
			}

		}
	}
	return nil
}
