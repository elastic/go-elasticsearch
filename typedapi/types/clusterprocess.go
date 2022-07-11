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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// ClusterProcess type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/cluster/stats/types.ts#L246-L249
type ClusterProcess struct {
	Cpu                 ClusterProcessCpu                 `json:"cpu"`
	OpenFileDescriptors ClusterProcessOpenFileDescriptors `json:"open_file_descriptors"`
}

// ClusterProcessBuilder holds ClusterProcess struct and provides a builder API.
type ClusterProcessBuilder struct {
	v *ClusterProcess
}

// NewClusterProcess provides a builder for the ClusterProcess struct.
func NewClusterProcessBuilder() *ClusterProcessBuilder {
	r := ClusterProcessBuilder{
		&ClusterProcess{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterProcess struct
func (rb *ClusterProcessBuilder) Build() ClusterProcess {
	return *rb.v
}

func (rb *ClusterProcessBuilder) Cpu(cpu *ClusterProcessCpuBuilder) *ClusterProcessBuilder {
	v := cpu.Build()
	rb.v.Cpu = v
	return rb
}

func (rb *ClusterProcessBuilder) OpenFileDescriptors(openfiledescriptors *ClusterProcessOpenFileDescriptorsBuilder) *ClusterProcessBuilder {
	v := openfiledescriptors.Build()
	rb.v.OpenFileDescriptors = v
	return rb
}
