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

// CustomServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L904-L972
type CustomServiceSettings struct {
	// Headers Specifies the HTTPS header parameters – such as `Authentication` or
	// `Contet-Type` – that are required to access the custom service.
	// For example:
	// ```
	//
	//	"headers":{
	//	  "Authorization": "Bearer ${api_key}",
	//	  "Content-Type": "application/json;charset=utf-8"
	//	}
	//
	// ```
	Headers json.RawMessage `json:"headers,omitempty"`
	// InputType Specifies the input type translation values that are used to replace the
	// `${input_type}` template in the request body.
	// For example:
	// ```
	//
	//	"input_type": {
	//	  "translation": {
	//	    "ingest": "do_ingest",
	//	    "search": "do_search"
	//	  },
	//	  "default": "a_default"
	//	},
	//
	// ```
	// If the subsequent inference requests come from a search context, the `search`
	// key will be used and the template will be replaced with `do_search`.
	// If it comes from the ingest context `do_ingest` is used. If it's a different
	// context that is not specified, the default value will be used. If no default
	// is specified an empty string is used.
	// `translation` can be:
	// * `classification`
	// * `clustering`
	// * `ingest`
	// * `search`
	InputType json.RawMessage `json:"input_type,omitempty"`
	// QueryParameters Specifies the query parameters as a list of tuples. The arrays inside the
	// `query_parameters` must have two items, a key and a value.
	// For example:
	// ```
	// "query_parameters":[
	//
	//	["param_key", "some_value"],
	//	["param_key", "another_value"],
	//	["other_key", "other_value"]
	//
	// ]
	// ```
	// If the base url is `https://www.elastic.co` it results in:
	// `https://www.elastic.co?param_key=some_value&param_key=another_value&other_key=other_value`.
	QueryParameters json.RawMessage `json:"query_parameters,omitempty"`
	// Request The request configuration object.
	Request CustomRequestParams `json:"request"`
	// Response The response configuration object.
	Response CustomResponseParams `json:"response"`
	// SecretParameters Specifies secret parameters, like `api_key` or `api_token`, that are required
	// to access the custom service.
	// For example:
	// ```
	//
	//	"secret_parameters":{
	//	  "api_key":"<api_key>"
	//	}
	//
	// ```
	SecretParameters json.RawMessage `json:"secret_parameters,omitempty"`
	// Url The URL endpoint to use for the requests.
	Url *string `json:"url,omitempty"`
}

func (s *CustomServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "headers":
			if err := dec.Decode(&s.Headers); err != nil {
				return fmt.Errorf("%s | %w", "Headers", err)
			}

		case "input_type":
			if err := dec.Decode(&s.InputType); err != nil {
				return fmt.Errorf("%s | %w", "InputType", err)
			}

		case "query_parameters":
			if err := dec.Decode(&s.QueryParameters); err != nil {
				return fmt.Errorf("%s | %w", "QueryParameters", err)
			}

		case "request":
			if err := dec.Decode(&s.Request); err != nil {
				return fmt.Errorf("%s | %w", "Request", err)
			}

		case "response":
			if err := dec.Decode(&s.Response); err != nil {
				return fmt.Errorf("%s | %w", "Response", err)
			}

		case "secret_parameters":
			if err := dec.Decode(&s.SecretParameters); err != nil {
				return fmt.Errorf("%s | %w", "SecretParameters", err)
			}

		case "url":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Url", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Url = &o

		}
	}
	return nil
}

// NewCustomServiceSettings returns a CustomServiceSettings.
func NewCustomServiceSettings() *CustomServiceSettings {
	r := &CustomServiceSettings{}

	return r
}
