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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ShardRecovery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/recovery/types.ts#L118-L135
type ShardRecovery struct {
	Id                int64                `json:"id"`
	Index             RecoveryIndexStatus  `json:"index"`
	Primary           bool                 `json:"primary"`
	Source            RecoveryOrigin       `json:"source"`
	Stage             string               `json:"stage"`
	Start             *RecoveryStartStatus `json:"start,omitempty"`
	StartTime         DateTime             `json:"start_time,omitempty"`
	StartTimeInMillis int64                `json:"start_time_in_millis"`
	StopTime          DateTime             `json:"stop_time,omitempty"`
	StopTimeInMillis  *int64               `json:"stop_time_in_millis,omitempty"`
	Target            RecoveryOrigin       `json:"target"`
	TotalTime         Duration             `json:"total_time,omitempty"`
	TotalTimeInMillis int64                `json:"total_time_in_millis"`
	Translog          TranslogStatus       `json:"translog"`
	Type              string               `json:"type"`
	VerifyIndex       VerifyIndex          `json:"verify_index"`
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

		case "source":
			if err := dec.Decode(&s.Source); err != nil {
				return fmt.Errorf("%s | %w", "Source", err)
			}

		case "stage":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Stage", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Stage = o

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
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = o

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
