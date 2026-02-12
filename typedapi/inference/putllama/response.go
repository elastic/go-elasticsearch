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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package putllama

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tasktypellama"
)

// Response holds the response body struct for the package putllama
//
// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/inference/put_llama/PutLlamaResponse.ts#L22-L25
type Response struct {
	// ChunkingSettings The chunking configuration object.
	// Applies only to the `embedding`, `sparse_embedding` and `text_embedding` task
	// types.
	// Not applicable to the `rerank`, `completion`, or `chat_completion` task
	// types.
	ChunkingSettings *types.InferenceChunkingSettings `json:"chunking_settings,omitempty"`
	// InferenceId The inference Id
	InferenceId string `json:"inference_id"`
	// Service The service type
	Service string `json:"service"`
	// ServiceSettings Settings specific to the service
	ServiceSettings json.RawMessage `json:"service_settings"`
	// TaskSettings Task settings specific to the service and task type
	TaskSettings json.RawMessage `json:"task_settings,omitempty"`
	// TaskType The task type
	TaskType tasktypellama.TaskTypeLlama `json:"task_type"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
