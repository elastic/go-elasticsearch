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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Package allocationexplaindecision
package allocationexplaindecision

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/allocation_explain/types.ts#L32-L37
type AllocationExplainDecision struct {
	Name string
}

var (
	NO = AllocationExplainDecision{"NO"}

	YES = AllocationExplainDecision{"YES"}

	THROTTLE = AllocationExplainDecision{"THROTTLE"}

	ALWAYS = AllocationExplainDecision{"ALWAYS"}
)

func (a AllocationExplainDecision) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *AllocationExplainDecision) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "NO":
		*a = NO
	case "YES":
		*a = YES
	case "THROTTLE":
		*a = THROTTLE
	case "ALWAYS":
		*a = ALWAYS
	default:
		*a = AllocationExplainDecision{string(text)}
	}

	return nil
}

func (a AllocationExplainDecision) String() string {
	return a.Name
}
