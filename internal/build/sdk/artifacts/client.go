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
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/spf13/afero"
)

const (
	defaultReleasesBaseURL = "https://elastic-release-api.s3.us-west-2.amazonaws.com/"
	defaultSnapshotBaseURL = "https://artifacts-snapshot.elastic.co/"
)

type config struct {
	releasesBaseURL string
	snapshotBaseURL string
	fs              afero.Fs
}

// Client is a client for interacting with the Elastic Artifacts API.
type Client struct {
	releasesBaseURL string
	snapshotBaseURL string
	client          *retryablehttp.Client
	fs              afero.Fs
}

// Option is a functional option for configuring a Client.
type Option func(*config)

// WithReleasesBaseURL sets the base URL for release artifacts.
func WithReleasesBaseURL(url string) Option {
	return func(c *config) {
		c.releasesBaseURL = url
	}
}

// WithSnapshotBaseURL sets the base URL for snapshot artifacts.
func WithSnapshotBaseURL(url string) Option {
	return func(c *config) {
		c.snapshotBaseURL = url
	}
}

// WithFS sets the filesystem to use for downloading artifacts.
// Defaults to afero.NewOsFs().
// This is useful for testing.
func WithFS(fs afero.Fs) Option {
	return func(c *config) {
		c.fs = fs
	}
}

// NewClient returns a new Client.
func NewClient(options ...Option) *Client {
	c := &config{
		releasesBaseURL: defaultReleasesBaseURL,
		snapshotBaseURL: defaultSnapshotBaseURL,
		fs:              afero.NewOsFs(),
	}
	for _, opt := range options {
		opt(c)
	}
	return &Client{
		releasesBaseURL: c.releasesBaseURL,
		snapshotBaseURL: c.snapshotBaseURL,
		client:          retryablehttp.NewClient(),
		fs:              c.fs,
	}
}

func doGet[T any](ctx context.Context, c *Client, url string) (*T, error) {
	request, err := retryablehttp.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var response T
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

func doDownload(ctx context.Context, c *Client, url string) ([]byte, error) {
	request, err := retryablehttp.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept-Encoding", "gzip")
	resp, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
