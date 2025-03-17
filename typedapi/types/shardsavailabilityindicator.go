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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indicatorhealthstatus"
)

// ShardsAvailabilityIndicator type.
//
// https://github.com/elastic/elasticsearch-specification/blob/c75a0abec670d027d13eb8d6f23374f86621c76b/specification/_global/health_report/types.ts#L106-L110
type ShardsAvailabilityIndicator struct {
	Details   *ShardsAvailabilityIndicatorDetails         `json:"details,omitempty"`
	Diagnosis []Diagnosis                                 `json:"diagnosis,omitempty"`
	Impacts   []Impact                                    `json:"impacts,omitempty"`
	Status    indicatorhealthstatus.IndicatorHealthStatus `json:"status"`
	Symptom   string                                      `json:"symptom"`
}

func (s *ShardsAvailabilityIndicator) UnmarshalJSON(data []byte) error {

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

		case "details":
			if err := dec.Decode(&s.Details); err != nil {
				return fmt.Errorf("%s | %w", "Details", err)
			}

		case "diagnosis":
			if err := dec.Decode(&s.Diagnosis); err != nil {
				return fmt.Errorf("%s | %w", "Diagnosis", err)
			}

		case "impacts":
			if err := dec.Decode(&s.Impacts); err != nil {
				return fmt.Errorf("%s | %w", "Impacts", err)
			}

		case "status":
			if err := dec.Decode(&s.Status); err != nil {
				return fmt.Errorf("%s | %w", "Status", err)
			}

		case "symptom":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Symptom", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Symptom = o

		}
	}
	return nil
}

// NewShardsAvailabilityIndicator returns a ShardsAvailabilityIndicator.
func NewShardsAvailabilityIndicator() *ShardsAvailabilityIndicator {
	r := &ShardsAvailabilityIndicator{}

	return r
}

// false
