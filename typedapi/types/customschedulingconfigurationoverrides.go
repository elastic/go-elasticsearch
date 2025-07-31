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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CustomSchedulingConfigurationOverrides type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/connector/_types/Connector.ts#L112-L118
type CustomSchedulingConfigurationOverrides struct {
	DomainAllowlist          []string `json:"domain_allowlist,omitempty"`
	MaxCrawlDepth            *int     `json:"max_crawl_depth,omitempty"`
	SeedUrls                 []string `json:"seed_urls,omitempty"`
	SitemapDiscoveryDisabled *bool    `json:"sitemap_discovery_disabled,omitempty"`
	SitemapUrls              []string `json:"sitemap_urls,omitempty"`
}

func (s *CustomSchedulingConfigurationOverrides) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "domain_allowlist":
			if err := dec.Decode(&s.DomainAllowlist); err != nil {
				return fmt.Errorf("%s | %w", "DomainAllowlist", err)
			}

		case "max_crawl_depth":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxCrawlDepth", err)
				}
				s.MaxCrawlDepth = &value
			case float64:
				f := int(v)
				s.MaxCrawlDepth = &f
			}

		case "seed_urls":
			if err := dec.Decode(&s.SeedUrls); err != nil {
				return fmt.Errorf("%s | %w", "SeedUrls", err)
			}

		case "sitemap_discovery_disabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SitemapDiscoveryDisabled", err)
				}
				s.SitemapDiscoveryDisabled = &value
			case bool:
				s.SitemapDiscoveryDisabled = &v
			}

		case "sitemap_urls":
			if err := dec.Decode(&s.SitemapUrls); err != nil {
				return fmt.Errorf("%s | %w", "SitemapUrls", err)
			}

		}
	}
	return nil
}

// NewCustomSchedulingConfigurationOverrides returns a CustomSchedulingConfigurationOverrides.
func NewCustomSchedulingConfigurationOverrides() *CustomSchedulingConfigurationOverrides {
	r := &CustomSchedulingConfigurationOverrides{}

	return r
}
