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

// Package enrichpolicyphase
package enrichpolicyphase

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/enrich/execute_policy/types.ts#L24-L29
type EnrichPolicyPhase struct {
	Name string
}

var (
	SCHEDULED = EnrichPolicyPhase{"SCHEDULED"}

	RUNNING = EnrichPolicyPhase{"RUNNING"}

	COMPLETE = EnrichPolicyPhase{"COMPLETE"}

	FAILED = EnrichPolicyPhase{"FAILED"}
)

func (e EnrichPolicyPhase) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *EnrichPolicyPhase) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "SCHEDULED":
		*e = SCHEDULED
	case "RUNNING":
		*e = RUNNING
	case "COMPLETE":
		*e = COMPLETE
	case "FAILED":
		*e = FAILED
	default:
		*e = EnrichPolicyPhase{string(text)}
	}

	return nil
}

func (e EnrichPolicyPhase) String() string {
	return e.Name
}
