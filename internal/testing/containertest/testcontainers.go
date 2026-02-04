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

//go:build integration && !multinode
// +build integration,!multinode

package containertest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/testcontainers/testcontainers-go"
	tces "github.com/testcontainers/testcontainers-go/modules/elasticsearch"
)

const (
	// Docker registry URLs for Elastic
	elasticDockerAuthURL = "https://docker-auth.elastic.co/auth"
	elasticDockerAPIURL  = "https://docker.elastic.co/v2"
	elasticESRepository  = "elasticsearch/elasticsearch"
)

// ElasticsearchService represents an Elasticsearch service, storing
// the container and configuration options.
type ElasticsearchService struct {
	container     *tces.ElasticsearchContainer
	configOptions tces.Options
}

// ESConfig returns the Elasticsearch client configuration.
func (es *ElasticsearchService) ESConfig() elasticsearch.Config {
	return elasticsearch.Config{
		Addresses: []string{
			es.configOptions.Address,
		},
		Username: "elastic",
		Password: es.configOptions.Password,
		CACert:   es.configOptions.CACert,
	}
}

// Terminate terminates the Elasticsearch container.
func (es *ElasticsearchService) Terminate(ctx context.Context) error {
	return es.container.Terminate(ctx)
}

// elasticsearchServiceConfig holds configuration for creating an ElasticsearchService.
type elasticsearchServiceConfig struct {
	resolveLatestPatch bool
}

// ElasticsearchServiceOption is a functional option for configuring ElasticsearchService.
type ElasticsearchServiceOption func(*elasticsearchServiceConfig)

// WithResolveLatestPatch enables automatic resolution of the latest available
// patch version for the given major.minor version from the Docker registry.
// When enabled, if the exact version doesn't exist, it will find and use
// the latest patch version that has an available Docker image.
func WithResolveLatestPatch(resolve bool) ElasticsearchServiceOption {
	return func(c *elasticsearchServiceConfig) {
		c.resolveLatestPatch = resolve
	}
}

// NewElasticsearchService creates a new Elasticsearch container service.
// The stackVersion parameter specifies the Elasticsearch version to use.
// Optional configuration can be provided via ElasticsearchServiceOption functions.
func NewElasticsearchService(stackVersion string, opts ...ElasticsearchServiceOption) (*ElasticsearchService, error) {
	cfg := &elasticsearchServiceConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	version := stackVersion
	if cfg.resolveLatestPatch {
		resolved, err := ResolveLatestPatchVersion(stackVersion)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve latest patch version: %w", err)
		}
		version = resolved
	}

	elasticsearchContainer, err := tces.RunContainer(
		context.Background(),
		testcontainers.WithImage("docker.elastic.co/elasticsearch/elasticsearch:"+version),
		tces.WithPassword("changeme"),
	)
	if err != nil {
		return nil, err
	}

	return &ElasticsearchService{
		container:     elasticsearchContainer,
		configOptions: elasticsearchContainer.Settings,
	}, nil
}

// tokenResponse represents the response from the Docker auth service.
type tokenResponse struct {
	Token string `json:"token"`
}

// tagsResponse represents the response from the Docker registry tags list endpoint.
type tagsResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// getRegistryToken obtains an anonymous bearer token for accessing the Elastic Docker registry.
func getRegistryToken() (string, error) {
	url := fmt.Sprintf("%s?service=token-service&scope=repository:%s:pull", elasticDockerAuthURL, elasticESRepository)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to get auth token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("auth request failed with status %d", resp.StatusCode)
	}

	var tokenResp tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	return tokenResp.Token, nil
}

// FetchAvailableTags retrieves all available tags for the Elasticsearch Docker image
// from the Elastic Docker registry.
func FetchAvailableTags() ([]string, error) {
	token, err := getRegistryToken()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/tags/list", elasticDockerAPIURL, elasticESRepository)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tags: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tags request failed with status %d", resp.StatusCode)
	}

	var tagsResp tagsResponse
	if err := json.NewDecoder(resp.Body).Decode(&tagsResp); err != nil {
		return nil, fmt.Errorf("failed to decode tags response: %w", err)
	}

	return tagsResp.Tags, nil
}

// parsedVersion represents a parsed semantic version for sorting purposes.
type parsedVersion struct {
	major      int
	minor      int
	patch      int
	preRelease string
	original   string
}

// parseVersion parses a version string into its components.
func parseVersion(version string) (parsedVersion, error) {
	// Pattern to match versions like "9.3.0" or "9.3.0-SNAPSHOT"
	pattern := regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)(?:-(.+))?$`)
	matches := pattern.FindStringSubmatch(version)
	if matches == nil {
		return parsedVersion{}, fmt.Errorf("invalid version format: %s", version)
	}

	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch, _ := strconv.Atoi(matches[3])

	preRelease := ""
	if len(matches) > 4 {
		preRelease = matches[4]
	}

	return parsedVersion{
		major:      major,
		minor:      minor,
		patch:      patch,
		preRelease: preRelease,
		original:   version,
	}, nil
}

// ResolveLatestPatchVersion finds the latest available patch version for the given
// version's major.minor from the Docker registry.
// If the input version has a -SNAPSHOT suffix, it will only match SNAPSHOT versions.
// Otherwise, it will only match stable (non-SNAPSHOT) versions.
func ResolveLatestPatchVersion(inputVersion string) (string, error) {
	parsed, err := parseVersion(inputVersion)
	if err != nil {
		return "", fmt.Errorf("failed to parse input version: %w", err)
	}

	tags, err := FetchAvailableTags()
	if err != nil {
		return "", err
	}

	isSnapshot := strings.HasSuffix(inputVersion, "-SNAPSHOT")

	// Filter tags matching major.minor and snapshot preference
	var candidates []parsedVersion
	for _, tag := range tags {
		tagParsed, err := parseVersion(tag)
		if err != nil {
			continue // Skip tags that don't match version pattern
		}

		// Must match major.minor
		if tagParsed.major != parsed.major || tagParsed.minor != parsed.minor {
			continue
		}

		// Match snapshot preference
		tagIsSnapshot := strings.HasSuffix(tag, "-SNAPSHOT")
		if isSnapshot {
			// For SNAPSHOT input, only match clean SNAPSHOT tags (no commit hash)
			if !tagIsSnapshot || strings.Contains(tagParsed.preRelease, "-") {
				continue
			}
		} else {
			// For stable input, only match non-prerelease versions
			if tagParsed.preRelease != "" {
				continue
			}
		}

		candidates = append(candidates, tagParsed)
	}

	if len(candidates) == 0 {
		return "", fmt.Errorf("no matching version found for %s", inputVersion)
	}

	// Sort by patch version descending to get the latest
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].patch > candidates[j].patch
	})

	return candidates[0].original, nil
}
