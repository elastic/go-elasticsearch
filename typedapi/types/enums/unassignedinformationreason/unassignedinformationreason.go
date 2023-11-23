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

// Package unassignedinformationreason
package unassignedinformationreason

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/allocation_explain/types.ts#L127-L146
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

	case "INDEX_CREATED":
		*u = INDEXCREATED
	case "CLUSTER_RECOVERED":
		*u = CLUSTERRECOVERED
	case "INDEX_REOPENED":
		*u = INDEXREOPENED
	case "DANGLING_INDEX_IMPORTED":
		*u = DANGLINGINDEXIMPORTED
	case "NEW_INDEX_RESTORED":
		*u = NEWINDEXRESTORED
	case "EXISTING_INDEX_RESTORED":
		*u = EXISTINGINDEXRESTORED
	case "REPLICA_ADDED":
		*u = REPLICAADDED
	case "ALLOCATION_FAILED":
		*u = ALLOCATIONFAILED
	case "NODE_LEFT":
		*u = NODELEFT
	case "REROUTE_CANCELLED":
		*u = REROUTECANCELLED
	case "REINITIALIZED":
		*u = REINITIALIZED
	case "REALLOCATED_REPLICA":
		*u = REALLOCATEDREPLICA
	case "PRIMARY_FAILED":
		*u = PRIMARYFAILED
	case "FORCED_EMPTY_PRIMARY":
		*u = FORCEDEMPTYPRIMARY
	case "MANUAL_ALLOCATION":
		*u = MANUALALLOCATION
	default:
		*u = UnassignedInformationReason{string(text)}
	}

	return nil
}

func (u UnassignedInformationReason) String() string {
	return u.Name
}
