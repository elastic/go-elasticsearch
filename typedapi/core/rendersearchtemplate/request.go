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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package rendersearchtemplate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Request holds the request body struct for the package rendersearchtemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/render_search_template/RenderSearchTemplateRequest.ts#L26-L77
type Request struct {
	File *string `json:"file,omitempty"`
	// Id The ID of the search template to render.
	// If no `source` is specified, this or the `<template-id>` request path
	// parameter is required.
	// If you specify both this parameter and the `<template-id>` parameter, the API
	// uses only `<template-id>`.
	Id *string `json:"id,omitempty"`
	// Params Key-value pairs used to replace Mustache variables in the template.
	// The key is the variable name.
	// The value is the variable value.
	Params map[string]json.RawMessage `json:"params,omitempty"`
	// Source An inline search template.
	// It supports the same parameters as the search API's request body.
	// These parameters also support Mustache variables.
	// If no `id` or `<templated-id>` is specified, this parameter is required.
	Source types.ScriptSource `json:"source,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Params: make(map[string]json.RawMessage, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Rendersearchtemplate request: %w", err)
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

		case "file":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "File", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.File = &o

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "params":
			if s.Params == nil {
				s.Params = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Params); err != nil {
				return fmt.Errorf("%s | %w", "Params", err)
			}

		case "source":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Source", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		source_field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Source", err)
				}

				switch t {

				case "aggregations", "collapse", "docvalue_fields", "explain", "ext", "fields", "from", "highlight", "indices_boost", "knn", "min_score", "pit", "post_filter", "profile", "query", "rank", "rescore", "retriever", "runtime_mappings", "script_fields", "search_after", "seq_no_primary_term", "size", "slice", "sort", "_source", "stats", "stored_fields", "suggest", "terminate_after", "timeout", "track_scores", "track_total_hits", "version":
					o := types.NewSearchRequestBody()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Source", err)
					}
					s.Source = o
					break source_field

				}
			}
			if s.Source == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Source); err != nil {
					return fmt.Errorf("%s | %w", "Source", err)
				}
			}

		}
	}
	return nil
}
