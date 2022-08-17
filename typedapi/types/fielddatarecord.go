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

// FielddataRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/fielddata/types.ts#L20-L48
type FielddataRecord struct {
	// Field field name
	Field *string `json:"field,omitempty"`
	// Host host name
	Host *string `json:"host,omitempty"`
	// Id node id
	Id *string `json:"id,omitempty"`
	// Ip ip address
	Ip *string `json:"ip,omitempty"`
	// Node node name
	Node *string `json:"node,omitempty"`
	// Size field data usage
	Size *string `json:"size,omitempty"`
}

// FielddataRecordBuilder holds FielddataRecord struct and provides a builder API.
type FielddataRecordBuilder struct {
	v *FielddataRecord
}

// NewFielddataRecord provides a builder for the FielddataRecord struct.
func NewFielddataRecordBuilder() *FielddataRecordBuilder {
	r := FielddataRecordBuilder{
		&FielddataRecord{},
	}

	return &r
}

// Build finalize the chain and returns the FielddataRecord struct
func (rb *FielddataRecordBuilder) Build() FielddataRecord {
	return *rb.v
}

// Field field name

func (rb *FielddataRecordBuilder) Field(field string) *FielddataRecordBuilder {
	rb.v.Field = &field
	return rb
}

// Host host name

func (rb *FielddataRecordBuilder) Host(host string) *FielddataRecordBuilder {
	rb.v.Host = &host
	return rb
}

// Id node id

func (rb *FielddataRecordBuilder) Id(id string) *FielddataRecordBuilder {
	rb.v.Id = &id
	return rb
}

// Ip ip address

func (rb *FielddataRecordBuilder) Ip(ip string) *FielddataRecordBuilder {
	rb.v.Ip = &ip
	return rb
}

// Node node name

func (rb *FielddataRecordBuilder) Node(node string) *FielddataRecordBuilder {
	rb.v.Node = &node
	return rb
}

// Size field data usage

func (rb *FielddataRecordBuilder) Size(size string) *FielddataRecordBuilder {
	rb.v.Size = &size
	return rb
}
