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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/followerindexstatus"
)

// FollowerIndex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ccr/follow_info/types.ts#L24-L35
type FollowerIndex struct {
	// FollowerIndex The name of the follower index.
	FollowerIndex string `json:"follower_index"`
	// LeaderIndex The name of the index in the leader cluster that is followed.
	LeaderIndex string `json:"leader_index"`
	// Parameters An object that encapsulates cross-cluster replication parameters. If the
	// follower index's status is paused, this object is omitted.
	Parameters *FollowerIndexParameters `json:"parameters,omitempty"`
	// RemoteCluster The remote cluster that contains the leader index.
	RemoteCluster string `json:"remote_cluster"`
	// Status The status of the index following: `active` or `paused`.
	Status followerindexstatus.FollowerIndexStatus `json:"status"`
}

func (s *FollowerIndex) UnmarshalJSON(data []byte) error {

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

		case "follower_index":
			if err := dec.Decode(&s.FollowerIndex); err != nil {
				return fmt.Errorf("%s | %w", "FollowerIndex", err)
			}

		case "leader_index":
			if err := dec.Decode(&s.LeaderIndex); err != nil {
				return fmt.Errorf("%s | %w", "LeaderIndex", err)
			}

		case "parameters":
			if err := dec.Decode(&s.Parameters); err != nil {
				return fmt.Errorf("%s | %w", "Parameters", err)
			}

		case "remote_cluster":
			if err := dec.Decode(&s.RemoteCluster); err != nil {
				return fmt.Errorf("%s | %w", "RemoteCluster", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		}
	}
	return nil
}

// NewFollowerIndex returns a FollowerIndex.
func NewFollowerIndex() *FollowerIndex {
	r := &FollowerIndex{}

	return r
}
