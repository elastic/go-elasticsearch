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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/followerindexstatus"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// FollowerIndex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ccr/follow_info/types.ts#L22-L28
type FollowerIndex struct {
	FollowerIndex string                                  `json:"follower_index"`
	LeaderIndex   string                                  `json:"leader_index"`
	Parameters    *FollowerIndexParameters                `json:"parameters,omitempty"`
	RemoteCluster string                                  `json:"remote_cluster"`
	Status        followerindexstatus.FollowerIndexStatus `json:"status"`
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
				return err
			}

		case "leader_index":
			if err := dec.Decode(&s.LeaderIndex); err != nil {
				return err
			}

		case "parameters":
			if err := dec.Decode(&s.Parameters); err != nil {
				return err
			}

		case "remote_cluster":
			if err := dec.Decode(&s.RemoteCluster); err != nil {
				return err
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return err
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
