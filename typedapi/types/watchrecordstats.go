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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/executionphase"
)

// WatchRecordStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/watcher/stats/types.ts#L54-L60
type WatchRecordStats struct {
	ExecutedActions []string                      `json:"executed_actions,omitempty"`
	ExecutionPhase  executionphase.ExecutionPhase `json:"execution_phase"`
	ExecutionTime   DateTime                      `json:"execution_time"`
	TriggeredTime   DateTime                      `json:"triggered_time"`
	WatchId         string                        `json:"watch_id"`
	WatchRecordId   string                        `json:"watch_record_id"`
}

func (s *WatchRecordStats) UnmarshalJSON(data []byte) error {

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

		case "executed_actions":
			if err := dec.Decode(&s.ExecutedActions); err != nil {
				return err
			}

		case "execution_phase":
			if err := dec.Decode(&s.ExecutionPhase); err != nil {
				return err
			}

		case "execution_time":
			if err := dec.Decode(&s.ExecutionTime); err != nil {
				return err
			}

		case "triggered_time":
			if err := dec.Decode(&s.TriggeredTime); err != nil {
				return err
			}

		case "watch_id":
			if err := dec.Decode(&s.WatchId); err != nil {
				return err
			}

		case "watch_record_id":
			if err := dec.Decode(&s.WatchRecordId); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewWatchRecordStats returns a WatchRecordStats.
func NewWatchRecordStats() *WatchRecordStats {
	r := &WatchRecordStats{}

	return r
}
