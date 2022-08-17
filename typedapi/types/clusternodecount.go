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

// ClusterNodeCount type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L176-L192
type ClusterNodeCount struct {
	CoordinatingOnly    int  `json:"coordinating_only"`
	Data                int  `json:"data"`
	DataCold            int  `json:"data_cold"`
	DataContent         int  `json:"data_content"`
	DataFrozen          *int `json:"data_frozen,omitempty"`
	DataHot             int  `json:"data_hot"`
	DataWarm            int  `json:"data_warm"`
	Ingest              int  `json:"ingest"`
	Master              int  `json:"master"`
	Ml                  int  `json:"ml"`
	RemoteClusterClient int  `json:"remote_cluster_client"`
	Total               int  `json:"total"`
	Transform           int  `json:"transform"`
	VotingOnly          int  `json:"voting_only"`
}

// ClusterNodeCountBuilder holds ClusterNodeCount struct and provides a builder API.
type ClusterNodeCountBuilder struct {
	v *ClusterNodeCount
}

// NewClusterNodeCount provides a builder for the ClusterNodeCount struct.
func NewClusterNodeCountBuilder() *ClusterNodeCountBuilder {
	r := ClusterNodeCountBuilder{
		&ClusterNodeCount{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterNodeCount struct
func (rb *ClusterNodeCountBuilder) Build() ClusterNodeCount {
	return *rb.v
}

func (rb *ClusterNodeCountBuilder) CoordinatingOnly(coordinatingonly int) *ClusterNodeCountBuilder {
	rb.v.CoordinatingOnly = coordinatingonly
	return rb
}

func (rb *ClusterNodeCountBuilder) Data(data int) *ClusterNodeCountBuilder {
	rb.v.Data = data
	return rb
}

func (rb *ClusterNodeCountBuilder) DataCold(datacold int) *ClusterNodeCountBuilder {
	rb.v.DataCold = datacold
	return rb
}

func (rb *ClusterNodeCountBuilder) DataContent(datacontent int) *ClusterNodeCountBuilder {
	rb.v.DataContent = datacontent
	return rb
}

func (rb *ClusterNodeCountBuilder) DataFrozen(datafrozen int) *ClusterNodeCountBuilder {
	rb.v.DataFrozen = &datafrozen
	return rb
}

func (rb *ClusterNodeCountBuilder) DataHot(datahot int) *ClusterNodeCountBuilder {
	rb.v.DataHot = datahot
	return rb
}

func (rb *ClusterNodeCountBuilder) DataWarm(datawarm int) *ClusterNodeCountBuilder {
	rb.v.DataWarm = datawarm
	return rb
}

func (rb *ClusterNodeCountBuilder) Ingest(ingest int) *ClusterNodeCountBuilder {
	rb.v.Ingest = ingest
	return rb
}

func (rb *ClusterNodeCountBuilder) Master(master int) *ClusterNodeCountBuilder {
	rb.v.Master = master
	return rb
}

func (rb *ClusterNodeCountBuilder) Ml(ml int) *ClusterNodeCountBuilder {
	rb.v.Ml = ml
	return rb
}

func (rb *ClusterNodeCountBuilder) RemoteClusterClient(remoteclusterclient int) *ClusterNodeCountBuilder {
	rb.v.RemoteClusterClient = remoteclusterclient
	return rb
}

func (rb *ClusterNodeCountBuilder) Total(total int) *ClusterNodeCountBuilder {
	rb.v.Total = total
	return rb
}

func (rb *ClusterNodeCountBuilder) Transform(transform int) *ClusterNodeCountBuilder {
	rb.v.Transform = transform
	return rb
}

func (rb *ClusterNodeCountBuilder) VotingOnly(votingonly int) *ClusterNodeCountBuilder {
	rb.v.VotingOnly = votingonly
	return rb
}
