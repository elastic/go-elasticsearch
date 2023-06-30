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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// MoreLikeThisQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/query_dsl/specialized.ts#L62-L89
type MoreLikeThisQuery struct {
	Analyzer               *string                  `json:"analyzer,omitempty"`
	Boost                  *float32                 `json:"boost,omitempty"`
	BoostTerms             *Float64                 `json:"boost_terms,omitempty"`
	FailOnUnsupportedField *bool                    `json:"fail_on_unsupported_field,omitempty"`
	Fields                 []string                 `json:"fields,omitempty"`
	Include                *bool                    `json:"include,omitempty"`
	Like                   []Like                   `json:"like"`
	MaxDocFreq             *int                     `json:"max_doc_freq,omitempty"`
	MaxQueryTerms          *int                     `json:"max_query_terms,omitempty"`
	MaxWordLength          *int                     `json:"max_word_length,omitempty"`
	MinDocFreq             *int                     `json:"min_doc_freq,omitempty"`
	MinTermFreq            *int                     `json:"min_term_freq,omitempty"`
	MinWordLength          *int                     `json:"min_word_length,omitempty"`
	MinimumShouldMatch     MinimumShouldMatch       `json:"minimum_should_match,omitempty"`
	PerFieldAnalyzer       map[string]string        `json:"per_field_analyzer,omitempty"`
	QueryName_             *string                  `json:"_name,omitempty"`
	Routing                *string                  `json:"routing,omitempty"`
	StopWords              []string                 `json:"stop_words,omitempty"`
	Unlike                 []Like                   `json:"unlike,omitempty"`
	Version                *int64                   `json:"version,omitempty"`
	VersionType            *versiontype.VersionType `json:"version_type,omitempty"`
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
