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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// InnerHits type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/search/_types/hits.ts#L106-L140
type InnerHits struct {
	Collapse       *FieldCollapse   `json:"collapse,omitempty"`
	DocvalueFields []FieldAndFormat `json:"docvalue_fields,omitempty"`
	Explain        *bool            `json:"explain,omitempty"`
	Fields         []string         `json:"fields,omitempty"`
	// From Inner hit starting document offset.
	From           *int       `json:"from,omitempty"`
	Highlight      *Highlight `json:"highlight,omitempty"`
	IgnoreUnmapped *bool      `json:"ignore_unmapped,omitempty"`
	// Name The name for the particular inner hit definition in the response.
	// Useful when a search request contains multiple inner hits.
	Name             *string                `json:"name,omitempty"`
	ScriptFields     map[string]ScriptField `json:"script_fields,omitempty"`
	SeqNoPrimaryTerm *bool                  `json:"seq_no_primary_term,omitempty"`
	// Size The maximum number of hits to return per `inner_hits`.
	Size *int `json:"size,omitempty"`
	// Sort How the inner hits should be sorted per `inner_hits`.
	// By default, inner hits are sorted by score.
	Sort        []SortCombinations `json:"sort,omitempty"`
	Source_     SourceConfig       `json:"_source,omitempty"`
	StoredField []string           `json:"stored_field,omitempty"`
	TrackScores *bool              `json:"track_scores,omitempty"`
	Version     *bool              `json:"version,omitempty"`
}

func (s *InnerHits) UnmarshalJSON(data []byte) error {

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

		case "collapse":
			if err := dec.Decode(&s.Collapse); err != nil {
				return err
			}

		case "docvalue_fields":
			if err := dec.Decode(&s.DocvalueFields); err != nil {
				return err
			}

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

		case "fields":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Fields = append(s.Fields, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Fields); err != nil {
					return err
				}
			}

		case "from":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.From = &value
			case float64:
				f := int(v)
				s.From = &f
			}

		case "highlight":
			if err := dec.Decode(&s.Highlight); err != nil {
				return err
			}

		case "ignore_unmapped":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreUnmapped = &value
			case bool:
				s.IgnoreUnmapped = &v
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "script_fields":
			if s.ScriptFields == nil {
				s.ScriptFields = make(map[string]ScriptField, 0)
			}
			if err := dec.Decode(&s.ScriptFields); err != nil {
				return err
			}

		case "seq_no_primary_term":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.SeqNoPrimaryTerm = &value
			case bool:
				s.SeqNoPrimaryTerm = &v
			}

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

		case "sort":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(SortCombinations)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Sort = append(s.Sort, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Sort); err != nil {
					return err
				}
			}

		case "_source":
			if err := dec.Decode(&s.Source_); err != nil {
				return err
			}

		case "stored_field":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.StoredField = append(s.StoredField, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.StoredField); err != nil {
					return err
				}
			}

		case "track_scores":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.TrackScores = &value
			case bool:
				s.TrackScores = &v
			}

		case "version":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Version = &value
			case bool:
				s.Version = &v
			}

		}
	}
	return nil
}

// NewInnerHits returns a InnerHits.
func NewInnerHits() *InnerHits {
	r := &InnerHits{
		ScriptFields: make(map[string]ScriptField, 0),
	}

	return r
}
