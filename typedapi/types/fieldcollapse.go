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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// FieldCollapse type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/search/_types/FieldCollapse.ts#L24-L41
type FieldCollapse struct {
	Collapse *FieldCollapse `json:"collapse,omitempty"`
	// Field The field to collapse the result set on
	Field string `json:"field"`
	// InnerHits The number of inner hits and their sort order
	InnerHits []InnerHits `json:"inner_hits,omitempty"`
	// MaxConcurrentGroupSearches The number of concurrent requests allowed to retrieve the inner_hits per
	// group
	MaxConcurrentGroupSearches *int `json:"max_concurrent_group_searches,omitempty"`
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
				return fmt.Errorf("%s | %w", "Collapse", err)
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "inner_hits":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewInnerHits()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "InnerHits", err)
				}

				s.InnerHits = append(s.InnerHits, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.InnerHits); err != nil {
					return fmt.Errorf("%s | %w", "InnerHits", err)
				}
			}

		case "max_concurrent_group_searches":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxConcurrentGroupSearches", err)
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

type FieldCollapseVariant interface {
	FieldCollapseCaster() *FieldCollapse
}

func (s *FieldCollapse) FieldCollapseCaster() *FieldCollapse {
	return s
}
