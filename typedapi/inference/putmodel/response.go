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
// https://github.com/elastic/elasticsearch-specification/tree/17ac39c7f9266bc303baa029f90194aecb1c3b7c

package putmodel

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/tasktype"
)

// Response holds the response body struct for the package putmodel
//
// https://github.com/elastic/elasticsearch-specification/blob/17ac39c7f9266bc303baa029f90194aecb1c3b7c/specification/inference/put_model/PutModelResponse.ts#L22-L24
type Response struct {

	// ModelId The model Id
	ModelId string `json:"model_id"`
	// Service The service type
	Service string `json:"service"`
	// ServiceSettings Settings specific to the service
	ServiceSettings json.RawMessage `json:"service_settings"`
	// TaskSettings Task settings specific to the service and model
	TaskSettings json.RawMessage `json:"task_settings"`
	// TaskType The model's task type
	TaskType tasktype.TaskType `json:"task_type"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
