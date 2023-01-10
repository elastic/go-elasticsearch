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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// ClusterNodeCount type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cluster/stats/types.ts#L176-L192
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

// NewClusterNodeCount returns a ClusterNodeCount.
func NewClusterNodeCount() *ClusterNodeCount {
	r := &ClusterNodeCount{}

	return r
}
