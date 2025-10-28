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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

// Package catallocationcolumn
package catallocationcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d520d9e8cf14cad487de5e0654007686c395b494/specification/cat/_types/CatBase.ts#L1184-L1249
type CatAllocationColumn struct {
	Name string
}

var (
	Shards = CatAllocationColumn{"shards"}

	Shardsundesired = CatAllocationColumn{"shards.undesired"}

	Writeloadforecast = CatAllocationColumn{"write_load.forecast"}

	Diskindicesforecast = CatAllocationColumn{"disk.indices.forecast"}

	Diskindices = CatAllocationColumn{"disk.indices"}

	Diskused = CatAllocationColumn{"disk.used"}

	Diskavail = CatAllocationColumn{"disk.avail"}

	Disktotal = CatAllocationColumn{"disk.total"}

	Diskpercent = CatAllocationColumn{"disk.percent"}

	Host = CatAllocationColumn{"host"}

	Ip = CatAllocationColumn{"ip"}

	Node = CatAllocationColumn{"node"}

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
