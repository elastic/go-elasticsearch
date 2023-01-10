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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestmode"
)

// DirectGenerator type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_global/search/_types/suggester.ts#L166-L178
type DirectGenerator struct {
	Field          string                   `json:"field"`
	MaxEdits       *int                     `json:"max_edits,omitempty"`
	MaxInspections *float32                 `json:"max_inspections,omitempty"`
	MaxTermFreq    *float32                 `json:"max_term_freq,omitempty"`
	MinDocFreq     *float32                 `json:"min_doc_freq,omitempty"`
	MinWordLength  *int                     `json:"min_word_length,omitempty"`
	PostFilter     *string                  `json:"post_filter,omitempty"`
	PreFilter      *string                  `json:"pre_filter,omitempty"`
	PrefixLength   *int                     `json:"prefix_length,omitempty"`
	Size           *int                     `json:"size,omitempty"`
	SuggestMode    *suggestmode.SuggestMode `json:"suggest_mode,omitempty"`
}

// NewDirectGenerator returns a DirectGenerator.
func NewDirectGenerator() *DirectGenerator {
	r := &DirectGenerator{}

	return r
}
