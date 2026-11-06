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

// Package recoverymode
package recoverymode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/data_recovery/start_recovery/types.ts#L20-L25
type RecoveryMode struct {
	Name string
}

var (

	// Replace Replace each source with its state at the recovery point, in the place where
	// it currently is.
	Replace = RecoveryMode{"replace"}

	// Rename Recover each source under a new name, which is derived from its name by the
	// `rename_strategy`.
	Rename = RecoveryMode{"rename"}
)

func (r RecoveryMode) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RecoveryMode) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "replace":
		*r = Replace
	case "rename":
		*r = Rename
	default:
		*r = RecoveryMode{string(text)}
	}

	return nil
}

func (r RecoveryMode) String() string {
	return r.Name
}
