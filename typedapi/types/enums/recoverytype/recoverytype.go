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

// Package recoverytype
package recoverytype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/eb2e22fb2ac404e676d19bcc7bb089647f029026/specification/indices/recovery/types.ts#L133-L144
type RecoveryType struct {
	Name string
}

var (

	// EMPTYSTORE An empty store. Indicates a new primary shard or the forced allocation of an
	// empty primary shard using the cluster reroute API.
	EMPTYSTORE = RecoveryType{"EMPTY_STORE"}

	// EXISTINGSTORE The store of an existing primary shard. Indicates recovery is related to node
	// startup or the allocation of an existing primary shard.
	EXISTINGSTORE = RecoveryType{"EXISTING_STORE"}

	// LOCALSHARDS Shards of another index on the same node. Indicates recovery is related to a
	// clone, shrink, or split operation.
	LOCALSHARDS = RecoveryType{"LOCAL_SHARDS"}

	// PEER A primary shard on another node. Indicates recovery is related to shard
	// replication.
	PEER = RecoveryType{"PEER"}

	// SNAPSHOT A snapshot. Indicates recovery is related to a snapshot restore operation.
	SNAPSHOT = RecoveryType{"SNAPSHOT"}
)

func (r RecoveryType) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RecoveryType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "empty_store":
		*r = EMPTYSTORE
	case "existing_store":
		*r = EXISTINGSTORE
	case "local_shards":
		*r = LOCALSHARDS
	case "peer":
		*r = PEER
	case "snapshot":
		*r = SNAPSHOT
	default:
		*r = RecoveryType{string(text)}
	}

	return nil
}

func (r RecoveryType) String() string {
	return r.Name
}
