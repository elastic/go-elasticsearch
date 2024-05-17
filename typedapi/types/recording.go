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
// https://github.com/elastic/elasticsearch-specification/tree/9a0362eb2579c6604966a8fb307caee92de04270

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Recording type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9a0362eb2579c6604966a8fb307caee92de04270/specification/nodes/_types/Stats.ts#L225-L230
type Recording struct {
	CumulativeExecutionCount      *int64   `json:"cumulative_execution_count,omitempty"`
	CumulativeExecutionTime       Duration `json:"cumulative_execution_time,omitempty"`
	CumulativeExecutionTimeMillis *int64   `json:"cumulative_execution_time_millis,omitempty"`
	Name                          *string  `json:"name,omitempty"`
}

func (s *Recording) UnmarshalJSON(data []byte) error {

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

		case "cumulative_execution_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CumulativeExecutionCount", err)
				}
				s.CumulativeExecutionCount = &value
			case float64:
				f := int64(v)
				s.CumulativeExecutionCount = &f
			}

		case "cumulative_execution_time":
			if err := dec.Decode(&s.CumulativeExecutionTime); err != nil {
				return fmt.Errorf("%s | %w", "CumulativeExecutionTime", err)
			}

		case "cumulative_execution_time_millis":
			if err := dec.Decode(&s.CumulativeExecutionTimeMillis); err != nil {
				return fmt.Errorf("%s | %w", "CumulativeExecutionTimeMillis", err)
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		}
	}
	return nil
}

// NewRecording returns a Recording.
func NewRecording() *Recording {
	r := &Recording{}

	return r
}
