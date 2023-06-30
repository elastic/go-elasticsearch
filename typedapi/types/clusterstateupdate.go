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

// ClusterStateUpdate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/nodes/_types/Stats.ts#L126-L142
type ClusterStateUpdate struct {
	CommitTime                    Duration `json:"commit_time,omitempty"`
	CommitTimeMillis              *int64   `json:"commit_time_millis,omitempty"`
	CompletionTime                Duration `json:"completion_time,omitempty"`
	CompletionTimeMillis          *int64   `json:"completion_time_millis,omitempty"`
	ComputationTime               Duration `json:"computation_time,omitempty"`
	ComputationTimeMillis         *int64   `json:"computation_time_millis,omitempty"`
	ContextConstructionTime       Duration `json:"context_construction_time,omitempty"`
	ContextConstructionTimeMillis *int64   `json:"context_construction_time_millis,omitempty"`
	Count                         int64    `json:"count"`
	MasterApplyTime               Duration `json:"master_apply_time,omitempty"`
	MasterApplyTimeMillis         *int64   `json:"master_apply_time_millis,omitempty"`
	NotificationTime              Duration `json:"notification_time,omitempty"`
	NotificationTimeMillis        *int64   `json:"notification_time_millis,omitempty"`
	PublicationTime               Duration `json:"publication_time,omitempty"`
	PublicationTimeMillis         *int64   `json:"publication_time_millis,omitempty"`
}

func (s *ClusterStateUpdate) UnmarshalJSON(data []byte) error {

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

		case "commit_time":
			if err := dec.Decode(&s.CommitTime); err != nil {
				return err
			}

		case "commit_time_millis":
			if err := dec.Decode(&s.CommitTimeMillis); err != nil {
				return err
			}

		case "completion_time":
			if err := dec.Decode(&s.CompletionTime); err != nil {
				return err
			}

		case "completion_time_millis":
			if err := dec.Decode(&s.CompletionTimeMillis); err != nil {
				return err
			}

		case "computation_time":
			if err := dec.Decode(&s.ComputationTime); err != nil {
				return err
			}

		case "computation_time_millis":
			if err := dec.Decode(&s.ComputationTimeMillis); err != nil {
				return err
			}

		case "context_construction_time":
			if err := dec.Decode(&s.ContextConstructionTime); err != nil {
				return err
			}

		case "context_construction_time_millis":
			if err := dec.Decode(&s.ContextConstructionTimeMillis); err != nil {
				return err
			}

		case "count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "master_apply_time":
			if err := dec.Decode(&s.MasterApplyTime); err != nil {
				return err
			}

		case "master_apply_time_millis":
			if err := dec.Decode(&s.MasterApplyTimeMillis); err != nil {
				return err
			}

		case "notification_time":
			if err := dec.Decode(&s.NotificationTime); err != nil {
				return err
			}

		case "notification_time_millis":
			if err := dec.Decode(&s.NotificationTimeMillis); err != nil {
				return err
			}

		case "publication_time":
			if err := dec.Decode(&s.PublicationTime); err != nil {
				return err
			}

		case "publication_time_millis":
			if err := dec.Decode(&s.PublicationTimeMillis); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewClusterStateUpdate returns a ClusterStateUpdate.
func NewClusterStateUpdate() *ClusterStateUpdate {
	r := &ClusterStateUpdate{}

	return r
}
