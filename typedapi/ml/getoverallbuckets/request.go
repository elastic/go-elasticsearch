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

package getoverallbuckets

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package getoverallbuckets
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/get_overall_buckets/MlGetOverallBucketsRequest.ts#L25-L143
type Request struct {

	// AllowNoMatch Refer to the description for the `allow_no_match` query parameter.
	AllowNoMatch *bool `json:"allow_no_match,omitempty"`
	// BucketSpan Refer to the description for the `bucket_span` query parameter.
	BucketSpan types.Duration `json:"bucket_span,omitempty"`
	// End Refer to the description for the `end` query parameter.
	End types.DateTime `json:"end,omitempty"`
	// ExcludeInterim Refer to the description for the `exclude_interim` query parameter.
	ExcludeInterim *bool `json:"exclude_interim,omitempty"`
	// OverallScore Refer to the description for the `overall_score` query parameter.
	OverallScore string `json:"overall_score,omitempty"`
	// Start Refer to the description for the `start` query parameter.
	Start types.DateTime `json:"start,omitempty"`
	// TopN Refer to the description for the `top_n` query parameter.
	TopN *int `json:"top_n,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Getoverallbuckets request: %w", err)
	}

	return &req, nil
}
