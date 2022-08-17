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


package estimatemodelmemory

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package estimatemodelmemory
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/estimate_model_memory/MlEstimateModelMemoryRequest.ts#L26-L61
type Request struct {

	// AnalysisConfig For a list of the properties that you can specify in the
	// `analysis_config` component of the body of this API.
	AnalysisConfig *types.AnalysisConfig `json:"analysis_config,omitempty"`

	// MaxBucketCardinality Estimates of the highest cardinality in a single bucket that is observed
	// for influencer fields over the time period that the job analyzes data.
	// To produce a good answer, values must be provided for all influencer
	// fields. Providing values for fields that are not listed as `influencers`
	// has no effect on the estimation.
	MaxBucketCardinality map[types.Field]int64 `json:"max_bucket_cardinality,omitempty"`

	// OverallCardinality Estimates of the cardinality that is observed for fields over the whole
	// time period that the job analyzes data. To produce a good answer, values
	// must be provided for fields referenced in the `by_field_name`,
	// `over_field_name` and `partition_field_name` of any detectors. Providing
	// values for other fields has no effect on the estimation. It can be
	// omitted from the request if no detectors have a `by_field_name`,
	// `over_field_name` or `partition_field_name`.
	OverallCardinality map[types.Field]int64 `json:"overall_cardinality,omitempty"`
}

// RequestBuilder is the builder API for the estimatemodelmemory.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			MaxBucketCardinality: make(map[types.Field]int64, 0),
			OverallCardinality:   make(map[types.Field]int64, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Estimatemodelmemory request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) AnalysisConfig(analysisconfig *types.AnalysisConfigBuilder) *RequestBuilder {
	v := analysisconfig.Build()
	rb.v.AnalysisConfig = &v
	return rb
}

func (rb *RequestBuilder) MaxBucketCardinality(value map[types.Field]int64) *RequestBuilder {
	rb.v.MaxBucketCardinality = value
	return rb
}

func (rb *RequestBuilder) OverallCardinality(value map[types.Field]int64) *RequestBuilder {
	rb.v.OverallCardinality = value
	return rb
}
