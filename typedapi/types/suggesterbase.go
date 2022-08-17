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

// SuggesterBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/suggester.ts#L119-L123
type SuggesterBase struct {
	Analyzer *string `json:"analyzer,omitempty"`
	Field    Field   `json:"field"`
	Size     *int    `json:"size,omitempty"`
}

// SuggesterBaseBuilder holds SuggesterBase struct and provides a builder API.
type SuggesterBaseBuilder struct {
	v *SuggesterBase
}

// NewSuggesterBase provides a builder for the SuggesterBase struct.
func NewSuggesterBaseBuilder() *SuggesterBaseBuilder {
	r := SuggesterBaseBuilder{
		&SuggesterBase{},
	}

	return &r
}

// Build finalize the chain and returns the SuggesterBase struct
func (rb *SuggesterBaseBuilder) Build() SuggesterBase {
	return *rb.v
}

func (rb *SuggesterBaseBuilder) Analyzer(analyzer string) *SuggesterBaseBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *SuggesterBaseBuilder) Field(field Field) *SuggesterBaseBuilder {
	rb.v.Field = field
	return rb
}

func (rb *SuggesterBaseBuilder) Size(size int) *SuggesterBaseBuilder {
	rb.v.Size = &size
	return rb
}
