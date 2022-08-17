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

// RoleDescriptor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/_types/RoleDescriptor.ts#L27-L36
type RoleDescriptor struct {
	Applications      []ApplicationPrivileges  `json:"applications,omitempty"`
	Cluster           []string                 `json:"cluster,omitempty"`
	Global            []GlobalPrivilege        `json:"global,omitempty"`
	Indices           []IndicesPrivileges      `json:"indices,omitempty"`
	Metadata          *Metadata                `json:"metadata,omitempty"`
	RunAs             []string                 `json:"run_as,omitempty"`
	TransientMetadata *TransientMetadataConfig `json:"transient_metadata,omitempty"`
}

// RoleDescriptorBuilder holds RoleDescriptor struct and provides a builder API.
type RoleDescriptorBuilder struct {
	v *RoleDescriptor
}

// NewRoleDescriptor provides a builder for the RoleDescriptor struct.
func NewRoleDescriptorBuilder() *RoleDescriptorBuilder {
	r := RoleDescriptorBuilder{
		&RoleDescriptor{},
	}

	return &r
}

// Build finalize the chain and returns the RoleDescriptor struct
func (rb *RoleDescriptorBuilder) Build() RoleDescriptor {
	return *rb.v
}

func (rb *RoleDescriptorBuilder) Applications(applications []ApplicationPrivilegesBuilder) *RoleDescriptorBuilder {
	tmp := make([]ApplicationPrivileges, len(applications))
	for _, value := range applications {
		tmp = append(tmp, value.Build())
	}
	rb.v.Applications = tmp
	return rb
}

func (rb *RoleDescriptorBuilder) Cluster(cluster ...string) *RoleDescriptorBuilder {
	rb.v.Cluster = cluster
	return rb
}

func (rb *RoleDescriptorBuilder) Global(arg []GlobalPrivilege) *RoleDescriptorBuilder {
	rb.v.Global = arg
	return rb
}

func (rb *RoleDescriptorBuilder) Indices(indices []IndicesPrivilegesBuilder) *RoleDescriptorBuilder {
	tmp := make([]IndicesPrivileges, len(indices))
	for _, value := range indices {
		tmp = append(tmp, value.Build())
	}
	rb.v.Indices = tmp
	return rb
}

func (rb *RoleDescriptorBuilder) Metadata(metadata *MetadataBuilder) *RoleDescriptorBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

func (rb *RoleDescriptorBuilder) RunAs(run_as ...string) *RoleDescriptorBuilder {
	rb.v.RunAs = run_as
	return rb
}

func (rb *RoleDescriptorBuilder) TransientMetadata(transientmetadata *TransientMetadataConfigBuilder) *RoleDescriptorBuilder {
	v := transientmetadata.Build()
	rb.v.TransientMetadata = &v
	return rb
}
