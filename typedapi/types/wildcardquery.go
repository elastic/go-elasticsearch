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

// WildcardQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/query_dsl/term.ts#L301-L321
type WildcardQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// CaseInsensitive Allows case insensitive matching of the pattern with the indexed field values
	// when set to true. Default is false which means the case sensitivity of
	// matching depends on the underlying field’s mapping.
	CaseInsensitive *bool   `json:"case_insensitive,omitempty"`
	QueryName_      *string `json:"_name,omitempty"`
	// Rewrite Method used to rewrite the query.
	Rewrite *string `json:"rewrite,omitempty"`
	// Value Wildcard pattern for terms you wish to find in the provided field. Required,
	// when wildcard is not set.
	Value *string `json:"value,omitempty"`
	// Wildcard Wildcard pattern for terms you wish to find in the provided field. Required,
	// when value is not set.
	Wildcard *string `json:"wildcard,omitempty"`
}

func (s *WildcardQuery) UnmarshalJSON(data []byte) error {

	if !bytes.HasPrefix(data, []byte(`{`)) {
		if !bytes.HasPrefix(data, []byte(`"`)) {
			data = append([]byte{'"'}, data...)
			data = append(data, []byte{'"'}...)
		}
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&s.Value)
		if err != nil {
			return err
		}
		return nil
	}

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

		case "case_insensitive":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CaseInsensitive", err)
				}
				s.CaseInsensitive = &value
			case bool:
				s.CaseInsensitive = &v
			}

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

		case "rewrite":
			if err := dec.Decode(&s.Rewrite); err != nil {
				return fmt.Errorf("%s | %w", "Rewrite", err)
			}

		case "value":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Value", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Value = &o

		case "wildcard":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Wildcard", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Wildcard = &o

		}
	}
	return nil
}

// NewWildcardQuery returns a WildcardQuery.
func NewWildcardQuery() *WildcardQuery {
	r := &WildcardQuery{}

	return r
}

type WildcardQueryVariant interface {
	WildcardQueryCaster() *WildcardQuery
}

func (s *WildcardQuery) WildcardQueryCaster() *WildcardQuery {
	return s
}
