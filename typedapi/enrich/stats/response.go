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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package stats

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package stats
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/enrich/stats/EnrichStatsResponse.ts#L22-L39

type Response struct {

	// CacheStats Objects containing information about the enrich cache stats on each ingest
	// node.
	CacheStats []types.CacheStats `json:"cache_stats,omitempty"`
	// CoordinatorStats Objects containing information about each coordinating ingest node for
	// configured enrich processors.
	CoordinatorStats []types.CoordinatorStats `json:"coordinator_stats"`
	// ExecutingPolicies Objects containing information about each enrich policy that is currently
	// executing.
	ExecutingPolicies []types.ExecutingPolicy `json:"executing_policies"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
