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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Package noderole
package noderole

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/Node.ts#L76-L94
type NodeRole struct {
	Name string
}

var (
	Master = NodeRole{"master"}

	Data = NodeRole{"data"}

	Datacold = NodeRole{"data_cold"}

	Datacontent = NodeRole{"data_content"}

	Datafrozen = NodeRole{"data_frozen"}

	Datahot = NodeRole{"data_hot"}

	Datawarm = NodeRole{"data_warm"}

	Client = NodeRole{"client"}

	Ingest = NodeRole{"ingest"}

	Ml = NodeRole{"ml"}

	Votingonly = NodeRole{"voting_only"}

	Transform = NodeRole{"transform"}

	Remoteclusterclient = NodeRole{"remote_cluster_client"}

	Coordinatingonly = NodeRole{"coordinating_only"}
)

func (n NodeRole) MarshalText() (text []byte, err error) {
	return []byte(n.String()), nil
}

func (n *NodeRole) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "master":
		*n = Master
	case "data":
		*n = Data
	case "data_cold":
		*n = Datacold
	case "data_content":
		*n = Datacontent
	case "data_frozen":
		*n = Datafrozen
	case "data_hot":
		*n = Datahot
	case "data_warm":
		*n = Datawarm
	case "client":
		*n = Client
	case "ingest":
		*n = Ingest
	case "ml":
		*n = Ml
	case "voting_only":
		*n = Votingonly
	case "transform":
		*n = Transform
	case "remote_cluster_client":
		*n = Remoteclusterclient
	case "coordinating_only":
		*n = Coordinatingonly
	default:
		*n = NodeRole{string(text)}
	}

	return nil
}

func (n NodeRole) String() string {
	return n.Name
}
