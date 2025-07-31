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
)

// CCSUsageTimeValue type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L846-L853
type CCSUsageTimeValue struct {
	// Avg The average time taken to execute a request, in milliseconds.
	Avg int64 `json:"avg"`
	// Max The maximum time taken to execute a request, in milliseconds.
	Max int64 `json:"max"`
	// P90 The 90th percentile of the time taken to execute requests, in milliseconds.
	P90 int64 `json:"p90"`
}

func (s *CCSUsageTimeValue) UnmarshalJSON(data []byte) error {

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

		case "avg":
			if err := dec.Decode(&s.Avg); err != nil {
				return fmt.Errorf("%s | %w", "Avg", err)
			}

		case "max":
			if err := dec.Decode(&s.Max); err != nil {
				return fmt.Errorf("%s | %w", "Max", err)
			}

		case "p90":
			if err := dec.Decode(&s.P90); err != nil {
				return fmt.Errorf("%s | %w", "P90", err)
			}

		}
	}
	return nil
}

// NewCCSUsageTimeValue returns a CCSUsageTimeValue.
func NewCCSUsageTimeValue() *CCSUsageTimeValue {
	r := &CCSUsageTimeValue{}

	return r
}
