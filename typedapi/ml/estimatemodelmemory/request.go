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


package estimatemodelmemory

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package estimatemodelmemory
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/estimate_model_memory/MlEstimateModelMemoryRequest.ts#L26-L61
type Request struct {

	// AnalysisConfig For a list of the properties that you can specify in the
	// `analysis_config` component of the body of this API.
	AnalysisConfig *types.AnalysisConfig `json:"analysis_config,omitempty"`
	// MaxBucketCardinality Estimates of the highest cardinality in a single bucket that is observed
	// for influencer fields over the time period that the job analyzes data.
	// To produce a good answer, values must be provided for all influencer
	// fields. Providing values for fields that are not listed as `influencers`
	// has no effect on the estimation.
	MaxBucketCardinality map[string]int64 `json:"max_bucket_cardinality,omitempty"`
	// OverallCardinality Estimates of the cardinality that is observed for fields over the whole
	// time period that the job analyzes data. To produce a good answer, values
	// must be provided for fields referenced in the `by_field_name`,
	// `over_field_name` and `partition_field_name` of any detectors. Providing
	// values for other fields has no effect on the estimation. It can be
	// omitted from the request if no detectors have a `by_field_name`,
	// `over_field_name` or `partition_field_name`.
	OverallCardinality map[string]int64 `json:"overall_cardinality,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		MaxBucketCardinality: make(map[string]int64, 0),
		OverallCardinality:   make(map[string]int64, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Estimatemodelmemory request: %w", err)
	}

	return &req, nil
}
