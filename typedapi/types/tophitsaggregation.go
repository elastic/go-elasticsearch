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

// TopHitsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/metric.ts#L369-L429
type TopHitsAggregation struct {
	// DocvalueFields Fields for which to return doc values.
	DocvalueFields []FieldAndFormat `json:"docvalue_fields,omitempty"`
	// Explain If `true`, returns detailed information about score computation as part of a
	// hit.
	Explain *bool `json:"explain,omitempty"`
	// Field The field on which to run the aggregation.
	Field *string `json:"field,omitempty"`
	// Fields Array of wildcard (*) patterns. The request returns values for field names
	// matching these patterns in the hits.fields property of the response.
	Fields []FieldAndFormat `json:"fields,omitempty"`
	// From Starting document offset.
	From *int `json:"from,omitempty"`
	// Highlight Specifies the highlighter to use for retrieving highlighted snippets from one
	// or more fields in the search results.
	Highlight *Highlight `json:"highlight,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
	// ScriptFields Returns the result of one or more script evaluations for each hit.
	ScriptFields map[string]ScriptField `json:"script_fields,omitempty"`
	// SeqNoPrimaryTerm If `true`, returns sequence number and primary term of the last modification
	// of each hit.
	SeqNoPrimaryTerm *bool `json:"seq_no_primary_term,omitempty"`
	// Size The maximum number of top matching hits to return per bucket.
	Size *int `json:"size,omitempty"`
	// Sort Sort order of the top matching hits.
	// By default, the hits are sorted by the score of the main query.
	Sort []SortCombinations `json:"sort,omitempty"`
	// Source_ Selects the fields of the source that are returned.
	Source_ SourceConfig `json:"_source,omitempty"`
	// StoredFields Returns values for the specified stored fields (fields that use the `store`
	// mapping option).
	StoredFields []string `json:"stored_fields,omitempty"`
	// TrackScores If `true`, calculates and returns document scores, even if the scores are not
	// used for sorting.
	TrackScores *bool `json:"track_scores,omitempty"`
	// Version If `true`, returns document version as part of a hit.
	Version *bool `json:"version,omitempty"`
}

func (s *TopHitsAggregation) UnmarshalJSON(data []byte) error {

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

		case "explain":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Explain", err)
				}
				s.Explain = &value
			case bool:
				s.Explain = &v
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
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

		case "highlight":
			if err := dec.Decode(&s.Highlight); err != nil {
				return fmt.Errorf("%s | %w", "Highlight", err)
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		case "script_fields":
			if s.ScriptFields == nil {
				s.ScriptFields = make(map[string]ScriptField, 0)
			}
			if err := dec.Decode(&s.ScriptFields); err != nil {
				return fmt.Errorf("%s | %w", "ScriptFields", err)
			}

		case "seq_no_primary_term":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SeqNoPrimaryTerm", err)
				}
				s.SeqNoPrimaryTerm = &value
			case bool:
				s.SeqNoPrimaryTerm = &v
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
				o := new(SortCombinations)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Sort", err)
				}

				s.Sort = append(s.Sort, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Sort); err != nil {
					return fmt.Errorf("%s | %w", "Sort", err)
				}
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
					o := NewSourceFilter()
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

		case "track_scores":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TrackScores", err)
				}
				s.TrackScores = &value
			case bool:
				s.TrackScores = &v
			}

		case "version":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Version", err)
				}
				s.Version = &value
			case bool:
				s.Version = &v
			}

		}
	}
	return nil
}

// NewTopHitsAggregation returns a TopHitsAggregation.
func NewTopHitsAggregation() *TopHitsAggregation {
	r := &TopHitsAggregation{
		ScriptFields: make(map[string]ScriptField),
	}

	return r
}
