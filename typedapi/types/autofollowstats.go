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

// AutoFollowStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ccr/stats/types.ts.ts#L32-L47
type AutoFollowStats struct {
	AutoFollowedClusters []AutoFollowedCluster `json:"auto_followed_clusters"`
	// NumberOfFailedFollowIndices The number of indices that the auto-follow coordinator failed to
	// automatically follow.
	// The causes of recent failures are captured in the logs of the elected master
	// node and in the `auto_follow_stats.recent_auto_follow_errors` field.
	NumberOfFailedFollowIndices int64 `json:"number_of_failed_follow_indices"`
	// NumberOfFailedRemoteClusterStateRequests The number of times that the auto-follow coordinator failed to retrieve the
	// cluster state from a remote cluster registered in a collection of auto-follow
	// patterns.
	NumberOfFailedRemoteClusterStateRequests int64 `json:"number_of_failed_remote_cluster_state_requests"`
	// NumberOfSuccessfulFollowIndices The number of indices that the auto-follow coordinator successfully followed.
	NumberOfSuccessfulFollowIndices int64 `json:"number_of_successful_follow_indices"`
	// RecentAutoFollowErrors An array of objects representing failures by the auto-follow coordinator.
	RecentAutoFollowErrors []ErrorCause `json:"recent_auto_follow_errors"`
}

func (s *AutoFollowStats) UnmarshalJSON(data []byte) error {

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

		case "auto_followed_clusters":
			if err := dec.Decode(&s.AutoFollowedClusters); err != nil {
				return fmt.Errorf("%s | %w", "AutoFollowedClusters", err)
			}

		case "number_of_failed_follow_indices":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfFailedFollowIndices", err)
				}
				s.NumberOfFailedFollowIndices = value
			case float64:
				f := int64(v)
				s.NumberOfFailedFollowIndices = f
			}

		case "number_of_failed_remote_cluster_state_requests":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfFailedRemoteClusterStateRequests", err)
				}
				s.NumberOfFailedRemoteClusterStateRequests = value
			case float64:
				f := int64(v)
				s.NumberOfFailedRemoteClusterStateRequests = f
			}

		case "number_of_successful_follow_indices":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfSuccessfulFollowIndices", err)
				}
				s.NumberOfSuccessfulFollowIndices = value
			case float64:
				f := int64(v)
				s.NumberOfSuccessfulFollowIndices = f
			}

		case "recent_auto_follow_errors":
			if err := dec.Decode(&s.RecentAutoFollowErrors); err != nil {
				return fmt.Errorf("%s | %w", "RecentAutoFollowErrors", err)
			}

		}
	}
	return nil
}

// NewAutoFollowStats returns a AutoFollowStats.
func NewAutoFollowStats() *AutoFollowStats {
	r := &AutoFollowStats{}

	return r
}
