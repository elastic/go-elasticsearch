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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

// DateRangeQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/query_dsl/term.ts#L72-L81
type DateRangeQuery struct {
	Boost      *float32                     `json:"boost,omitempty"`
	Format     *string                      `json:"format,omitempty"`
	From       string                       `json:"from,omitempty"`
	Gt         *string                      `json:"gt,omitempty"`
	Gte        *string                      `json:"gte,omitempty"`
	Lt         *string                      `json:"lt,omitempty"`
	Lte        *string                      `json:"lte,omitempty"`
	QueryName_ *string                      `json:"_name,omitempty"`
	Relation   *rangerelation.RangeRelation `json:"relation,omitempty"`
	TimeZone   *string                      `json:"time_zone,omitempty"`
	To         string                       `json:"to,omitempty"`
}

func (s *DateRangeQuery) UnmarshalJSON(data []byte) error {

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
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.Boost = &f
			case float64:
				f := float32(v)
				s.Boost = &f
			}

		case "format":
			if err := dec.Decode(&s.Format); err != nil {
				return err
			}

		case "from":
			if err := dec.Decode(&s.From); err != nil {
				return err
			}

		case "gt":
			if err := dec.Decode(&s.Gt); err != nil {
				return err
			}

		case "gte":
			if err := dec.Decode(&s.Gte); err != nil {
				return err
			}

		case "lt":
			if err := dec.Decode(&s.Lt); err != nil {
				return err
			}

		case "lte":
			if err := dec.Decode(&s.Lte); err != nil {
				return err
			}

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		case "relation":
			if err := dec.Decode(&s.Relation); err != nil {
				return err
			}

		case "time_zone":
			if err := dec.Decode(&s.TimeZone); err != nil {
				return err
			}

		case "to":
			if err := dec.Decode(&s.To); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDateRangeQuery returns a DateRangeQuery.
func NewDateRangeQuery() *DateRangeQuery {
	r := &DateRangeQuery{}

	return r
}
