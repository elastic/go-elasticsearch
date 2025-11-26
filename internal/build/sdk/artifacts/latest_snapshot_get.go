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
	"net/url"
)

// GetLatestSnapshotRequest represents the request parameters for the GetLatestSnapshot API.
type GetLatestSnapshotRequest struct {
	Branch string
}

// Validate checks the request parameters for errors.
func (r GetLatestSnapshotRequest) Validate() error {
	if len(r.Branch) == 0 {
		return fmt.Errorf("branch is required")
	}
	return nil
}

// GetLatestSnapshotResponse represents the response from the GetLatestSnapshot API.
type GetLatestSnapshotResponse struct {
	Version     string `json:"version"`
	BuildID     string `json:"build_id"`
	ManifestURL string `json:"manifest_url"`
	SummaryURL  string `json:"summary_url"`
}

// GetLatestSnapshot gets the latest snapshot build for a given branch.
func (c *Client) GetLatestSnapshot(ctx context.Context, req *GetLatestSnapshotRequest) (*GetLatestSnapshotResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	requestUrl, err := url.JoinPath(c.snapshotBaseURL, fmt.Sprintf("elasticsearch/latest/%s.json", req.Branch))
	if err != nil {
		return nil, err
	}

	return doGet[GetLatestSnapshotResponse](ctx, c, requestUrl)
}
