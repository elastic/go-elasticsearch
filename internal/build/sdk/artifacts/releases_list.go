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
