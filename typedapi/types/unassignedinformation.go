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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/unassignedinformationreason"
)

// UnassignedInformation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cluster/allocation_explain/types.ts#L117-L125
type UnassignedInformation struct {
	AllocationStatus         *string                                                 `json:"allocation_status,omitempty"`
	At                       DateTime                                                `json:"at"`
	Delayed                  *bool                                                   `json:"delayed,omitempty"`
	Details                  *string                                                 `json:"details,omitempty"`
	FailedAllocationAttempts *int                                                    `json:"failed_allocation_attempts,omitempty"`
	LastAllocationStatus     *string                                                 `json:"last_allocation_status,omitempty"`
	Reason                   unassignedinformationreason.UnassignedInformationReason `json:"reason"`
}

func (s *UnassignedInformation) UnmarshalJSON(data []byte) error {

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

		case "allocation_status":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AllocationStatus = &o

		case "at":
			if err := dec.Decode(&s.At); err != nil {
				return err
			}

		case "delayed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Delayed = &value
			case bool:
				s.Delayed = &v
			}

		case "details":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Details = &o

		case "failed_allocation_attempts":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FailedAllocationAttempts = &value
			case float64:
				f := int(v)
				s.FailedAllocationAttempts = &f
			}

		case "last_allocation_status":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LastAllocationStatus = &o

		case "reason":
			if err := dec.Decode(&s.Reason); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewUnassignedInformation returns a UnassignedInformation.
func NewUnassignedInformation() *UnassignedInformation {
	r := &UnassignedInformation{}

	return r
}
