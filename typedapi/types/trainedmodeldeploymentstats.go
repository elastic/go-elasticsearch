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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deploymentstate"
)

// TrainedModelDeploymentStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/TrainedModel.ts#L62-L102
type TrainedModelDeploymentStats struct {
	// AllocationStatus The detailed allocation status for the deployment.
	AllocationStatus TrainedModelDeploymentAllocationStatus `json:"allocation_status"`
	CacheSize        ByteSize                               `json:"cache_size,omitempty"`
	// DeploymentId The unique identifier for the trained model deployment.
	DeploymentId string `json:"deployment_id"`
	// ErrorCount The sum of `error_count` for all nodes in the deployment.
	ErrorCount int `json:"error_count"`
	// InferenceCount The sum of `inference_count` for all nodes in the deployment.
	InferenceCount int `json:"inference_count"`
	// ModelId The unique identifier for the trained model.
	ModelId string `json:"model_id"`
	// Nodes The deployment stats for each node that currently has the model allocated.
	// In serverless, stats are reported for a single unnamed virtual node.
	Nodes TrainedModelDeploymentNodesStats `json:"nodes"`
	// NumberOfAllocations The number of allocations requested.
	NumberOfAllocations int `json:"number_of_allocations"`
	// QueueCapacity The number of inference requests that can be queued before new requests are
	// rejected.
	QueueCapacity int `json:"queue_capacity"`
	// Reason The reason for the current deployment state. Usually only populated when
	// the model is not deployed to a node.
	Reason string `json:"reason"`
	// RejectedExecutionCount The sum of `rejected_execution_count` for all nodes in the deployment.
	// Individual nodes reject an inference request if the inference queue is full.
	// The queue size is controlled by the `queue_capacity` setting in the start
	// trained model deployment API.
	RejectedExecutionCount int `json:"rejected_execution_count"`
	// StartTime The epoch timestamp when the deployment started.
	StartTime int64 `json:"start_time"`
	// State The overall state of the deployment.
	State deploymentstate.DeploymentState `json:"state"`
	// ThreadsPerAllocation The number of threads used be each allocation during inference.
	ThreadsPerAllocation int `json:"threads_per_allocation"`
	// TimeoutCount The sum of `timeout_count` for all nodes in the deployment.
	TimeoutCount int `json:"timeout_count"`
}

func (s *TrainedModelDeploymentStats) UnmarshalJSON(data []byte) error {

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

		case "allocation_status":
			if err := dec.Decode(&s.AllocationStatus); err != nil {
				return err
			}

		case "cache_size":
			if err := dec.Decode(&s.CacheSize); err != nil {
				return err
			}

		case "deployment_id":
			if err := dec.Decode(&s.DeploymentId); err != nil {
				return err
			}

		case "error_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ErrorCount = value
			case float64:
				f := int(v)
				s.ErrorCount = f
			}

		case "inference_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.InferenceCount = value
			case float64:
				f := int(v)
				s.InferenceCount = f
			}

		case "model_id":
			if err := dec.Decode(&s.ModelId); err != nil {
				return err
			}

		case "nodes":
			if err := dec.Decode(&s.Nodes); err != nil {
				return err
			}

		case "number_of_allocations":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NumberOfAllocations = value
			case float64:
				f := int(v)
				s.NumberOfAllocations = f
			}

		case "queue_capacity":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.QueueCapacity = value
			case float64:
				f := int(v)
				s.QueueCapacity = f
			}

		case "reason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = o

		case "rejected_execution_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.RejectedExecutionCount = value
			case float64:
				f := int(v)
				s.RejectedExecutionCount = f
			}

		case "start_time":
			if err := dec.Decode(&s.StartTime); err != nil {
				return err
			}

		case "state":
			if err := dec.Decode(&s.State); err != nil {
				return err
			}

		case "threads_per_allocation":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ThreadsPerAllocation = value
			case float64:
				f := int(v)
				s.ThreadsPerAllocation = f
			}

		case "timeout_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TimeoutCount = value
			case float64:
				f := int(v)
				s.TimeoutCount = f
			}

		}
	}
	return nil
}

// NewTrainedModelDeploymentStats returns a TrainedModelDeploymentStats.
func NewTrainedModelDeploymentStats() *TrainedModelDeploymentStats {
	r := &TrainedModelDeploymentStats{}

	return r
}
