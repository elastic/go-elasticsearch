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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// CompletionSuggestOption type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/suggester.ts#L73-L84
type CompletionSuggestOption struct {
	CollateMatch *bool                  `json:"collate_match,omitempty"`
	Contexts     map[string][]Context   `json:"contexts,omitempty"`
	Fields       map[string]interface{} `json:"fields,omitempty"`
	Id_          *string                `json:"_id,omitempty"`
	Index_       *IndexName             `json:"_index,omitempty"`
	Routing_     *Routing               `json:"_routing,omitempty"`
	Score        *float64               `json:"score,omitempty"`
	Score_       *float64               `json:"_score,omitempty"`
	Source_      interface{}            `json:"_source,omitempty"`
	Text         string                 `json:"text"`
}

// CompletionSuggestOptionBuilder holds CompletionSuggestOption struct and provides a builder API.
type CompletionSuggestOptionBuilder struct {
	v *CompletionSuggestOption
}

// NewCompletionSuggestOption provides a builder for the CompletionSuggestOption struct.
func NewCompletionSuggestOptionBuilder() *CompletionSuggestOptionBuilder {
	r := CompletionSuggestOptionBuilder{
		&CompletionSuggestOption{
			Contexts: make(map[string][]Context, 0),
			Fields:   make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the CompletionSuggestOption struct
func (rb *CompletionSuggestOptionBuilder) Build() CompletionSuggestOption {
	return *rb.v
}

func (rb *CompletionSuggestOptionBuilder) CollateMatch(collatematch bool) *CompletionSuggestOptionBuilder {
	rb.v.CollateMatch = &collatematch
	return rb
}

func (rb *CompletionSuggestOptionBuilder) Contexts(value map[string][]Context) *CompletionSuggestOptionBuilder {
	rb.v.Contexts = value
	return rb
}

func (rb *CompletionSuggestOptionBuilder) Fields(value map[string]interface{}) *CompletionSuggestOptionBuilder {
	rb.v.Fields = value
	return rb
}

func (rb *CompletionSuggestOptionBuilder) Id_(id_ string) *CompletionSuggestOptionBuilder {
	rb.v.Id_ = &id_
	return rb
}

func (rb *CompletionSuggestOptionBuilder) Index_(index_ IndexName) *CompletionSuggestOptionBuilder {
	rb.v.Index_ = &index_
	return rb
}

func (rb *CompletionSuggestOptionBuilder) Routing_(routing_ Routing) *CompletionSuggestOptionBuilder {
	rb.v.Routing_ = &routing_
	return rb
}

func (rb *CompletionSuggestOptionBuilder) Score(score float64) *CompletionSuggestOptionBuilder {
	rb.v.Score = &score
	return rb
}

func (rb *CompletionSuggestOptionBuilder) Score_(score_ float64) *CompletionSuggestOptionBuilder {
	rb.v.Score_ = &score_
	return rb
}

func (rb *CompletionSuggestOptionBuilder) Source_(source_ interface{}) *CompletionSuggestOptionBuilder {
	rb.v.Source_ = source_
	return rb
}

func (rb *CompletionSuggestOptionBuilder) Text(text string) *CompletionSuggestOptionBuilder {
	rb.v.Text = text
	return rb
}
