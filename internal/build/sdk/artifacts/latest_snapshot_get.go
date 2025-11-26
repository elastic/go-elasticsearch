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
