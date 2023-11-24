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

package rankeval

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package rankeval
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/rank_eval/RankEvalResponse.ts#L26-L34
type Response struct {

	// Details The details section contains one entry for every query in the original
	// requests section, keyed by the search request id
	Details  map[string]types.RankEvalMetricDetail `json:"details"`
	Failures map[string]json.RawMessage            `json:"failures"`
	// MetricScore The overall evaluation quality calculated by the defined metric
	MetricScore types.Float64 `json:"metric_score"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Details:  make(map[string]types.RankEvalMetricDetail, 0),
		Failures: make(map[string]json.RawMessage, 0),
	}
	return r
}
