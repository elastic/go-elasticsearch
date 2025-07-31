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
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/query_api_keys/QueryApiKeysRequest.ts#L26-L124
type Request struct {

	// Aggregations Any aggregations to run over the corpus of returned API keys.
	// Aggregations and queries work together. Aggregations are computed only on the
	// API keys that match the query.
	// This supports only a subset of aggregation types, namely: `terms`, `range`,
	// `date_range`, `missing`,
	// `cardinality`, `value_count`, `composite`, `filter`, and `filters`.
	// Additionally, aggregations only run over the same subset of fields that query
	// works with.
	Aggregations map[string]types.ApiKeyAggregationContainer `json:"aggregations,omitempty"`
	// From The starting document offset.
	// It must not be negative.
	// By default, you cannot page through more than 10,000 hits using the `from`
	// and `size` parameters.
	// To page through more hits, use the `search_after` parameter.
	From *int `json:"from,omitempty"`
	// Query A query to filter which API keys to return.
	// If the query parameter is missing, it is equivalent to a `match_all` query.
	// The query supports a subset of query types, including `match_all`, `bool`,
	// `term`, `terms`, `match`,
	// `ids`, `prefix`, `wildcard`, `exists`, `range`, and `simple_query_string`.
	// You can query the following public information associated with an API key:
	// `id`, `type`, `name`,
	// `creation`, `expiration`, `invalidated`, `invalidation`, `username`, `realm`,
	// and `metadata`.
	//
	// NOTE: The queryable string values associated with API keys are internally
	// mapped as keywords.
	// Consequently, if no `analyzer` parameter is specified for a `match` query,
	// then the provided match query string is interpreted as a single keyword
	// value.
	// Such a match query is hence equivalent to a `term` query.
	Query *types.ApiKeyQueryContainer `json:"query,omitempty"`
	// SearchAfter The search after definition.
	SearchAfter []types.FieldValue `json:"search_after,omitempty"`
	// Size The number of hits to return.
	// It must not be negative.
	// The `size` parameter can be set to `0`, in which case no API key matches are
	// returned, only the aggregation results.
	// By default, you cannot page through more than 10,000 hits using the `from`
	// and `size` parameters.
	// To page through more hits, use the `search_after` parameter.
	Size *int `json:"size,omitempty"`
	// Sort The sort definition.
	// Other than `id`, all public fields of an API key are eligible for sorting.
	// In addition, sort can also be applied to the `_doc` field to sort by index
	// order.
	Sort []types.SortCombinations `json:"sort,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Aggregations: make(map[string]types.ApiKeyAggregationContainer, 0),
	}

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

		case "aggregations", "aggs":
			if s.Aggregations == nil {
				s.Aggregations = make(map[string]types.ApiKeyAggregationContainer, 0)
			}
			if err := dec.Decode(&s.Aggregations); err != nil {
				return fmt.Errorf("%s | %w", "Aggregations", err)
			}

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
