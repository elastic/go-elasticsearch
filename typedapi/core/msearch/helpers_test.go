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

package msearch

import (
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

func TestMsearch_AddSearch(t *testing.T) {
	type args struct {
		header types.MultisearchHeader
		body   types.MultisearchBody
	}
	tests := []struct {
		name    string
		args    args
		expects string
		wantErr bool
	}{
		{
			name: "nominal test case",
			args: args{
				header: types.MultisearchHeader{Index: []string{"foo"}},
				body:   types.MultisearchBody{Query: &types.Query{Term: map[string]types.TermQuery{"bar": {Value: "baz"}}}},
			},
			expects: `{"index":["foo"]}
{"query":{"term":{"bar":{"value":"baz"}}}}
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New(nil)
			if err := r.AddSearch(tt.args.header, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("AddSearch() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !strings.EqualFold(r.buf.String(), tt.expects) {
				t.Errorf("wrong body")
			}
		})
	}
}
