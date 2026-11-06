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
	"strconv"
)

// OsqSegmentAutoCalibration type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/indices/segments/types.ts#L107-L111
type OsqSegmentAutoCalibration struct {
	// Calibrated If `true`, the segment has been calibrated, and `parameters` is present.
	Calibrated bool `json:"calibrated"`
	// NumberOfVectors The number of vectors of the field in the segment.
	NumberOfVectors int64 `json:"number_of_vectors"`
	// Parameters The calibration parameters.
	Parameters *OsqSegmentAutoCalibrationParameters `json:"parameters,omitempty"`
	Size       ByteSize                             `json:"size,omitempty"`
	// SizeInBytes The size of the vectors of the field in the segment, in bytes.
	SizeInBytes int64  `json:"size_in_bytes"`
	Type        string `json:"type,omitempty"`
}

func (s *OsqSegmentAutoCalibration) UnmarshalJSON(data []byte) error {

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

		case "calibrated":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Calibrated", err)
				}
				s.Calibrated = value
			case bool:
				s.Calibrated = v
			}

		case "number_of_vectors":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "NumberOfVectors", err)
				}
				s.NumberOfVectors = value
			case float64:
				f := int64(v)
				s.NumberOfVectors = f
			}

		case "parameters":
			if err := dec.Decode(&s.Parameters); err != nil {
				return fmt.Errorf("%s | %w", "Parameters", err)
			}

		case "size":
			if err := dec.Decode(&s.Size); err != nil {
				return fmt.Errorf("%s | %w", "Size", err)
			}

		case "size_in_bytes":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SizeInBytes", err)
				}
				s.SizeInBytes = value
			case float64:
				f := int64(v)
				s.SizeInBytes = f
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s OsqSegmentAutoCalibration) MarshalJSON() ([]byte, error) {
	type innerOsqSegmentAutoCalibration OsqSegmentAutoCalibration
	tmp := innerOsqSegmentAutoCalibration{
		Calibrated:      s.Calibrated,
		NumberOfVectors: s.NumberOfVectors,
		Parameters:      s.Parameters,
		Size:            s.Size,
		SizeInBytes:     s.SizeInBytes,
		Type:            s.Type,
	}

	tmp.Type = "osq"

	return json.Marshal(tmp)
}

// NewOsqSegmentAutoCalibration returns a OsqSegmentAutoCalibration.
func NewOsqSegmentAutoCalibration() *OsqSegmentAutoCalibration {
	r := &OsqSegmentAutoCalibration{}

	return r
}
