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

// NodeInfoSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L68-L84
type NodeInfoSettings struct {
	Action       *NodeInfoAction           `json:"action,omitempty"`
	Bootstrap    *NodeInfoBootstrap        `json:"bootstrap,omitempty"`
	Client       NodeInfoClient            `json:"client"`
	Cluster      NodeInfoSettingsCluster   `json:"cluster"`
	Discovery    *NodeInfoDiscover         `json:"discovery,omitempty"`
	Http         NodeInfoSettingsHttp      `json:"http"`
	Ingest       *NodeInfoSettingsIngest   `json:"ingest,omitempty"`
	Network      *NodeInfoSettingsNetwork  `json:"network,omitempty"`
	Node         NodeInfoSettingsNode      `json:"node"`
	Path         NodeInfoPath              `json:"path"`
	Repositories *NodeInfoRepositories     `json:"repositories,omitempty"`
	Script       *NodeInfoScript           `json:"script,omitempty"`
	Search       *NodeInfoSearch           `json:"search,omitempty"`
	Transport    NodeInfoSettingsTransport `json:"transport"`
	Xpack        *NodeInfoXpack            `json:"xpack,omitempty"`
}

// NodeInfoSettingsBuilder holds NodeInfoSettings struct and provides a builder API.
type NodeInfoSettingsBuilder struct {
	v *NodeInfoSettings
}

// NewNodeInfoSettings provides a builder for the NodeInfoSettings struct.
func NewNodeInfoSettingsBuilder() *NodeInfoSettingsBuilder {
	r := NodeInfoSettingsBuilder{
		&NodeInfoSettings{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoSettings struct
func (rb *NodeInfoSettingsBuilder) Build() NodeInfoSettings {
	return *rb.v
}

func (rb *NodeInfoSettingsBuilder) Action(action *NodeInfoActionBuilder) *NodeInfoSettingsBuilder {
	v := action.Build()
	rb.v.Action = &v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Bootstrap(bootstrap *NodeInfoBootstrapBuilder) *NodeInfoSettingsBuilder {
	v := bootstrap.Build()
	rb.v.Bootstrap = &v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Client(client *NodeInfoClientBuilder) *NodeInfoSettingsBuilder {
	v := client.Build()
	rb.v.Client = v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Cluster(cluster *NodeInfoSettingsClusterBuilder) *NodeInfoSettingsBuilder {
	v := cluster.Build()
	rb.v.Cluster = v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Discovery(discovery *NodeInfoDiscoverBuilder) *NodeInfoSettingsBuilder {
	v := discovery.Build()
	rb.v.Discovery = &v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Http(http *NodeInfoSettingsHttpBuilder) *NodeInfoSettingsBuilder {
	v := http.Build()
	rb.v.Http = v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Ingest(ingest *NodeInfoSettingsIngestBuilder) *NodeInfoSettingsBuilder {
	v := ingest.Build()
	rb.v.Ingest = &v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Network(network *NodeInfoSettingsNetworkBuilder) *NodeInfoSettingsBuilder {
	v := network.Build()
	rb.v.Network = &v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Node(node *NodeInfoSettingsNodeBuilder) *NodeInfoSettingsBuilder {
	v := node.Build()
	rb.v.Node = v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Path(path *NodeInfoPathBuilder) *NodeInfoSettingsBuilder {
	v := path.Build()
	rb.v.Path = v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Repositories(repositories *NodeInfoRepositoriesBuilder) *NodeInfoSettingsBuilder {
	v := repositories.Build()
	rb.v.Repositories = &v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Script(script *NodeInfoScriptBuilder) *NodeInfoSettingsBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Search(search *NodeInfoSearchBuilder) *NodeInfoSettingsBuilder {
	v := search.Build()
	rb.v.Search = &v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Transport(transport *NodeInfoSettingsTransportBuilder) *NodeInfoSettingsBuilder {
	v := transport.Build()
	rb.v.Transport = v
	return rb
}

func (rb *NodeInfoSettingsBuilder) Xpack(xpack *NodeInfoXpackBuilder) *NodeInfoSettingsBuilder {
	v := xpack.Build()
	rb.v.Xpack = &v
	return rb
}
