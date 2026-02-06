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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package getsamplestats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package getsamplestats
//
// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/indices/get_sample_stats/GetRandomSampleStatsResponse.ts#L23-L44
type Response struct {
	LastException                        *string        `json:"last_exception,omitempty"`
	PotentialSamples                     int64          `json:"potential_samples"`
	SamplesAccepted                      int64          `json:"samples_accepted"`
	SamplesRejectedForCondition          int64          `json:"samples_rejected_for_condition"`
	SamplesRejectedForException          int64          `json:"samples_rejected_for_exception"`
	SamplesRejectedForMaxSamplesExceeded int64          `json:"samples_rejected_for_max_samples_exceeded"`
	SamplesRejectedForRate               int64          `json:"samples_rejected_for_rate"`
	SamplesRejectedForSize               int64          `json:"samples_rejected_for_size"`
	TimeCompilingCondition               types.Duration `json:"time_compiling_condition,omitempty"`
	TimeCompilingConditionMillis         int64          `json:"time_compiling_condition_millis"`
	TimeEvaluatingCondition              types.Duration `json:"time_evaluating_condition,omitempty"`
	TimeEvaluatingConditionMillis        int64          `json:"time_evaluating_condition_millis"`
	TimeSampling                         types.Duration `json:"time_sampling,omitempty"`
	TimeSamplingMillis                   int64          `json:"time_sampling_millis"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
