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

package termvectors

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// Request holds the request body struct for the package termvectors
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/termvectors/TermVectorsRequest.ts#L33-L239
type Request struct {

	// Doc An artificial document (a document not present in the index) for which you
	// want to retrieve term vectors.
	Doc json.RawMessage `json:"doc,omitempty"`
	// FieldStatistics If `true`, the response includes:
	//
	// * The document count (how many documents contain this field).
	// * The sum of document frequencies (the sum of document frequencies for all
	// terms in this field).
	// * The sum of total term frequencies (the sum of total term frequencies of
	// each term in this field).
	FieldStatistics *bool `json:"field_statistics,omitempty"`
	// Fields A list of fields to include in the statistics.
	// It is used as the default list unless a specific field list is provided in
	// the `completion_fields` or `fielddata_fields` parameters.
	Fields []string `json:"fields,omitempty"`
	// Filter Filter terms based on their tf-idf scores.
	// This could be useful in order find out a good characteristic vector of a
	// document.
	// This feature works in a similar manner to the second phase of the More Like
	// This Query.
	Filter *types.TermVectorsFilter `json:"filter,omitempty"`
	// Offsets If `true`, the response includes term offsets.
	Offsets *bool `json:"offsets,omitempty"`
	// Payloads If `true`, the response includes term payloads.
	Payloads *bool `json:"payloads,omitempty"`
	// PerFieldAnalyzer Override the default per-field analyzer.
	// This is useful in order to generate term vectors in any fashion, especially
	// when using artificial documents.
	// When providing an analyzer for a field that already stores term vectors, the
	// term vectors will be regenerated.
	PerFieldAnalyzer map[string]string `json:"per_field_analyzer,omitempty"`
	// Positions If `true`, the response includes term positions.
	Positions *bool `json:"positions,omitempty"`
	// Routing A custom value that is used to route operations to a specific shard.
	Routing *string `json:"routing,omitempty"`
	// TermStatistics If `true`, the response includes:
	//
	// * The total term frequency (how often a term occurs in all documents).
	// * The document frequency (the number of documents containing the current
	// term).
	//
	// By default these values are not returned since term statistics can have a
	// serious performance impact.
	TermStatistics *bool `json:"term_statistics,omitempty"`
	// Version If `true`, returns the document version as part of a hit.
	Version *int64 `json:"version,omitempty"`
	// VersionType The version type.
	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		PerFieldAnalyzer: make(map[string]string, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Termvectors request: %w", err)
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

		case "doc":
			if err := dec.Decode(&s.Doc); err != nil {
				return fmt.Errorf("%s | %w", "Doc", err)
			}

		case "field_statistics":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "FieldStatistics", err)
				}
				s.FieldStatistics = &value
			case bool:
				s.FieldStatistics = &v
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
			if err := dec.Decode(&s.Filter); err != nil {
				return fmt.Errorf("%s | %w", "Filter", err)
			}

		case "offsets":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Offsets", err)
				}
				s.Offsets = &value
			case bool:
				s.Offsets = &v
			}

		case "payloads":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Payloads", err)
				}
				s.Payloads = &value
			case bool:
				s.Payloads = &v
			}

		case "per_field_analyzer":
			if s.PerFieldAnalyzer == nil {
				s.PerFieldAnalyzer = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.PerFieldAnalyzer); err != nil {
				return fmt.Errorf("%s | %w", "PerFieldAnalyzer", err)
			}

		case "positions":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Positions", err)
				}
				s.Positions = &value
			case bool:
				s.Positions = &v
			}

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
			}

		case "term_statistics":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TermStatistics", err)
				}
				s.TermStatistics = &value
			case bool:
				s.TermStatistics = &v
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		case "version_type":
			if err := dec.Decode(&s.VersionType); err != nil {
				return fmt.Errorf("%s | %w", "VersionType", err)
			}

		}
	}
	return nil
}
