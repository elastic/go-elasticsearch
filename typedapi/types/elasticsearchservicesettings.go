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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ElasticsearchServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L1215-L1249
type ElasticsearchServiceSettings struct {
	// AdaptiveAllocations Adaptive allocations configuration details.
	// If `enabled` is true, the number of allocations of the model is set based on
	// the current load the process gets.
	// When the load is high, a new model allocation is automatically created,
	// respecting the value of `max_number_of_allocations` if it's set.
	// When the load is low, a model allocation is automatically removed, respecting
	// the value of `min_number_of_allocations` if it's set.
	// If `enabled` is true, do not set the number of allocations manually.
	AdaptiveAllocations *AdaptiveAllocations `json:"adaptive_allocations,omitempty"`
	// DeploymentId The deployment identifier for a trained model deployment.
	// When `deployment_id` is used the `model_id` is optional.
	DeploymentId *string `json:"deployment_id,omitempty"`
	// ModelId The name of the model to use for the inference task.
	// It can be the ID of a built-in model (for example, `.multilingual-e5-small`
	// for E5) or a text embedding model that was uploaded by using the Eland
	// client.
	ModelId string `json:"model_id"`
	// NumAllocations The total number of allocations that are assigned to the model across machine
	// learning nodes.
	// Increasing this value generally increases the throughput.
	// If adaptive allocations are enabled, do not set this value because it's
	// automatically set.
	NumAllocations *int `json:"num_allocations,omitempty"`
	// NumThreads The number of threads used by each model allocation during inference.
	// This setting generally increases the speed per inference request.
	// The inference process is a compute-bound process; `threads_per_allocations`
	// must not exceed the number of available allocated processors per node.
	// The value must be a power of 2.
	// The maximum value is 32.
	NumThreads int `json:"num_threads"`
}

func (s *ElasticsearchServiceSettings) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "adaptive_allocations":
			if err := dec.Decode(&s.AdaptiveAllocations); err != nil {
				return fmt.Errorf("%s | %w", "AdaptiveAllocations", err)
			}

		case "deployment_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "DeploymentId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DeploymentId = &o

		case "model_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ModelId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelId = o

		case "num_allocations":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumAllocations", err)
				}
				s.NumAllocations = &value
			case float64:
				f := int(v)
				s.NumAllocations = &f
			}

		case "num_threads":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumThreads", err)
				}
				s.NumThreads = value
			case float64:
				f := int(v)
				s.NumThreads = f
			}

		}
	}
	return nil
}

// NewElasticsearchServiceSettings returns a ElasticsearchServiceSettings.
func NewElasticsearchServiceSettings() *ElasticsearchServiceSettings {
	r := &ElasticsearchServiceSettings{}

	return r
}
