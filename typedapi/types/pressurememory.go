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

// PressureMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/nodes/_types/Stats.ts#L144-L201
type PressureMemory struct {
	// All Memory consumed by indexing requests in the coordinating, primary, or replica
	// stage.
	All ByteSize `json:"all,omitempty"`
	// AllInBytes Memory consumed, in bytes, by indexing requests in the coordinating, primary,
	// or replica stage.
	AllInBytes *int64 `json:"all_in_bytes,omitempty"`
	// CombinedCoordinatingAndPrimary Memory consumed by indexing requests in the coordinating or primary stage.
	// This value is not the sum of coordinating and primary as a node can reuse the
	// coordinating memory if the primary stage is executed locally.
	CombinedCoordinatingAndPrimary ByteSize `json:"combined_coordinating_and_primary,omitempty"`
	// CombinedCoordinatingAndPrimaryInBytes Memory consumed, in bytes, by indexing requests in the coordinating or
	// primary stage.
	// This value is not the sum of coordinating and primary as a node can reuse the
	// coordinating memory if the primary stage is executed locally.
	CombinedCoordinatingAndPrimaryInBytes *int64 `json:"combined_coordinating_and_primary_in_bytes,omitempty"`
	// Coordinating Memory consumed by indexing requests in the coordinating stage.
	Coordinating ByteSize `json:"coordinating,omitempty"`
	// CoordinatingInBytes Memory consumed, in bytes, by indexing requests in the coordinating stage.
	CoordinatingInBytes *int64 `json:"coordinating_in_bytes,omitempty"`
	// CoordinatingRejections Number of indexing requests rejected in the coordinating stage.
	CoordinatingRejections   *int64 `json:"coordinating_rejections,omitempty"`
	LargeOperationRejections *int64 `json:"large_operation_rejections,omitempty"`
	// Primary Memory consumed by indexing requests in the primary stage.
	Primary                   ByteSize `json:"primary,omitempty"`
	PrimaryDocumentRejections *int64   `json:"primary_document_rejections,omitempty"`
	// PrimaryInBytes Memory consumed, in bytes, by indexing requests in the primary stage.
	PrimaryInBytes *int64 `json:"primary_in_bytes,omitempty"`
	// PrimaryRejections Number of indexing requests rejected in the primary stage.
	PrimaryRejections *int64 `json:"primary_rejections,omitempty"`
	// Replica Memory consumed by indexing requests in the replica stage.
	Replica ByteSize `json:"replica,omitempty"`
	// ReplicaInBytes Memory consumed, in bytes, by indexing requests in the replica stage.
	ReplicaInBytes *int64 `json:"replica_in_bytes,omitempty"`
	// ReplicaRejections Number of indexing requests rejected in the replica stage.
	ReplicaRejections *int64 `json:"replica_rejections,omitempty"`
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
				return fmt.Errorf("%s | %w", "All", err)
			}

		case "all_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "AllInBytes", err)
				}
				s.AllInBytes = &value
			case float64:
				f := int64(v)
				s.AllInBytes = &f
			}

		case "combined_coordinating_and_primary":
			if err := dec.Decode(&s.CombinedCoordinatingAndPrimary); err != nil {
				return fmt.Errorf("%s | %w", "CombinedCoordinatingAndPrimary", err)
			}

		case "combined_coordinating_and_primary_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CombinedCoordinatingAndPrimaryInBytes", err)
				}
				s.CombinedCoordinatingAndPrimaryInBytes = &value
			case float64:
				f := int64(v)
				s.CombinedCoordinatingAndPrimaryInBytes = &f
			}

		case "coordinating":
			if err := dec.Decode(&s.Coordinating); err != nil {
				return fmt.Errorf("%s | %w", "Coordinating", err)
			}

		case "coordinating_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CoordinatingInBytes", err)
				}
				s.CoordinatingInBytes = &value
			case float64:
				f := int64(v)
				s.CoordinatingInBytes = &f
			}

		case "coordinating_rejections":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "CoordinatingRejections", err)
				}
				s.CoordinatingRejections = &value
			case float64:
				f := int64(v)
				s.CoordinatingRejections = &f
			}

		case "large_operation_rejections":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "LargeOperationRejections", err)
				}
				s.LargeOperationRejections = &value
			case float64:
				f := int64(v)
				s.LargeOperationRejections = &f
			}

		case "primary":
			if err := dec.Decode(&s.Primary); err != nil {
				return fmt.Errorf("%s | %w", "Primary", err)
			}

		case "primary_document_rejections":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryDocumentRejections", err)
				}
				s.PrimaryDocumentRejections = &value
			case float64:
				f := int64(v)
				s.PrimaryDocumentRejections = &f
			}

		case "primary_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryInBytes", err)
				}
				s.PrimaryInBytes = &value
			case float64:
				f := int64(v)
				s.PrimaryInBytes = &f
			}

		case "primary_rejections":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "PrimaryRejections", err)
				}
				s.PrimaryRejections = &value
			case float64:
				f := int64(v)
				s.PrimaryRejections = &f
			}

		case "replica":
			if err := dec.Decode(&s.Replica); err != nil {
				return fmt.Errorf("%s | %w", "Replica", err)
			}

		case "replica_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReplicaInBytes", err)
				}
				s.ReplicaInBytes = &value
			case float64:
				f := int64(v)
				s.ReplicaInBytes = &f
			}

		case "replica_rejections":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReplicaRejections", err)
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
