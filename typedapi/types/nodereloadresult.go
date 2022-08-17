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

// NodeReloadResult holds the union for the following types:
//
//	NodeReloadError
//	Stats
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/NodeReloadResult.ts#L29-L30
type NodeReloadResult interface{}

// NodeReloadResultBuilder holds NodeReloadResult struct and provides a builder API.
type NodeReloadResultBuilder struct {
	v NodeReloadResult
}

// NewNodeReloadResult provides a builder for the NodeReloadResult struct.
func NewNodeReloadResultBuilder() *NodeReloadResultBuilder {
	return &NodeReloadResultBuilder{}
}

// Build finalize the chain and returns the NodeReloadResult struct
func (u *NodeReloadResultBuilder) Build() NodeReloadResult {
	return u.v
}

func (u *NodeReloadResultBuilder) NodeReloadError(nodereloaderror *NodeReloadErrorBuilder) *NodeReloadResultBuilder {
	v := nodereloaderror.Build()
	u.v = &v
	return u
}

func (u *NodeReloadResultBuilder) Stats(stats *StatsBuilder) *NodeReloadResultBuilder {
	v := stats.Build()
	u.v = &v
	return u
}
