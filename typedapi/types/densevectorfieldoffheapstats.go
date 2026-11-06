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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// The off-heap stats of a `dense_vector` field.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/cluster/stats/types.ts#L216-L231
type DenseVectorFieldOffHeapStats struct {
	// AutoCalibration The auto-calibration information of the field, aggregated by distinct set of
	// parameters. Only returned when `include_auto_calibration` is `true`.
	AutoCalibration              []DenseVectorAutoCalibration `json:"auto_calibration,omitempty"`
	DenseVectorFieldOffHeapStats map[string]int64             `json:"-"`
}

func (s *DenseVectorFieldOffHeapStats) UnmarshalJSON(data []byte) error {

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

		case "auto_calibration":

			buf := []json.RawMessage{}
			dec.Decode(&buf)
			for _, rawMsg := range buf {

				source := bytes.NewReader(rawMsg)
				localDec := json.NewDecoder(source)
				kind := make(map[string]string, 0)
				localDec.Decode(&kind)
				source.Seek(0, io.SeekStart)

				switch kind["type"] {

				case "osq":
					o := NewOsqDenseVectorAutoCalibration()
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("%s | %w", "osq", err)
					}
					s.AutoCalibration = append(s.AutoCalibration, *o)
				default:
					o := new(any)
					if err := localDec.Decode(&o); err != nil {
						return fmt.Errorf("AutoCalibration | %w", err)
					}
					s.AutoCalibration = append(s.AutoCalibration, *o)
				}
			}

		default:

			if key, ok := t.(string); ok {
				if s.DenseVectorFieldOffHeapStats == nil {
					s.DenseVectorFieldOffHeapStats = make(map[string]int64, 0)
				}
				raw := new(int64)
				if err := dec.Decode(&raw); err != nil {
					return fmt.Errorf("%s | %w", "DenseVectorFieldOffHeapStats", err)
				}
				if raw != nil {
					s.DenseVectorFieldOffHeapStats[key] = *raw
				}
			}

		}
	}
	return nil
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s DenseVectorFieldOffHeapStats) MarshalJSON() ([]byte, error) {
	type opt DenseVectorFieldOffHeapStats
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]json.RawMessage, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.DenseVectorFieldOffHeapStats {
		marshaled, err := json.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal additional property %q: %w", key, err)
		}
		tmp[fmt.Sprintf("%s", key)] = marshaled
	}
	delete(tmp, "DenseVectorFieldOffHeapStats")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewDenseVectorFieldOffHeapStats returns a DenseVectorFieldOffHeapStats.
func NewDenseVectorFieldOffHeapStats() *DenseVectorFieldOffHeapStats {
	r := &DenseVectorFieldOffHeapStats{
		DenseVectorFieldOffHeapStats: make(map[string]int64),
	}

	return r
}
