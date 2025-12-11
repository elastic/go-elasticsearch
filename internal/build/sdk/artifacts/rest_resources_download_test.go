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
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v9/internal/build/sdk/ref"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/spf13/afero"
)

// createTestZip creates a zip file in memory with the given files
func createTestZip(files map[string]string) ([]byte, error) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	for name, content := range files {
		f, err := w.Create(name)
		if err != nil {
			return nil, err
		}
		if _, err := f.Write([]byte(content)); err != nil {
			return nil, err
		}
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func TestClient_DownloadRestResources(t *testing.T) {
	// Create test zip content
	testZipFiles := map[string]string{
		"rest-api-spec/api/search.json": `{"name": "search"}`,
		"rest-api-spec/api/index.json":  `{"name": "index"}`,
		"rest-api-spec/test/test1.yaml": "test: 1",
	}
	testZipData, err := createTestZip(testZipFiles)
	if err != nil {
		t.Fatalf("failed to create test zip: %v", err)
	}

	tests := []struct {
		name       string
		ref        string
		dest       string
		responses  map[string]string
		wantErr    bool
		wantErrMsg string
		wantFiles  []string
	}{
		{
			name: "success with snapshot",
			ref:  "main",
			dest: "/output",
			responses: map[string]string{
				"https://artifacts-snapshot.elastic.co/elasticsearch/latest/main.json": `{
					"version": "9.3.0-SNAPSHOT",
					"build_id": "9.3.0-abc123",
					"manifest_url": "https://artifacts-snapshot.elastic.co/elasticsearch/9.3.0-abc123/manifest-9.3.0-SNAPSHOT.json",
					"summary_url": "https://artifacts-snapshot.elastic.co/elasticsearch/9.3.0-abc123/summary.html"
				}`,
				"https://artifacts-snapshot.elastic.co/elasticsearch/9.3.0-abc123/manifest-9.3.0-SNAPSHOT.json": `{
					"branch": "main",
					"version": "9.3.0-SNAPSHOT",
					"build_id": "9.3.0-abc123",
					"start_time": "2025-05-22T12:43:21+00:00",
					"projects": {
						"elasticsearch": {
							"branch": "main",
							"commit_hash": "abc123def456",
							"packages": {
								"rest-resources-zip-9.3.0-SNAPSHOT.zip": {
									"url": "https://artifacts-snapshot.elastic.co/elasticsearch/rest-resources.zip",
									"type": "zip"
								}
							}
						}
					}
				}`,
			},
			wantErr: false,
			wantFiles: []string{
				"/output/rest-api-spec/api/search.json",
				"/output/rest-api-spec/api/index.json",
				"/output/rest-api-spec/test/test1.yaml",
				"/output/elasticsearch.json",
			},
		},
		{
			name: "success with release",
			ref:  "8.15.0",
			dest: "/output",
			responses: map[string]string{
				"https://artifacts.elastic.co/releases/stack.json": `{
					"releases": [
						{
							"version": "8.15.0",
							"manifest": "https://artifacts.elastic.co/releases/8.15.0/manifest.json"
						}
					]
				}`,
				"https://artifacts.elastic.co/releases/8.15.0/manifest.json": `{
					"branch": "8.15",
					"version": "8.15.0",
					"build_id": "8.15.0-release",
					"start_time": "2025-01-15T10:00:00+00:00",
					"projects": {
						"elasticsearch": {
							"branch": "8.15",
							"commit_hash": "release123",
							"packages": {
								"rest-resources-zip-8.15.0.zip": {
									"url": "https://artifacts.elastic.co/releases/rest-resources.zip",
									"type": "zip"
								}
							}
						}
					}
				}`,
			},
			wantErr: false,
			wantFiles: []string{
				"/output/rest-api-spec/api/search.json",
				"/output/rest-api-spec/api/index.json",
				"/output/rest-api-spec/test/test1.yaml",
				"/output/elasticsearch.json",
			},
		},
		{
			name:       "empty ref",
			ref:        "",
			dest:       "/output",
			wantErr:    true,
			wantErrMsg: "empty ref",
		},
		{
			name:       "empty destination",
			ref:        "main",
			dest:       "",
			wantErr:    true,
			wantErrMsg: "empty destination",
		},
		{
			name: "release not found",
			ref:  "99.99.99",
			dest: "/output",
			responses: map[string]string{
				"https://artifacts.elastic.co/releases/stack.json": `{
					"releases": [
						{
							"version": "8.15.0",
							"manifest": "https://artifacts.elastic.co/releases/8.15.0/manifest.json"
						}
					]
				}`,
			},
			wantErr:    true,
			wantErrMsg: "release 99.99.99 not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memFs := afero.NewMemMapFs()

			httpClient := retryablehttp.NewClient()
			httpClient.Logger = nil
			httpClient.RetryMax = 0
			httpClient.HTTPClient.Transport = &mockRoundTripper{
				RoundTripFunc: func(r *http.Request) (*http.Response, error) {
					url := r.URL.String()

					// Check for zip download
					if strings.HasSuffix(url, "rest-resources.zip") {
						return &http.Response{
							Request:    r.Clone(context.Background()),
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(bytes.NewReader(testZipData)),
						}, nil
					}

					// Check for JSON responses
					if response, ok := tt.responses[url]; ok {
						return &http.Response{
							Request:    r.Clone(context.Background()),
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(response)),
						}, nil
					}

					return &http.Response{
						Request:    r.Clone(context.Background()),
						StatusCode: http.StatusNotFound,
						Body:       io.NopCloser(strings.NewReader("Not Found")),
					}, nil
				},
			}

			client := &Client{
				releasesBaseURL: defaultReleasesBaseURL,
				snapshotBaseURL: defaultSnapshotBaseURL,
				client:          httpClient,
				fs:              memFs,
			}

			parsedRef, _ := ref.Parse(tt.ref)
			err := client.DownloadRestResources(context.Background(), parsedRef, tt.dest)

			if tt.wantErr {
				if err == nil {
					t.Errorf("DownloadRestResources() expected error, got nil")
					return
				}
				if tt.wantErrMsg != "" && !strings.Contains(err.Error(), tt.wantErrMsg) {
					t.Errorf("DownloadRestResources() error = %v, want error containing %q", err, tt.wantErrMsg)
				}
				return
			}

			if err != nil {
				t.Errorf("DownloadRestResources() unexpected error: %v", err)
				return
			}

			// Verify expected files exist
			for _, wantFile := range tt.wantFiles {
				exists, err := afero.Exists(memFs, wantFile)
				if err != nil {
					t.Errorf("error checking file %s: %v", wantFile, err)
					continue
				}
				if !exists {
					t.Errorf("expected file %s does not exist", wantFile)
				}
			}

			// Verify elasticsearch.json content
			if tt.dest != "" {
				elasticsearchJSON, err := afero.ReadFile(memFs, tt.dest+"/elasticsearch.json")
				if err != nil {
					t.Errorf("cannot read elasticsearch.json: %v", err)
					return
				}

				var build Build
				if err := json.Unmarshal(elasticsearchJSON, &build); err != nil {
					t.Errorf("cannot unmarshal elasticsearch.json: %v", err)
					return
				}

				if build.Version == "" {
					t.Error("elasticsearch.json has empty version")
				}
				if build.Projects.Elasticsearch.CommitHash == "" {
					t.Error("elasticsearch.json has empty commit hash")
				}
			}
		})
	}
}

func TestManifest_ToBuild(t *testing.T) {
	manifest := &Manifest{
		Branch:    "main",
		Version:   "9.3.0-SNAPSHOT",
		BuildID:   "9.3.0-abc123",
		StartTime: "2025-05-22T12:43:21+00:00",
		Projects: struct {
			Elasticsearch Project `json:"elasticsearch"`
		}{
			Elasticsearch: Project{
				Branch:     "main",
				CommitHash: "abc123def456",
				Packages: map[string]Package{
					"rest-resources.zip": {
						URL:  "https://example.com/rest-resources.zip",
						Type: "zip",
					},
				},
			},
		},
	}

	build := manifest.ToBuild()

	if build.Version != "9.3.0-SNAPSHOT" {
		t.Errorf("Version = %q, want %q", build.Version, "9.3.0-SNAPSHOT")
	}

	if build.BuildID != "9.3.0-abc123" {
		t.Errorf("BuildID = %q, want %q", build.BuildID, "9.3.0-abc123")
	}

	// Check start time is kept in RFC3339 format (matches original time.Time serialization)
	if build.StartTime != "2025-05-22T12:43:21Z" {
		t.Errorf("StartTime = %q, want %q", build.StartTime, "2025-05-22T12:43:21Z")
	}

	if build.Projects.Elasticsearch.Branch != "main" {
		t.Errorf("Branch = %q, want %q", build.Projects.Elasticsearch.Branch, "main")
	}

	if build.Projects.Elasticsearch.CommitHash != "abc123def456" {
		t.Errorf("CommitHash = %q, want %q", build.Projects.Elasticsearch.CommitHash, "abc123def456")
	}

	if len(build.Projects.Elasticsearch.Packages) != 1 {
		t.Errorf("Packages count = %d, want 1", len(build.Projects.Elasticsearch.Packages))
	}

	pkg, ok := build.Projects.Elasticsearch.Packages["rest-resources.zip"]
	if !ok {
		t.Error("rest-resources.zip package not found")
	} else {
		if pkg.URL != "https://example.com/rest-resources.zip" {
			t.Errorf("Package URL = %q, want %q", pkg.URL, "https://example.com/rest-resources.zip")
		}
		if pkg.Type != "zip" {
			t.Errorf("Package Type = %q, want %q", pkg.Type, "zip")
		}
	}
}
