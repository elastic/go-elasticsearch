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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package knnsearch

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package knnsearch
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/knn_search/KnnSearchResponse.ts#L26-L54
type Response struct {

	// Fields Contains field values for the documents. These fields
	// must be specified in the request using the `fields` parameter.
	Fields map[string]json.RawMessage `json:"fields,omitempty"`
	// Hits Contains returned documents and metadata.
	Hits types.HitsMetadata `json:"hits"`
	// MaxScore Highest returned document score. This value is null for requests
	// that do not sort by score.
	MaxScore *types.Float64 `json:"max_score,omitempty"`
	// Shards_ Contains a count of shards used for the request.
	Shards_ types.ShardStatistics `json:"_shards"`
	// TimedOut If true, the request timed out before completion;
	// returned results may be partial or empty.
	TimedOut bool `json:"timed_out"`
	// Took Milliseconds it took Elasticsearch to execute the request.
	Took int64 `json:"took"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Fields: make(map[string]json.RawMessage, 0),
	}
	return r
}
