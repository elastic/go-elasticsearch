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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package clusterstatemetric
package clusterstatemetric

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cluster/state/ClusterStateRequest.ts#L123-L133
type ClusterStateMetric struct {
	Name string
}

var (
	All = ClusterStateMetric{"_all"}

	Version = ClusterStateMetric{"version"}

	Masternode = ClusterStateMetric{"master_node"}

	Blocks = ClusterStateMetric{"blocks"}

	Nodes = ClusterStateMetric{"nodes"}

	Metadata = ClusterStateMetric{"metadata"}

	Routingtable = ClusterStateMetric{"routing_table"}

	Routingnodes = ClusterStateMetric{"routing_nodes"}

	Customs = ClusterStateMetric{"customs"}
)

func (c ClusterStateMetric) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *ClusterStateMetric) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "_all":
		*c = All
	case "version":
		*c = Version
	case "master_node":
		*c = Masternode
	case "blocks":
		*c = Blocks
	case "nodes":
		*c = Nodes
	case "metadata":
		*c = Metadata
	case "routing_table":
		*c = Routingtable
	case "routing_nodes":
		*c = Routingnodes
	case "customs":
		*c = Customs
	default:
		*c = ClusterStateMetric{string(text)}
	}

	return nil
}

func (c ClusterStateMetric) String() string {
	return c.Name
}
