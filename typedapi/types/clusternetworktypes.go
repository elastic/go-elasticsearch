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

package types

// ClusterNetworkTypes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/stats/types.ts#L337-L346
type ClusterNetworkTypes struct {
	// HttpTypes Contains statistics about the HTTP network types used by selected nodes.
	HttpTypes map[string]int `json:"http_types"`
	// TransportTypes Contains statistics about the transport network types used by selected nodes.
	TransportTypes map[string]int `json:"transport_types"`
}

// NewClusterNetworkTypes returns a ClusterNetworkTypes.
func NewClusterNetworkTypes() *ClusterNetworkTypes {
	r := &ClusterNetworkTypes{
		HttpTypes:      make(map[string]int, 0),
		TransportTypes: make(map[string]int, 0),
	}

	return r
}
