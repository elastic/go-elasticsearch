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

// ShapeQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/query_dsl/specialized.ts#L367-L381
type ShapeQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// IgnoreUnmapped When set to `true` the query ignores an unmapped field and will not match any
	// documents.
	IgnoreUnmapped *bool                      `json:"ignore_unmapped,omitempty"`
	QueryName_     *string                    `json:"_name,omitempty"`
	ShapeQuery     map[string]ShapeFieldQuery `json:"-"`
}

func (s *ShapeQuery) UnmarshalJSON(data []byte) error {

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

		case "ignore_unmapped":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreUnmapped", err)
				}
				s.IgnoreUnmapped = &value
			case bool:
				s.IgnoreUnmapped = &v
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

		default:

			if key, ok := t.(string); ok {
				if s.ShapeQuery == nil {
					s.ShapeQuery = make(map[string]ShapeFieldQuery, 0)
				}
				raw := NewShapeFieldQuery()
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "ShapeQuery", err)
				}
				s.ShapeQuery[key] = *raw
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s ShapeQuery) MarshalJSON() ([]byte, error) {
	type opt ShapeQuery
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.ShapeQuery {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "ShapeQuery")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewShapeQuery returns a ShapeQuery.
func NewShapeQuery() *ShapeQuery {
	r := &ShapeQuery{
		ShapeQuery: make(map[string]ShapeFieldQuery),
	}

	return r
}

type ShapeQueryVariant interface {
	ShapeQueryCaster() *ShapeQuery
}

func (s *ShapeQuery) ShapeQueryCaster() *ShapeQuery {
	return s
}
