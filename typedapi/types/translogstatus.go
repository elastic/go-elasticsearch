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

// TranslogStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/indices/recovery/types.ts#L102-L109
type TranslogStatus struct {
	Percent           Percentage `json:"percent"`
	Recovered         int64      `json:"recovered"`
	Total             int64      `json:"total"`
	TotalOnStart      int64      `json:"total_on_start"`
	TotalTime         Duration   `json:"total_time,omitempty"`
	TotalTimeInMillis int64      `json:"total_time_in_millis"`
}

func (s *TranslogStatus) UnmarshalJSON(data []byte) error {

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

		case "percent":
			if err := dec.Decode(&s.Percent); err != nil {
				return fmt.Errorf("%s | %w", "Percent", err)
			}

		case "recovered":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Recovered", err)
				}
				s.Recovered = value
			case float64:
				f := int64(v)
				s.Recovered = f
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

		case "total_on_start":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalOnStart", err)
				}
				s.TotalOnStart = value
			case float64:
				f := int64(v)
				s.TotalOnStart = f
			}

		case "total_time":
			if err := dec.Decode(&s.TotalTime); err != nil {
				return fmt.Errorf("%s | %w", "TotalTime", err)
			}

		case "total_time_in_millis":
			if err := dec.Decode(&s.TotalTimeInMillis); err != nil {
				return fmt.Errorf("%s | %w", "TotalTimeInMillis", err)
			}

		}
	}
	return nil
}

// NewTranslogStatus returns a TranslogStatus.
func NewTranslogStatus() *TranslogStatus {
	r := &TranslogStatus{}

	return r
}
