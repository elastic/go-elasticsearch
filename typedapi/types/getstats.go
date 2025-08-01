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

// GetStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/Stats.ts#L155-L166
type GetStats struct {
	Current             int64    `json:"current"`
	ExistsTime          Duration `json:"exists_time,omitempty"`
	ExistsTimeInMillis  int64    `json:"exists_time_in_millis"`
	ExistsTotal         int64    `json:"exists_total"`
	MissingTime         Duration `json:"missing_time,omitempty"`
	MissingTimeInMillis int64    `json:"missing_time_in_millis"`
	MissingTotal        int64    `json:"missing_total"`
	Time                Duration `json:"time,omitempty"`
	TimeInMillis        int64    `json:"time_in_millis"`
	Total               int64    `json:"total"`
}

func (s *GetStats) UnmarshalJSON(data []byte) error {

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

		case "current":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Current", err)
				}
				s.Current = value
			case float64:
				f := int64(v)
				s.Current = f
			}

		case "exists_time":
			if err := dec.Decode(&s.ExistsTime); err != nil {
				return fmt.Errorf("%s | %w", "ExistsTime", err)
			}

		case "exists_time_in_millis":
			if err := dec.Decode(&s.ExistsTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "ExistsTimeInMillis", err)
			}

		case "exists_total":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ExistsTotal", err)
				}
				s.ExistsTotal = value
			case float64:
				f := int64(v)
				s.ExistsTotal = f
			}

		case "missing_time":
			if err := dec.Decode(&s.MissingTime); err != nil {
				return fmt.Errorf("%s | %w", "MissingTime", err)
			}

		case "missing_time_in_millis":
			if err := dec.Decode(&s.MissingTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "MissingTimeInMillis", err)
			}

		case "missing_total":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MissingTotal", err)
				}
				s.MissingTotal = value
			case float64:
				f := int64(v)
				s.MissingTotal = f
			}

		case "time":
			if err := dec.Decode(&s.Time); err != nil {
				return fmt.Errorf("%s | %w", "Time", err)
			}

		case "time_in_millis":
			if err := dec.Decode(&s.TimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "TimeInMillis", err)
			}

		case "total":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Total", err)
				}
				s.Total = value
			case float64:
				f := int64(v)
				s.Total = f
			}

		}
	}
	return nil
}

// NewGetStats returns a GetStats.
func NewGetStats() *GetStats {
	r := &GetStats{}

	return r
}
