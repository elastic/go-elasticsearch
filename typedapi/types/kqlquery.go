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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Returns documents matching a provided Kibana Query Language (KQL) expression.
// The expression is parsed and rewritten into standard Query DSL.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_types/query_dsl/KqlQuery.ts#L24-L48
type KqlQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query. Boost values are relative to the default value of 1.0. A boost
	// value between 0 and 1.0 decreases the relevance score. A value greater than
	// 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// CaseInsensitive If `true`, performs case-insensitive matching for field names and keyword or
	// text terms.
	CaseInsensitive *bool `json:"case_insensitive,omitempty"`
	// DefaultField Default field, or field pattern with wildcards, to target when a bare term
	// does not specify a field. Defaults to the `index.query.default_field` index
	// setting, which has a default value of `*`.
	DefaultField *string `json:"default_field,omitempty"`
	// Query The KQL expression to parse.
	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`
	// TimeZone Coordinated Universal Time (UTC) offset or IANA time zone used to interpret
	// date literals in the expression.
	TimeZone *string `json:"time_zone,omitempty"`
}

func (s *KqlQuery) UnmarshalJSON(data []byte) error {

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

		case "default_field":
			if err := dec.Decode(&s.DefaultField); err != nil {
				return fmt.Errorf("%s | %w", "DefaultField", err)
			}

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Query", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Query = o

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

		case "time_zone":
			if err := dec.Decode(&s.TimeZone); err != nil {
				return fmt.Errorf("%s | %w", "TimeZone", err)
			}

		}
	}
	return nil
}

// NewKqlQuery returns a KqlQuery.
func NewKqlQuery() *KqlQuery {
	r := &KqlQuery{}

	return r
}

type KqlQueryVariant interface {
	KqlQueryCaster() *KqlQuery
}

func (s *KqlQuery) KqlQueryCaster() *KqlQuery {
	return s
}
