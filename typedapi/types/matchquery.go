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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

// MatchQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/fulltext.ts#L282-L347
type MatchQuery struct {
	// Analyzer Analyzer used to convert the text in the query value into tokens.
	Analyzer *string `json:"analyzer,omitempty"`
	// AutoGenerateSynonymsPhraseQuery If `true`, match phrase queries are automatically created for multi-term
	// synonyms.
	AutoGenerateSynonymsPhraseQuery *bool `json:"auto_generate_synonyms_phrase_query,omitempty"`
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost           *float32 `json:"boost,omitempty"`
	CutoffFrequency *Float64 `json:"cutoff_frequency,omitempty"`
	// Fuzziness Maximum edit distance allowed for matching.
	Fuzziness Fuzziness `json:"fuzziness,omitempty"`
	// FuzzyRewrite Method used to rewrite the query.
	FuzzyRewrite *string `json:"fuzzy_rewrite,omitempty"`
	// FuzzyTranspositions If `true`, edits for fuzzy matching include transpositions of two adjacent
	// characters (for example, `ab` to `ba`).
	FuzzyTranspositions *bool `json:"fuzzy_transpositions,omitempty"`
	// Lenient If `true`, format-based errors, such as providing a text query value for a
	// numeric field, are ignored.
	Lenient *bool `json:"lenient,omitempty"`
	// MaxExpansions Maximum number of terms to which the query will expand.
	MaxExpansions *int `json:"max_expansions,omitempty"`
	// MinimumShouldMatch Minimum number of clauses that must match for a document to be returned.
	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`
	// Operator Boolean logic used to interpret text in the query value.
	Operator *operator.Operator `json:"operator,omitempty"`
	// PrefixLength Number of beginning characters left unchanged for fuzzy matching.
	PrefixLength *int `json:"prefix_length,omitempty"`
	// Query Text, number, boolean value or date you wish to find in the provided field.
	Query      string  `json:"query"`
	QueryName_ *string `json:"_name,omitempty"`
	// ZeroTermsQuery Indicates whether no documents are returned if the `analyzer` removes all
	// tokens, such as when using a `stop` filter.
	ZeroTermsQuery *zerotermsquery.ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

func (s *MatchQuery) UnmarshalJSON(data []byte) error {

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
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Analyzer = &o

		case "auto_generate_synonyms_phrase_query":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AutoGenerateSynonymsPhraseQuery = &value
			case bool:
				s.AutoGenerateSynonymsPhraseQuery = &v
			}

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

		case "cutoff_frequency":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.CutoffFrequency = &f
			case float64:
				f := Float64(v)
				s.CutoffFrequency = &f
			}

		case "fuzziness":
			if err := dec.Decode(&s.Fuzziness); err != nil {
				return err
			}

		case "fuzzy_rewrite":
			if err := dec.Decode(&s.FuzzyRewrite); err != nil {
				return err
			}

		case "fuzzy_transpositions":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.FuzzyTranspositions = &value
			case bool:
				s.FuzzyTranspositions = &v
			}

		case "lenient":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Lenient = &value
			case bool:
				s.Lenient = &v
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

		case "minimum_should_match":
			if err := dec.Decode(&s.MinimumShouldMatch); err != nil {
				return err
			}

		case "operator":
			if err := dec.Decode(&s.Operator); err != nil {
				return err
			}

		case "prefix_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PrefixLength = &value
			case float64:
				f := int(v)
				s.PrefixLength = &f
			}

		case "query":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
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
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueryName_ = &o

		case "zero_terms_query":
			if err := dec.Decode(&s.ZeroTermsQuery); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewMatchQuery returns a MatchQuery.
func NewMatchQuery() *MatchQuery {
	r := &MatchQuery{}

	return r
}
