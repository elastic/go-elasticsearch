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
	"fmt"
	"strings"
)

// GetManifestRequest represents the request parameters for the GetManifest API.
type GetManifestRequest struct {
	URL string
}

// Validate checks the request parameters for errors.
func (r GetManifestRequest) Validate() error {
	if len(r.URL) == 0 {
		return fmt.Errorf("url is required")
	}

	return nil
}

// Package represents a package in the manifest.
type Package struct {
	URL          string   `json:"url"`
	ShaURL       string   `json:"sha_url"`
	Type         string   `json:"type"`
	Architecture string   `json:"architecture"`
	OS           []string `json:"os"`
}

// Dependency represents a dependency in the manifest.
type Dependency struct {
	Prefix   string `json:"prefix"`
	BuildURI string `json:"build_uri"`
}

// Project represents a project in the manifest.
type Project struct {
	Branch               string             `json:"branch"`
	CommitHash           string             `json:"commit_hash"`
	CommitURL            string             `json:"commit_url"`
	BuildDurationSeconds int                `json:"build_duration_seconds"`
	Packages             map[string]Package `json:"packages"`
	Dependencies         []Dependency       `json:"dependencies"`
}

// Manifest represents the manifest for a build.
type Manifest struct {
	Branch               string `json:"branch"`
	ReleaseBranch        string `json:"release_branch"`
	Version              string `json:"version"`
	BuildID              string `json:"build_id"`
	StartTime            string `json:"start_time"`
	EndTime              string `json:"end_time"`
	BuildDurationSeconds int    `json:"build_duration_seconds"`
	ManifestVersion      string `json:"manifest_version"`
	Prefix               string `json:"prefix"`
	Projects             struct {
		Elasticsearch Project `json:"elasticsearch"`
	} `json:"projects"`
}

// GetManifestResponse represents the response from the GetManifest API.
type GetManifestResponse struct {
	Manifest *Manifest
}

// Build represents the legacy build format for backward compatibility with elasticsearch.json
type Build struct {
	StartTime string `json:"start_time"`
	Version   string `json:"version"`
	BuildID   string `json:"build_id"`
	Projects  struct {
		Elasticsearch BuildProject `json:"elasticsearch"`
	} `json:"projects"`
}

// BuildProject represents the legacy build project format for backward compatibility with elasticsearch.json
type BuildProject struct {
	Branch     string                  `json:"branch"`
	CommitHash string                  `json:"commit_hash"`
	Packages   map[string]BuildPackage `json:"packages"`
}

// BuildPackage represents the legacy build package format for backward compatibility with elasticsearch.json
type BuildPackage struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

// ToBuild converts a Manifest to the legacy Build format for backward compatibility
func (m *Manifest) ToBuild() Build {
	var build Build

	build.StartTime = m.StartTime
	build.Version = m.Version
	build.BuildID = m.BuildID
	build.Projects.Elasticsearch.Branch = m.Projects.Elasticsearch.Branch
	build.Projects.Elasticsearch.CommitHash = m.Projects.Elasticsearch.CommitHash
	build.Projects.Elasticsearch.Packages = make(map[string]BuildPackage)

	for key, pkg := range m.Projects.Elasticsearch.Packages {
		build.Projects.Elasticsearch.Packages[key] = BuildPackage{
			URL:  pkg.URL,
			Type: pkg.Type,
		}
	}

	return build
}

// GetManifest retrieves the manifest for a build.
func (c *Client) GetManifest(ctx context.Context, req *GetManifestRequest) (*GetManifestResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if !strings.HasPrefix(req.URL, c.snapshotBaseURL) && !strings.HasPrefix(req.URL, c.releasesBaseURL) {
		return nil, fmt.Errorf("invalid url: %s, must be child of: %s or %s", req.URL, c.snapshotBaseURL, c.releasesBaseURL)
	}

	manifest, err := doGet[Manifest](ctx, c, req.URL)
	if err != nil {
		return nil, err
	}
	return &GetManifestResponse{Manifest: manifest}, nil
}
