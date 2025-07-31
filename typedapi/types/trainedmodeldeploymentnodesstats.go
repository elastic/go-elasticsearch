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

// TrainedModelDeploymentNodesStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/TrainedModel.ts#L156-L201
type TrainedModelDeploymentNodesStats struct {
	// AverageInferenceTimeMs The average time for each inference call to complete on this node.
	AverageInferenceTimeMs Float64 `json:"average_inference_time_ms,omitempty"`
	// AverageInferenceTimeMsExcludingCacheHits The average time for each inference call to complete on this node, excluding
	// cache
	AverageInferenceTimeMsExcludingCacheHits Float64 `json:"average_inference_time_ms_excluding_cache_hits,omitempty"`
	AverageInferenceTimeMsLastMinute         Float64 `json:"average_inference_time_ms_last_minute,omitempty"`
	// ErrorCount The number of errors when evaluating the trained model.
	ErrorCount                       *int   `json:"error_count,omitempty"`
	InferenceCacheHitCount           *int64 `json:"inference_cache_hit_count,omitempty"`
	InferenceCacheHitCountLastMinute *int64 `json:"inference_cache_hit_count_last_minute,omitempty"`
	// InferenceCount The total number of inference calls made against this node for this model.
	InferenceCount *int64 `json:"inference_count,omitempty"`
	// LastAccess The epoch time stamp of the last inference call for the model on this node.
	LastAccess *int64 `json:"last_access,omitempty"`
	// Node Information pertaining to the node.
	Node DiscoveryNode `json:"node,omitempty"`
	// NumberOfAllocations The number of allocations assigned to this node.
	NumberOfAllocations *int `json:"number_of_allocations,omitempty"`
	// NumberOfPendingRequests The number of inference requests queued to be processed.
	NumberOfPendingRequests *int  `json:"number_of_pending_requests,omitempty"`
	PeakThroughputPerMinute int64 `json:"peak_throughput_per_minute"`
	// RejectedExecutionCount The number of inference requests that were not processed because the queue
	// was full.
	RejectedExecutionCount *int `json:"rejected_execution_count,omitempty"`
	// RoutingState The current routing state and reason for the current routing state for this
	// allocation.
	RoutingState TrainedModelAssignmentRoutingStateAndReason `json:"routing_state"`
	// StartTime The epoch timestamp when the allocation started.
	StartTime *int64 `json:"start_time,omitempty"`
	// ThreadsPerAllocation The number of threads used by each allocation during inference.
	ThreadsPerAllocation *int `json:"threads_per_allocation,omitempty"`
	ThroughputLastMinute int  `json:"throughput_last_minute"`
	// TimeoutCount The number of inference requests that timed out before being processed.
	TimeoutCount *int `json:"timeout_count,omitempty"`
}

func (s *TrainedModelDeploymentNodesStats) UnmarshalJSON(data []byte) error {

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

		case "average_inference_time_ms":
			if err := dec.Decode(&s.AverageInferenceTimeMs); err != nil {
				return fmt.Errorf("%s | %w", "AverageInferenceTimeMs", err)
			}

		case "average_inference_time_ms_excluding_cache_hits":
			if err := dec.Decode(&s.AverageInferenceTimeMsExcludingCacheHits); err != nil {
				return fmt.Errorf("%s | %w", "AverageInferenceTimeMsExcludingCacheHits", err)
			}

		case "average_inference_time_ms_last_minute":
			if err := dec.Decode(&s.AverageInferenceTimeMsLastMinute); err != nil {
				return fmt.Errorf("%s | %w", "AverageInferenceTimeMsLastMinute", err)
			}

		case "error_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ErrorCount", err)
				}
				s.ErrorCount = &value
			case float64:
				f := int(v)
				s.ErrorCount = &f
			}

		case "inference_cache_hit_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "InferenceCacheHitCount", err)
				}
				s.InferenceCacheHitCount = &value
			case float64:
				f := int64(v)
				s.InferenceCacheHitCount = &f
			}

		case "inference_cache_hit_count_last_minute":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "InferenceCacheHitCountLastMinute", err)
				}
				s.InferenceCacheHitCountLastMinute = &value
			case float64:
				f := int64(v)
				s.InferenceCacheHitCountLastMinute = &f
			}

		case "inference_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "InferenceCount", err)
				}
				s.InferenceCount = &value
			case float64:
				f := int64(v)
				s.InferenceCount = &f
			}

		case "last_access":
			if err := dec.Decode(&s.LastAccess); err != nil {
				return fmt.Errorf("%s | %w", "LastAccess", err)
			}

		case "node":
			if err := dec.Decode(&s.Node); err != nil {
				return fmt.Errorf("%s | %w", "Node", err)
			}

		case "number_of_allocations":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfAllocations", err)
				}
				s.NumberOfAllocations = &value
			case float64:
				f := int(v)
				s.NumberOfAllocations = &f
			}

		case "number_of_pending_requests":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfPendingRequests", err)
				}
				s.NumberOfPendingRequests = &value
			case float64:
				f := int(v)
				s.NumberOfPendingRequests = &f
			}

		case "peak_throughput_per_minute":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PeakThroughputPerMinute", err)
				}
				s.PeakThroughputPerMinute = value
			case float64:
				f := int64(v)
				s.PeakThroughputPerMinute = f
			}

		case "rejected_execution_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RejectedExecutionCount", err)
				}
				s.RejectedExecutionCount = &value
			case float64:
				f := int(v)
				s.RejectedExecutionCount = &f
			}

		case "routing_state":
			if err := dec.Decode(&s.RoutingState); err != nil {
				return fmt.Errorf("%s | %w", "RoutingState", err)
			}

		case "start_time":
			if err := dec.Decode(&s.StartTime); err != nil {
				return fmt.Errorf("%s | %w", "StartTime", err)
			}

		case "threads_per_allocation":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ThreadsPerAllocation", err)
				}
				s.ThreadsPerAllocation = &value
			case float64:
				f := int(v)
				s.ThreadsPerAllocation = &f
			}

		case "throughput_last_minute":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ThroughputLastMinute", err)
				}
				s.ThroughputLastMinute = value
			case float64:
				f := int(v)
				s.ThroughputLastMinute = f
			}

		case "timeout_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TimeoutCount", err)
				}
				s.TimeoutCount = &value
			case float64:
				f := int(v)
				s.TimeoutCount = &f
			}

		}
	}
	return nil
}

// NewTrainedModelDeploymentNodesStats returns a TrainedModelDeploymentNodesStats.
func NewTrainedModelDeploymentNodesStats() *TrainedModelDeploymentNodesStats {
	r := &TrainedModelDeploymentNodesStats{}

	return r
}
