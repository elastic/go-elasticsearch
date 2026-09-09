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

//go:build !integration
// +build !integration

package esapi

import "testing"

func TestEscapePathPart(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"plain", "plain", "plain"},
		{"slash", "a/b", "a%2Fb"},
		{"date-math", "<my-index-{now/d}>", "%3Cmy-index-%7Bnow%2Fd%7D%3E"},
		{"commas preserved", "idx1,idx2", "idx1,idx2"},
		{"slash in list elements", "a/b,c/d", "a%2Fb,c%2Fd"},
		{"date-math inside list", "<idx-{now/d}>,plain", "%3Cidx-%7Bnow%2Fd%7D%3E,plain"},
		{"already encoded double-escapes", "%2F", "%252F"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := escapePathPart(tc.in)
			if got != tc.want {
				t.Errorf("escapePathPart(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
