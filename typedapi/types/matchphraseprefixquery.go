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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// MatchPhrasePrefixQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/query_dsl/fulltext.ts#L182-L189
type MatchPhrasePrefixQuery struct {
	Analyzer       *string                        `json:"analyzer,omitempty"`
	Boost          *float32                       `json:"boost,omitempty"`
	MaxExpansions  *int                           `json:"max_expansions,omitempty"`
	Query          string                         `json:"query"`
	QueryName_     *string                        `json:"_name,omitempty"`
	Slop           *int                           `json:"slop,omitempty"`
	ZeroTermsQuery *zerotermsquery.ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

func (s *MatchPhrasePrefixQuery) UnmarshalJSON(data []byte) error {

	if !bytes.HasPrefix(data, []byte(`{`)) {
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&s.Query)
		return err
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

		case "analyzer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Analyzer = &o

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

		case "max_expansions":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxExpansions = &value
			case float64:
				f := int(v)
				s.MaxExpansions = &f
			}

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Query = o

		case "_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.QueryName_ = &o

		case "slop":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Slop = &value
			case float64:
				f := int(v)
				s.Slop = &f
			}

		case "zero_terms_query":
			if err := dec.Decode(&s.ZeroTermsQuery); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewMatchPhrasePrefixQuery returns a MatchPhrasePrefixQuery.
func NewMatchPhrasePrefixQuery() *MatchPhrasePrefixQuery {
	r := &MatchPhrasePrefixQuery{}

	return r
}
