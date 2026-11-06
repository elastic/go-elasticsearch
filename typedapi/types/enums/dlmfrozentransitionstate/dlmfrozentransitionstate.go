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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

// Package dlmfrozentransitionstate
package dlmfrozentransitionstate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/health_report/types.ts#L281-L290
type DlmFrozenTransitionState struct {
	Name string
}

var (

	// Unmarked The index is eligible for a frozen-tier transition but the data stream
	// lifecycle has not marked it for conversion.
	Unmarked = DlmFrozenTransitionState{"unmarked"}

	// Marked The index is marked for conversion but has not been submitted to the
	// transition executor.
	Marked = DlmFrozenTransitionState{"marked"}

	// Queued The index has been submitted to the transition executor and is waiting in its
	// queue for a thread.
	Queued = DlmFrozenTransitionState{"queued"}

	// Running The transition of the index is in progress.
	Running = DlmFrozenTransitionState{"running"}
)

func (d DlmFrozenTransitionState) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DlmFrozenTransitionState) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "unmarked":
		*d = Unmarked
	case "marked":
		*d = Marked
	case "queued":
		*d = Queued
	case "running":
		*d = Running
	default:
		*d = DlmFrozenTransitionState{string(text)}
	}

	return nil
}

func (d DlmFrozenTransitionState) String() string {
	return d.Name
}
