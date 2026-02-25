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

// Package catallocationcolumn
package catallocationcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L1196-L1261
type CatAllocationColumn struct {
	Name string
}

var (

	// Shards The number of shards on the node.
	Shards = CatAllocationColumn{"shards"}

	// Shardsundesired The number of shards scheduled to be moved elsewhere in the cluster.
	Shardsundesired = CatAllocationColumn{"shards.undesired"}

	// Writeloadforecast The sum of index write load forecasts.
	Writeloadforecast = CatAllocationColumn{"write_load.forecast"}

	// Diskindicesforecast The sum of shard size forecasts.
	Diskindicesforecast = CatAllocationColumn{"disk.indices.forecast"}

	// Diskindices The disk space used by Elasticsearch indices.
	Diskindices = CatAllocationColumn{"disk.indices"}

	// Diskused The total disk space used on the node.
	Diskused = CatAllocationColumn{"disk.used"}

	// Diskavail The available disk space on the node.
	Diskavail = CatAllocationColumn{"disk.avail"}

	// Disktotal The total disk capacity of all volumes on the node.
	Disktotal = CatAllocationColumn{"disk.total"}

	// Diskpercent The percentage of disk space used on the node.
	Diskpercent = CatAllocationColumn{"disk.percent"}

	// Host IThe host of the node.
	Host = CatAllocationColumn{"host"}

	// Ip The IP address of the node.
	Ip = CatAllocationColumn{"ip"}

	// Node The name of the node.
	Node = CatAllocationColumn{"node"}

	// Noderole The roles assigned to the node.
	Noderole = CatAllocationColumn{"node.role"}
)

func (c CatAllocationColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatAllocationColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "shards":
		*c = Shards
	case "shards.undesired":
		*c = Shardsundesired
	case "write_load.forecast":
		*c = Writeloadforecast
	case "disk.indices.forecast":
		*c = Diskindicesforecast
	case "disk.indices":
		*c = Diskindices
	case "disk.used":
		*c = Diskused
	case "disk.avail":
		*c = Diskavail
	case "disk.total":
		*c = Disktotal
	case "disk.percent":
		*c = Diskpercent
	case "host":
		*c = Host
	case "ip":
		*c = Ip
	case "node":
		*c = Node
	case "node.role":
		*c = Noderole
	default:
		*c = CatAllocationColumn{string(text)}
	}

	return nil
}

func (c CatAllocationColumn) String() string {
	return c.Name
}
