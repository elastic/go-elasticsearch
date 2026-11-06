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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package knneval

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/knnevalvectoropskind"
)

// Response holds the response body struct for the package knneval
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/_global/knn_eval/KnnEvalResponse.ts#L32-L75
type Response struct {
	// Baseline The settings of the baseline search. Not returned when `wait_for_completion`
	// is `false`.
	Baseline *types.KnnEvalReportedSettings `json:"baseline,omitempty"`
	// BaselineTookMs The time the baseline searches took, in milliseconds. Not returned when
	// `wait_for_completion` is `false`.
	BaselineTookMs *int64 `json:"baseline_took_ms,omitempty"`
	// BaselineVectorOps The number of vector operations of the baseline searches. Not returned when
	// `wait_for_completion` is `false`.
	BaselineVectorOps *int64 `json:"baseline_vector_ops,omitempty"`
	// BaselineVectorOpsKind What `baseline_vector_ops` counts. Not returned when `wait_for_completion` is
	// `false`.
	BaselineVectorOpsKind *knnevalvectoropskind.KnnEvalVectorOpsKind `json:"baseline_vector_ops_kind,omitempty"`
	// Failures The queries that failed, by query identifier. Not returned when
	// `wait_for_completion` is `false`.
	Failures map[string]types.KnnEvalFailure `json:"failures,omitempty"`
	// QuerySource The source of the query vectors. Not returned when `wait_for_completion` is
	// `false`.
	QuerySource *types.KnnEvalQuerySourceSummary `json:"query_source,omitempty"`
	// Results The results of each `knn_settings` entry, in the order of the request. Not
	// returned when `wait_for_completion` is `false`.
	Results []types.KnnEvalResult `json:"results,omitempty"`
	// Task The identifier of the task that runs the evaluation. Only returned when
	// `wait_for_completion` is `false`.
	Task *string `json:"task,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{
		Failures: make(map[string]types.KnnEvalFailure, 0),
	}
	return r
}
