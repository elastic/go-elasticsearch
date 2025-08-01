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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// DataframeEvaluationRegressionMetricsMsle type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/DataframeEvaluation.ts#L112-L115
type DataframeEvaluationRegressionMetricsMsle struct {
	// Offset Defines the transition point at which you switch from minimizing quadratic
	// error to minimizing quadratic log error. Defaults to 1.
	Offset *Float64 `json:"offset,omitempty"`
}

func (s *DataframeEvaluationRegressionMetricsMsle) UnmarshalJSON(data []byte) error {

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

		case "offset":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Offset", err)
				}
				f := Float64(value)
				s.Offset = &f
			case float64:
				f := Float64(v)
				s.Offset = &f
			}

		}
	}
	return nil
}

// NewDataframeEvaluationRegressionMetricsMsle returns a DataframeEvaluationRegressionMetricsMsle.
func NewDataframeEvaluationRegressionMetricsMsle() *DataframeEvaluationRegressionMetricsMsle {
	r := &DataframeEvaluationRegressionMetricsMsle{}

	return r
}

type DataframeEvaluationRegressionMetricsMsleVariant interface {
	DataframeEvaluationRegressionMetricsMsleCaster() *DataframeEvaluationRegressionMetricsMsle
}

func (s *DataframeEvaluationRegressionMetricsMsle) DataframeEvaluationRegressionMetricsMsleCaster() *DataframeEvaluationRegressionMetricsMsle {
	return s
}
