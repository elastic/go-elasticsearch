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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package previewdatafeed

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package previewdatafeed
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/preview_datafeed/MlPreviewDatafeedRequest.ts#L25-L64
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

// RequestBuilder is the builder API for the previewdatafeed.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Previewdatafeed request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) DatafeedConfig(datafeedconfig *types.DatafeedConfigBuilder) *RequestBuilder {
	v := datafeedconfig.Build()
	rb.v.DatafeedConfig = &v
	return rb
}

func (rb *RequestBuilder) JobConfig(jobconfig *types.JobConfigBuilder) *RequestBuilder {
	v := jobconfig.Build()
	rb.v.JobConfig = &v
	return rb
}
