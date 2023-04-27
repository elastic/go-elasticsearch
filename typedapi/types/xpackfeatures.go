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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

// XpackFeatures type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/xpack/info/types.ts#L42-L72
type XpackFeatures struct {
	AggregateMetric     XpackFeature  `json:"aggregate_metric"`
	Analytics           XpackFeature  `json:"analytics"`
	Archive             XpackFeature  `json:"archive"`
	Ccr                 XpackFeature  `json:"ccr"`
	DataFrame           *XpackFeature `json:"data_frame,omitempty"`
	DataScience         *XpackFeature `json:"data_science,omitempty"`
	DataStreams         XpackFeature  `json:"data_streams"`
	DataTiers           XpackFeature  `json:"data_tiers"`
	Enrich              XpackFeature  `json:"enrich"`
	Eql                 XpackFeature  `json:"eql"`
	Flattened           *XpackFeature `json:"flattened,omitempty"`
	FrozenIndices       XpackFeature  `json:"frozen_indices"`
	Graph               XpackFeature  `json:"graph"`
	Ilm                 XpackFeature  `json:"ilm"`
	Logstash            XpackFeature  `json:"logstash"`
	Ml                  XpackFeature  `json:"ml"`
	Monitoring          XpackFeature  `json:"monitoring"`
	Rollup              XpackFeature  `json:"rollup"`
	RuntimeFields       *XpackFeature `json:"runtime_fields,omitempty"`
	SearchableSnapshots XpackFeature  `json:"searchable_snapshots"`
	Security            XpackFeature  `json:"security"`
	Slm                 XpackFeature  `json:"slm"`
	Spatial             XpackFeature  `json:"spatial"`
	Sql                 XpackFeature  `json:"sql"`
	Transform           XpackFeature  `json:"transform"`
	Vectors             *XpackFeature `json:"vectors,omitempty"`
	VotingOnly          XpackFeature  `json:"voting_only"`
	Watcher             XpackFeature  `json:"watcher"`
}

// NewXpackFeatures returns a XpackFeatures.
func NewXpackFeatures() *XpackFeatures {
	r := &XpackFeatures{}

	return r
}
