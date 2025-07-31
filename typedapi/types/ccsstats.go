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

// CCSStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L769-L784
type CCSStats struct {
	// Clusters Contains remote cluster settings and metrics collected from them.
	// The keys are cluster names, and the values are per-cluster data.
	// Only present if `include_remotes` option is set to true.
	Clusters map[string]RemoteClusterInfo `json:"clusters,omitempty"`
	// Esql_ Information about ES|QL cross-cluster query usage.
	Esql_ *CCSUsageStats `json:"_esql,omitempty"`
	// Search_ Information about cross-cluster search usage.
	Search_ CCSUsageStats `json:"_search"`
}

// NewCCSStats returns a CCSStats.
func NewCCSStats() *CCSStats {
	r := &CCSStats{
		Clusters: make(map[string]RemoteClusterInfo),
	}

	return r
}
