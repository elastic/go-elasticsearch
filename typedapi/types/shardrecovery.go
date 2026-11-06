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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/recoverypriority"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/recoverystage"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/recoverytype"
)

// ShardRecovery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/indices/recovery/types.ts#L171-L237
type ShardRecovery struct {
	// BlockedForMillis The elapsed time in milliseconds recovery is blocked for.
	//
	// Only present with `gate`.
	BlockedForMillis *int64 `json:"blocked_for_millis,omitempty"`
	// Gate The name of the recovery gate that blocked recovery on the target node.
	//
	// Only present for queued recoveries in the `CREATED` stage that are blocked by
	// a recovery gate.
	Gate  *string             `json:"gate,omitempty"`
	Id    int64               `json:"id"`
	Index RecoveryIndexStatus `json:"index"`
	// LocalRetries The number of times this recovery has failed in a way which is retried
	// locally (i.e. on the data node).
	LocalRetries *int `json:"local_retries,omitempty"`
	Primary      bool `json:"primary"`
	// Priority The recovery priority.
	Priority *recoverypriority.RecoveryPriority `json:"priority,omitempty"`
	Source   RecoveryOrigin                     `json:"source"`
	// Stage The recovery stage.
	Stage recoverystage.RecoveryStage `json:"stage"`
	Start *RecoveryStartStatus        `json:"start,omitempty"`
	// StartTime The time the recovery started. For recoveries in the `CREATED` stage (not yet
	// started), this value is the Unix epoch (1970-01-01T00:00:00.000Z).
	StartTime DateTime `json:"start_time,omitempty"`
	// StartTimeInMillis The time the recovery started, in milliseconds since the Unix epoch. For
	// recoveries in the `CREATED` stage (not yet started), this value is 0.
	StartTimeInMillis int64 `json:"start_time_in_millis"`
	// StopTime The time the recovery completed. Only present for completed recoveries
	// (`DONE` stage).
	StopTime DateTime `json:"stop_time,omitempty"`
	// StopTimeInMillis The time the recovery completed, in milliseconds since the Unix epoch. Only
	// present for completed recoveries (`DONE` stage).
	StopTimeInMillis *int64         `json:"stop_time_in_millis,omitempty"`
	Target           RecoveryOrigin `json:"target"`
	TotalTime        Duration       `json:"total_time,omitempty"`
	// TotalTimeInMillis The total elapsed recovery time in milliseconds. For recoveries in the
	// `CREATED` stage (not yet started), this value is 0.
	TotalTimeInMillis int64          `json:"total_time_in_millis"`
	Translog          TranslogStatus `json:"translog"`
	// Type The recovery source type.
	Type        recoverytype.RecoveryType `json:"type"`
	VerifyIndex VerifyIndex               `json:"verify_index"`
}

func (s *ShardRecovery) UnmarshalJSON(data []byte) error {

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

		case "blocked_for_millis":
			if err := dec.Decode(&s.BlockedForMillis); err != nil {
				return fmt.Errorf("%s | %w", "BlockedForMillis", err)
			}

		case "gate":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Gate", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Gate = &o

		case "id":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Id", err)
				}
				s.Id = value
			case float64:
				f := int64(v)
				s.Id = f
			}

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return fmt.Errorf("%s | %w", "Index", err)
			}

		case "local_retries":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "LocalRetries", err)
				}
				s.LocalRetries = &value
			case float64:
				f := int(v)
				s.LocalRetries = &f
			}

		case "primary":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Primary", err)
				}
				s.Primary = value
			case bool:
				s.Primary = v
			}

		case "priority":
			if err := dec.Decode(&s.Priority); err != nil {
				return fmt.Errorf("%s | %w", "Priority", err)
			}

		case "source":
			if err := dec.Decode(&s.Source); err != nil {
				return fmt.Errorf("%s | %w", "Source", err)
			}

		case "stage":
			if err := dec.Decode(&s.Stage); err != nil {
				return fmt.Errorf("%s | %w", "Stage", err)
			}

		case "start":
			if err := dec.Decode(&s.Start); err != nil {
				return fmt.Errorf("%s | %w", "Start", err)
			}

		case "start_time":
			if err := dec.Decode(&s.StartTime); err != nil {
				return fmt.Errorf("%s | %w", "StartTime", err)
			}

		case "start_time_in_millis":
			if err := dec.Decode(&s.StartTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "StartTimeInMillis", err)
			}

		case "stop_time":
			if err := dec.Decode(&s.StopTime); err != nil {
				return fmt.Errorf("%s | %w", "StopTime", err)
			}

		case "stop_time_in_millis":
			if err := dec.Decode(&s.StopTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "StopTimeInMillis", err)
			}

		case "target":
			if err := dec.Decode(&s.Target); err != nil {
				return fmt.Errorf("%s | %w", "Target", err)
			}

		case "total_time":
			if err := dec.Decode(&s.TotalTime); err != nil {
				return fmt.Errorf("%s | %w", "TotalTime", err)
			}

		case "total_time_in_millis":
			if err := dec.Decode(&s.TotalTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "TotalTimeInMillis", err)
			}

		case "translog":
			if err := dec.Decode(&s.Translog); err != nil {
				return fmt.Errorf("%s | %w", "Translog", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "verify_index":
			if err := dec.Decode(&s.VerifyIndex); err != nil {
				return fmt.Errorf("%s | %w", "VerifyIndex", err)
			}

		}
	}
	return nil
}

// NewShardRecovery returns a ShardRecovery.
func NewShardRecovery() *ShardRecovery {
	r := &ShardRecovery{}

	return r
}
