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

// FieldCollapse type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_global/search/_types/FieldCollapse.ts#L24-L29
type FieldCollapse struct {
	Collapse                   *FieldCollapse `json:"collapse,omitempty"`
	Field                      string         `json:"field"`
	InnerHits                  []InnerHits    `json:"inner_hits,omitempty"`
	MaxConcurrentGroupSearches *int           `json:"max_concurrent_group_searches,omitempty"`
}

func (s *FieldCollapse) UnmarshalJSON(data []byte) error {

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

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "inner_hits":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewInnerHits()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.InnerHits = append(s.InnerHits, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.InnerHits); err != nil {
					return err
				}
			}

		case "max_concurrent_group_searches":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxConcurrentGroupSearches = &value
			case float64:
				f := int(v)
				s.MaxConcurrentGroupSearches = &f
			}

		}
	}
	return nil
}

// NewFieldCollapse returns a FieldCollapse.
func NewFieldCollapse() *FieldCollapse {
	r := &FieldCollapse{}

	return r
}
