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
// https://github.com/elastic/elasticsearch-specification/tree/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6

// Package recoverypriority
package recoverypriority

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6/specification/indices/recovery/types.ts#L146-L162
type RecoveryPriority struct {
	Name string
}

var (

	// UNASSIGNEDNEWPRIMARY A primary shard which is unassigned because it is newly created.
	UNASSIGNEDNEWPRIMARY = RecoveryPriority{"UNASSIGNED_NEW_PRIMARY"}

	// UNASSIGNEDUNEXPECTED A shard which is unassigned because of an unexpected condition, i.e. some
	// kind of failure.
	UNASSIGNEDUNEXPECTED = RecoveryPriority{"UNASSIGNED_UNEXPECTED"}

	// UNASSIGNEDEXPECTED A shard which is unassigned for an expected condition, i.e. because of a user
	// operation such as opening or restoring an index, but which is not an
	// UNASSIGNED_NEW_PRIMARY.
	UNASSIGNEDEXPECTED = RecoveryPriority{"UNASSIGNED_EXPECTED"}

	// RELOCATIONCANREMAINNO A shard which is assigned, and is being relocated because it cannot remain on
	// its current node according to the allocation deciders.
	RELOCATIONCANREMAINNO = RecoveryPriority{"RELOCATION_CAN_REMAIN_NO"}

	// RELOCATIONCANREMAINNOTPREFERRED A shard which is assigned, and is being relocated because it is not preferred
	// for it to remain on its current node according to the allocation deciders.
	RELOCATIONCANREMAINNOTPREFERRED = RecoveryPriority{"RELOCATION_CAN_REMAIN_NOT_PREFERRED"}

	// RELOCATEREBALANCING A shard which is assigned, and is being relocated for rebalancing.
	RELOCATEREBALANCING = RecoveryPriority{"RELOCATE_REBALANCING"}

	// UNKNOWN Placeholder value for unknown priorities.
	UNKNOWN = RecoveryPriority{"UNKNOWN"}
)

func (r RecoveryPriority) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RecoveryPriority) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "unassigned_new_primary":
		*r = UNASSIGNEDNEWPRIMARY
	case "unassigned_unexpected":
		*r = UNASSIGNEDUNEXPECTED
	case "unassigned_expected":
		*r = UNASSIGNEDEXPECTED
	case "relocation_can_remain_no":
		*r = RELOCATIONCANREMAINNO
	case "relocation_can_remain_not_preferred":
		*r = RELOCATIONCANREMAINNOTPREFERRED
	case "relocate_rebalancing":
		*r = RELOCATEREBALANCING
	case "unknown":
		*r = UNKNOWN
	default:
		*r = RecoveryPriority{string(text)}
	}

	return nil
}

func (r RecoveryPriority) String() string {
	return r.Name
}
