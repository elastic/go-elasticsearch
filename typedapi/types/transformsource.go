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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// TransformSource type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/transform/_types/Transform.ts#L145-L163
type TransformSource struct {
	// Index The source indices for the transform. It can be a single index, an index
	// pattern (for example, `"my-index-*""`), an
	// array of indices (for example, `["my-index-000001", "my-index-000002"]`), or
	// an array of index patterns (for
	// example, `["my-index-*", "my-other-index-*"]`. For remote indices use the
	// syntax `"remote_name:index_name"`. If
	// any indices are in remote clusters then the master node and at least one
	// transform node must have the `remote_cluster_client` node role.
	Index []string `json:"index"`
	// Query A query clause that retrieves a subset of data from the source index.
	Query *Query `json:"query,omitempty"`
	// RuntimeMappings Definitions of search-time runtime fields that can be used by the transform.
	// For search runtime fields all data
	// nodes, including remote nodes, must be 7.12 or later.
	RuntimeMappings map[string]RuntimeField `json:"runtime_mappings,omitempty"`
}

// NewTransformSource returns a TransformSource.
func NewTransformSource() *TransformSource {
	r := &TransformSource{}

	return r
}
