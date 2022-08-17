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

// UpdateByQueryRethrottleNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/update_by_query_rethrottle/UpdateByQueryRethrottleNode.ts#L25-L27
type UpdateByQueryRethrottleNode struct {
	Attributes       map[string]string   `json:"attributes"`
	Host             Host                `json:"host"`
	Ip               Ip                  `json:"ip"`
	Name             Name                `json:"name"`
	Roles            *NodeRoles          `json:"roles,omitempty"`
	Tasks            map[TaskId]TaskInfo `json:"tasks"`
	TransportAddress TransportAddress    `json:"transport_address"`
}

// UpdateByQueryRethrottleNodeBuilder holds UpdateByQueryRethrottleNode struct and provides a builder API.
type UpdateByQueryRethrottleNodeBuilder struct {
	v *UpdateByQueryRethrottleNode
}

// NewUpdateByQueryRethrottleNode provides a builder for the UpdateByQueryRethrottleNode struct.
func NewUpdateByQueryRethrottleNodeBuilder() *UpdateByQueryRethrottleNodeBuilder {
	r := UpdateByQueryRethrottleNodeBuilder{
		&UpdateByQueryRethrottleNode{
			Attributes: make(map[string]string, 0),
			Tasks:      make(map[TaskId]TaskInfo, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the UpdateByQueryRethrottleNode struct
func (rb *UpdateByQueryRethrottleNodeBuilder) Build() UpdateByQueryRethrottleNode {
	return *rb.v
}

func (rb *UpdateByQueryRethrottleNodeBuilder) Attributes(value map[string]string) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.Attributes = value
	return rb
}

func (rb *UpdateByQueryRethrottleNodeBuilder) Host(host Host) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.Host = host
	return rb
}

func (rb *UpdateByQueryRethrottleNodeBuilder) Ip(ip Ip) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.Ip = ip
	return rb
}

func (rb *UpdateByQueryRethrottleNodeBuilder) Name(name Name) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.Name = name
	return rb
}

func (rb *UpdateByQueryRethrottleNodeBuilder) Roles(roles *NodeRolesBuilder) *UpdateByQueryRethrottleNodeBuilder {
	v := roles.Build()
	rb.v.Roles = &v
	return rb
}

func (rb *UpdateByQueryRethrottleNodeBuilder) Tasks(values map[TaskId]*TaskInfoBuilder) *UpdateByQueryRethrottleNodeBuilder {
	tmp := make(map[TaskId]TaskInfo, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Tasks = tmp
	return rb
}

func (rb *UpdateByQueryRethrottleNodeBuilder) TransportAddress(transportaddress TransportAddress) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}
