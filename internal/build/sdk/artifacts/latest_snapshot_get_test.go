package artifacts

import (
	"context"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/go-retryablehttp"
)

type mockRoundTripper struct {
	RoundTripFunc func(*http.Request) (*http.Response, error)
}

func (m mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.RoundTripFunc == nil {
		return nil, errors.New("the RoundTripFunc of the mock is nil")
	}
	return m.RoundTripFunc(req)
}

func TestClient_GetLatestInfo(t *testing.T) {
	type fields struct {
		baseURL string
		client  *retryablehttp.Client
	}
	type args struct {
		ctx context.Context
		req *GetLatestSnapshotRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetLatestSnapshotResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				baseURL: defaultSnapshotBaseURL,
				client: func() *retryablehttp.Client {
					c := retryablehttp.NewClient()
					c.Logger = nil
					c.HTTPClient.Transport = &mockRoundTripper{
						RoundTripFunc: func(r *http.Request) (*http.Response, error) {
							return &http.Response{
								Request:    r.Clone(context.Background()),
								StatusCode: http.StatusOK,
								Body: io.NopCloser(strings.NewReader(`{
    "version": "9.3.0-SNAPSHOT",
	"build_id": "9.3.0-dd19c56f",
	"manifest_url": "https://artifacts-snapshot.elastic.co/elasticsearch/9.3.0-dd19c56f/manifest-9.3.0-SNAPSHOT.json",
	"summary_url": "https://artifacts-snapshot.elastic.co/elasticsearch/9.3.0-dd19c56f/summary-9.3.0-SNAPSHOT.html"
}`)),
							}, nil
						},
					}
					return c
				}(),
			},
			args: args{
				ctx: context.Background(),
				req: &GetLatestSnapshotRequest{
					Branch: "master",
				},
			},
			wantErr: false,
			want: &GetLatestSnapshotResponse{
				Version:     "9.3.0-SNAPSHOT",
				BuildID:     "9.3.0-dd19c56f",
				ManifestURL: "https://artifacts-snapshot.elastic.co/elasticsearch/9.3.0-dd19c56f/manifest-9.3.0-SNAPSHOT.json",
				SummaryURL:  "https://artifacts-snapshot.elastic.co/elasticsearch/9.3.0-dd19c56f/summary-9.3.0-SNAPSHOT.html",
			},
		},
		{
			name: "404",
			fields: fields{
				baseURL: defaultSnapshotBaseURL,
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
				req: &GetLatestSnapshotRequest{
					Branch: "master",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				snapshotBaseURL: tt.fields.baseURL,
				client:          tt.fields.client,
			}
			got, err := c.GetLatestSnapshot(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatestSnapshot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLatestSnapshot() got = %v, want %v", got, tt.want)
			}
		})
	}
}
