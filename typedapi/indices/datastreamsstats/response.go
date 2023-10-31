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

package datastreamsstats

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package datastreamsstats
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/indices/data_streams_stats/IndicesDataStreamsStatsResponse.ts#L25-L43

type Response struct {

	// BackingIndices Total number of backing indices for the selected data streams.
	BackingIndices int `json:"backing_indices"`
	// DataStreamCount Total number of selected data streams.
	DataStreamCount int `json:"data_stream_count"`
	// DataStreams Contains statistics for the selected data streams.
	DataStreams []types.DataStreamsStatsItem `json:"data_streams"`
	// Shards_ Contains information about shards that attempted to execute the request.
	Shards_ types.ShardStatistics `json:"_shards"`
	// TotalStoreSizeBytes Total size, in bytes, of all shards for the selected data streams.
	TotalStoreSizeBytes int `json:"total_store_size_bytes"`
	// TotalStoreSizes Total size of all shards for the selected data streams.
	// This property is included only if the `human` query parameter is `true`
	TotalStoreSizes types.ByteSize `json:"total_store_sizes,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
