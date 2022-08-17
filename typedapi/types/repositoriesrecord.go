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

// RepositoriesRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/repositories/types.ts#L20-L31
type RepositoriesRecord struct {
	// Id unique repository id
	Id *string `json:"id,omitempty"`
	// Type repository type
	Type *string `json:"type,omitempty"`
}

// RepositoriesRecordBuilder holds RepositoriesRecord struct and provides a builder API.
type RepositoriesRecordBuilder struct {
	v *RepositoriesRecord
}

// NewRepositoriesRecord provides a builder for the RepositoriesRecord struct.
func NewRepositoriesRecordBuilder() *RepositoriesRecordBuilder {
	r := RepositoriesRecordBuilder{
		&RepositoriesRecord{},
	}

	return &r
}

// Build finalize the chain and returns the RepositoriesRecord struct
func (rb *RepositoriesRecordBuilder) Build() RepositoriesRecord {
	return *rb.v
}

// Id unique repository id

func (rb *RepositoriesRecordBuilder) Id(id string) *RepositoriesRecordBuilder {
	rb.v.Id = &id
	return rb
}

// Type repository type

func (rb *RepositoriesRecordBuilder) Type_(type_ string) *RepositoriesRecordBuilder {
	rb.v.Type = &type_
	return rb
}
