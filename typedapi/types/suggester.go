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

import (
	"encoding/json"
)

// Suggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/suggester.ts#L98-L101
type Suggester struct {
	Suggesters map[string]FieldSuggester `json:"-"`
	// Text Global suggest text, to avoid repetition when the same text is used in
	// several suggesters
	Text *string `json:"text,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s Suggester) MarshalJSON() ([]byte, error) {
	type opt Suggester
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.Suggesters {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// SuggesterBuilder holds Suggester struct and provides a builder API.
type SuggesterBuilder struct {
	v *Suggester
}

// NewSuggester provides a builder for the Suggester struct.
func NewSuggesterBuilder() *SuggesterBuilder {
	r := SuggesterBuilder{
		&Suggester{
			Suggesters: make(map[string]FieldSuggester, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Suggester struct
func (rb *SuggesterBuilder) Build() Suggester {
	return *rb.v
}

func (rb *SuggesterBuilder) Suggesters(values map[string]*FieldSuggesterBuilder) *SuggesterBuilder {
	tmp := make(map[string]FieldSuggester, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Suggesters = tmp
	return rb
}

// Text Global suggest text, to avoid repetition when the same text is used in
// several suggesters

func (rb *SuggesterBuilder) Text(text string) *SuggesterBuilder {
	rb.v.Text = &text
	return rb
}
