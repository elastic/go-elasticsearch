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
// https://github.com/elastic/elasticsearch-specification/tree/1ed5f4795fc7c4d9875601f883b8d5fb9023c526

package info

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package info
//
// https://github.com/elastic/elasticsearch-specification/blob/1ed5f4795fc7c4d9875601f883b8d5fb9023c526/specification/nodes/info/NodesInfoResponse.ts#L30-L32
type Response struct {
	ClusterName string `json:"cluster_name"`
	// NodeStats Contains statistics about the number of nodes selected by the request’s node
	// filters.
	NodeStats *types.NodeStatistics     `json:"_nodes,omitempty"`
	Nodes     map[string]types.NodeInfo `json:"nodes"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Nodes: make(map[string]types.NodeInfo, 0),
	}
	return r
}
