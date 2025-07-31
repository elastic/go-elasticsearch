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

// StagnatingBackingIndices type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/health_report/types.ts#L158-L162
type StagnatingBackingIndices struct {
	FirstOccurrenceTimestamp int64  `json:"first_occurrence_timestamp"`
	IndexName                string `json:"index_name"`
	RetryCount               int    `json:"retry_count"`
}

func (s *StagnatingBackingIndices) UnmarshalJSON(data []byte) error {

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

		case "first_occurrence_timestamp":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "FirstOccurrenceTimestamp", err)
				}
				s.FirstOccurrenceTimestamp = value
			case float64:
				f := int64(v)
				s.FirstOccurrenceTimestamp = f
			}

		case "index_name":
			if err := dec.Decode(&s.IndexName); err != nil {
				return fmt.Errorf("%s | %w", "IndexName", err)
			}

		case "retry_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "RetryCount", err)
				}
				s.RetryCount = value
			case float64:
				f := int(v)
				s.RetryCount = f
			}

		}
	}
	return nil
}

// NewStagnatingBackingIndices returns a StagnatingBackingIndices.
func NewStagnatingBackingIndices() *StagnatingBackingIndices {
	r := &StagnatingBackingIndices{}

	return r
}
