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

func ptr[T any](v T) *T {
	return &v
}

func TestClient_ListReleases(t *testing.T) {
	type fields struct {
		releasesBaseURL string
		client          *retryablehttp.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ListReleasesResponse
		wantErr bool
	}{
		{
			name: "success with multiple releases",
			fields: fields{
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
	"releases": [
		{
			"version": "8.15.0",
			"public_release_date": "2024-08-08",
			"is_end_of_support": false,
			"end_of_support_date": "2026-02-08",
			"is_end_of_maintenance": false,
			"end_of_maintenance_date": "2025-02-08",
			"is_retired": false,
			"retired_date": null,
			"manifest": "https://artifacts.elastic.co/releases/8.15.0/manifest.json"
		},
		{
			"version": "8.14.3",
			"public_release_date": "2024-07-10",
			"is_end_of_support": false,
			"end_of_support_date": "2026-01-10",
			"is_end_of_maintenance": true,
			"end_of_maintenance_date": "2024-08-08",
			"is_retired": false,
			"retired_date": null,
			"manifest": "https://artifacts.elastic.co/releases/8.14.3/manifest.json"
		},
		{
			"version": "7.17.22",
			"public_release_date": "2024-06-01",
			"is_end_of_support": true,
			"end_of_support_date": "2023-08-01",
			"is_end_of_maintenance": true,
			"end_of_maintenance_date": "2023-02-01",
			"is_retired": true,
			"retired_date": "2024-01-01",
			"manifest": null
		}
	]
}`)),
							}, nil
						},
					}
					return c
				}(),
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
			want: &ListReleasesResponse{
				Releases: []Release{
					{
						Version:              "8.15.0",
						PublicReleaseDate:    ptr("2024-08-08"),
						IsEndOfSupport:       ptr(false),
						EndOfSupportDate:     ptr("2026-02-08"),
						IsEndOfMaintenance:   ptr(false),
						EndOfMaintenanceDate: ptr("2025-02-08"),
						IsRetired:            ptr(false),
						RetiredDate:          nil,
						Manifest:             ptr("https://artifacts.elastic.co/releases/8.15.0/manifest.json"),
					},
					{
						Version:              "8.14.3",
						PublicReleaseDate:    ptr("2024-07-10"),
						IsEndOfSupport:       ptr(false),
						EndOfSupportDate:     ptr("2026-01-10"),
						IsEndOfMaintenance:   ptr(true),
						EndOfMaintenanceDate: ptr("2024-08-08"),
						IsRetired:            ptr(false),
						RetiredDate:          nil,
						Manifest:             ptr("https://artifacts.elastic.co/releases/8.14.3/manifest.json"),
					},
					{
						Version:              "7.17.22",
						PublicReleaseDate:    ptr("2024-06-01"),
						IsEndOfSupport:       ptr(true),
						EndOfSupportDate:     ptr("2023-08-01"),
						IsEndOfMaintenance:   ptr(true),
						EndOfMaintenanceDate: ptr("2023-02-01"),
						IsRetired:            ptr(true),
						RetiredDate:          ptr("2024-01-01"),
						Manifest:             nil,
					},
				},
			},
		},
		{
			name: "404",
			fields: fields{
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
			},
			wantErr: true,
		},
		{
			name: "empty releases array",
			fields: fields{
				releasesBaseURL: defaultReleasesBaseURL,
				client: func() *retryablehttp.Client {
					c := retryablehttp.NewClient()
					c.Logger = nil
					c.HTTPClient.Transport = &mockRoundTripper{
						RoundTripFunc: func(r *http.Request) (*http.Response, error) {
							return &http.Response{
								Request:    r.Clone(context.Background()),
								StatusCode: http.StatusOK,
								Body:       io.NopCloser(strings.NewReader(`{"releases": []}`)),
							}, nil
						},
					}
					return c
				}(),
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
			want: &ListReleasesResponse{
				Releases: []Release{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				releasesBaseURL: tt.fields.releasesBaseURL,
				client:          tt.fields.client,
			}
			got, err := c.ListReleases(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListReleases() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListReleases() got = %v, want %v", got, tt.want)
			}
		})
	}
}
