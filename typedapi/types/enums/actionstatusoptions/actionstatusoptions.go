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
// https://github.com/elastic/elasticsearch-specification/tree/7fd0bd13eaf28bd179fc906f57da09e852eb818e

// Package actionstatusoptions
package actionstatusoptions

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7fd0bd13eaf28bd179fc906f57da09e852eb818e/specification/watcher/_types/Action.ts#L96-L108
type ActionStatusOptions struct {
	Name string
}

var (
	Success = ActionStatusOptions{"success"}

	Failure = ActionStatusOptions{"failure"}

	Partialfailure = ActionStatusOptions{"partial_failure"}

	Acknowledged = ActionStatusOptions{"acknowledged"}

	Throttled = ActionStatusOptions{"throttled"}

	Conditionfailed = ActionStatusOptions{"condition_failed"}

	Simulated = ActionStatusOptions{"simulated"}
)

func (a ActionStatusOptions) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *ActionStatusOptions) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "success":
		*a = Success
	case "failure":
		*a = Failure
	case "partial_failure":
		*a = Partialfailure
	case "acknowledged":
		*a = Acknowledged
	case "throttled":
		*a = Throttled
	case "condition_failed":
		*a = Conditionfailed
	case "simulated":
		*a = Simulated
	default:
		*a = ActionStatusOptions{string(text)}
	}

	return nil
}

func (a ActionStatusOptions) String() string {
	return a.Name
}
