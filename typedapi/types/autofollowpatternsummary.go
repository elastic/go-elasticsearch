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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// AutoFollowPatternSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ccr/get_auto_follow_pattern/types.ts#L28-L52
type AutoFollowPatternSummary struct {
	Active bool `json:"active"`
	// FollowIndexPattern The name of follower index.
	FollowIndexPattern *string `json:"follow_index_pattern,omitempty"`
	// LeaderIndexExclusionPatterns An array of simple index patterns that can be used to exclude indices from
	// being auto-followed.
	LeaderIndexExclusionPatterns []string `json:"leader_index_exclusion_patterns"`
	// LeaderIndexPatterns An array of simple index patterns to match against indices in the remote
	// cluster specified by the remote_cluster field.
	LeaderIndexPatterns []string `json:"leader_index_patterns"`
	// MaxOutstandingReadRequests The maximum number of outstanding reads requests from the remote cluster.
	MaxOutstandingReadRequests int `json:"max_outstanding_read_requests"`
	// RemoteCluster The remote cluster containing the leader indices to match against.
	RemoteCluster string `json:"remote_cluster"`
}

func (s *AutoFollowPatternSummary) UnmarshalJSON(data []byte) error {

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

		case "active":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Active = value
			case bool:
				s.Active = v
			}

		case "follow_index_pattern":
			if err := dec.Decode(&s.FollowIndexPattern); err != nil {
				return err
			}

		case "leader_index_exclusion_patterns":
			if err := dec.Decode(&s.LeaderIndexExclusionPatterns); err != nil {
				return err
			}

		case "leader_index_patterns":
			if err := dec.Decode(&s.LeaderIndexPatterns); err != nil {
				return err
			}

		case "max_outstanding_read_requests":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxOutstandingReadRequests = value
			case float64:
				f := int(v)
				s.MaxOutstandingReadRequests = f
			}

		case "remote_cluster":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RemoteCluster = o

		}
	}
	return nil
}

// NewAutoFollowPatternSummary returns a AutoFollowPatternSummary.
func NewAutoFollowPatternSummary() *AutoFollowPatternSummary {
	r := &AutoFollowPatternSummary{}

	return r
}
