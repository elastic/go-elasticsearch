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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package getrepositoriesmeteringinfo

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package getrepositoriesmeteringinfo
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/nodes/get_repositories_metering_info/GetRepositoriesMeteringInfoResponse.ts#L36-L38

type Response struct {

	// ClusterName Name of the cluster. Based on the [Cluster name
	// setting](https://www.elastic.co/guide/en/elasticsearch/reference/current/important-settings.html#cluster-name).
	ClusterName string `json:"cluster_name"`
	// NodeStats Contains statistics about the number of nodes selected by the request’s node
	// filters.
	NodeStats *types.NodeStatistics `json:"_nodes,omitempty"`
	// Nodes Contains repositories metering information for the nodes selected by the
	// request.
	Nodes map[string]types.RepositoryMeteringInformation `json:"nodes"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Nodes: make(map[string]types.RepositoryMeteringInformation, 0),
	}
	return r
}
