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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// AutoscalingDecider type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/autoscaling/get_autoscaling_capacity/GetAutoscalingCapacityResponse.ts#L52-L56
type AutoscalingDecider struct {
	ReasonDetails    json.RawMessage     `json:"reason_details,omitempty"`
	ReasonSummary    *string             `json:"reason_summary,omitempty"`
	RequiredCapacity AutoscalingCapacity `json:"required_capacity"`
}

func (s *AutoscalingDecider) UnmarshalJSON(data []byte) error {

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

		case "reason_details":
			if err := dec.Decode(&s.ReasonDetails); err != nil {
				return fmt.Errorf("%s | %w", "ReasonDetails", err)
			}

		case "reason_summary":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ReasonSummary", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ReasonSummary = &o

		case "required_capacity":
			if err := dec.Decode(&s.RequiredCapacity); err != nil {
				return fmt.Errorf("%s | %w", "RequiredCapacity", err)
			}

		}
	}
	return nil
}

// NewAutoscalingDecider returns a AutoscalingDecider.
func NewAutoscalingDecider() *AutoscalingDecider {
	r := &AutoscalingDecider{}

	return r
}
