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

// Role type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/get_role/types.ts#L29-L39
type Role struct {
	Applications      []ApplicationPrivileges                   `json:"applications"`
	Cluster           []string                                  `json:"cluster"`
	Global            map[string]map[string]map[string][]string `json:"global,omitempty"`
	Indices           []IndicesPrivileges                       `json:"indices"`
	Metadata          Metadata                                  `json:"metadata"`
	RoleTemplates     []RoleTemplate                            `json:"role_templates,omitempty"`
	RunAs             []string                                  `json:"run_as"`
	TransientMetadata TransientMetadataConfig                   `json:"transient_metadata"`
}

// RoleBuilder holds Role struct and provides a builder API.
type RoleBuilder struct {
	v *Role
}

// NewRole provides a builder for the Role struct.
func NewRoleBuilder() *RoleBuilder {
	r := RoleBuilder{
		&Role{
			Global: make(map[string]map[string]map[string][]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Role struct
func (rb *RoleBuilder) Build() Role {
	return *rb.v
}

func (rb *RoleBuilder) Applications(applications []ApplicationPrivilegesBuilder) *RoleBuilder {
	tmp := make([]ApplicationPrivileges, len(applications))
	for _, value := range applications {
		tmp = append(tmp, value.Build())
	}
	rb.v.Applications = tmp
	return rb
}

func (rb *RoleBuilder) Cluster(cluster ...string) *RoleBuilder {
	rb.v.Cluster = cluster
	return rb
}

func (rb *RoleBuilder) Global(value map[string]map[string]map[string][]string) *RoleBuilder {
	rb.v.Global = value
	return rb
}

func (rb *RoleBuilder) Indices(indices []IndicesPrivilegesBuilder) *RoleBuilder {
	tmp := make([]IndicesPrivileges, len(indices))
	for _, value := range indices {
		tmp = append(tmp, value.Build())
	}
	rb.v.Indices = tmp
	return rb
}

func (rb *RoleBuilder) Metadata(metadata *MetadataBuilder) *RoleBuilder {
	v := metadata.Build()
	rb.v.Metadata = v
	return rb
}

func (rb *RoleBuilder) RoleTemplates(role_templates []RoleTemplateBuilder) *RoleBuilder {
	tmp := make([]RoleTemplate, len(role_templates))
	for _, value := range role_templates {
		tmp = append(tmp, value.Build())
	}
	rb.v.RoleTemplates = tmp
	return rb
}

func (rb *RoleBuilder) RunAs(run_as ...string) *RoleBuilder {
	rb.v.RunAs = run_as
	return rb
}

func (rb *RoleBuilder) TransientMetadata(transientmetadata *TransientMetadataConfigBuilder) *RoleBuilder {
	v := transientmetadata.Build()
	rb.v.TransientMetadata = v
	return rb
}
