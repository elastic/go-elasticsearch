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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package getbuckets

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package getbuckets
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/get_buckets/MlGetBucketsRequest.ts#L26-L133
type Request struct {

	// AnomalyScore Refer to the description for the `anomaly_score` query parameter.
	AnomalyScore *float64 `json:"anomaly_score,omitempty"`
	// Desc Refer to the description for the `desc` query parameter.
	Desc *bool `json:"desc,omitempty"`
	// End Refer to the description for the `end` query parameter.
	End *types.DateTime `json:"end,omitempty"`
	// ExcludeInterim Refer to the description for the `exclude_interim` query parameter.
	ExcludeInterim *bool `json:"exclude_interim,omitempty"`
	// Expand Refer to the description for the `expand` query parameter.
	Expand *bool       `json:"expand,omitempty"`
	Page   *types.Page `json:"page,omitempty"`
	// Sort Refer to the desription for the `sort` query parameter.
	Sort *string `json:"sort,omitempty"`
	// Start Refer to the description for the `start` query parameter.
	Start *types.DateTime `json:"start,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Getbuckets request: %w", err)
	}

	return &req, nil
}
