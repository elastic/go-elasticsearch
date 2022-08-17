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

// Repository type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/_types/SnapshotRepository.ts#L23-L27
type Repository struct {
	Settings RepositorySettings `json:"settings"`
	Type     string             `json:"type"`
	Uuid     *Uuid              `json:"uuid,omitempty"`
}

// RepositoryBuilder holds Repository struct and provides a builder API.
type RepositoryBuilder struct {
	v *Repository
}

// NewRepository provides a builder for the Repository struct.
func NewRepositoryBuilder() *RepositoryBuilder {
	r := RepositoryBuilder{
		&Repository{},
	}

	return &r
}

// Build finalize the chain and returns the Repository struct
func (rb *RepositoryBuilder) Build() Repository {
	return *rb.v
}

func (rb *RepositoryBuilder) Settings(settings *RepositorySettingsBuilder) *RepositoryBuilder {
	v := settings.Build()
	rb.v.Settings = v
	return rb
}

func (rb *RepositoryBuilder) Type_(type_ string) *RepositoryBuilder {
	rb.v.Type = type_
	return rb
}

func (rb *RepositoryBuilder) Uuid(uuid Uuid) *RepositoryBuilder {
	rb.v.Uuid = &uuid
	return rb
}
