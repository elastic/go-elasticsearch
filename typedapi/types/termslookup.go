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

// TermsLookup type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/term.ts#L132-L137
type TermsLookup struct {
	Id      Id        `json:"id"`
	Index   IndexName `json:"index"`
	Path    Field     `json:"path"`
	Routing *Routing  `json:"routing,omitempty"`
}

// TermsLookupBuilder holds TermsLookup struct and provides a builder API.
type TermsLookupBuilder struct {
	v *TermsLookup
}

// NewTermsLookup provides a builder for the TermsLookup struct.
func NewTermsLookupBuilder() *TermsLookupBuilder {
	r := TermsLookupBuilder{
		&TermsLookup{},
	}

	return &r
}

// Build finalize the chain and returns the TermsLookup struct
func (rb *TermsLookupBuilder) Build() TermsLookup {
	return *rb.v
}

func (rb *TermsLookupBuilder) Id(id Id) *TermsLookupBuilder {
	rb.v.Id = id
	return rb
}

func (rb *TermsLookupBuilder) Index(index IndexName) *TermsLookupBuilder {
	rb.v.Index = index
	return rb
}

func (rb *TermsLookupBuilder) Path(path Field) *TermsLookupBuilder {
	rb.v.Path = path
	return rb
}

func (rb *TermsLookupBuilder) Routing(routing Routing) *TermsLookupBuilder {
	rb.v.Routing = &routing
	return rb
}
