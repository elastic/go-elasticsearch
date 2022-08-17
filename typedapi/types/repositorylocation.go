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

// RepositoryLocation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/RepositoryMeteringInformation.ts#L68-L74
type RepositoryLocation struct {
	BasePath string `json:"base_path"`
	// Bucket Bucket name (GCP, S3)
	Bucket *string `json:"bucket,omitempty"`
	// Container Container name (Azure)
	Container *string `json:"container,omitempty"`
}

// RepositoryLocationBuilder holds RepositoryLocation struct and provides a builder API.
type RepositoryLocationBuilder struct {
	v *RepositoryLocation
}

// NewRepositoryLocation provides a builder for the RepositoryLocation struct.
func NewRepositoryLocationBuilder() *RepositoryLocationBuilder {
	r := RepositoryLocationBuilder{
		&RepositoryLocation{},
	}

	return &r
}

// Build finalize the chain and returns the RepositoryLocation struct
func (rb *RepositoryLocationBuilder) Build() RepositoryLocation {
	return *rb.v
}

func (rb *RepositoryLocationBuilder) BasePath(basepath string) *RepositoryLocationBuilder {
	rb.v.BasePath = basepath
	return rb
}

// Bucket Bucket name (GCP, S3)

func (rb *RepositoryLocationBuilder) Bucket(bucket string) *RepositoryLocationBuilder {
	rb.v.Bucket = &bucket
	return rb
}

// Container Container name (Azure)

func (rb *RepositoryLocationBuilder) Container(container string) *RepositoryLocationBuilder {
	rb.v.Container = &container
	return rb
}
