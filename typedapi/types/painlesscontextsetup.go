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

// PainlessContextSetup type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/scripts_painless_execute/types.ts#L25-L29
type PainlessContextSetup struct {
	Document interface{}    `json:"document,omitempty"`
	Index    IndexName      `json:"index"`
	Query    QueryContainer `json:"query"`
}

// PainlessContextSetupBuilder holds PainlessContextSetup struct and provides a builder API.
type PainlessContextSetupBuilder struct {
	v *PainlessContextSetup
}

// NewPainlessContextSetup provides a builder for the PainlessContextSetup struct.
func NewPainlessContextSetupBuilder() *PainlessContextSetupBuilder {
	r := PainlessContextSetupBuilder{
		&PainlessContextSetup{},
	}

	return &r
}

// Build finalize the chain and returns the PainlessContextSetup struct
func (rb *PainlessContextSetupBuilder) Build() PainlessContextSetup {
	return *rb.v
}

func (rb *PainlessContextSetupBuilder) Document(document interface{}) *PainlessContextSetupBuilder {
	rb.v.Document = document
	return rb
}

func (rb *PainlessContextSetupBuilder) Index(index IndexName) *PainlessContextSetupBuilder {
	rb.v.Index = index
	return rb
}

func (rb *PainlessContextSetupBuilder) Query(query *QueryContainerBuilder) *PainlessContextSetupBuilder {
	v := query.Build()
	rb.v.Query = v
	return rb
}
