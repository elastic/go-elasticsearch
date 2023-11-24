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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

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
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/term.ts#L116-L143
type DateRangeQuery struct {
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// Format Date format used to convert `date` values in the query.
	Format *string `json:"format,omitempty"`
	From   string  `json:"from,omitempty"`
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
	// TimeZone Coordinated Universal Time (UTC) offset or IANA time zone used to convert
	// `date` values in the query to UTC.
	TimeZone *string `json:"time_zone,omitempty"`
	To       string  `json:"to,omitempty"`
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
