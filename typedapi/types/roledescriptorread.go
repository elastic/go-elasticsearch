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

// RoleDescriptorRead type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/_types/RoleDescriptor.ts#L38-L47
type RoleDescriptorRead struct {
	Applications      []ApplicationPrivileges  `json:"applications,omitempty"`
	Cluster           []string                 `json:"cluster"`
	Global            []GlobalPrivilege        `json:"global,omitempty"`
	Indices           []IndicesPrivileges      `json:"indices"`
	Metadata          *Metadata                `json:"metadata,omitempty"`
	RunAs             []string                 `json:"run_as,omitempty"`
	TransientMetadata *TransientMetadataConfig `json:"transient_metadata,omitempty"`
}

// RoleDescriptorReadBuilder holds RoleDescriptorRead struct and provides a builder API.
type RoleDescriptorReadBuilder struct {
	v *RoleDescriptorRead
}

// NewRoleDescriptorRead provides a builder for the RoleDescriptorRead struct.
func NewRoleDescriptorReadBuilder() *RoleDescriptorReadBuilder {
	r := RoleDescriptorReadBuilder{
		&RoleDescriptorRead{},
	}

	return &r
}

// Build finalize the chain and returns the RoleDescriptorRead struct
func (rb *RoleDescriptorReadBuilder) Build() RoleDescriptorRead {
	return *rb.v
}

func (rb *RoleDescriptorReadBuilder) Applications(applications []ApplicationPrivilegesBuilder) *RoleDescriptorReadBuilder {
	tmp := make([]ApplicationPrivileges, len(applications))
	for _, value := range applications {
		tmp = append(tmp, value.Build())
	}
	rb.v.Applications = tmp
	return rb
}

func (rb *RoleDescriptorReadBuilder) Cluster(cluster ...string) *RoleDescriptorReadBuilder {
	rb.v.Cluster = cluster
	return rb
}

func (rb *RoleDescriptorReadBuilder) Global(arg []GlobalPrivilege) *RoleDescriptorReadBuilder {
	rb.v.Global = arg
	return rb
}

func (rb *RoleDescriptorReadBuilder) Indices(indices []IndicesPrivilegesBuilder) *RoleDescriptorReadBuilder {
	tmp := make([]IndicesPrivileges, len(indices))
	for _, value := range indices {
		tmp = append(tmp, value.Build())
	}
	rb.v.Indices = tmp
	return rb
}

func (rb *RoleDescriptorReadBuilder) Metadata(metadata *MetadataBuilder) *RoleDescriptorReadBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

func (rb *RoleDescriptorReadBuilder) RunAs(run_as ...string) *RoleDescriptorReadBuilder {
	rb.v.RunAs = run_as
	return rb
}

func (rb *RoleDescriptorReadBuilder) TransientMetadata(transientmetadata *TransientMetadataConfigBuilder) *RoleDescriptorReadBuilder {
	v := transientmetadata.Build()
	rb.v.TransientMetadata = &v
	return rb
}
