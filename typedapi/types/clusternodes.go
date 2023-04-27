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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// ClusterNodes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/cluster/stats/types.ts#L201-L228
type ClusterNodes struct {
	// Count Contains counts for nodes selected by the request’s node filters.
	Count ClusterNodeCount `json:"count"`
	// DiscoveryTypes Contains statistics about the discovery types used by selected nodes.
	DiscoveryTypes map[string]int `json:"discovery_types"`
	// Fs Contains statistics about file stores by selected nodes.
	Fs               ClusterFileSystem       `json:"fs"`
	IndexingPressure ClusterIndexingPressure `json:"indexing_pressure"`
	Ingest           ClusterIngest           `json:"ingest"`
	// Jvm Contains statistics about the Java Virtual Machines (JVMs) used by selected
	// nodes.
	Jvm ClusterJvm `json:"jvm"`
	// NetworkTypes Contains statistics about the transport and HTTP networks used by selected
	// nodes.
	NetworkTypes ClusterNetworkTypes `json:"network_types"`
	// Os Contains statistics about the operating systems used by selected nodes.
	Os ClusterOperatingSystem `json:"os"`
	// PackagingTypes Contains statistics about Elasticsearch distributions installed on selected
	// nodes.
	PackagingTypes []NodePackagingType `json:"packaging_types"`
	// Plugins Contains statistics about installed plugins and modules by selected nodes.
	Plugins []PluginStats `json:"plugins"`
	// Process Contains statistics about processes used by selected nodes.
	Process ClusterProcess `json:"process"`
	// Versions Array of Elasticsearch versions used on selected nodes.
	Versions []string `json:"versions"`
}

func (s *ClusterNodes) UnmarshalJSON(data []byte) error {

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

		case "count":
			if err := dec.Decode(&s.Count); err != nil {
				return err
			}

		case "discovery_types":
			if s.DiscoveryTypes == nil {
				s.DiscoveryTypes = make(map[string]int, 0)
			}
			if err := dec.Decode(&s.DiscoveryTypes); err != nil {
				return err
			}

		case "fs":
			if err := dec.Decode(&s.Fs); err != nil {
				return err
			}

		case "indexing_pressure":
			if err := dec.Decode(&s.IndexingPressure); err != nil {
				return err
			}

		case "ingest":
			if err := dec.Decode(&s.Ingest); err != nil {
				return err
			}

		case "jvm":
			if err := dec.Decode(&s.Jvm); err != nil {
				return err
			}

		case "network_types":
			if err := dec.Decode(&s.NetworkTypes); err != nil {
				return err
			}

		case "os":
			if err := dec.Decode(&s.Os); err != nil {
				return err
			}

		case "packaging_types":
			if err := dec.Decode(&s.PackagingTypes); err != nil {
				return err
			}

		case "plugins":
			if err := dec.Decode(&s.Plugins); err != nil {
				return err
			}

		case "process":
			if err := dec.Decode(&s.Process); err != nil {
				return err
			}

		case "versions":
			if err := dec.Decode(&s.Versions); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewClusterNodes returns a ClusterNodes.
func NewClusterNodes() *ClusterNodes {
	r := &ClusterNodes{
		DiscoveryTypes: make(map[string]int, 0),
	}

	return r
}
