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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

// TermRangeQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/query_dsl/term.ts#L170-L170
type TermRangeQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// Gt Greater than.
	Gt *string `json:"gt,omitempty"`
	// Gte Greater than or equal to.
	Gte *string `json:"gte,omitempty"`
	// Lt Less than.
	Lt *string `json:"lt,omitempty"`
	// Lte Less than or equal to.
	Lte        *string `json:"lte,omitempty"`
	QueryName_ *string `json:"_name,omitempty"`
	// Relation Indicates how the range query matches values for `range` fields.
	Relation *rangerelation.RangeRelation `json:"relation,omitempty"`
}

func (s *TermRangeQuery) UnmarshalJSON(data []byte) error {

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

		case "boost":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Boost", err)
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "gt":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Gt", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Gt = &o

		case "gte":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Gte", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Gte = &o

		case "lt":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Lt", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Lt = &o

		case "lte":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Lte", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Lte = &o

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "QueryName_", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		case "relation":
			if err := dec.Decode(&s.Relation); err != nil {
				return fmt.Errorf("%s | %w", "Relation", err)
			}

		}
	}
	return nil
}

// NewTermRangeQuery returns a TermRangeQuery.
func NewTermRangeQuery() *TermRangeQuery {
	r := &TermRangeQuery{}

	return r
}

type TermRangeQueryVariant interface {
	TermRangeQueryCaster() *TermRangeQuery
}

func (s *TermRangeQuery) TermRangeQueryCaster() *TermRangeQuery {
	return s
}

func (s *TermRangeQuery) RangeQueryCaster() *RangeQuery {
	o := RangeQuery(s)
	return &o
}
