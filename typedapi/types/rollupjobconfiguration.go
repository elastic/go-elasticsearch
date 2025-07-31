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

// RollupJobConfiguration type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/rollup/get_jobs/types.ts#L45-L54
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
				return fmt.Errorf("%s | %w", "Cron", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Cron = o

		case "groups":
			if err := dec.Decode(&s.Groups); err != nil {
				return fmt.Errorf("%s | %w", "Groups", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "index_pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "IndexPattern", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.IndexPattern = o

		case "metrics":
			if err := dec.Decode(&s.Metrics); err != nil {
				return fmt.Errorf("%s | %w", "Metrics", err)
			}

		case "page_size":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PageSize", err)
				}
				s.PageSize = value
			case float64:
				f := int64(v)
				s.PageSize = f
			}

		case "rollup_index":
			if err := dec.Decode(&s.RollupIndex); err != nil {
				return fmt.Errorf("%s | %w", "RollupIndex", err)
			}

		case "timeout":
			if err := dec.Decode(&s.Timeout); err != nil {
				return fmt.Errorf("%s | %w", "Timeout", err)
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
