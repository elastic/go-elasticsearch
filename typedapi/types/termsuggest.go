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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// TermSuggest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_global/search/_types/suggester.ts#L62-L67
type TermSuggest struct {
	Length  int               `json:"length"`
	Offset  int               `json:"offset"`
	Options TermSuggestOption `json:"options"`
	Text    string            `json:"text"`
}

// TermSuggestBuilder holds TermSuggest struct and provides a builder API.
type TermSuggestBuilder struct {
	v *TermSuggest
}

// NewTermSuggest provides a builder for the TermSuggest struct.
func NewTermSuggestBuilder() *TermSuggestBuilder {
	r := TermSuggestBuilder{
		&TermSuggest{},
	}

	return &r
}

// Build finalize the chain and returns the TermSuggest struct
func (rb *TermSuggestBuilder) Build() TermSuggest {
	return *rb.v
}

func (rb *TermSuggestBuilder) Length(length int) *TermSuggestBuilder {
	rb.v.Length = length
	return rb
}

func (rb *TermSuggestBuilder) Offset(offset int) *TermSuggestBuilder {
	rb.v.Offset = offset
	return rb
}

func (rb *TermSuggestBuilder) Options(options *TermSuggestOptionBuilder) *TermSuggestBuilder {
	v := options.Build()
	rb.v.Options = v
	return rb
}

func (rb *TermSuggestBuilder) Text(text string) *TermSuggestBuilder {
	rb.v.Text = text
	return rb
}
