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

// DataFrameAnalyticsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cat/ml_data_frame_analytics/types.ts#L22-L102
type DataFrameAnalyticsRecord struct {
	// AssignmentExplanation Messages related to the selection of a node.
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// CreateTime The time when the job was created.
	CreateTime *string `json:"create_time,omitempty"`
	// Description A description of the job.
	Description *string `json:"description,omitempty"`
	// DestIndex The name of the destination index.
	DestIndex *string `json:"dest_index,omitempty"`
	// FailureReason Messages about the reason why the job failed.
	FailureReason *string `json:"failure_reason,omitempty"`
	// Id The identifier for the job.
	Id *string `json:"id,omitempty"`
	// ModelMemoryLimit The approximate maximum amount of memory resources that are permitted for the
	// job.
	ModelMemoryLimit *string `json:"model_memory_limit,omitempty"`
	// NodeAddress The network address of the assigned node.
	NodeAddress *string `json:"node.address,omitempty"`
	// NodeEphemeralId The ephemeral identifier of the assigned node.
	NodeEphemeralId *string `json:"node.ephemeral_id,omitempty"`
	// NodeId The unique identifier of the assigned node.
	NodeId *string `json:"node.id,omitempty"`
	// NodeName The name of the assigned node.
	NodeName *string `json:"node.name,omitempty"`
	// Progress The progress report for the job by phase.
	Progress *string `json:"progress,omitempty"`
	// SourceIndex The name of the source index.
	SourceIndex *string `json:"source_index,omitempty"`
	// State The current status of the job.
	State *string `json:"state,omitempty"`
	// Type The type of analysis that the job performs.
	Type *string `json:"type,omitempty"`
	// Version The version of Elasticsearch when the job was created.
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
