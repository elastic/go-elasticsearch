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

// GetResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/get/types.ts#L25-L35
type GetResult struct {
	Fields       map[string]interface{} `json:"fields,omitempty"`
	Found        bool                   `json:"found"`
	Id_          Id                     `json:"_id"`
	Index_       IndexName              `json:"_index"`
	PrimaryTerm_ *int64                 `json:"_primary_term,omitempty"`
	Routing_     *string                `json:"_routing,omitempty"`
	SeqNo_       *SequenceNumber        `json:"_seq_no,omitempty"`
	Source_      interface{}            `json:"_source,omitempty"`
	Version_     *VersionNumber         `json:"_version,omitempty"`
}

// GetResultBuilder holds GetResult struct and provides a builder API.
type GetResultBuilder struct {
	v *GetResult
}

// NewGetResult provides a builder for the GetResult struct.
func NewGetResultBuilder() *GetResultBuilder {
	r := GetResultBuilder{
		&GetResult{
			Fields: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GetResult struct
func (rb *GetResultBuilder) Build() GetResult {
	return *rb.v
}

func (rb *GetResultBuilder) Fields(value map[string]interface{}) *GetResultBuilder {
	rb.v.Fields = value
	return rb
}

func (rb *GetResultBuilder) Found(found bool) *GetResultBuilder {
	rb.v.Found = found
	return rb
}

func (rb *GetResultBuilder) Id_(id_ Id) *GetResultBuilder {
	rb.v.Id_ = id_
	return rb
}

func (rb *GetResultBuilder) Index_(index_ IndexName) *GetResultBuilder {
	rb.v.Index_ = index_
	return rb
}

func (rb *GetResultBuilder) PrimaryTerm_(primaryterm_ int64) *GetResultBuilder {
	rb.v.PrimaryTerm_ = &primaryterm_
	return rb
}

func (rb *GetResultBuilder) Routing_(routing_ string) *GetResultBuilder {
	rb.v.Routing_ = &routing_
	return rb
}

func (rb *GetResultBuilder) SeqNo_(seqno_ SequenceNumber) *GetResultBuilder {
	rb.v.SeqNo_ = &seqno_
	return rb
}

func (rb *GetResultBuilder) Source_(source_ interface{}) *GetResultBuilder {
	rb.v.Source_ = source_
	return rb
}

func (rb *GetResultBuilder) Version_(version_ VersionNumber) *GetResultBuilder {
	rb.v.Version_ = &version_
	return rb
}
