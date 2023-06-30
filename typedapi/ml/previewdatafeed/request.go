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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package previewdatafeed

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package previewdatafeed
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/preview_datafeed/MlPreviewDatafeedRequest.ts#L26-L69
type Request struct {

	// DatafeedConfig The datafeed definition to preview.
	DatafeedConfig *types.DatafeedConfig `json:"datafeed_config,omitempty"`
	// JobConfig The configuration details for the anomaly detection job that is associated
	// with the datafeed. If the
	// `datafeed_config` object does not include a `job_id` that references an
	// existing anomaly detection job, you must
	// supply this `job_config` object. If you include both a `job_id` and a
	// `job_config`, the latter information is
	// used. You cannot specify a `job_config` object unless you also supply a
	// `datafeed_config` object.
	JobConfig *types.JobConfig `json:"job_config,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Previewdatafeed request: %w", err)
	}

	return &req, nil
}
