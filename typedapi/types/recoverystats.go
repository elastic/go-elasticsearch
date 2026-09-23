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
// https://github.com/elastic/elasticsearch-specification/tree/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// RecoveryStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6/specification/_types/Stats.ts#L258-L279
type RecoveryStats struct {
	CurrentAsSource        int64    `json:"current_as_source"`
	CurrentAsSourceQueued  *int64   `json:"current_as_source_queued,omitempty"`
	CurrentAsTarget        int64    `json:"current_as_target"`
	CurrentAsTargetQueued  *int64   `json:"current_as_target_queued,omitempty"`
	CurrentFromStore       *int64   `json:"current_from_store,omitempty"`
	CurrentFromStoreQueued *int64   `json:"current_from_store_queued,omitempty"`
	ThrottleTime           Duration `json:"throttle_time,omitempty"`
	ThrottleTimeInMillis   int64    `json:"throttle_time_in_millis"`
}

func (s *RecoveryStats) UnmarshalJSON(data []byte) error {

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

		case "current_as_source":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CurrentAsSource", err)
				}
				s.CurrentAsSource = value
			case float64:
				f := int64(v)
				s.CurrentAsSource = f
			}

		case "current_as_source_queued":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CurrentAsSourceQueued", err)
				}
				s.CurrentAsSourceQueued = &value
			case float64:
				f := int64(v)
				s.CurrentAsSourceQueued = &f
			}

		case "current_as_target":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CurrentAsTarget", err)
				}
				s.CurrentAsTarget = value
			case float64:
				f := int64(v)
				s.CurrentAsTarget = f
			}

		case "current_as_target_queued":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CurrentAsTargetQueued", err)
				}
				s.CurrentAsTargetQueued = &value
			case float64:
				f := int64(v)
				s.CurrentAsTargetQueued = &f
			}

		case "current_from_store":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CurrentFromStore", err)
				}
				s.CurrentFromStore = &value
			case float64:
				f := int64(v)
				s.CurrentFromStore = &f
			}

		case "current_from_store_queued":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CurrentFromStoreQueued", err)
				}
				s.CurrentFromStoreQueued = &value
			case float64:
				f := int64(v)
				s.CurrentFromStoreQueued = &f
			}

		case "throttle_time":
			if err := dec.Decode(&s.ThrottleTime); err != nil {
				return fmt.Errorf("%s | %w", "ThrottleTime", err)
			}

		case "throttle_time_in_millis":
			if err := dec.Decode(&s.ThrottleTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "ThrottleTimeInMillis", err)
			}

		}
	}
	return nil
}

// NewRecoveryStats returns a RecoveryStats.
func NewRecoveryStats() *RecoveryStats {
	r := &RecoveryStats{}

	return r
}
