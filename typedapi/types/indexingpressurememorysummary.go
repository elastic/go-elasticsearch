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

// IndexingPressureMemorySummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/cluster/stats/types.ts#L309-L318
type IndexingPressureMemorySummary struct {
	AllInBytes                            int64  `json:"all_in_bytes"`
	CombinedCoordinatingAndPrimaryInBytes int64  `json:"combined_coordinating_and_primary_in_bytes"`
	CoordinatingInBytes                   int64  `json:"coordinating_in_bytes"`
	CoordinatingRejections                *int64 `json:"coordinating_rejections,omitempty"`
	PrimaryInBytes                        int64  `json:"primary_in_bytes"`
	PrimaryRejections                     *int64 `json:"primary_rejections,omitempty"`
	ReplicaInBytes                        int64  `json:"replica_in_bytes"`
	ReplicaRejections                     *int64 `json:"replica_rejections,omitempty"`
}

func (s *IndexingPressureMemorySummary) UnmarshalJSON(data []byte) error {

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

		case "all_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.AllInBytes = value
			case float64:
				f := int64(v)
				s.AllInBytes = f
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
				s.CombinedCoordinatingAndPrimaryInBytes = value
			case float64:
				f := int64(v)
				s.CombinedCoordinatingAndPrimaryInBytes = f
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
				s.CoordinatingInBytes = value
			case float64:
				f := int64(v)
				s.CoordinatingInBytes = f
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

		case "primary_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.PrimaryInBytes = value
			case float64:
				f := int64(v)
				s.PrimaryInBytes = f
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

		case "replica_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ReplicaInBytes = value
			case float64:
				f := int64(v)
				s.ReplicaInBytes = f
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

// NewIndexingPressureMemorySummary returns a IndexingPressureMemorySummary.
func NewIndexingPressureMemorySummary() *IndexingPressureMemorySummary {
	r := &IndexingPressureMemorySummary{}

	return r
}
