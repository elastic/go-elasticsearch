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

package artifacts

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
)

func TestClient_GetManifest(t *testing.T) {
	type fields struct {
		snapshotBaseURL string
		releasesBaseURL string
		client          *retryablehttp.Client
	}
	type args struct {
		ctx context.Context
		req *GetManifestRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetManifestResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				snapshotBaseURL: defaultSnapshotBaseURL,
				releasesBaseURL: defaultReleasesBaseURL,
				client: func() *retryablehttp.Client {
					c := retryablehttp.NewClient()
					c.Logger = nil
					c.HTTPClient.Transport = &mockRoundTripper{
						RoundTripFunc: func(r *http.Request) (*http.Response, error) {
							return &http.Response{
								Request:    r.Clone(context.Background()),
								StatusCode: http.StatusOK,
								Body: io.NopCloser(strings.NewReader(`{
	"branch": "main",
	"release_branch": "main",
	"version": "9.3.0-SNAPSHOT",
	"build_id": "9.3.0-abc123",
	"start_time": "2025-05-22T12:43:21+00:00",
	"end_time": "2025-05-22T13:00:00+00:00",
	"build_duration_seconds": 1000,
	"manifest_version": "1.0",
	"prefix": "elasticsearch",
	"projects": {
		"elasticsearch": {
			"branch": "main",
			"commit_hash": "abc123def456",
			"commit_url": "https://github.com/elastic/elasticsearch/commit/abc123def456",
			"build_duration_seconds": 500,
			"packages": {
				"rest-resources-zip-9.3.0-SNAPSHOT.zip": {
					"url": "https://artifacts-snapshot.elastic.co/elasticsearch/rest-resources.zip",
					"sha_url": "https://artifacts-snapshot.elastic.co/elasticsearch/rest-resources.zip.sha512",
					"type": "zip",
					"architecture": "any",
					"os": ["any"]
				}
			},
			"dependencies": []
		}
	}
}`)),
							}, nil
						},
					}
					return c
				}(),
			},
			args: args{
				ctx: context.Background(),
				req: &GetManifestRequest{
					URL: "https://artifacts-snapshot.elastic.co/elasticsearch/9.3.0-abc123/manifest.json",
				},
			},
			wantErr: false,
			want: &GetManifestResponse{
				Manifest: &Manifest{
					Branch:               "main",
					ReleaseBranch:        "main",
					Version:              "9.3.0-SNAPSHOT",
					BuildID:              "9.3.0-abc123",
					StartTime:            "2025-05-22T12:43:21+00:00",
					EndTime:              "2025-05-22T13:00:00+00:00",
					BuildDurationSeconds: 1000,
					ManifestVersion:      "1.0",
					Prefix:               "elasticsearch",
					Projects: struct {
						Elasticsearch Project `json:"elasticsearch"`
					}{
						Elasticsearch: Project{
							Branch:               "main",
							CommitHash:           "abc123def456",
							CommitURL:            "https://github.com/elastic/elasticsearch/commit/abc123def456",
							BuildDurationSeconds: 500,
							Packages: map[string]Package{
								"rest-resources-zip-9.3.0-SNAPSHOT.zip": {
									URL:          "https://artifacts-snapshot.elastic.co/elasticsearch/rest-resources.zip",
									ShaURL:       "https://artifacts-snapshot.elastic.co/elasticsearch/rest-resources.zip.sha512",
									Type:         "zip",
									Architecture: "any",
									OS:           []string{"any"},
								},
							},
							Dependencies: []Dependency{},
						},
					},
				},
			},
		},
		{
			name: "404",
			fields: fields{
				snapshotBaseURL: defaultSnapshotBaseURL,
				releasesBaseURL: defaultReleasesBaseURL,
				client: func() *retryablehttp.Client {
					c := retryablehttp.NewClient()
					c.Logger = nil
					c.HTTPClient.Transport = &mockRoundTripper{
						RoundTripFunc: func(r *http.Request) (*http.Response, error) {
							return &http.Response{
								Request:    r.Clone(context.Background()),
								StatusCode: http.StatusNotFound,
								Body:       io.NopCloser(strings.NewReader(`Not Found`)),
							}, nil
						},
					}
					return c
				}(),
			},
			args: args{
				ctx: context.Background(),
				req: &GetManifestRequest{
					URL: "https://artifacts-snapshot.elastic.co/elasticsearch/not-found/manifest.json",
				},
			},
			wantErr: true,
		},
		{
			name: "validation error - empty URL",
			fields: fields{
				snapshotBaseURL: defaultSnapshotBaseURL,
				releasesBaseURL: defaultReleasesBaseURL,
				client:          retryablehttp.NewClient(),
			},
			args: args{
				ctx: context.Background(),
				req: &GetManifestRequest{
					URL: "",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid URL - not child of base URLs",
			fields: fields{
				snapshotBaseURL: defaultSnapshotBaseURL,
				releasesBaseURL: defaultReleasesBaseURL,
				client:          retryablehttp.NewClient(),
			},
			args: args{
				ctx: context.Background(),
				req: &GetManifestRequest{
					URL: "https://malicious-site.com/manifest.json",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				snapshotBaseURL: tt.fields.snapshotBaseURL,
				releasesBaseURL: tt.fields.releasesBaseURL,
				client:          tt.fields.client,
			}
			got, err := c.GetManifest(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetManifest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetManifest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
