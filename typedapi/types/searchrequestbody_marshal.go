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

package types

import "encoding/json"

// MarshalJSON filters nil values from SearchAfter before encoding.
//
// Elasticsearch rejects a search_after array that contains a null element
// (returning a 500 "all shards failed" error). Callers that build sort
// cursors with a nil FieldValue to mean "no cursor" would otherwise trigger
// this error silently. Nil entries are stripped here so that:
//   - []FieldValue{nil}       → search_after field omitted entirely
//   - []FieldValue{nil, "x"}  → search_after: ["x"]
func (s SearchRequestBody) MarshalJSON() ([]byte, error) {
	// Filter nil values from SearchAfter.
	if len(s.SearchAfter) > 0 {
		filtered := make([]FieldValue, 0, len(s.SearchAfter))
		for _, v := range s.SearchAfter {
			if v != nil {
				filtered = append(filtered, v)
			}
		}
		if len(filtered) == 0 {
			s.SearchAfter = nil
		} else {
			s.SearchAfter = filtered
		}
	}

	// Use a type alias to invoke default JSON encoding without recursion.
	type Plain SearchRequestBody
	return json.Marshal(Plain(s))
}
