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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package explaindataframeanalytics

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package explaindataframeanalytics
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/explain_data_frame_analytics/MlExplainDataFrameAnalyticsResponse.ts#L25-L32
type Response struct {

	// FieldSelection An array of objects that explain selection for each field, sorted by the
	// field names.
	FieldSelection []types.DataframeAnalyticsFieldSelection `json:"field_selection"`
	// MemoryEstimation An array of objects that explain selection for each field, sorted by the
	// field names.
	MemoryEstimation types.DataframeAnalyticsMemoryEstimation `json:"memory_estimation"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
