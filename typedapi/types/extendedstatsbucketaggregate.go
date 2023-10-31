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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ExtendedStatsBucketAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/Aggregate.ts#L298-L299
type ExtendedStatsBucketAggregate struct {
	Avg                        Float64                          `json:"avg,omitempty"`
	AvgAsString                *string                          `json:"avg_as_string,omitempty"`
	Count                      int64                            `json:"count"`
	Max                        Float64                          `json:"max,omitempty"`
	MaxAsString                *string                          `json:"max_as_string,omitempty"`
	Meta                       Metadata                         `json:"meta,omitempty"`
	Min                        Float64                          `json:"min,omitempty"`
	MinAsString                *string                          `json:"min_as_string,omitempty"`
	StdDeviation               Float64                          `json:"std_deviation,omitempty"`
	StdDeviationAsString       *string                          `json:"std_deviation_as_string,omitempty"`
	StdDeviationBounds         *StandardDeviationBounds         `json:"std_deviation_bounds,omitempty"`
	StdDeviationBoundsAsString *StandardDeviationBoundsAsString `json:"std_deviation_bounds_as_string,omitempty"`
	StdDeviationPopulation     Float64                          `json:"std_deviation_population,omitempty"`
	StdDeviationSampling       Float64                          `json:"std_deviation_sampling,omitempty"`
	Sum                        Float64                          `json:"sum"`
	SumAsString                *string                          `json:"sum_as_string,omitempty"`
	SumOfSquares               Float64                          `json:"sum_of_squares,omitempty"`
	SumOfSquaresAsString       *string                          `json:"sum_of_squares_as_string,omitempty"`
	Variance                   Float64                          `json:"variance,omitempty"`
	VarianceAsString           *string                          `json:"variance_as_string,omitempty"`
	VariancePopulation         Float64                          `json:"variance_population,omitempty"`
	VariancePopulationAsString *string                          `json:"variance_population_as_string,omitempty"`
	VarianceSampling           Float64                          `json:"variance_sampling,omitempty"`
	VarianceSamplingAsString   *string                          `json:"variance_sampling_as_string,omitempty"`
}

func (s *ExtendedStatsBucketAggregate) UnmarshalJSON(data []byte) error {

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

		case "avg":
			if err := dec.Decode(&s.Avg); err != nil {
				return err
			}

		case "avg_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AvgAsString = &o

		case "count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Count = value
			case float64:
				f := int64(v)
				s.Count = f
			}

		case "max":
			if err := dec.Decode(&s.Max); err != nil {
				return err
			}

		case "max_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxAsString = &o

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "min":
			if err := dec.Decode(&s.Min); err != nil {
				return err
			}

		case "min_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MinAsString = &o

		case "std_deviation":
			if err := dec.Decode(&s.StdDeviation); err != nil {
				return err
			}

		case "std_deviation_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StdDeviationAsString = &o

		case "std_deviation_bounds":
			if err := dec.Decode(&s.StdDeviationBounds); err != nil {
				return err
			}

		case "std_deviation_bounds_as_string":
			if err := dec.Decode(&s.StdDeviationBoundsAsString); err != nil {
				return err
			}

		case "std_deviation_population":
			if err := dec.Decode(&s.StdDeviationPopulation); err != nil {
				return err
			}

		case "std_deviation_sampling":
			if err := dec.Decode(&s.StdDeviationSampling); err != nil {
				return err
			}

		case "sum":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Sum = f
			case float64:
				f := Float64(v)
				s.Sum = f
			}

		case "sum_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SumAsString = &o

		case "sum_of_squares":
			if err := dec.Decode(&s.SumOfSquares); err != nil {
				return err
			}

		case "sum_of_squares_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SumOfSquaresAsString = &o

		case "variance":
			if err := dec.Decode(&s.Variance); err != nil {
				return err
			}

		case "variance_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.VarianceAsString = &o

		case "variance_population":
			if err := dec.Decode(&s.VariancePopulation); err != nil {
				return err
			}

		case "variance_population_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.VariancePopulationAsString = &o

		case "variance_sampling":
			if err := dec.Decode(&s.VarianceSampling); err != nil {
				return err
			}

		case "variance_sampling_as_string":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.VarianceSamplingAsString = &o

		}
	}
	return nil
}

// NewExtendedStatsBucketAggregate returns a ExtendedStatsBucketAggregate.
func NewExtendedStatsBucketAggregate() *ExtendedStatsBucketAggregate {
	r := &ExtendedStatsBucketAggregate{}

	return r
}
