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

// SnapshotsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cat/snapshots/types.ts#L24-L96
type SnapshotsRecord struct {
	// Duration The time it took the snapshot process to complete, in time units.
	Duration Duration `json:"duration,omitempty"`
	// EndEpoch The Unix epoch time (seconds since 1970-01-01 00:00:00) at which the snapshot
	// process ended.
	EndEpoch StringifiedEpochTimeUnitSeconds `json:"end_epoch,omitempty"`
	// EndTime The time (HH:MM:SS) at which the snapshot process ended.
	EndTime *string `json:"end_time,omitempty"`
	// FailedShards The number of failed shards in the snapshot.
	FailedShards *string `json:"failed_shards,omitempty"`
	// Id The unique identifier for the snapshot.
	Id *string `json:"id,omitempty"`
	// Indices The number of indices in the snapshot.
	Indices *string `json:"indices,omitempty"`
	// Reason The reason for any snapshot failures.
	Reason *string `json:"reason,omitempty"`
	// Repository The repository name.
	Repository *string `json:"repository,omitempty"`
	// StartEpoch The Unix epoch time (seconds since 1970-01-01 00:00:00) at which the snapshot
	// process started.
	StartEpoch StringifiedEpochTimeUnitSeconds `json:"start_epoch,omitempty"`
	// StartTime The time (HH:MM:SS) at which the snapshot process started.
	StartTime ScheduleTimeOfDay `json:"start_time,omitempty"`
	// Status The state of the snapshot process.
	// Returned values include:
	// `FAILED`: The snapshot process failed.
	// `INCOMPATIBLE`: The snapshot process is incompatible with the current cluster
	// version.
	// `IN_PROGRESS`: The snapshot process started but has not completed.
	// `PARTIAL`: The snapshot process completed with a partial success.
	// `SUCCESS`: The snapshot process completed with a full success.
	Status *string `json:"status,omitempty"`
	// SuccessfulShards The number of successful shards in the snapshot.
	SuccessfulShards *string `json:"successful_shards,omitempty"`
	// TotalShards The total number of shards in the snapshot.
	TotalShards *string `json:"total_shards,omitempty"`
}

func (s *SnapshotsRecord) UnmarshalJSON(data []byte) error {

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

		case "duration", "dur":
			if err := dec.Decode(&s.Duration); err != nil {
				return err
			}

		case "end_epoch", "ete", "endEpoch":
			if err := dec.Decode(&s.EndEpoch); err != nil {
				return err
			}

		case "end_time", "eti", "endTime":
			if err := dec.Decode(&s.EndTime); err != nil {
				return err
			}

		case "failed_shards", "fs":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FailedShards = &o

		case "id", "snapshot":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Id = &o

		case "indices", "i":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Indices = &o

		case "reason", "r":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Reason = &o

		case "repository", "re", "repo":
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

		case "start_epoch", "ste", "startEpoch":
			if err := dec.Decode(&s.StartEpoch); err != nil {
				return err
			}

		case "start_time", "sti", "startTime":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := NewHourAndMinute()
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.StartTime = *o

			default:
				if err := localDec.Decode(&s.StartTime); err != nil {
					return err
				}
			}

		case "status", "s":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Status = &o

		case "successful_shards", "ss":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SuccessfulShards = &o

		case "total_shards", "ts":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TotalShards = &o

		}
	}
	return nil
}

// NewSnapshotsRecord returns a SnapshotsRecord.
func NewSnapshotsRecord() *SnapshotsRecord {
	r := &SnapshotsRecord{}

	return r
}
