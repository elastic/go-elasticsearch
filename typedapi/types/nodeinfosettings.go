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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

// NodeInfoSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/nodes/info/types.ts#L79-L95
type NodeInfoSettings struct {
	Action       *NodeInfoAction           `json:"action,omitempty"`
	Bootstrap    *NodeInfoBootstrap        `json:"bootstrap,omitempty"`
	Client       *NodeInfoClient           `json:"client,omitempty"`
	Cluster      NodeInfoSettingsCluster   `json:"cluster"`
	Discovery    *NodeInfoDiscover         `json:"discovery,omitempty"`
	Http         NodeInfoSettingsHttp      `json:"http"`
	Ingest       *NodeInfoSettingsIngest   `json:"ingest,omitempty"`
	Network      *NodeInfoSettingsNetwork  `json:"network,omitempty"`
	Node         NodeInfoSettingsNode      `json:"node"`
	Path         *NodeInfoPath             `json:"path,omitempty"`
	Repositories *NodeInfoRepositories     `json:"repositories,omitempty"`
	Script       *NodeInfoScript           `json:"script,omitempty"`
	Search       *NodeInfoSearch           `json:"search,omitempty"`
	Transport    NodeInfoSettingsTransport `json:"transport"`
	Xpack        *NodeInfoXpack            `json:"xpack,omitempty"`
}

// NewNodeInfoSettings returns a NodeInfoSettings.
func NewNodeInfoSettings() *NodeInfoSettings {
	r := &NodeInfoSettings{}

	return r
}
