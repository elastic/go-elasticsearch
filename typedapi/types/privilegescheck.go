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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/clusterprivilege"
)

// PrivilegesCheck type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/has_privileges_user_profile/types.ts#L26-L33
type PrivilegesCheck struct {
	Application []ApplicationPrivilegesCheck `json:"application,omitempty"`
	// Cluster A list of the cluster privileges that you want to check.
	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster,omitempty"`
	Index   []IndexPrivilegesCheck              `json:"index,omitempty"`
}

// PrivilegesCheckBuilder holds PrivilegesCheck struct and provides a builder API.
type PrivilegesCheckBuilder struct {
	v *PrivilegesCheck
}

// NewPrivilegesCheck provides a builder for the PrivilegesCheck struct.
func NewPrivilegesCheckBuilder() *PrivilegesCheckBuilder {
	r := PrivilegesCheckBuilder{
		&PrivilegesCheck{},
	}

	return &r
}

// Build finalize the chain and returns the PrivilegesCheck struct
func (rb *PrivilegesCheckBuilder) Build() PrivilegesCheck {
	return *rb.v
}

func (rb *PrivilegesCheckBuilder) Application(application []ApplicationPrivilegesCheckBuilder) *PrivilegesCheckBuilder {
	tmp := make([]ApplicationPrivilegesCheck, len(application))
	for _, value := range application {
		tmp = append(tmp, value.Build())
	}
	rb.v.Application = tmp
	return rb
}

// Cluster A list of the cluster privileges that you want to check.

func (rb *PrivilegesCheckBuilder) Cluster(cluster ...clusterprivilege.ClusterPrivilege) *PrivilegesCheckBuilder {
	rb.v.Cluster = cluster
	return rb
}

func (rb *PrivilegesCheckBuilder) Index(index []IndexPrivilegesCheckBuilder) *PrivilegesCheckBuilder {
	tmp := make([]IndexPrivilegesCheck, len(index))
	for _, value := range index {
		tmp = append(tmp, value.Build())
	}
	rb.v.Index = tmp
	return rb
}
