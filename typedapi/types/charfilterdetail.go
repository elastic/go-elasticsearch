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

// CharFilterDetail type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/analyze/types.ts#L46-L49
type CharFilterDetail struct {
	FilteredText []string `json:"filtered_text"`
	Name         string   `json:"name"`
}

// CharFilterDetailBuilder holds CharFilterDetail struct and provides a builder API.
type CharFilterDetailBuilder struct {
	v *CharFilterDetail
}

// NewCharFilterDetail provides a builder for the CharFilterDetail struct.
func NewCharFilterDetailBuilder() *CharFilterDetailBuilder {
	r := CharFilterDetailBuilder{
		&CharFilterDetail{},
	}

	return &r
}

// Build finalize the chain and returns the CharFilterDetail struct
func (rb *CharFilterDetailBuilder) Build() CharFilterDetail {
	return *rb.v
}

func (rb *CharFilterDetailBuilder) FilteredText(filtered_text ...string) *CharFilterDetailBuilder {
	rb.v.FilteredText = filtered_text
	return rb
}

func (rb *CharFilterDetailBuilder) Name(name string) *CharFilterDetailBuilder {
	rb.v.Name = name
	return rb
}
