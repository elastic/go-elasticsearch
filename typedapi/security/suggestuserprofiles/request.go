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

package suggestuserprofiles

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package suggestuserprofiles
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/suggest_user_profiles/Request.ts#L24-L81
type Request struct {

	// Data A comma-separated list of filters for the `data` field of the profile
	// document.
	// To return all content use `data=*`.
	// To return a subset of content, use `data=<key>` to retrieve content nested
	// under the specified `<key>`.
	// By default, the API returns no `data` content.
	// It is an error to specify `data` as both the query parameter and the request
	// body field.
	Data []string `json:"data,omitempty"`
	// Hint Extra search criteria to improve relevance of the suggestion result.
	// Profiles matching the spcified hint are ranked higher in the response.
	// Profiles not matching the hint aren't excluded from the response as long as
	// the profile matches the `name` field query.
	Hint *types.Hint `json:"hint,omitempty"`
	// Name A query string used to match name-related fields in user profile documents.
	// Name-related fields are the user's `username`, `full_name`, and `email`.
	Name *string `json:"name,omitempty"`
	// Size The number of profiles to return.
	Size *int64 `json:"size,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Suggestuserprofiles request: %w", err)
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

		case "data":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Data", err)
				}

				s.Data = append(s.Data, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Data); err != nil {
					return fmt.Errorf("%s | %w", "Data", err)
				}
			}

		case "hint":
			if err := dec.Decode(&s.Hint); err != nil {
				return fmt.Errorf("%s | %w", "Hint", err)
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "size":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Size", err)
				}
				s.Size = &value
			case float64:
				f := int64(v)
				s.Size = &f
			}

		}
	}
	return nil
}
