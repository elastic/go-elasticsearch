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

package termsenum

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package termsenum
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/_global/terms_enum/TermsEnumRequest.ts#L26-L65
type Request struct {

	// CaseInsensitive When true the provided search string is matched against index terms without
	// case sensitivity.
	CaseInsensitive *bool `json:"case_insensitive,omitempty"`
	// Field The string to match at the start of indexed terms. If not provided, all terms
	// in the field are considered.
	Field string `json:"field"`
	// IndexFilter Allows to filter an index shard if the provided query rewrites to match_none.
	IndexFilter *types.Query `json:"index_filter,omitempty"`
	SearchAfter *string      `json:"search_after,omitempty"`
	// Size How many matching terms to return.
	Size *int `json:"size,omitempty"`
	// String The string after which terms in the index should be returned. Allows for a
	// form of pagination if the last result from one request is passed as the
	// search_after parameter for a subsequent request.
	String *string `json:"string,omitempty"`
	// Timeout The maximum length of time to spend collecting results. Defaults to "1s" (one
	// second). If the timeout is exceeded the complete flag set to false in the
	// response and the results may be partial or empty.
	Timeout types.Duration `json:"timeout,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Termsenum request: %w", err)
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

		case "case_insensitive":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.CaseInsensitive = &value
			case bool:
				s.CaseInsensitive = &v
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "index_filter":
			if err := dec.Decode(&s.IndexFilter); err != nil {
				return err
			}

		case "search_after":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SearchAfter = &o

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

		case "string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.String = &o

		case "timeout":
			if err := dec.Decode(&s.Timeout); err != nil {
				return err
			}

		}
	}
	return nil
}
