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

package gettrainedmodelsstats

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package gettrainedmodelsstats
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/get_trained_models_stats/MlGetTrainedModelStatsResponse.ts#L23-L33

type Response struct {

	// Count The total number of trained model statistics that matched the requested ID
	// patterns. Could be higher than the number of items in the trained_model_stats
	// array as the size of the array is restricted by the supplied size parameter.
	Count int `json:"count"`
	// TrainedModelStats An array of trained model statistics, which are sorted by the model_id value
	// in ascending order.
	TrainedModelStats []types.TrainedModelStats `json:"trained_model_stats"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
