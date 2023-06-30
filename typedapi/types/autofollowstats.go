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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// AutoFollowStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ccr/stats/types.ts.ts#L33-L39
type AutoFollowStats struct {
	AutoFollowedClusters                     []AutoFollowedCluster `json:"auto_followed_clusters"`
	NumberOfFailedFollowIndices              int64                 `json:"number_of_failed_follow_indices"`
	NumberOfFailedRemoteClusterStateRequests int64                 `json:"number_of_failed_remote_cluster_state_requests"`
	NumberOfSuccessfulFollowIndices          int64                 `json:"number_of_successful_follow_indices"`
	RecentAutoFollowErrors                   []ErrorCause          `json:"recent_auto_follow_errors"`
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
				return err
			}

		case "number_of_failed_follow_indices":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NumberOfFailedFollowIndices = value
			case float64:
				f := int64(v)
				s.NumberOfFailedFollowIndices = f
			}

		case "number_of_failed_remote_cluster_state_requests":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NumberOfFailedRemoteClusterStateRequests = value
			case float64:
				f := int64(v)
				s.NumberOfFailedRemoteClusterStateRequests = f
			}

		case "number_of_successful_follow_indices":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.NumberOfSuccessfulFollowIndices = value
			case float64:
				f := int64(v)
				s.NumberOfSuccessfulFollowIndices = f
			}

		case "recent_auto_follow_errors":
			if err := dec.Decode(&s.RecentAutoFollowErrors); err != nil {
				return err
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
