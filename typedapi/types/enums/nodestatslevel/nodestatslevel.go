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

// Package nodestatslevel
package nodestatslevel

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d520d9e8cf14cad487de5e0654007686c395b494/specification/_types/common.ts#L253-L257
type NodeStatsLevel struct {
	Name string
}

var (
	Node = NodeStatsLevel{"node"}

	Indices = NodeStatsLevel{"indices"}

	Shards = NodeStatsLevel{"shards"}
)

func (n NodeStatsLevel) MarshalText() (text []byte, err error) {
	return []byte(n.String()), nil
}

func (n *NodeStatsLevel) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "node":
		*n = Node
	case "indices":
		*n = Indices
	case "shards":
		*n = Shards
	default:
		*n = NodeStatsLevel{string(text)}
	}

	return nil
}

func (n NodeStatsLevel) String() string {
	return n.Name
}
