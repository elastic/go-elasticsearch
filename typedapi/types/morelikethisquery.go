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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// MoreLikeThisQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/query_dsl/specialized.ts#L62-L89
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

// NewMoreLikeThisQuery returns a MoreLikeThisQuery.
func NewMoreLikeThisQuery() *MoreLikeThisQuery {
	r := &MoreLikeThisQuery{
		PerFieldAnalyzer: make(map[string]string, 0),
	}

	return r
}
