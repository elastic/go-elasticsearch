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

// RecoveryRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cat/recovery/types.ts#L24-L155
type RecoveryRecord struct {
	// Bytes The number of bytes to recover.
	Bytes *string `json:"bytes,omitempty"`
	// BytesPercent The ratio of bytes recovered.
	BytesPercent Percentage `json:"bytes_percent,omitempty"`
	// BytesRecovered The bytes recovered.
	BytesRecovered *string `json:"bytes_recovered,omitempty"`
	// BytesTotal The total number of bytes.
	BytesTotal *string `json:"bytes_total,omitempty"`
	// Files The number of files to recover.
	Files *string `json:"files,omitempty"`
	// FilesPercent The ratio of files recovered.
	FilesPercent Percentage `json:"files_percent,omitempty"`
	// FilesRecovered The files recovered.
	FilesRecovered *string `json:"files_recovered,omitempty"`
	// FilesTotal The total number of files.
	FilesTotal *string `json:"files_total,omitempty"`
	// Index The index name.
	Index *string `json:"index,omitempty"`
	// Repository The repository name.
	Repository *string `json:"repository,omitempty"`
	// Shard The shard name.
	Shard *string `json:"shard,omitempty"`
	// Snapshot The snapshot name.
	Snapshot *string `json:"snapshot,omitempty"`
	// SourceHost The source host.
	SourceHost *string `json:"source_host,omitempty"`
	// SourceNode The source node name.
	SourceNode *string `json:"source_node,omitempty"`
	// Stage The recovery stage.
	Stage *string `json:"stage,omitempty"`
	// StartTime The recovery start time.
	StartTime DateTime `json:"start_time,omitempty"`
	// StartTimeMillis The recovery start time in epoch milliseconds.
	StartTimeMillis *int64 `json:"start_time_millis,omitempty"`
	// StopTime The recovery stop time.
	StopTime DateTime `json:"stop_time,omitempty"`
	// StopTimeMillis The recovery stop time in epoch milliseconds.
	StopTimeMillis *int64 `json:"stop_time_millis,omitempty"`
	// TargetHost The target host.
	TargetHost *string `json:"target_host,omitempty"`
	// TargetNode The target node name.
	TargetNode *string `json:"target_node,omitempty"`
	// Time The recovery time.
	Time Duration `json:"time,omitempty"`
	// TranslogOps The number of translog operations to recover.
	TranslogOps *string `json:"translog_ops,omitempty"`
	// TranslogOpsPercent The ratio of translog operations recovered.
	TranslogOpsPercent Percentage `json:"translog_ops_percent,omitempty"`
	// TranslogOpsRecovered The translog operations recovered.
	TranslogOpsRecovered *string `json:"translog_ops_recovered,omitempty"`
	// Type The recovery type.
	Type *string `json:"type,omitempty"`
}

func (s *RecoveryRecord) UnmarshalJSON(data []byte) error {

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

		case "bytes", "b":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Bytes = &o

		case "bytes_percent", "bp":
			if err := dec.Decode(&s.BytesPercent); err != nil {
				return err
			}

		case "bytes_recovered", "br":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BytesRecovered = &o

		case "bytes_total", "tb":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BytesTotal = &o

		case "files", "f":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Files = &o

		case "files_percent", "fp":
			if err := dec.Decode(&s.FilesPercent); err != nil {
				return err
			}

		case "files_recovered", "fr":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FilesRecovered = &o

		case "files_total", "tf":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FilesTotal = &o

		case "index", "i", "idx":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "repository", "rep":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Repository = &o

		case "shard", "s", "sh":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Shard = &o

		case "snapshot", "snap":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Snapshot = &o

		case "source_host", "shost":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SourceHost = &o

		case "source_node", "snode":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SourceNode = &o

		case "stage", "st":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Stage = &o

		case "start_time", "start":
			if err := dec.Decode(&s.StartTime); err != nil {
				return err
			}

		case "start_time_millis", "start_millis":
			if err := dec.Decode(&s.StartTimeMillis); err != nil {
				return err
			}

		case "stop_time", "stop":
			if err := dec.Decode(&s.StopTime); err != nil {
				return err
			}

		case "stop_time_millis", "stop_millis":
			if err := dec.Decode(&s.StopTimeMillis); err != nil {
				return err
			}

		case "target_host", "thost":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TargetHost = &o

		case "target_node", "tnode":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TargetNode = &o

		case "time", "t", "ti":
			if err := dec.Decode(&s.Time); err != nil {
				return err
			}

		case "translog_ops", "to":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TranslogOps = &o

		case "translog_ops_percent", "top":
			if err := dec.Decode(&s.TranslogOpsPercent); err != nil {
				return err
			}

		case "translog_ops_recovered", "tor":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TranslogOpsRecovered = &o

		case "type", "ty":
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

		}
	}
	return nil
}

// NewRecoveryRecord returns a RecoveryRecord.
func NewRecoveryRecord() *RecoveryRecord {
	r := &RecoveryRecord{}

	return r
}
