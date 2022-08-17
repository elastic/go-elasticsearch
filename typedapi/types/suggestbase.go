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

// SuggestBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/suggester.ts#L42-L46
type SuggestBase struct {
	Length int    `json:"length"`
	Offset int    `json:"offset"`
	Text   string `json:"text"`
}

// SuggestBaseBuilder holds SuggestBase struct and provides a builder API.
type SuggestBaseBuilder struct {
	v *SuggestBase
}

// NewSuggestBase provides a builder for the SuggestBase struct.
func NewSuggestBaseBuilder() *SuggestBaseBuilder {
	r := SuggestBaseBuilder{
		&SuggestBase{},
	}

	return &r
}

// Build finalize the chain and returns the SuggestBase struct
func (rb *SuggestBaseBuilder) Build() SuggestBase {
	return *rb.v
}

func (rb *SuggestBaseBuilder) Length(length int) *SuggestBaseBuilder {
	rb.v.Length = length
	return rb
}

func (rb *SuggestBaseBuilder) Offset(offset int) *SuggestBaseBuilder {
	rb.v.Offset = offset
	return rb
}

func (rb *SuggestBaseBuilder) Text(text string) *SuggestBaseBuilder {
	rb.v.Text = text
	return rb
}
