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

// OsqSegmentAutoCalibrationParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/indices/segments/types.ts#L113-L122
type OsqSegmentAutoCalibrationParameters struct {
	// Bits The number of bits used to quantize the vectors.
	Bits int `json:"bits"`
	// Oversample The oversampling factor applied when rescoring.
	Oversample float32 `json:"oversample"`
	// Precondition If `true`, the vectors are preconditioned before quantization.
	Precondition bool `json:"precondition"`
	// QueryBits The number of bits used to quantize the query vectors.
	QueryBits int `json:"query_bits"`
}

func (s *OsqSegmentAutoCalibrationParameters) UnmarshalJSON(data []byte) error {

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

		case "bits":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Bits", err)
				}
				s.Bits = value
			case float64:
				f := int(v)
				s.Bits = f
			}

		case "oversample":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Oversample", err)
				}
				f := float32(value)
				s.Oversample = f
			case float64:
				f := float32(v)
				s.Oversample = f
			}

		case "precondition":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Precondition", err)
				}
				s.Precondition = value
			case bool:
				s.Precondition = v
			}

		case "query_bits":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "QueryBits", err)
				}
				s.QueryBits = value
			case float64:
				f := int(v)
				s.QueryBits = f
			}

		}
	}
	return nil
}

// NewOsqSegmentAutoCalibrationParameters returns a OsqSegmentAutoCalibrationParameters.
func NewOsqSegmentAutoCalibrationParameters() *OsqSegmentAutoCalibrationParameters {
	r := &OsqSegmentAutoCalibrationParameters{}

	return r
}
