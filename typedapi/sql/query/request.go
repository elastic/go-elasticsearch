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

package query

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package query
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/sql/query/QuerySqlRequest.ts#L27-L151
type Request struct {

	// AllowPartialSearchResults If `true`, the response has partial results when there are shard request
	// timeouts or shard failures.
	// If `false`, the API returns an error with no partial results.
	AllowPartialSearchResults *bool `json:"allow_partial_search_results,omitempty"`
	// Catalog The default catalog (cluster) for queries.
	// If unspecified, the queries execute on the data in the local cluster only.
	Catalog *string `json:"catalog,omitempty"`
	// Columnar If `true`, the results are in a columnar fashion: one row represents all the
	// values of a certain column from the current page of results.
	// The API supports this parameter only for CBOR, JSON, SMILE, and YAML
	// responses.
	Columnar *bool `json:"columnar,omitempty"`
	// Cursor The cursor used to retrieve a set of paginated results.
	// If you specify a cursor, the API only uses the `columnar` and `time_zone`
	// request body parameters.
	// It ignores other request body parameters.
	Cursor *string `json:"cursor,omitempty"`
	// FetchSize The maximum number of rows (or entries) to return in one response.
	FetchSize *int `json:"fetch_size,omitempty"`
	// FieldMultiValueLeniency If `false`, the API returns an exception when encountering multiple values
	// for a field.
	// If `true`, the API is lenient and returns the first value from the array with
	// no guarantee of consistent results.
	FieldMultiValueLeniency *bool `json:"field_multi_value_leniency,omitempty"`
	// Filter The Elasticsearch query DSL for additional filtering.
	Filter *types.Query `json:"filter,omitempty"`
	// IndexUsingFrozen If `true`, the search can run on frozen indices.
	IndexUsingFrozen *bool `json:"index_using_frozen,omitempty"`
	// KeepAlive The retention period for an async or saved synchronous search.
	KeepAlive types.Duration `json:"keep_alive,omitempty"`
	// KeepOnCompletion If `true`, Elasticsearch stores synchronous searches if you also specify the
	// `wait_for_completion_timeout` parameter.
	// If `false`, Elasticsearch only stores async searches that don't finish before
	// the `wait_for_completion_timeout`.
	KeepOnCompletion *bool `json:"keep_on_completion,omitempty"`
	// PageTimeout The minimum retention period for the scroll cursor.
	// After this time period, a pagination request might fail because the scroll
	// cursor is no longer available.
	// Subsequent scroll requests prolong the lifetime of the scroll cursor by the
	// duration of `page_timeout` in the scroll request.
	PageTimeout types.Duration `json:"page_timeout,omitempty"`
	// Params The values for parameters in the query.
	Params []json.RawMessage `json:"params,omitempty"`
	// Query The SQL query to run.
	Query *string `json:"query,omitempty"`
	// RequestTimeout The timeout before the request fails.
	RequestTimeout types.Duration `json:"request_timeout,omitempty"`
	// RuntimeMappings One or more runtime fields for the search request.
	// These fields take precedence over mapped fields with the same name.
	RuntimeMappings types.RuntimeFields `json:"runtime_mappings,omitempty"`
	// TimeZone The ISO-8601 time zone ID for the search.
	TimeZone *string `json:"time_zone,omitempty"`
	// WaitForCompletionTimeout The period to wait for complete results.
	// It defaults to no timeout, meaning the request waits for complete search
	// results.
	// If the search doesn't finish within this period, the search becomes async.
	//
	// To save a synchronous search, you must specify this parameter and the
	// `keep_on_completion` parameter.
	WaitForCompletionTimeout types.Duration `json:"wait_for_completion_timeout,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Query request: %w", err)
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

		case "allow_partial_search_results":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllowPartialSearchResults", err)
				}
				s.AllowPartialSearchResults = &value
			case bool:
				s.AllowPartialSearchResults = &v
			}

		case "catalog":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Catalog", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Catalog = &o

		case "columnar":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Columnar", err)
				}
				s.Columnar = &value
			case bool:
				s.Columnar = &v
			}

		case "cursor":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Cursor", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Cursor = &o

		case "fetch_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "FetchSize", err)
				}
				s.FetchSize = &value
			case float64:
				f := int(v)
				s.FetchSize = &f
			}

		case "field_multi_value_leniency":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "FieldMultiValueLeniency", err)
				}
				s.FieldMultiValueLeniency = &value
			case bool:
				s.FieldMultiValueLeniency = &v
			}

		case "filter":
			if err := dec.Decode(&s.Filter); err != nil {
				return fmt.Errorf("%s | %w", "Filter", err)
			}

		case "index_using_frozen":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IndexUsingFrozen", err)
				}
				s.IndexUsingFrozen = &value
			case bool:
				s.IndexUsingFrozen = &v
			}

		case "keep_alive":
			if err := dec.Decode(&s.KeepAlive); err != nil {
				return fmt.Errorf("%s | %w", "KeepAlive", err)
			}

		case "keep_on_completion":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "KeepOnCompletion", err)
				}
				s.KeepOnCompletion = &value
			case bool:
				s.KeepOnCompletion = &v
			}

		case "page_timeout":
			if err := dec.Decode(&s.PageTimeout); err != nil {
				return fmt.Errorf("%s | %w", "PageTimeout", err)
			}

		case "params":
			if err := dec.Decode(&s.Params); err != nil {
				return fmt.Errorf("%s | %w", "Params", err)
			}

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Query = &o

		case "request_timeout":
			if err := dec.Decode(&s.RequestTimeout); err != nil {
				return fmt.Errorf("%s | %w", "RequestTimeout", err)
			}

		case "runtime_mappings":
			if err := dec.Decode(&s.RuntimeMappings); err != nil {
				return fmt.Errorf("%s | %w", "RuntimeMappings", err)
			}

		case "time_zone":
			if err := dec.Decode(&s.TimeZone); err != nil {
				return fmt.Errorf("%s | %w", "TimeZone", err)
			}

		case "wait_for_completion_timeout":
			if err := dec.Decode(&s.WaitForCompletionTimeout); err != nil {
				return fmt.Errorf("%s | %w", "WaitForCompletionTimeout", err)
			}

		}
	}
	return nil
}
