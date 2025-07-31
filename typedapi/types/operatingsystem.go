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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// OperatingSystem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/nodes/_types/Stats.ts#L1023-L1029
type OperatingSystem struct {
	Cgroup    *Cgroup              `json:"cgroup,omitempty"`
	Cpu       *Cpu                 `json:"cpu,omitempty"`
	Mem       *ExtendedMemoryStats `json:"mem,omitempty"`
	Swap      *MemoryStats         `json:"swap,omitempty"`
	Timestamp *int64               `json:"timestamp,omitempty"`
}

func (s *OperatingSystem) UnmarshalJSON(data []byte) error {

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

		case "cgroup":
			if err := dec.Decode(&s.Cgroup); err != nil {
				return fmt.Errorf("%s | %w", "Cgroup", err)
			}

		case "cpu":
			if err := dec.Decode(&s.Cpu); err != nil {
				return fmt.Errorf("%s | %w", "Cpu", err)
			}

		case "mem":
			if err := dec.Decode(&s.Mem); err != nil {
				return fmt.Errorf("%s | %w", "Mem", err)
			}

		case "swap":
			if err := dec.Decode(&s.Swap); err != nil {
				return fmt.Errorf("%s | %w", "Swap", err)
			}

		case "timestamp":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Timestamp", err)
				}
				s.Timestamp = &value
			case float64:
				f := int64(v)
				s.Timestamp = &f
			}

		}
	}
	return nil
}

// NewOperatingSystem returns a OperatingSystem.
func NewOperatingSystem() *OperatingSystem {
	r := &OperatingSystem{}

	return r
}
