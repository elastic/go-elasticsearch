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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// Operation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/mtermvectors/types.ts#L35-L49
type Operation struct {
	Doc             interface{}              `json:"doc,omitempty"`
	FieldStatistics *bool                    `json:"field_statistics,omitempty"`
	Fields          *Fields                  `json:"fields,omitempty"`
	Filter          *Filter                  `json:"filter,omitempty"`
	Id_             Id                       `json:"_id"`
	Index_          *IndexName               `json:"_index,omitempty"`
	Offsets         *bool                    `json:"offsets,omitempty"`
	Payloads        *bool                    `json:"payloads,omitempty"`
	Positions       *bool                    `json:"positions,omitempty"`
	Routing         *Routing                 `json:"routing,omitempty"`
	TermStatistics  *bool                    `json:"term_statistics,omitempty"`
	Version         *VersionNumber           `json:"version,omitempty"`
	VersionType     *versiontype.VersionType `json:"version_type,omitempty"`
}

// OperationBuilder holds Operation struct and provides a builder API.
type OperationBuilder struct {
	v *Operation
}

// NewOperation provides a builder for the Operation struct.
func NewOperationBuilder() *OperationBuilder {
	r := OperationBuilder{
		&Operation{},
	}

	return &r
}

// Build finalize the chain and returns the Operation struct
func (rb *OperationBuilder) Build() Operation {
	return *rb.v
}

func (rb *OperationBuilder) Doc(doc interface{}) *OperationBuilder {
	rb.v.Doc = doc
	return rb
}

func (rb *OperationBuilder) FieldStatistics(fieldstatistics bool) *OperationBuilder {
	rb.v.FieldStatistics = &fieldstatistics
	return rb
}

func (rb *OperationBuilder) Fields(fields *FieldsBuilder) *OperationBuilder {
	v := fields.Build()
	rb.v.Fields = &v
	return rb
}

func (rb *OperationBuilder) Filter(filter *FilterBuilder) *OperationBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *OperationBuilder) Id_(id_ Id) *OperationBuilder {
	rb.v.Id_ = id_
	return rb
}

func (rb *OperationBuilder) Index_(index_ IndexName) *OperationBuilder {
	rb.v.Index_ = &index_
	return rb
}

func (rb *OperationBuilder) Offsets(offsets bool) *OperationBuilder {
	rb.v.Offsets = &offsets
	return rb
}

func (rb *OperationBuilder) Payloads(payloads bool) *OperationBuilder {
	rb.v.Payloads = &payloads
	return rb
}

func (rb *OperationBuilder) Positions(positions bool) *OperationBuilder {
	rb.v.Positions = &positions
	return rb
}

func (rb *OperationBuilder) Routing(routing Routing) *OperationBuilder {
	rb.v.Routing = &routing
	return rb
}

func (rb *OperationBuilder) TermStatistics(termstatistics bool) *OperationBuilder {
	rb.v.TermStatistics = &termstatistics
	return rb
}

func (rb *OperationBuilder) Version(version VersionNumber) *OperationBuilder {
	rb.v.Version = &version
	return rb
}

func (rb *OperationBuilder) VersionType(versiontype versiontype.VersionType) *OperationBuilder {
	rb.v.VersionType = &versiontype
	return rb
}
