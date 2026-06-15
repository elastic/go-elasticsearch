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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

// Package clusterstatemetric
package clusterstatemetric

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/eb2e22fb2ac404e676d19bcc7bb089647f029026/specification/cluster/state/ClusterStateRequest.ts#L130-L149
type ClusterStateMetric struct {
	Name string
}

var (

	// All Shows all metrics.
	All = ClusterStateMetric{"_all"}

	// Version Shows the cluster state version.
	Version = ClusterStateMetric{"version"}

	// Masternode Shows the elected `master_node` part of the response.
	Masternode = ClusterStateMetric{"master_node"}

	// Blocks Shows the `blocks` part of the response.
	Blocks = ClusterStateMetric{"blocks"}

	// Nodes Shows the `nodes` part of the response.
	Nodes = ClusterStateMetric{"nodes"}

	// Metadata Shows the `metadata` part of the response. If you supply a comma-separated
	// list of indices, the returned output will only contain metadata for these
	// indices.
	Metadata = ClusterStateMetric{"metadata"}

	// Routingtable Shows the `routing_table` part of the response. If you supply a
	// comma-separated list of indices, the returned output will only contain the
	// routing table for these indices.
	Routingtable = ClusterStateMetric{"routing_table"}

	// Routingnodes Shows the `routing_nodes` part of the response.
	Routingnodes = ClusterStateMetric{"routing_nodes"}

	// Customs Shows the `customs` part of the response, which includes custom cluster state
	// information.
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
