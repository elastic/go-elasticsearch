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

// MatrixStatsFields type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/Aggregate.ts#L877-L886
type MatrixStatsFields struct {
	Correlation map[string]Float64 `json:"correlation"`
	Count       int64              `json:"count"`
	Covariance  map[string]Float64 `json:"covariance"`
	Kurtosis    Float64            `json:"kurtosis"`
	Mean        Float64            `json:"mean"`
	Name        string             `json:"name"`
	Skewness    Float64            `json:"skewness"`
	Variance    Float64            `json:"variance"`
}

func (s *MatrixStatsFields) UnmarshalJSON(data []byte) error {

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

		case "correlation":
			if s.Correlation == nil {
				s.Correlation = make(map[string]Float64, 0)
			}
			if err := dec.Decode(&s.Correlation); err != nil {
				return fmt.Errorf("%s | %w", "Correlation", err)
			}

		case "count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Count", err)
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "covariance":
			if s.Covariance == nil {
				s.Covariance = make(map[string]Float64, 0)
			}
			if err := dec.Decode(&s.Covariance); err != nil {
				return fmt.Errorf("%s | %w", "Covariance", err)
			}

		case "kurtosis":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Kurtosis", err)
				}
				f := Float64(value)
				s.Kurtosis = f
			case float64:
				f := Float64(v)
				s.Kurtosis = f
			}

		case "mean":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Mean", err)
				}
				f := Float64(value)
				s.Mean = f
			case float64:
				f := Float64(v)
				s.Mean = f
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "skewness":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Skewness", err)
				}
				f := Float64(value)
				s.Skewness = f
			case float64:
				f := Float64(v)
				s.Skewness = f
			}

		case "variance":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Variance", err)
				}
				f := Float64(value)
				s.Variance = f
			case float64:
				f := Float64(v)
				s.Variance = f
			}

		}
	}
	return nil
}

// NewMatrixStatsFields returns a MatrixStatsFields.
func NewMatrixStatsFields() *MatrixStatsFields {
	r := &MatrixStatsFields{
		Correlation: make(map[string]Float64),
		Covariance:  make(map[string]Float64),
	}

	return r
}
