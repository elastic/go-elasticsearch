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

// DataStreamLifecycleDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/health_report/types.ts#L153-L157
type DataStreamLifecycleDetails struct {
	StagnatingBackingIndices      []StagnatingBackingIndices `json:"stagnating_backing_indices,omitempty"`
	StagnatingBackingIndicesCount int                        `json:"stagnating_backing_indices_count"`
	TotalBackingIndicesInError    int                        `json:"total_backing_indices_in_error"`
}

func (s *DataStreamLifecycleDetails) UnmarshalJSON(data []byte) error {

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

		case "stagnating_backing_indices":
			if err := dec.Decode(&s.StagnatingBackingIndices); err != nil {
				return fmt.Errorf("%s | %w", "StagnatingBackingIndices", err)
			}

		case "stagnating_backing_indices_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "StagnatingBackingIndicesCount", err)
				}
				s.StagnatingBackingIndicesCount = value
			case float64:
				f := int(v)
				s.StagnatingBackingIndicesCount = f
			}

		case "total_backing_indices_in_error":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalBackingIndicesInError", err)
				}
				s.TotalBackingIndicesInError = value
			case float64:
				f := int(v)
				s.TotalBackingIndicesInError = f
			}

		}
	}
	return nil
}

// NewDataStreamLifecycleDetails returns a DataStreamLifecycleDetails.
func NewDataStreamLifecycleDetails() *DataStreamLifecycleDetails {
	r := &DataStreamLifecycleDetails{}

	return r
}
