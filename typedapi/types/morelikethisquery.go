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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// MoreLikeThisQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/query_dsl/specialized.ts#L87-L172
type MoreLikeThisQuery struct {
	// Analyzer The analyzer that is used to analyze the free form text.
	// Defaults to the analyzer associated with the first field in fields.
	Analyzer *string `json:"analyzer,omitempty"`
	// Boost Floating point number used to decrease or increase the relevance scores of
	// the query.
	// Boost values are relative to the default value of 1.0.
	// A boost value between 0 and 1.0 decreases the relevance score.
	// A value greater than 1.0 increases the relevance score.
	Boost *float32 `json:"boost,omitempty"`
	// BoostTerms Each term in the formed query could be further boosted by their tf-idf score.
	// This sets the boost factor to use when using this feature.
	// Defaults to deactivated (0).
	BoostTerms *Float64 `json:"boost_terms,omitempty"`
	// FailOnUnsupportedField Controls whether the query should fail (throw an exception) if any of the
	// specified fields are not of the supported types (`text` or `keyword`).
	FailOnUnsupportedField *bool `json:"fail_on_unsupported_field,omitempty"`
	// Fields A list of fields to fetch and analyze the text from.
	// Defaults to the `index.query.default_field` index setting, which has a
	// default value of `*`.
	Fields []string `json:"fields,omitempty"`
	// Include Specifies whether the input documents should also be included in the search
	// results returned.
	Include *bool `json:"include,omitempty"`
	// Like Specifies free form text and/or a single or multiple documents for which you
	// want to find similar documents.
	Like []Like `json:"like"`
	// MaxDocFreq The maximum document frequency above which the terms are ignored from the
	// input document.
	MaxDocFreq *int `json:"max_doc_freq,omitempty"`
	// MaxQueryTerms The maximum number of query terms that can be selected.
	MaxQueryTerms *int `json:"max_query_terms,omitempty"`
	// MaxWordLength The maximum word length above which the terms are ignored.
	// Defaults to unbounded (`0`).
	MaxWordLength *int `json:"max_word_length,omitempty"`
	// MinDocFreq The minimum document frequency below which the terms are ignored from the
	// input document.
	MinDocFreq *int `json:"min_doc_freq,omitempty"`
	// MinTermFreq The minimum term frequency below which the terms are ignored from the input
	// document.
	MinTermFreq *int `json:"min_term_freq,omitempty"`
	// MinWordLength The minimum word length below which the terms are ignored.
	MinWordLength *int `json:"min_word_length,omitempty"`
	// MinimumShouldMatch After the disjunctive query has been formed, this parameter controls the
	// number of terms that must match.
	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`
	QueryName_         *string            `json:"_name,omitempty"`
	Routing            *string            `json:"routing,omitempty"`
	// StopWords An array of stop words.
	// Any word in this set is ignored.
	StopWords StopWords `json:"stop_words,omitempty"`
	// Unlike Used in combination with `like` to exclude documents that match a set of
	// terms.
	Unlike      []Like                   `json:"unlike,omitempty"`
	Version     *int64                   `json:"version,omitempty"`
	VersionType *versiontype.VersionType `json:"version_type,omitempty"`
}

func (s *MoreLikeThisQuery) UnmarshalJSON(data []byte) error {

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

		case "boost_terms":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "BoostTerms", err)
				}
				f := Float64(value)
				s.BoostTerms = &f
			case float64:
				f := Float64(v)
				s.BoostTerms = &f
			}

		case "fail_on_unsupported_field":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "FailOnUnsupportedField", err)
				}
				s.FailOnUnsupportedField = &value
			case bool:
				s.FailOnUnsupportedField = &v
			}

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return fmt.Errorf("%s | %w", "Fields", err)
			}

		case "include":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Include", err)
				}
				s.Include = &value
			case bool:
				s.Include = &v
			}

		case "like":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(Like)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Like", err)
				}

				s.Like = append(s.Like, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Like); err != nil {
					return fmt.Errorf("%s | %w", "Like", err)
				}
			}

		case "max_doc_freq":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxDocFreq", err)
				}
				s.MaxDocFreq = &value
			case float64:
				f := int(v)
				s.MaxDocFreq = &f
			}

		case "max_query_terms":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxQueryTerms", err)
				}
				s.MaxQueryTerms = &value
			case float64:
				f := int(v)
				s.MaxQueryTerms = &f
			}

		case "max_word_length":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxWordLength", err)
				}
				s.MaxWordLength = &value
			case float64:
				f := int(v)
				s.MaxWordLength = &f
			}

		case "min_doc_freq":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinDocFreq", err)
				}
				s.MinDocFreq = &value
			case float64:
				f := int(v)
				s.MinDocFreq = &f
			}

		case "min_term_freq":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinTermFreq", err)
				}
				s.MinTermFreq = &value
			case float64:
				f := int(v)
				s.MinTermFreq = &f
			}

		case "min_word_length":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinWordLength", err)
				}
				s.MinWordLength = &value
			case float64:
				f := int(v)
				s.MinWordLength = &f
			}

		case "minimum_should_match":
			if err := dec.Decode(&s.MinimumShouldMatch); err != nil {
				return fmt.Errorf("%s | %w", "MinimumShouldMatch", err)
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

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return fmt.Errorf("%s | %w", "Routing", err)
			}

		case "stop_words":
			if err := dec.Decode(&s.StopWords); err != nil {
				return fmt.Errorf("%s | %w", "StopWords", err)
			}

		case "unlike":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(Like)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Unlike", err)
				}

				s.Unlike = append(s.Unlike, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Unlike); err != nil {
					return fmt.Errorf("%s | %w", "Unlike", err)
				}
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		case "version_type":
			if err := dec.Decode(&s.VersionType); err != nil {
				return fmt.Errorf("%s | %w", "VersionType", err)
			}

		}
	}
	return nil
}

// NewMoreLikeThisQuery returns a MoreLikeThisQuery.
func NewMoreLikeThisQuery() *MoreLikeThisQuery {
	r := &MoreLikeThisQuery{}

	return r
}
