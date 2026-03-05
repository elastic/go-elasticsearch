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

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestSearchRequestBody_MarshalJSON_SearchAfterNil(t *testing.T) {
	tests := []struct {
		name        string
		searchAfter []FieldValue
		wantInJSON  string
		wantAbsent  string
	}{
		{
			name:        "nil slice omitted",
			searchAfter: nil,
			wantAbsent:  "search_after",
		},
		{
			name:        "empty slice omitted",
			searchAfter: []FieldValue{},
			wantAbsent:  "search_after",
		},
		{
			name:        "all-nil slice omitted",
			searchAfter: []FieldValue{nil, nil},
			wantAbsent:  "search_after",
		},
		{
			name:        "nil element stripped",
			searchAfter: []FieldValue{nil, "cursor-value"},
			wantInJSON:  `"search_after":["cursor-value"]`,
		},
		{
			name:        "non-nil values preserved",
			searchAfter: []FieldValue{"a", int64(42), true},
			wantInJSON:  `"search_after":["a",42,true]`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body := SearchRequestBody{SearchAfter: tc.searchAfter}
			data, err := json.Marshal(body)
			if err != nil {
				t.Fatalf("unexpected marshal error: %s", err)
			}
			got := string(data)

			if tc.wantInJSON != "" && !strings.Contains(got, tc.wantInJSON) {
				t.Errorf("expected JSON to contain %q, got: %s", tc.wantInJSON, got)
			}
			if tc.wantAbsent != "" && strings.Contains(got, tc.wantAbsent) {
				t.Errorf("expected JSON NOT to contain %q, got: %s", tc.wantAbsent, got)
			}
		})
	}
}
