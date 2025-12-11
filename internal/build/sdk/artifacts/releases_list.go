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
	"net/url"
)

// Release represents a release in the Stack API.
type Release struct {
	Version              string  `json:"version"`
	PublicReleaseDate    *string `json:"public_release_date"`
	IsEndOfSupport       *bool   `json:"is_end_of_support"`
	EndOfSupportDate     *string `json:"end_of_support_date"`
	IsEndOfMaintenance   *bool   `json:"is_end_of_maintenance"`
	EndOfMaintenanceDate *string `json:"end_of_maintenance_date"`
	IsRetired            *bool   `json:"is_retired"`
	RetiredDate          *string `json:"retired_date"`
	Manifest             *string `json:"manifest"`
}

// ListReleasesResponse represents the response from the ListReleases API.
type ListReleasesResponse struct {
	Releases []Release `json:"releases"`
}

// ListReleases lists all releases.
func (c *Client) ListReleases(ctx context.Context) (*ListReleasesResponse, error) {
	requestUrl, err := url.JoinPath(c.releasesBaseURL, "releases", "stack.json")
	if err != nil {
		return nil, err
	}

	return doGet[ListReleasesResponse](ctx, c, requestUrl)
}
