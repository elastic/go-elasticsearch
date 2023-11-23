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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ShardsAvailabilityIndicatorDetails type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/health_report/types.ts#L109-L119
type ShardsAvailabilityIndicatorDetails struct {
	CreatingPrimaries     int64 `json:"creating_primaries"`
	InitializingPrimaries int64 `json:"initializing_primaries"`
	InitializingReplicas  int64 `json:"initializing_replicas"`
	RestartingPrimaries   int64 `json:"restarting_primaries"`
	RestartingReplicas    int64 `json:"restarting_replicas"`
	StartedPrimaries      int64 `json:"started_primaries"`
	StartedReplicas       int64 `json:"started_replicas"`
	UnassignedPrimaries   int64 `json:"unassigned_primaries"`
	UnassignedReplicas    int64 `json:"unassigned_replicas"`
}

func (s *ShardsAvailabilityIndicatorDetails) UnmarshalJSON(data []byte) error {

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

		case "creating_primaries":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CreatingPrimaries = value
			case float64:
				f := int64(v)
				s.CreatingPrimaries = f
			}

		case "initializing_primaries":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.InitializingPrimaries = value
			case float64:
				f := int64(v)
				s.InitializingPrimaries = f
			}

		case "initializing_replicas":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.InitializingReplicas = value
			case float64:
				f := int64(v)
				s.InitializingReplicas = f
			}

		case "restarting_primaries":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RestartingPrimaries = value
			case float64:
				f := int64(v)
				s.RestartingPrimaries = f
			}

		case "restarting_replicas":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RestartingReplicas = value
			case float64:
				f := int64(v)
				s.RestartingReplicas = f
			}

		case "started_primaries":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.StartedPrimaries = value
			case float64:
				f := int64(v)
				s.StartedPrimaries = f
			}

		case "started_replicas":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.StartedReplicas = value
			case float64:
				f := int64(v)
				s.StartedReplicas = f
			}

		case "unassigned_primaries":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.UnassignedPrimaries = value
			case float64:
				f := int64(v)
				s.UnassignedPrimaries = f
			}

		case "unassigned_replicas":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.UnassignedReplicas = value
			case float64:
				f := int64(v)
				s.UnassignedReplicas = f
			}

		}
	}
	return nil
}

// NewShardsAvailabilityIndicatorDetails returns a ShardsAvailabilityIndicatorDetails.
func NewShardsAvailabilityIndicatorDetails() *ShardsAvailabilityIndicatorDetails {
	r := &ShardsAvailabilityIndicatorDetails{}

	return r
}
