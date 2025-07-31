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

package knnsearch

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package knnsearch
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/knn_search/KnnSearchRequest.ts#L26-L112
type Request struct {

	// DocvalueFields The request returns doc values for field names matching these patterns
	// in the `hits.fields` property of the response.
	// It accepts wildcard (`*`) patterns.
	DocvalueFields []types.FieldAndFormat `json:"docvalue_fields,omitempty"`
	// Fields The request returns values for field names matching these patterns
	// in the `hits.fields` property of the response.
	// It accepts wildcard (`*`) patterns.
	Fields []string `json:"fields,omitempty"`
	// Filter A query to filter the documents that can match. The kNN search will return
	// the top
	// `k` documents that also match this filter. The value can be a single query or
	// a
	// list of queries. If `filter` isn't provided, all documents are allowed to
	// match.
	Filter []types.Query `json:"filter,omitempty"`
	// Knn The kNN query to run.
	Knn types.CoreKnnQuery `json:"knn"`
	// Source_ Indicates which source fields are returned for matching documents. These
	// fields are returned in the `hits._source` property of the search response.
	Source_ types.SourceConfig `json:"_source,omitempty"`
	// StoredFields A list of stored fields to return as part of a hit. If no fields are
	// specified,
	// no stored fields are included in the response. If this field is specified,
	// the `_source`
	// parameter defaults to `false`. You can pass `_source: true` to return both
	// source fields
	// and stored fields in the search response.
	StoredFields []string `json:"stored_fields,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Knnsearch request: %w", err)
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

		case "docvalue_fields":
			if err := dec.Decode(&s.DocvalueFields); err != nil {
				return fmt.Errorf("%s | %w", "DocvalueFields", err)
			}

		case "fields":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Fields", err)
				}

				s.Fields = append(s.Fields, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Fields); err != nil {
					return fmt.Errorf("%s | %w", "Fields", err)
				}
			}

		case "filter":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := types.NewQuery()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Filter", err)
				}

				s.Filter = append(s.Filter, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Filter); err != nil {
					return fmt.Errorf("%s | %w", "Filter", err)
				}
			}

		case "knn":
			if err := dec.Decode(&s.Knn); err != nil {
				return fmt.Errorf("%s | %w", "Knn", err)
			}

		case "_source":
			message := json.RawMessage{}
			if err := dec.Decode(&message); err != nil {
				return fmt.Errorf("%s | %w", "Source_", err)
			}
			keyDec := json.NewDecoder(bytes.NewReader(message))
		source__field:
			for {
				t, err := keyDec.Token()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return fmt.Errorf("%s | %w", "Source_", err)
				}

				switch t {

				case "exclude_vectors", "excludes", "includes":
					o := types.NewSourceFilter()
					localDec := json.NewDecoder(bytes.NewReader(message))
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "Source_", err)
					}
					s.Source_ = o
					break source__field

				}
			}
			if s.Source_ == nil {
				localDec := json.NewDecoder(bytes.NewReader(message))
				if err := localDec.Decode(&s.Source_); err != nil {
					return fmt.Errorf("%s | %w", "Source_", err)
				}
			}

		case "stored_fields":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "StoredFields", err)
				}

				s.StoredFields = append(s.StoredFields, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.StoredFields); err != nil {
					return fmt.Errorf("%s | %w", "StoredFields", err)
				}
			}

		}
	}
	return nil
}
