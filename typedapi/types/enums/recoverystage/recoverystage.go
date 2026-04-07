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
// https://github.com/elastic/elasticsearch-specification/tree/836fca874204ca4173ae5c36fb6b5107d28d2fc0

// Package recoverystage
package recoverystage

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/836fca874204ca4173ae5c36fb6b5107d28d2fc0/specification/indices/recovery/types.ts#L118-L131
type RecoveryStage struct {
	Name string
}

var (

	// INIT Recovery has not started.
	INIT = RecoveryStage{"INIT"}

	// INDEX Reading index metadata and copying bytes from source to destination.
	INDEX = RecoveryStage{"INDEX"}

	// VERIFYINDEX Verifying the integrity of the index.
	VERIFYINDEX = RecoveryStage{"VERIFY_INDEX"}

	// TRANSLOG Replaying the transaction log.
	TRANSLOG = RecoveryStage{"TRANSLOG"}

	// FINALIZE Cleanup.
	FINALIZE = RecoveryStage{"FINALIZE"}

	// DONE Complete.
	DONE = RecoveryStage{"DONE"}
)

func (r RecoveryStage) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RecoveryStage) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "init":
		*r = INIT
	case "index":
		*r = INDEX
	case "verify_index":
		*r = VERIFYINDEX
	case "translog":
		*r = TRANSLOG
	case "finalize":
		*r = FINALIZE
	case "done":
		*r = DONE
	default:
		*r = RecoveryStage{string(text)}
	}

	return nil
}

func (r RecoveryStage) String() string {
	return r.Name
}
