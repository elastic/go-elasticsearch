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

// ResponseBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/usage/NodesUsageResponse.ts#L25-L28
type ResponseBase struct {
	ClusterName Name `json:"cluster_name"`
	// NodeStats Contains statistics about the number of nodes selected by the request’s node
	// filters.
	NodeStats *NodeStatistics      `json:"_nodes,omitempty"`
	Nodes     map[string]NodeUsage `json:"nodes"`
}

// ResponseBaseBuilder holds ResponseBase struct and provides a builder API.
type ResponseBaseBuilder struct {
	v *ResponseBase
}

// NewResponseBase provides a builder for the ResponseBase struct.
func NewResponseBaseBuilder() *ResponseBaseBuilder {
	r := ResponseBaseBuilder{
		&ResponseBase{
			Nodes: make(map[string]NodeUsage, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ResponseBase struct
func (rb *ResponseBaseBuilder) Build() ResponseBase {
	return *rb.v
}

func (rb *ResponseBaseBuilder) ClusterName(clustername Name) *ResponseBaseBuilder {
	rb.v.ClusterName = clustername
	return rb
}

// NodeStats Contains statistics about the number of nodes selected by the request’s node
// filters.

func (rb *ResponseBaseBuilder) NodeStats(nodestats *NodeStatisticsBuilder) *ResponseBaseBuilder {
	v := nodestats.Build()
	rb.v.NodeStats = &v
	return rb
}

func (rb *ResponseBaseBuilder) Nodes(values map[string]*NodeUsageBuilder) *ResponseBaseBuilder {
	tmp := make(map[string]NodeUsage, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Nodes = tmp
	return rb
}
