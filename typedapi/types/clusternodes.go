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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

// ClusterNodes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/stats/types.ts#L369-L402
type ClusterNodes struct {
	// Count Contains counts for nodes selected by the request’s node filters.
	Count ClusterNodeCount `json:"count"`
	// DiscoveryTypes Contains statistics about the discovery types used by selected nodes.
	DiscoveryTypes map[string]int `json:"discovery_types"`
	// Fs Contains statistics about file stores by selected nodes.
	Fs               ClusterFileSystem       `json:"fs"`
	IndexingPressure ClusterIndexingPressure `json:"indexing_pressure"`
	Ingest           ClusterIngest           `json:"ingest"`
	// Jvm Contains statistics about the Java Virtual Machines (JVMs) used by selected
	// nodes.
	Jvm ClusterJvm `json:"jvm"`
	// NetworkTypes Contains statistics about the transport and HTTP networks used by selected
	// nodes.
	NetworkTypes ClusterNetworkTypes `json:"network_types"`
	// Os Contains statistics about the operating systems used by selected nodes.
	Os ClusterOperatingSystem `json:"os"`
	// PackagingTypes Contains statistics about Elasticsearch distributions installed on selected
	// nodes.
	PackagingTypes []NodePackagingType `json:"packaging_types"`
	// Plugins Contains statistics about installed plugins and modules by selected nodes.
	// If no plugins or modules are installed, this array is empty.
	Plugins []PluginStats `json:"plugins"`
	// Process Contains statistics about processes used by selected nodes.
	Process ClusterProcess `json:"process"`
	// Versions Array of Elasticsearch versions used on selected nodes.
	Versions []string `json:"versions"`
}

// NewClusterNodes returns a ClusterNodes.
func NewClusterNodes() *ClusterNodes {
	r := &ClusterNodes{
		DiscoveryTypes: make(map[string]int, 0),
	}

	return r
}
