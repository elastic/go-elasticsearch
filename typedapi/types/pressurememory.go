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
)

// PressureMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/nodes/_types/Stats.ts#L66-L80
type PressureMemory struct {
	All                                   ByteSize `json:"all,omitempty"`
	AllInBytes                            *int64   `json:"all_in_bytes,omitempty"`
	CombinedCoordinatingAndPrimary        ByteSize `json:"combined_coordinating_and_primary,omitempty"`
	CombinedCoordinatingAndPrimaryInBytes *int64   `json:"combined_coordinating_and_primary_in_bytes,omitempty"`
	Coordinating                          ByteSize `json:"coordinating,omitempty"`
	CoordinatingInBytes                   *int64   `json:"coordinating_in_bytes,omitempty"`
	CoordinatingRejections                *int64   `json:"coordinating_rejections,omitempty"`
	Primary                               ByteSize `json:"primary,omitempty"`
	PrimaryInBytes                        *int64   `json:"primary_in_bytes,omitempty"`
	PrimaryRejections                     *int64   `json:"primary_rejections,omitempty"`
	Replica                               ByteSize `json:"replica,omitempty"`
	ReplicaInBytes                        *int64   `json:"replica_in_bytes,omitempty"`
	ReplicaRejections                     *int64   `json:"replica_rejections,omitempty"`
}

func (s *PressureMemory) UnmarshalJSON(data []byte) error {

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

		case "all":
			if err := dec.Decode(&s.All); err != nil {
				return err
			}

		case "all_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.AllInBytes = &value
			case float64:
				f := int64(v)
				s.AllInBytes = &f
			}

		case "combined_coordinating_and_primary":
			if err := dec.Decode(&s.CombinedCoordinatingAndPrimary); err != nil {
				return err
			}

		case "combined_coordinating_and_primary_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CombinedCoordinatingAndPrimaryInBytes = &value
			case float64:
				f := int64(v)
				s.CombinedCoordinatingAndPrimaryInBytes = &f
			}

		case "coordinating":
			if err := dec.Decode(&s.Coordinating); err != nil {
				return err
			}

		case "coordinating_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CoordinatingInBytes = &value
			case float64:
				f := int64(v)
				s.CoordinatingInBytes = &f
			}

		case "coordinating_rejections":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CoordinatingRejections = &value
			case float64:
				f := int64(v)
				s.CoordinatingRejections = &f
			}

		case "primary":
			if err := dec.Decode(&s.Primary); err != nil {
				return err
			}

		case "primary_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PrimaryInBytes = &value
			case float64:
				f := int64(v)
				s.PrimaryInBytes = &f
			}

		case "primary_rejections":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PrimaryRejections = &value
			case float64:
				f := int64(v)
				s.PrimaryRejections = &f
			}

		case "replica":
			if err := dec.Decode(&s.Replica); err != nil {
				return err
			}

		case "replica_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ReplicaInBytes = &value
			case float64:
				f := int64(v)
				s.ReplicaInBytes = &f
			}

		case "replica_rejections":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ReplicaRejections = &value
			case float64:
				f := int64(v)
				s.ReplicaRejections = &f
			}

		}
	}
	return nil
}

// NewPressureMemory returns a PressureMemory.
func NewPressureMemory() *PressureMemory {
	r := &PressureMemory{}

	return r
}
