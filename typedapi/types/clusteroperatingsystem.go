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

// ClusterOperatingSystem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L228-L235
type ClusterOperatingSystem struct {
	AllocatedProcessors int                                  `json:"allocated_processors"`
	Architectures       []ClusterOperatingSystemArchitecture `json:"architectures,omitempty"`
	AvailableProcessors int                                  `json:"available_processors"`
	Mem                 OperatingSystemMemoryInfo            `json:"mem"`
	Names               []ClusterOperatingSystemName         `json:"names"`
	PrettyNames         []ClusterOperatingSystemPrettyName   `json:"pretty_names"`
}

// ClusterOperatingSystemBuilder holds ClusterOperatingSystem struct and provides a builder API.
type ClusterOperatingSystemBuilder struct {
	v *ClusterOperatingSystem
}

// NewClusterOperatingSystem provides a builder for the ClusterOperatingSystem struct.
func NewClusterOperatingSystemBuilder() *ClusterOperatingSystemBuilder {
	r := ClusterOperatingSystemBuilder{
		&ClusterOperatingSystem{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterOperatingSystem struct
func (rb *ClusterOperatingSystemBuilder) Build() ClusterOperatingSystem {
	return *rb.v
}

func (rb *ClusterOperatingSystemBuilder) AllocatedProcessors(allocatedprocessors int) *ClusterOperatingSystemBuilder {
	rb.v.AllocatedProcessors = allocatedprocessors
	return rb
}

func (rb *ClusterOperatingSystemBuilder) Architectures(architectures []ClusterOperatingSystemArchitectureBuilder) *ClusterOperatingSystemBuilder {
	tmp := make([]ClusterOperatingSystemArchitecture, len(architectures))
	for _, value := range architectures {
		tmp = append(tmp, value.Build())
	}
	rb.v.Architectures = tmp
	return rb
}

func (rb *ClusterOperatingSystemBuilder) AvailableProcessors(availableprocessors int) *ClusterOperatingSystemBuilder {
	rb.v.AvailableProcessors = availableprocessors
	return rb
}

func (rb *ClusterOperatingSystemBuilder) Mem(mem *OperatingSystemMemoryInfoBuilder) *ClusterOperatingSystemBuilder {
	v := mem.Build()
	rb.v.Mem = v
	return rb
}

func (rb *ClusterOperatingSystemBuilder) Names(names []ClusterOperatingSystemNameBuilder) *ClusterOperatingSystemBuilder {
	tmp := make([]ClusterOperatingSystemName, len(names))
	for _, value := range names {
		tmp = append(tmp, value.Build())
	}
	rb.v.Names = tmp
	return rb
}

func (rb *ClusterOperatingSystemBuilder) PrettyNames(pretty_names []ClusterOperatingSystemPrettyNameBuilder) *ClusterOperatingSystemBuilder {
	tmp := make([]ClusterOperatingSystemPrettyName, len(pretty_names))
	for _, value := range pretty_names {
		tmp = append(tmp, value.Build())
	}
	rb.v.PrettyNames = tmp
	return rb
}
