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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// MoreLikeThisQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/specialized.ts#L78-L163
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
	// PerFieldAnalyzer Overrides the default analyzer.
	PerFieldAnalyzer map[string]string `json:"per_field_analyzer,omitempty"`
	QueryName_       *string           `json:"_name,omitempty"`
	Routing          *string           `json:"routing,omitempty"`
	// StopWords An array of stop words.
	// Any word in this set is ignored.
	StopWords []string `json:"stop_words,omitempty"`
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
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
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

		case "boost_terms":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.BoostTerms = &f
			case float64:
				f := Float64(v)
				s.BoostTerms = &f
			}

		case "fail_on_unsupported_field":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.FailOnUnsupportedField = &value
			case bool:
				s.FailOnUnsupportedField = &v
			}

		case "fields":
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "include":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
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
					return err
				}

				s.Like = append(s.Like, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Like); err != nil {
					return err
				}
			}

		case "max_doc_freq":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxDocFreq = &value
			case float64:
				f := int(v)
				s.MaxDocFreq = &f
			}

		case "max_query_terms":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxQueryTerms = &value
			case float64:
				f := int(v)
				s.MaxQueryTerms = &f
			}

		case "max_word_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxWordLength = &value
			case float64:
				f := int(v)
				s.MaxWordLength = &f
			}

		case "min_doc_freq":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MinDocFreq = &value
			case float64:
				f := int(v)
				s.MinDocFreq = &f
			}

		case "min_term_freq":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MinTermFreq = &value
			case float64:
				f := int(v)
				s.MinTermFreq = &f
			}

		case "min_word_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MinWordLength = &value
			case float64:
				f := int(v)
				s.MinWordLength = &f
			}

		case "minimum_should_match":
			if err := dec.Decode(&s.MinimumShouldMatch); err != nil {
				return err
			}

		case "per_field_analyzer":
			if s.PerFieldAnalyzer == nil {
				s.PerFieldAnalyzer = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.PerFieldAnalyzer); err != nil {
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

		case "routing":
			if err := dec.Decode(&s.Routing); err != nil {
				return err
			}

		case "stop_words":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.StopWords = append(s.StopWords, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.StopWords); err != nil {
					return err
				}
			}

		case "unlike":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(Like)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Unlike = append(s.Unlike, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Unlike); err != nil {
					return err
				}
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		case "version_type":
			if err := dec.Decode(&s.VersionType); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewMoreLikeThisQuery returns a MoreLikeThisQuery.
func NewMoreLikeThisQuery() *MoreLikeThisQuery {
	r := &MoreLikeThisQuery{
		PerFieldAnalyzer: make(map[string]string, 0),
	}

	return r
}
