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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// RollupJobConfiguration type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/rollup/get_jobs/types.ts#L34-L43
type RollupJobConfiguration struct {
	Cron         string        `json:"cron"`
	Groups       Groupings     `json:"groups"`
	Id           string        `json:"id"`
	IndexPattern string        `json:"index_pattern"`
	Metrics      []FieldMetric `json:"metrics"`
	PageSize     int64         `json:"page_size"`
	RollupIndex  string        `json:"rollup_index"`
	Timeout      Duration      `json:"timeout"`
}

func (s *RollupJobConfiguration) UnmarshalJSON(data []byte) error {

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

		case "cron":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Cron = o

		case "groups":
			if err := dec.Decode(&s.Groups); err != nil {
				return err
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "index_pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.IndexPattern = o

		case "metrics":
			if err := dec.Decode(&s.Metrics); err != nil {
				return err
			}

		case "page_size":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PageSize = value
			case float64:
				f := int64(v)
				s.PageSize = f
			}

		case "rollup_index":
			if err := dec.Decode(&s.RollupIndex); err != nil {
				return err
			}

		case "timeout":
			if err := dec.Decode(&s.Timeout); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewRollupJobConfiguration returns a RollupJobConfiguration.
func NewRollupJobConfiguration() *RollupJobConfiguration {
	r := &RollupJobConfiguration{}

	return r
}
