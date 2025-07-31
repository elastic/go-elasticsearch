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

// DenseVectorOffHeapStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cluster/stats/types.ts#L166-L178
type DenseVectorOffHeapStats struct {
	Fielddata         map[string]map[string]int64 `json:"fielddata,omitempty"`
	TotalSize         ByteSize                    `json:"total_size,omitempty"`
	TotalSizeBytes    int64                       `json:"total_size_bytes"`
	TotalVebSize      ByteSize                    `json:"total_veb_size,omitempty"`
	TotalVebSizeBytes int64                       `json:"total_veb_size_bytes"`
	TotalVecSize      ByteSize                    `json:"total_vec_size,omitempty"`
	TotalVecSizeBytes int64                       `json:"total_vec_size_bytes"`
	TotalVeqSize      ByteSize                    `json:"total_veq_size,omitempty"`
	TotalVeqSizeBytes int64                       `json:"total_veq_size_bytes"`
	TotalVexSize      ByteSize                    `json:"total_vex_size,omitempty"`
	TotalVexSizeBytes int64                       `json:"total_vex_size_bytes"`
}

func (s *DenseVectorOffHeapStats) UnmarshalJSON(data []byte) error {

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

		case "fielddata":
			if s.Fielddata == nil {
				s.Fielddata = make(map[string]map[string]int64, 0)
			}
			if err := dec.Decode(&s.Fielddata); err != nil {
				return fmt.Errorf("%s | %w", "Fielddata", err)
			}

		case "total_size":
			if err := dec.Decode(&s.TotalSize); err != nil {
				return fmt.Errorf("%s | %w", "TotalSize", err)
			}

		case "total_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalSizeBytes", err)
				}
				s.TotalSizeBytes = value
			case float64:
				f := int64(v)
				s.TotalSizeBytes = f
			}

		case "total_veb_size":
			if err := dec.Decode(&s.TotalVebSize); err != nil {
				return fmt.Errorf("%s | %w", "TotalVebSize", err)
			}

		case "total_veb_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalVebSizeBytes", err)
				}
				s.TotalVebSizeBytes = value
			case float64:
				f := int64(v)
				s.TotalVebSizeBytes = f
			}

		case "total_vec_size":
			if err := dec.Decode(&s.TotalVecSize); err != nil {
				return fmt.Errorf("%s | %w", "TotalVecSize", err)
			}

		case "total_vec_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalVecSizeBytes", err)
				}
				s.TotalVecSizeBytes = value
			case float64:
				f := int64(v)
				s.TotalVecSizeBytes = f
			}

		case "total_veq_size":
			if err := dec.Decode(&s.TotalVeqSize); err != nil {
				return fmt.Errorf("%s | %w", "TotalVeqSize", err)
			}

		case "total_veq_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalVeqSizeBytes", err)
				}
				s.TotalVeqSizeBytes = value
			case float64:
				f := int64(v)
				s.TotalVeqSizeBytes = f
			}

		case "total_vex_size":
			if err := dec.Decode(&s.TotalVexSize); err != nil {
				return fmt.Errorf("%s | %w", "TotalVexSize", err)
			}

		case "total_vex_size_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalVexSizeBytes", err)
				}
				s.TotalVexSizeBytes = value
			case float64:
				f := int64(v)
				s.TotalVexSizeBytes = f
			}

		}
	}
	return nil
}

// NewDenseVectorOffHeapStats returns a DenseVectorOffHeapStats.
func NewDenseVectorOffHeapStats() *DenseVectorOffHeapStats {
	r := &DenseVectorOffHeapStats{
		Fielddata: make(map[string]map[string]int64),
	}

	return r
}
