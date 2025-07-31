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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package updatetrainedmodeldeployment

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package updatetrainedmodeldeployment
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/update_trained_model_deployment/MlUpdateTrainedModelDeploymentRequest.ts#L25-L78
type Request struct {

	// AdaptiveAllocations Adaptive allocations configuration. When enabled, the number of allocations
	// is set based on the current load.
	// If adaptive_allocations is enabled, do not set the number of allocations
	// manually.
	AdaptiveAllocations *types.AdaptiveAllocationsSettings `json:"adaptive_allocations,omitempty"`
	// NumberOfAllocations The number of model allocations on each node where the model is deployed.
	// All allocations on a node share the same copy of the model in memory but use
	// a separate set of threads to evaluate the model.
	// Increasing this value generally increases the throughput.
	// If this setting is greater than the number of hardware threads
	// it will automatically be changed to a value less than the number of hardware
	// threads.
	// If adaptive_allocations is enabled, do not set this value, because it’s
	// automatically set.
	NumberOfAllocations *int `json:"number_of_allocations,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Updatetrainedmodeldeployment request: %w", err)
	}

	return &req, nil
}
