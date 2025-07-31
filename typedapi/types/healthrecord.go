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

// HealthRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cat/health/types.ts#L23-L99
type HealthRecord struct {
	// ActiveShardsPercent active number of shards in percent
	ActiveShardsPercent *string `json:"active_shards_percent,omitempty"`
	// Cluster cluster name
	Cluster *string `json:"cluster,omitempty"`
	// Epoch seconds since 1970-01-01 00:00:00
	Epoch StringifiedEpochTimeUnitSeconds `json:"epoch,omitempty"`
	// Init number of initializing nodes
	Init *string `json:"init,omitempty"`
	// MaxTaskWaitTime wait time of longest task pending
	MaxTaskWaitTime *string `json:"max_task_wait_time,omitempty"`
	// NodeData number of nodes that can store data
	NodeData *string `json:"node.data,omitempty"`
	// NodeTotal total number of nodes
	NodeTotal *string `json:"node.total,omitempty"`
	// PendingTasks number of pending tasks
	PendingTasks *string `json:"pending_tasks,omitempty"`
	// Pri number of primary shards
	Pri *string `json:"pri,omitempty"`
	// Relo number of relocating nodes
	Relo *string `json:"relo,omitempty"`
	// Shards total number of shards
	Shards *string `json:"shards,omitempty"`
	// Status health status
	Status *string `json:"status,omitempty"`
	// Timestamp time in HH:MM:SS
	Timestamp *string `json:"timestamp,omitempty"`
	// Unassign number of unassigned shards
	Unassign *string `json:"unassign,omitempty"`
	// UnassignPri number of unassigned primary shards
	UnassignPri *string `json:"unassign.pri,omitempty"`
}

func (s *HealthRecord) UnmarshalJSON(data []byte) error {

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

		case "active_shards_percent", "asp", "activeShardsPercent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ActiveShardsPercent", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ActiveShardsPercent = &o

		case "cluster", "cl":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Cluster", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Cluster = &o

		case "epoch", "time":
			if err := dec.Decode(&s.Epoch); err != nil {
				return fmt.Errorf("%s | %w", "Epoch", err)
			}

		case "init", "i", "shards.initializing", "shardsInitializing":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Init", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Init = &o

		case "max_task_wait_time", "mtwt", "maxTaskWaitTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MaxTaskWaitTime", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxTaskWaitTime = &o

		case "node.data", "nd", "nodeData":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "NodeData", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeData = &o

		case "node.total", "nt", "nodeTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "NodeTotal", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeTotal = &o

		case "pending_tasks", "pt", "pendingTasks":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "PendingTasks", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PendingTasks = &o

		case "pri", "p", "shards.primary", "shardsPrimary":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Pri", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pri = &o

		case "relo", "r", "shards.relocating", "shardsRelocating":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Relo", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Relo = &o

		case "shards", "t", "sh", "shards.total", "shardsTotal":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Shards", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Shards = &o

		case "status", "st":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Status = &o

		case "timestamp", "ts", "hms", "hhmmss":
			if err := dec.Decode(&s.Timestamp); err != nil {
				return fmt.Errorf("%s | %w", "Timestamp", err)
			}

		case "unassign", "u", "shards.unassigned", "shardsUnassigned":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Unassign", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Unassign = &o

		case "unassign.pri", "up", "shards.unassigned.primary", "shardsUnassignedPrimary":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "UnassignPri", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UnassignPri = &o

		}
	}
	return nil
}

// NewHealthRecord returns a HealthRecord.
func NewHealthRecord() *HealthRecord {
	r := &HealthRecord{}

	return r
}
