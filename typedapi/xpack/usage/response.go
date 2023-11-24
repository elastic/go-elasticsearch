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

package usage

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package usage
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/xpack/usage/XPackUsageResponse.ts#L43-L79
type Response struct {
	AggregateMetric     types.Base                    `json:"aggregate_metric"`
	Analytics           types.Analytics               `json:"analytics"`
	Archive             types.Archive                 `json:"archive"`
	Ccr                 types.Ccr                     `json:"ccr"`
	DataFrame           *types.Base                   `json:"data_frame,omitempty"`
	DataScience         *types.Base                   `json:"data_science,omitempty"`
	DataStreams         *types.DataStreams            `json:"data_streams,omitempty"`
	DataTiers           types.DataTiers               `json:"data_tiers"`
	Enrich              *types.Base                   `json:"enrich,omitempty"`
	Eql                 types.Eql                     `json:"eql"`
	Flattened           *types.Flattened              `json:"flattened,omitempty"`
	FrozenIndices       types.FrozenIndices           `json:"frozen_indices"`
	Graph               types.Base                    `json:"graph"`
	HealthApi           *types.HealthStatistics       `json:"health_api,omitempty"`
	Ilm                 types.Ilm                     `json:"ilm"`
	Logstash            types.Base                    `json:"logstash"`
	Ml                  types.MachineLearning         `json:"ml"`
	Monitoring          types.Monitoring              `json:"monitoring"`
	Rollup              types.Base                    `json:"rollup"`
	RuntimeFields       *types.XpackRuntimeFieldTypes `json:"runtime_fields,omitempty"`
	SearchableSnapshots types.SearchableSnapshots     `json:"searchable_snapshots"`
	Security            types.Security                `json:"security"`
	Slm                 types.Slm                     `json:"slm"`
	Spatial             types.Base                    `json:"spatial"`
	Sql                 types.Sql                     `json:"sql"`
	Transform           types.Base                    `json:"transform"`
	Vectors             *types.Vector                 `json:"vectors,omitempty"`
	VotingOnly          types.Base                    `json:"voting_only"`
	Watcher             types.Watcher                 `json:"watcher"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
