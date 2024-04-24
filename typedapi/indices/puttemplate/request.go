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
// https://github.com/elastic/elasticsearch-specification/tree/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757

package puttemplate

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package puttemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/5fb8f1ce9c4605abcaa44aa0f17dbfc60497a757/specification/indices/put_template/IndicesPutTemplateRequest.ts#L29-L105
type Request struct {

	// Aliases Aliases for the index.
	Aliases map[string]types.Alias `json:"aliases,omitempty"`
	// IndexPatterns Array of wildcard expressions used to match the names
	// of indices during creation.
	IndexPatterns []string `json:"index_patterns,omitempty"`
	// Mappings Mapping for fields in the index.
	Mappings *types.TypeMapping `json:"mappings,omitempty"`
	// Order Order in which Elasticsearch applies this template if index
	// matches multiple templates.
	//
	// Templates with lower 'order' values are merged first. Templates with higher
	// 'order' values are merged later, overriding templates with lower values.
	Order *int `json:"order,omitempty"`
	// Settings Configuration options for the index.
	Settings map[string]json.RawMessage `json:"settings,omitempty"`
	// Version Version number used to manage index templates externally. This number
	// is not automatically generated by Elasticsearch.
	Version *int64 `json:"version,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Aliases:  make(map[string]types.Alias, 0),
		Settings: make(map[string]json.RawMessage, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Puttemplate request: %w", err)
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

		case "aliases":
			if s.Aliases == nil {
				s.Aliases = make(map[string]types.Alias, 0)
			}
			if err := dec.Decode(&s.Aliases); err != nil {
				return fmt.Errorf("%s | %w", "Aliases", err)
			}

		case "index_patterns":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "IndexPatterns", err)
				}

				s.IndexPatterns = append(s.IndexPatterns, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.IndexPatterns); err != nil {
					return fmt.Errorf("%s | %w", "IndexPatterns", err)
				}
			}

		case "mappings":
			if err := dec.Decode(&s.Mappings); err != nil {
				return fmt.Errorf("%s | %w", "Mappings", err)
			}

		case "order":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Order", err)
				}
				s.Order = &value
			case float64:
				f := int(v)
				s.Order = &f
			}

		case "settings":
			if s.Settings == nil {
				s.Settings = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}
