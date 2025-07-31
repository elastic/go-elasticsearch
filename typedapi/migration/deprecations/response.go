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

package deprecations

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package deprecations
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/migration/deprecations/DeprecationInfoResponse.ts#L23-L54
type Response struct {

	// ClusterSettings Cluster-level deprecation warnings.
	ClusterSettings []types.Deprecation            `json:"cluster_settings"`
	DataStreams     map[string][]types.Deprecation `json:"data_streams"`
	// IlmPolicies ILM policy warnings are sectioned off per policy.
	IlmPolicies map[string][]types.Deprecation `json:"ilm_policies"`
	// IndexSettings Index warnings are sectioned off per index and can be filtered using an
	// index-pattern in the query.
	// This section includes warnings for the backing indices of data streams
	// specified in the request path.
	IndexSettings map[string][]types.Deprecation `json:"index_settings"`
	// MlSettings Machine learning-related deprecation warnings.
	MlSettings []types.Deprecation `json:"ml_settings"`
	// NodeSettings Node-level deprecation warnings.
	// Since only a subset of your nodes might incorporate these settings, it is
	// important to read the details section for more information about which nodes
	// are affected.
	NodeSettings []types.Deprecation `json:"node_settings"`
	// Templates Template warnings are sectioned off per template and include deprecations for
	// both component templates and
	// index templates.
	Templates map[string][]types.Deprecation `json:"templates"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		DataStreams:   make(map[string][]types.Deprecation, 0),
		IlmPolicies:   make(map[string][]types.Deprecation, 0),
		IndexSettings: make(map[string][]types.Deprecation, 0),
		Templates:     make(map[string][]types.Deprecation, 0),
	}
	return r
}
