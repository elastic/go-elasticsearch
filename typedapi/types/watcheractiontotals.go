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
// https://github.com/elastic/elasticsearch-specification/tree/3ea9ce260df22d3244bff5bace485dd97ff4046d

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// WatcherActionTotals type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3ea9ce260df22d3244bff5bace485dd97ff4046d/specification/xpack/usage/types.ts#L422-L425
type WatcherActionTotals struct {
	Total         Duration `json:"total"`
	TotalTimeInMs int64    `json:"total_time_in_ms"`
}

func (s *WatcherActionTotals) UnmarshalJSON(data []byte) error {

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

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return fmt.Errorf("%s | %w", "Total", err)
			}

		case "total_time_in_ms":
			if err := dec.Decode(&s.TotalTimeInMs); err != nil {
				return fmt.Errorf("%s | %w", "TotalTimeInMs", err)
			}

		}
	}
	return nil
}

// NewWatcherActionTotals returns a WatcherActionTotals.
func NewWatcherActionTotals() *WatcherActionTotals {
	r := &WatcherActionTotals{}

	return r
}

// false
