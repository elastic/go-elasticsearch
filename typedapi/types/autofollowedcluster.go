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
)

// AutoFollowedCluster type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ccr/stats/types.ts.ts#L27-L31
type AutoFollowedCluster struct {
	ClusterName              string `json:"cluster_name"`
	LastSeenMetadataVersion  int64  `json:"last_seen_metadata_version"`
	TimeSinceLastCheckMillis int64  `json:"time_since_last_check_millis"`
}

func (s *AutoFollowedCluster) UnmarshalJSON(data []byte) error {

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

		case "cluster_name":
			if err := dec.Decode(&s.ClusterName); err != nil {
				return err
			}

		case "last_seen_metadata_version":
			if err := dec.Decode(&s.LastSeenMetadataVersion); err != nil {
				return err
			}

		case "time_since_last_check_millis":
			if err := dec.Decode(&s.TimeSinceLastCheckMillis); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewAutoFollowedCluster returns a AutoFollowedCluster.
func NewAutoFollowedCluster() *AutoFollowedCluster {
	r := &AutoFollowedCluster{}

	return r
}
