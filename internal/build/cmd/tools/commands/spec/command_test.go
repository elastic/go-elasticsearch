// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

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