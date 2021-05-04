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

package cmd

import (
	"encoding/json"
	"testing"
)

func TestBuild_zipfileUrl(t *testing.T) {
	tests := []struct {
		name   string
		json   string
		want   string
	}{
		{
			name: "Simple test with valid url",
			json: `
			{
			  "start_time": "Mon, 19 Apr 2021 12:15:47 GMT",
			  "version": "8.0.0-SNAPSHOT",
			  "build_id": "8.0.0-ab7cd914",
			  "projects": {
				"elasticsearch": {
				  "branch": "master",
				  "commit_hash": "d3be79018b5b70a118ea5a897a539428b728df5a",
				  "Packages": {
					"rest-resources-zip-8.0.0-SNAPSHOT.zip": {
					  "url": "https://snapshots.elastic.co/8.0.0-ab7cd914/downloads/elasticsearch/rest-resources-zip-8.0.0-SNAPSHOT.zip",
					  "type": "zip"
					}
				  }
				}
			  }
			}
			`,
			want: "https://snapshots.elastic.co/8.0.0-ab7cd914/downloads/elasticsearch/rest-resources-zip-8.0.0-SNAPSHOT.zip",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Build{}
			if err := json.Unmarshal([]byte(tt.json), &b); err != nil {
				t.Fatalf(err.Error())
			}
			if got := b.zipfileUrl(); got != tt.want {
				t.Errorf("zipfileUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}