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

// DataFrameAnalyticsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cat/ml_data_frame_analytics/types.ts#L22-L102
type DataFrameAnalyticsRecord struct {
	// AssignmentExplanation why the job is or is not assigned to a node
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// CreateTime job creation time
	CreateTime *string `json:"create_time,omitempty"`
	// Description description
	Description *string `json:"description,omitempty"`
	// DestIndex destination index
	DestIndex *string `json:"dest_index,omitempty"`
	// FailureReason failure reason
	FailureReason *string `json:"failure_reason,omitempty"`
	// Id the id
	Id *string `json:"id,omitempty"`
	// ModelMemoryLimit model memory limit
	ModelMemoryLimit *string `json:"model_memory_limit,omitempty"`
	// NodeAddress network address of the assigned node
	NodeAddress *string `json:"node.address,omitempty"`
	// NodeEphemeralId ephemeral id of the assigned node
	NodeEphemeralId *string `json:"node.ephemeral_id,omitempty"`
	// NodeId id of the assigned node
	NodeId *string `json:"node.id,omitempty"`
	// NodeName name of the assigned node
	NodeName *string `json:"node.name,omitempty"`
	// Progress progress
	Progress *string `json:"progress,omitempty"`
	// SourceIndex source index
	SourceIndex *string `json:"source_index,omitempty"`
	// State job state
	State *string `json:"state,omitempty"`
	// Type analysis type
	Type *string `json:"type,omitempty"`
	// Version the version of Elasticsearch when the analytics was created
	Version *string `json:"version,omitempty"`
}

func (s *DataFrameAnalyticsRecord) UnmarshalJSON(data []byte) error {

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

		case "assignment_explanation", "ae", "assignmentExplanation":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AssignmentExplanation = &o

		case "create_time", "ct", "createTime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CreateTime = &o

		case "description", "d":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "dest_index", "di", "destIndex":
			if err := dec.Decode(&s.DestIndex); err != nil {
				return err
			}

		case "failure_reason", "fr", "failureReason":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FailureReason = &o

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "model_memory_limit", "mml", "modelMemoryLimit":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelMemoryLimit = &o

		case "node.address", "na", "nodeAddress":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeAddress = &o

		case "node.ephemeral_id", "ne", "nodeEphemeralId":
			if err := dec.Decode(&s.NodeEphemeralId); err != nil {
				return err
			}

		case "node.id", "ni", "nodeId":
			if err := dec.Decode(&s.NodeId); err != nil {
				return err
			}

		case "node.name", "nn", "nodeName":
			if err := dec.Decode(&s.NodeName); err != nil {
				return err
			}

		case "progress", "p":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Progress = &o

		case "source_index", "si", "sourceIndex":
			if err := dec.Decode(&s.SourceIndex); err != nil {
				return err
			}

		case "state", "s":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.State = &o

		case "type", "t":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = &o

		case "version", "v":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataFrameAnalyticsRecord returns a DataFrameAnalyticsRecord.
func NewDataFrameAnalyticsRecord() *DataFrameAnalyticsRecord {
	r := &DataFrameAnalyticsRecord{}

	return r
}
