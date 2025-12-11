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
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/elastic/go-elasticsearch/v9/internal/build/sdk/ref"
)

// DownloadRestResources downloads the rest-resources zip file for the given Elasticsearch version
// and extracts it to the destination directory.
// It also writes elasticsearch.json with build info for backward compatibility.
func (c *Client) DownloadRestResources(ctx context.Context, ref ref.Ref, dest string) error {
	if ref.IsEmpty() {
		return errors.New("empty ref")
	}
	if dest == "" {
		return errors.New("empty destination")
	}

	var manifestURL string

	// First, check if the latest snapshot at master matches our ref version
	latestMasterSnapshot, err := c.GetLatestSnapshot(ctx, &GetLatestSnapshotRequest{Branch: "master"})
	if err == nil && latestMasterSnapshot.Version == ref.String() {
		manifestURL = latestMasterSnapshot.ManifestURL
	} else {
		// Try to find in releases using the base version (without prerelease suffix)
		baseVersion := ref.BaseVersion()
		releases, err := c.ListReleases(ctx)
		if err != nil {
			return fmt.Errorf("cannot list releases: %w", err)
		}

		var release *Release
		for i := range releases.Releases {
			if releases.Releases[i].Version == baseVersion {
				release = &releases.Releases[i]
				break
			}
		}

		if release != nil && release.Manifest != nil {
			manifestURL = *release.Manifest
		} else if ref.IsPreRelease() {
			// Fall back to getting snapshot from target branch
			latestSnapshot, err := c.GetLatestSnapshot(ctx, &GetLatestSnapshotRequest{Branch: ref.TargetBranch()})
			if err != nil {
				return fmt.Errorf("cannot get latest snapshot: %w", err)
			}
			manifestURL = latestSnapshot.ManifestURL
		} else {
			return fmt.Errorf("release %s not found", ref.String())
		}
	}

	manifest, err := c.GetManifest(ctx, &GetManifestRequest{URL: manifestURL})
	if err != nil {
		return fmt.Errorf("cannot get manifest: %w", err)
	}

	elasticsearchProject := manifest.Manifest.Projects.Elasticsearch

	var zipURL string

	for key, value := range elasticsearchProject.Packages {
		if strings.Contains(key, "rest-resources") {
			zipURL = value.URL
			break
		}
	}

	if zipURL == "" {
		return errors.New("rest-resources package not found in manifest")
	}

	// Download the zip file
	zipData, err := doDownload(ctx, c, zipURL)
	if err != nil {
		return fmt.Errorf("cannot download rest-resources zip: %w", err)
	}

	// Extract zip to destination
	if err := c.extractZipToDestination(zipData, dest); err != nil {
		return fmt.Errorf("cannot extract zip: %w", err)
	}

	// Write elasticsearch.json with build info for backward compatibility
	build := manifest.Manifest.ToBuild()
	buildJSON, err := json.MarshalIndent(build, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal build info: %w", err)
	}

	if err := c.writeFile(filepath.Join(dest, "elasticsearch.json"), buildJSON); err != nil {
		return fmt.Errorf("cannot write elasticsearch.json: %w", err)
	}

	return nil
}

// extractZipToDestination extracts zip data to the destination directory using the client's filesystem
func (c *Client) extractZipToDestination(data []byte, dest string) error {
	zipReader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return fmt.Errorf("cannot read zip: %w", err)
	}

	if err := c.fs.MkdirAll(dest, 0755); err != nil {
		return fmt.Errorf("cannot create destination directory: %w", err)
	}

	for _, file := range zipReader.File {
		destPath := filepath.Join(dest, file.Name)

		if file.FileInfo().IsDir() {
			if err := c.fs.MkdirAll(destPath, 0755); err != nil {
				return fmt.Errorf("cannot create directory %s: %w", destPath, err)
			}
			continue
		}

		// Ensure parent directory exists
		if err := c.fs.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return fmt.Errorf("cannot create parent directory for %s: %w", destPath, err)
		}

		if err := c.extractZipFile(file, destPath); err != nil {
			return err
		}
	}

	return nil
}

// extractZipFile extracts a single file from the zip archive
func (c *Client) extractZipFile(file *zip.File, destPath string) error {
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("cannot open file in zip: %w", err)
	}
	defer func() { _ = src.Close() }()

	dst, err := c.fs.Create(destPath)
	if err != nil {
		return fmt.Errorf("cannot create file %s: %w", destPath, err)
	}
	defer func() { _ = dst.Close() }()

	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("cannot write file %s: %w", destPath, err)
	}

	return nil
}

// writeFile writes data to a file using the client's filesystem
func (c *Client) writeFile(path string, data []byte) error {
	f, err := c.fs.Create(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = f.Write(data)
	return err
}
