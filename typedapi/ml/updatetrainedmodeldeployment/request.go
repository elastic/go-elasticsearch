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
// https://github.com/elastic/elasticsearch-specification/tree/07bf82537a186562d8699685e3704ea338b268ef

package updatetrainedmodeldeployment

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package updatetrainedmodeldeployment
//
// https://github.com/elastic/elasticsearch-specification/blob/07bf82537a186562d8699685e3704ea338b268ef/specification/ml/update_trained_model_deployment/MlUpdateTrainedModelDeploymentRequest.ts#L24-L50
type Request struct {

	// NumberOfAllocations The number of model allocations on each node where the model is deployed.
	// All allocations on a node share the same copy of the model in memory but use
	// a separate set of threads to evaluate the model.
	// Increasing this value generally increases the throughput.
	// If this setting is greater than the number of hardware threads
	// it will automatically be changed to a value less than the number of hardware
	// threads.
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
