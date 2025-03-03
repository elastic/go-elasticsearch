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

import "testing"

func ptr[T any](t T) *T { return &t }

func TestElasticsearchError(t *testing.T) {
	tests := []struct {
		in   ElasticsearchError
		want string
	}{
		{
			ElasticsearchError{
				Status: 400,
				ErrorCause: ErrorCause{
					Reason: ptr("[1:77] [bool] failed to parse field [filter]"),
					Type:   "x_content_parse_exception",
					CausedBy: &ErrorCause{
						Reason: ptr("[term] query doesn't support multiple fields, found [collection_id] and [entities]"),
						Type:   "parsing_exception",
					},
					RootCause: []ErrorCause{{
						Reason: ptr("[term] query doesn't support multiple fields, found [collection_id] and [entities]"),
						Type:   "parsing_exception",
					}},
				},
			},
			"status: 400, failed: [x_content_parse_exception], reason: [1:77] [bool] failed to parse field [filter]: [term] query doesn't support multiple fields, found [collection_id] and [entities]"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			have := tt.in.Error()
			if have != tt.want {
				t.Errorf("\nhave: %q\nwant: %q", have, tt.want)
			}
		})
	}
}
