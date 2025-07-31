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

// JobUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/xpack/usage/types.ts#L374-L380
type JobUsage struct {
	Count     int              `json:"count"`
	CreatedBy map[string]int64 `json:"created_by"`
	Detectors JobStatistics    `json:"detectors"`
	Forecasts MlJobForecasts   `json:"forecasts"`
	ModelSize JobStatistics    `json:"model_size"`
}

func (s *JobUsage) UnmarshalJSON(data []byte) error {

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

		case "count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int(v)
				s.Count = f
			}

		case "created_by":
			if s.CreatedBy == nil {
				s.CreatedBy = make(map[string]int64, 0)
			}
			if err := dec.Decode(&s.CreatedBy); err != nil {
				return fmt.Errorf("%s | %w", "CreatedBy", err)
			}

		case "detectors":
			if err := dec.Decode(&s.Detectors); err != nil {
				return fmt.Errorf("%s | %w", "Detectors", err)
			}

		case "forecasts":
			if err := dec.Decode(&s.Forecasts); err != nil {
				return fmt.Errorf("%s | %w", "Forecasts", err)
			}

		case "model_size":
			if err := dec.Decode(&s.ModelSize); err != nil {
				return fmt.Errorf("%s | %w", "ModelSize", err)
			}

		}
	}
	return nil
}

// NewJobUsage returns a JobUsage.
func NewJobUsage() *JobUsage {
	r := &JobUsage{
		CreatedBy: make(map[string]int64),
	}

	return r
}
