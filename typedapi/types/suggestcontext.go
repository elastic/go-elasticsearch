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

// SuggestContext type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/specialized.ts#L36-L41
type SuggestContext struct {
	Name      Name   `json:"name"`
	Path      *Field `json:"path,omitempty"`
	Precision string `json:"precision,omitempty"`
	Type      string `json:"type"`
}

// SuggestContextBuilder holds SuggestContext struct and provides a builder API.
type SuggestContextBuilder struct {
	v *SuggestContext
}

// NewSuggestContext provides a builder for the SuggestContext struct.
func NewSuggestContextBuilder() *SuggestContextBuilder {
	r := SuggestContextBuilder{
		&SuggestContext{},
	}

	return &r
}

// Build finalize the chain and returns the SuggestContext struct
func (rb *SuggestContextBuilder) Build() SuggestContext {
	return *rb.v
}

func (rb *SuggestContextBuilder) Name(name Name) *SuggestContextBuilder {
	rb.v.Name = name
	return rb
}

func (rb *SuggestContextBuilder) Path(path Field) *SuggestContextBuilder {
	rb.v.Path = &path
	return rb
}

func (rb *SuggestContextBuilder) Precision(arg string) *SuggestContextBuilder {
	rb.v.Precision = arg
	return rb
}

func (rb *SuggestContextBuilder) Type_(type_ string) *SuggestContextBuilder {
	rb.v.Type = type_
	return rb
}
