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

// ModelSnapshot type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/Model.ts#L25-L46
type ModelSnapshot struct {
	// Description An optional description of the job.
	Description *string `json:"description,omitempty"`
	// JobId A numerical character string that uniquely identifies the job that the
	// snapshot was created for.
	JobId string `json:"job_id"`
	// LatestRecordTimeStamp The timestamp of the latest processed record.
	LatestRecordTimeStamp *int `json:"latest_record_time_stamp,omitempty"`
	// LatestResultTimeStamp The timestamp of the latest bucket result.
	LatestResultTimeStamp *int `json:"latest_result_time_stamp,omitempty"`
	// MinVersion The minimum version required to be able to restore the model snapshot.
	MinVersion string `json:"min_version"`
	// ModelSizeStats Summary information describing the model.
	ModelSizeStats *ModelSizeStats `json:"model_size_stats,omitempty"`
	// Retain If true, this snapshot will not be deleted during automatic cleanup of
	// snapshots older than model_snapshot_retention_days. However, this snapshot
	// will be deleted when the job is deleted. The default value is false.
	Retain bool `json:"retain"`
	// SnapshotDocCount For internal use only.
	SnapshotDocCount int64 `json:"snapshot_doc_count"`
	// SnapshotId A numerical character string that uniquely identifies the model snapshot.
	SnapshotId string `json:"snapshot_id"`
	// Timestamp The creation timestamp for the snapshot.
	Timestamp int64 `json:"timestamp"`
}

func (s *ModelSnapshot) UnmarshalJSON(data []byte) error {

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

		case "description":
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

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "latest_record_time_stamp":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.LatestRecordTimeStamp = &value
			case float64:
				f := int(v)
				s.LatestRecordTimeStamp = &f
			}

		case "latest_result_time_stamp":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.LatestResultTimeStamp = &value
			case float64:
				f := int(v)
				s.LatestResultTimeStamp = &f
			}

		case "min_version":
			if err := dec.Decode(&s.MinVersion); err != nil {
				return err
			}

		case "model_size_stats":
			if err := dec.Decode(&s.ModelSizeStats); err != nil {
				return err
			}

		case "retain":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Retain = value
			case bool:
				s.Retain = v
			}

		case "snapshot_doc_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SnapshotDocCount = value
			case float64:
				f := int64(v)
				s.SnapshotDocCount = f
			}

		case "snapshot_id":
			if err := dec.Decode(&s.SnapshotId); err != nil {
				return err
			}

		case "timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Timestamp = value
			case float64:
				f := int64(v)
				s.Timestamp = f
			}

		}
	}
	return nil
}

// NewModelSnapshot returns a ModelSnapshot.
func NewModelSnapshot() *ModelSnapshot {
	r := &ModelSnapshot{}

	return r
}
