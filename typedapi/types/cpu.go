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
	"strconv"
)

// Cpu type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/nodes/_types/Stats.ts#L218-L227
type Cpu struct {
	LoadAverage   map[string]Float64 `json:"load_average,omitempty"`
	Percent       *int               `json:"percent,omitempty"`
	Sys           Duration           `json:"sys,omitempty"`
	SysInMillis   *int64             `json:"sys_in_millis,omitempty"`
	Total         Duration           `json:"total,omitempty"`
	TotalInMillis *int64             `json:"total_in_millis,omitempty"`
	User          Duration           `json:"user,omitempty"`
	UserInMillis  *int64             `json:"user_in_millis,omitempty"`
}

func (s *Cpu) UnmarshalJSON(data []byte) error {

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

		case "load_average":
			if s.LoadAverage == nil {
				s.LoadAverage = make(map[string]Float64, 0)
			}
			if err := dec.Decode(&s.LoadAverage); err != nil {
				return err
			}

		case "percent":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Percent = &value
			case float64:
				f := int(v)
				s.Percent = &f
			}

		case "sys":
			if err := dec.Decode(&s.Sys); err != nil {
				return err
			}

		case "sys_in_millis":
			if err := dec.Decode(&s.SysInMillis); err != nil {
				return err
			}

		case "total":
			if err := dec.Decode(&s.Total); err != nil {
				return err
			}

		case "total_in_millis":
			if err := dec.Decode(&s.TotalInMillis); err != nil {
				return err
			}

		case "user":
			if err := dec.Decode(&s.User); err != nil {
				return err
			}

		case "user_in_millis":
			if err := dec.Decode(&s.UserInMillis); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewCpu returns a Cpu.
func NewCpu() *Cpu {
	r := &Cpu{
		LoadAverage: make(map[string]Float64, 0),
	}

	return r
}
