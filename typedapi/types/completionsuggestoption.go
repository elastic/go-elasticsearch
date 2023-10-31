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

// CompletionSuggestOption type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/search/_types/suggester.ts#L73-L84
type CompletionSuggestOption struct {
	CollateMatch *bool                      `json:"collate_match,omitempty"`
	Contexts     map[string][]Context       `json:"contexts,omitempty"`
	Fields       map[string]json.RawMessage `json:"fields,omitempty"`
	Id_          *string                    `json:"_id,omitempty"`
	Index_       *string                    `json:"_index,omitempty"`
	Routing_     *string                    `json:"_routing,omitempty"`
	Score        *Float64                   `json:"score,omitempty"`
	Score_       *Float64                   `json:"_score,omitempty"`
	Source_      json.RawMessage            `json:"_source,omitempty"`
	Text         string                     `json:"text"`
}

func (s *CompletionSuggestOption) UnmarshalJSON(data []byte) error {

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

		case "collate_match":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.CollateMatch = &value
			case bool:
				s.CollateMatch = &v
			}

		case "contexts":
			if s.Contexts == nil {
				s.Contexts = make(map[string][]Context, 0)
			}
			if err := dec.Decode(&s.Contexts); err != nil {
				return err
			}

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Id_ = &o

		case "_index":
			if err := dec.Decode(&s.Index_); err != nil {
				return err
			}

		case "_routing":
			if err := dec.Decode(&s.Routing_); err != nil {
				return err
			}

		case "score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Score = &f
			case float64:
				f := Float64(v)
				s.Score = &f
			}

		case "_score":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Score_ = &f
			case float64:
				f := Float64(v)
				s.Score_ = &f
			}

		case "_source":
			if err := dec.Decode(&s.Source_); err != nil {
				return err
			}

		case "text":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Text = o

		}
	}
	return nil
}

// NewCompletionSuggestOption returns a CompletionSuggestOption.
func NewCompletionSuggestOption() *CompletionSuggestOption {
	r := &CompletionSuggestOption{
		Contexts: make(map[string][]Context, 0),
		Fields:   make(map[string]json.RawMessage, 0),
	}

	return r
}
