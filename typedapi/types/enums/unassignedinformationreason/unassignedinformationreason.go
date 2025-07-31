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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package unassignedinformationreason
package unassignedinformationreason

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/allocation_explain/types.ts#L138-L157
type UnassignedInformationReason struct {
	Name string
}

var (
	INDEXCREATED = UnassignedInformationReason{"INDEX_CREATED"}

	CLUSTERRECOVERED = UnassignedInformationReason{"CLUSTER_RECOVERED"}

	INDEXREOPENED = UnassignedInformationReason{"INDEX_REOPENED"}

	DANGLINGINDEXIMPORTED = UnassignedInformationReason{"DANGLING_INDEX_IMPORTED"}

	NEWINDEXRESTORED = UnassignedInformationReason{"NEW_INDEX_RESTORED"}

	EXISTINGINDEXRESTORED = UnassignedInformationReason{"EXISTING_INDEX_RESTORED"}

	REPLICAADDED = UnassignedInformationReason{"REPLICA_ADDED"}

	ALLOCATIONFAILED = UnassignedInformationReason{"ALLOCATION_FAILED"}

	NODELEFT = UnassignedInformationReason{"NODE_LEFT"}

	REROUTECANCELLED = UnassignedInformationReason{"REROUTE_CANCELLED"}

	REINITIALIZED = UnassignedInformationReason{"REINITIALIZED"}

	REALLOCATEDREPLICA = UnassignedInformationReason{"REALLOCATED_REPLICA"}

	PRIMARYFAILED = UnassignedInformationReason{"PRIMARY_FAILED"}

	FORCEDEMPTYPRIMARY = UnassignedInformationReason{"FORCED_EMPTY_PRIMARY"}

	MANUALALLOCATION = UnassignedInformationReason{"MANUAL_ALLOCATION"}
)

func (u UnassignedInformationReason) MarshalText() (text []byte, err error) {
	return []byte(u.String()), nil
}

func (u *UnassignedInformationReason) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "index_created":
		*u = INDEXCREATED
	case "cluster_recovered":
		*u = CLUSTERRECOVERED
	case "index_reopened":
		*u = INDEXREOPENED
	case "dangling_index_imported":
		*u = DANGLINGINDEXIMPORTED
	case "new_index_restored":
		*u = NEWINDEXRESTORED
	case "existing_index_restored":
		*u = EXISTINGINDEXRESTORED
	case "replica_added":
		*u = REPLICAADDED
	case "allocation_failed":
		*u = ALLOCATIONFAILED
	case "node_left":
		*u = NODELEFT
	case "reroute_cancelled":
		*u = REROUTECANCELLED
	case "reinitialized":
		*u = REINITIALIZED
	case "reallocated_replica":
		*u = REALLOCATEDREPLICA
	case "primary_failed":
		*u = PRIMARYFAILED
	case "forced_empty_primary":
		*u = FORCEDEMPTYPRIMARY
	case "manual_allocation":
		*u = MANUALALLOCATION
	default:
		*u = UnassignedInformationReason{string(text)}
	}

	return nil
}

func (u UnassignedInformationReason) String() string {
	return u.Name
}
