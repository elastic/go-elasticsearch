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

// CompletionSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/suggester.ts#L127-L133
type CompletionSuggester struct {
	Analyzer       *string                       `json:"analyzer,omitempty"`
	Contexts       map[Field][]CompletionContext `json:"contexts,omitempty"`
	Field          Field                         `json:"field"`
	Fuzzy          *SuggestFuzziness             `json:"fuzzy,omitempty"`
	Prefix         *string                       `json:"prefix,omitempty"`
	Regex          *string                       `json:"regex,omitempty"`
	Size           *int                          `json:"size,omitempty"`
	SkipDuplicates *bool                         `json:"skip_duplicates,omitempty"`
}

// CompletionSuggesterBuilder holds CompletionSuggester struct and provides a builder API.
type CompletionSuggesterBuilder struct {
	v *CompletionSuggester
}

// NewCompletionSuggester provides a builder for the CompletionSuggester struct.
func NewCompletionSuggesterBuilder() *CompletionSuggesterBuilder {
	r := CompletionSuggesterBuilder{
		&CompletionSuggester{
			Contexts: make(map[Field][]CompletionContext, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the CompletionSuggester struct
func (rb *CompletionSuggesterBuilder) Build() CompletionSuggester {
	return *rb.v
}

func (rb *CompletionSuggesterBuilder) Analyzer(analyzer string) *CompletionSuggesterBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *CompletionSuggesterBuilder) Contexts(value map[Field][]CompletionContext) *CompletionSuggesterBuilder {
	rb.v.Contexts = value
	return rb
}

func (rb *CompletionSuggesterBuilder) Field(field Field) *CompletionSuggesterBuilder {
	rb.v.Field = field
	return rb
}

func (rb *CompletionSuggesterBuilder) Fuzzy(fuzzy *SuggestFuzzinessBuilder) *CompletionSuggesterBuilder {
	v := fuzzy.Build()
	rb.v.Fuzzy = &v
	return rb
}

func (rb *CompletionSuggesterBuilder) Prefix(prefix string) *CompletionSuggesterBuilder {
	rb.v.Prefix = &prefix
	return rb
}

func (rb *CompletionSuggesterBuilder) Regex(regex string) *CompletionSuggesterBuilder {
	rb.v.Regex = &regex
	return rb
}

func (rb *CompletionSuggesterBuilder) Size(size int) *CompletionSuggesterBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *CompletionSuggesterBuilder) SkipDuplicates(skipduplicates bool) *CompletionSuggesterBuilder {
	rb.v.SkipDuplicates = &skipduplicates
	return rb
}
