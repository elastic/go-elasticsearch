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
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v9"

	"github.com/testcontainers/testcontainers-go"
	tces "github.com/testcontainers/testcontainers-go/modules/elasticsearch"
)

const (
	elasticDockerAuthURL = "https://docker-auth.elastic.co/auth"
	elasticDockerAPIURL  = "https://docker.elastic.co/v2"
	elasticESRepository  = "elasticsearch/elasticsearch"
)

// ElasticStackImage is the resolved Elasticsearch Docker image version.
// It is set at init time by resolving the latest available patch version
// for the configured major.minor (from STACK_VERSION env or library version).
var ElasticStackImage string

func init() {
	stackVersion := elasticsearch.Version
	if v := os.Getenv("STACK_VERSION"); v != "" {
		stackVersion = v
	}

	resolved, err := ResolveLatestPatchVersion(stackVersion)
	if err != nil {
		// Fall back to the original version if resolution fails
		ElasticStackImage = stackVersion
		return
	}
	ElasticStackImage = resolved
}

// ElasticsearchService wraps an Elasticsearch container and its configuration.
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

// ESOptions returns the Elasticsearch client options.
func (es *ElasticsearchService) ESOptions() []elasticsearch.Option {
	return []elasticsearch.Option{
		elasticsearch.WithAddresses(es.configOptions.Address),
		elasticsearch.WithBasicAuth("elastic", es.configOptions.Password),
		elasticsearch.WithCACert(es.configOptions.CACert),
	}
}

// Terminate terminates the Elasticsearch container.
func (es *ElasticsearchService) Terminate(ctx context.Context) error {
	return es.container.Terminate(ctx)
}

type elasticsearchServiceConfig struct {
	env map[string]string
}

// ElasticsearchServiceOption configures ElasticsearchService creation.
type ElasticsearchServiceOption func(*elasticsearchServiceConfig)

// WithEnv sets environment variables for the Elasticsearch container.
func WithEnv(env map[string]string) ElasticsearchServiceOption {
	return func(c *elasticsearchServiceConfig) {
		c.env = env
	}
}

// NewElasticsearchService creates a new Elasticsearch container service.
func NewElasticsearchService(stackVersion string, opts ...ElasticsearchServiceOption) (*ElasticsearchService, error) {
	cfg := &elasticsearchServiceConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	containerOpts := []testcontainers.ContainerCustomizer{
		testcontainers.WithImage("docker.elastic.co/elasticsearch/elasticsearch:" + stackVersion),
		tces.WithPassword("changeme"),
	}
	if cfg.env != nil {
		containerOpts = append(containerOpts, testcontainers.WithEnv(cfg.env))
	}

	elasticsearchContainer, err := tces.RunContainer(context.Background(), containerOpts...)
	if err != nil {
		return nil, err
	}

	return &ElasticsearchService{
		container:     elasticsearchContainer,
		configOptions: elasticsearchContainer.Settings,
	}, nil
}

type tokenResponse struct {
	Token string `json:"token"`
}

type tagsResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

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

// FetchAvailableTags retrieves all available Elasticsearch image tags from the Docker registry.
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

type parsedVersion struct {
	major      int
	minor      int
	patch      int
	preRelease string
	original   string
	isSnapshot bool
}

func parseVersion(version string) (parsedVersion, error) {
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
// major.minor from the Docker registry. Release builds are prioritized over SNAPSHOTs.
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

	var candidates []parsedVersion
	for _, tag := range tags {
		tagParsed, err := parseVersion(tag)
		if err != nil {
			continue
		}

		if tagParsed.major != parsed.major || tagParsed.minor != parsed.minor {
			continue
		}

		tagIsSnapshot := strings.HasSuffix(tag, "-SNAPSHOT")

		if isSnapshot {
			// Accept both release and clean SNAPSHOT tags (no commit hash)
			if tagParsed.preRelease != "" && tagParsed.preRelease != "SNAPSHOT" {
				continue
			}
		} else {
			if tagParsed.preRelease != "" {
				continue
			}
		}

		tagParsed.isSnapshot = tagIsSnapshot
		candidates = append(candidates, tagParsed)
	}

	if len(candidates) == 0 {
		return "", fmt.Errorf("no matching version found for %s", inputVersion)
	}

	// Sort: release before SNAPSHOT, then by patch descending
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].isSnapshot != candidates[j].isSnapshot {
			return !candidates[i].isSnapshot
		}
		return candidates[i].patch > candidates[j].patch
	})

	return candidates[0].original, nil
}
