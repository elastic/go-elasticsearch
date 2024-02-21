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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package stats

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package stats
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/indices/stats/IndicesStatsResponse.ts#L24-L30
type Response struct {
	All_    types.IndicesStats            `json:"_all"`
	Indices map[string]types.IndicesStats `json:"indices,omitempty"`
	Shards_ types.ShardStatistics         `json:"_shards"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Indices: make(map[string]types.IndicesStats, 0),
	}
	return r
}
