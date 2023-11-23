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
)

// ClusterStateUpdate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/nodes/_types/Stats.ts#L278-L343
type ClusterStateUpdate struct {
	// CommitTime The cumulative amount of time spent waiting for a successful cluster state
	// update to commit, which measures the time from the start of each publication
	// until a majority of the master-eligible nodes have written the state to disk
	// and confirmed the write to the elected master.
	CommitTime Duration `json:"commit_time,omitempty"`
	// CommitTimeMillis The cumulative amount of time, in milliseconds, spent waiting for a
	// successful cluster state update to commit, which measures the time from the
	// start of each publication until a majority of the master-eligible nodes have
	// written the state to disk and confirmed the write to the elected master.
	CommitTimeMillis *int64 `json:"commit_time_millis,omitempty"`
	// CompletionTime The cumulative amount of time spent waiting for a successful cluster state
	// update to complete, which measures the time from the start of each
	// publication until all the other nodes have notified the elected master that
	// they have applied the cluster state.
	CompletionTime Duration `json:"completion_time,omitempty"`
	// CompletionTimeMillis The cumulative amount of time, in milliseconds,  spent waiting for a
	// successful cluster state update to complete, which measures the time from the
	// start of each publication until all the other nodes have notified the elected
	// master that they have applied the cluster state.
	CompletionTimeMillis *int64 `json:"completion_time_millis,omitempty"`
	// ComputationTime The cumulative amount of time spent computing no-op cluster state updates
	// since the node started.
	ComputationTime Duration `json:"computation_time,omitempty"`
	// ComputationTimeMillis The cumulative amount of time, in milliseconds, spent computing no-op cluster
	// state updates since the node started.
	ComputationTimeMillis *int64 `json:"computation_time_millis,omitempty"`
	// ContextConstructionTime The cumulative amount of time spent constructing a publication context since
	// the node started for publications that ultimately succeeded.
	// This statistic includes the time spent computing the difference between the
	// current and new cluster state preparing a serialized representation of this
	// difference.
	ContextConstructionTime Duration `json:"context_construction_time,omitempty"`
	// ContextConstructionTimeMillis The cumulative amount of time, in milliseconds, spent constructing a
	// publication context since the node started for publications that ultimately
	// succeeded.
	// This statistic includes the time spent computing the difference between the
	// current and new cluster state preparing a serialized representation of this
	// difference.
	ContextConstructionTimeMillis *int64 `json:"context_construction_time_millis,omitempty"`
	// Count The number of cluster state update attempts that did not change the cluster
	// state since the node started.
	Count int64 `json:"count"`
	// MasterApplyTime The cumulative amount of time spent successfully applying cluster state
	// updates on the elected master since the node started.
	MasterApplyTime Duration `json:"master_apply_time,omitempty"`
	// MasterApplyTimeMillis The cumulative amount of time, in milliseconds, spent successfully applying
	// cluster state updates on the elected master since the node started.
	MasterApplyTimeMillis *int64 `json:"master_apply_time_millis,omitempty"`
	// NotificationTime The cumulative amount of time spent notifying listeners of a no-op cluster
	// state update since the node started.
	NotificationTime Duration `json:"notification_time,omitempty"`
	// NotificationTimeMillis The cumulative amount of time, in milliseconds, spent notifying listeners of
	// a no-op cluster state update since the node started.
	NotificationTimeMillis *int64 `json:"notification_time_millis,omitempty"`
	// PublicationTime The cumulative amount of time spent publishing cluster state updates which
	// ultimately succeeded, which includes everything from the start of the
	// publication (just after the computation of the new cluster state) until the
	// publication has finished and the master node is ready to start processing the
	// next state update.
	// This includes the time measured by `context_construction_time`,
	// `commit_time`, `completion_time` and `master_apply_time`.
	PublicationTime Duration `json:"publication_time,omitempty"`
	// PublicationTimeMillis The cumulative amount of time, in milliseconds, spent publishing cluster
	// state updates which ultimately succeeded, which includes everything from the
	// start of the publication (just after the computation of the new cluster
	// state) until the publication has finished and the master node is ready to
	// start processing the next state update.
	// This includes the time measured by `context_construction_time`,
	// `commit_time`, `completion_time` and `master_apply_time`.
	PublicationTimeMillis *int64 `json:"publication_time_millis,omitempty"`
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
