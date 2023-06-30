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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// SearchTemplateRequestBody type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/watcher/_types/Input.ts#L128-L145
type SearchTemplateRequestBody struct {
	Explain *bool `json:"explain,omitempty"`
	// Id ID of the search template to use. If no source is specified,
	// this parameter is required.
	Id      *string                    `json:"id,omitempty"`
	Params  map[string]json.RawMessage `json:"params,omitempty"`
	Profile *bool                      `json:"profile,omitempty"`
	// Source An inline search template. Supports the same parameters as the search API's
	// request body. Also supports Mustache variables. If no id is specified, this
	// parameter is required.
	Source *string `json:"source,omitempty"`
}

func (s *SearchTemplateRequestBody) UnmarshalJSON(data []byte) error {

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

		case "explain":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Explain = &value
			case bool:
				s.Explain = &v
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "params":
			if s.Params == nil {
				s.Params = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Params); err != nil {
				return err
			}

		case "profile":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Profile = &value
			case bool:
				s.Profile = &v
			}

		case "source":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Source = &o

		}
	}
	return nil
}

// NewSearchTemplateRequestBody returns a SearchTemplateRequestBody.
func NewSearchTemplateRequestBody() *SearchTemplateRequestBody {
	r := &SearchTemplateRequestBody{
		Params: make(map[string]json.RawMessage, 0),
	}

	return r
}
