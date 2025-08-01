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

// CgroupCpu type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/nodes/_types/Stats.ts#L525-L542
type CgroupCpu struct {
	// CfsPeriodMicros The period of time, in microseconds, for how regularly all tasks in the same
	// cgroup as the Elasticsearch process should have their access to CPU resources
	// reallocated.
	CfsPeriodMicros *int `json:"cfs_period_micros,omitempty"`
	// CfsQuotaMicros The total amount of time, in microseconds, for which all tasks in the same
	// cgroup as the Elasticsearch process can run during one period
	// `cfs_period_micros`.
	CfsQuotaMicros *int `json:"cfs_quota_micros,omitempty"`
	// ControlGroup The `cpu` control group to which the Elasticsearch process belongs.
	ControlGroup *string `json:"control_group,omitempty"`
	// Stat Contains CPU statistics for the node.
	Stat *CgroupCpuStat `json:"stat,omitempty"`
}

func (s *CgroupCpu) UnmarshalJSON(data []byte) error {

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

		case "cfs_period_micros":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CfsPeriodMicros", err)
				}
				s.CfsPeriodMicros = &value
			case float64:
				f := int(v)
				s.CfsPeriodMicros = &f
			}

		case "cfs_quota_micros":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CfsQuotaMicros", err)
				}
				s.CfsQuotaMicros = &value
			case float64:
				f := int(v)
				s.CfsQuotaMicros = &f
			}

		case "control_group":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ControlGroup", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ControlGroup = &o

		case "stat":
			if err := dec.Decode(&s.Stat); err != nil {
				return fmt.Errorf("%s | %w", "Stat", err)
			}

		}
	}
	return nil
}

// NewCgroupCpu returns a CgroupCpu.
func NewCgroupCpu() *CgroupCpu {
	r := &CgroupCpu{}

	return r
}
