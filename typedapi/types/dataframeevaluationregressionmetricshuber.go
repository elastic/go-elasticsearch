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

// DataframeEvaluationRegressionMetricsHuber type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/DataframeEvaluation.ts#L117-L120
type DataframeEvaluationRegressionMetricsHuber struct {
	// Delta Approximates 1/2 (prediction - actual)2 for values much less than delta and
	// approximates a straight line with slope delta for values much larger than
	// delta. Defaults to 1. Delta needs to be greater than 0.
	Delta *Float64 `json:"delta,omitempty"`
}

func (s *DataframeEvaluationRegressionMetricsHuber) UnmarshalJSON(data []byte) error {

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

		case "delta":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Delta = &f
			case float64:
				f := Float64(v)
				s.Delta = &f
			}

		}
	}
	return nil
}

// NewDataframeEvaluationRegressionMetricsHuber returns a DataframeEvaluationRegressionMetricsHuber.
func NewDataframeEvaluationRegressionMetricsHuber() *DataframeEvaluationRegressionMetricsHuber {
	r := &DataframeEvaluationRegressionMetricsHuber{}

	return r
}
