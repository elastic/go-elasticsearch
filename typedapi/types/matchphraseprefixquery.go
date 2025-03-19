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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

// MatchPhrasePrefixQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/c75a0abec670d027d13eb8d6f23374f86621c76b/specification/_types/query_dsl/fulltext.ts#L440-L469
type MatchPhrasePrefixQuery struct {
	// Analyzer Analyzer used to convert text in the query value into tokens.
	Analyzer *string `json:"analyzer,omitempty"`
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// MaxExpansions Maximum number of terms to which the last provided term of the query value
	// will expand.
	MaxExpansions *int `json:"max_expansions,omitempty"`
	// Query Text you wish to find in the provided field.
	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`
	// Slop Maximum number of positions allowed between matching tokens.
	Slop *int `json:"slop,omitempty"`
	// ZeroTermsQuery Indicates whether no documents are returned if the analyzer removes all
	// tokens, such as when using a `stop` filter.
	ZeroTermsQuery *zerotermsquery.ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

func (s *MatchPhrasePrefixQuery) UnmarshalJSON(data []byte) error {

	if !bytes.HasPrefix(data, []byte(`{`)) {
		if !bytes.HasPrefix(data, []byte(`"`)) {
			data = append([]byte{'"'}, data...)
			data = append(data, []byte{'"'}...)
		}
		err := json.NewDecoder(bytes.NewReader(data)).Decode(&s.Query)
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

		case "analyzer":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Analyzer", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Analyzer = &o

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

		case "max_expansions":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxExpansions", err)
				}
				s.MaxExpansions = &value
			case float64:
				f := int(v)
				s.MaxExpansions = &f
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

		case "slop":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Slop", err)
				}
				s.Slop = &value
			case float64:
				f := int(v)
				s.Slop = &f
			}

		case "zero_terms_query":
			if err := dec.Decode(&s.ZeroTermsQuery); err != nil {
				return fmt.Errorf("%s | %w", "ZeroTermsQuery", err)
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

// true

type MatchPhrasePrefixQueryVariant interface {
	MatchPhrasePrefixQueryCaster() *MatchPhrasePrefixQuery
}

func (s *MatchPhrasePrefixQuery) MatchPhrasePrefixQueryCaster() *MatchPhrasePrefixQuery {
	return s
}
