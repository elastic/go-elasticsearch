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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// CartesianCentroidAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/aa1459fbdcaf57c653729142b3b6e9982373bb1c/specification/_types/aggregations/metric.ts#L116-L116
type CartesianCentroidAggregation struct {
	// Field The field on which to run the aggregation.
	Field *string `json:"field,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *CartesianCentroidAggregation) UnmarshalJSON(data []byte) error {

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

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return fmt.Errorf("%s | %w", "Field", err)
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		}
	}
	return nil
}

// NewCartesianCentroidAggregation returns a CartesianCentroidAggregation.
func NewCartesianCentroidAggregation() *CartesianCentroidAggregation {
	r := &CartesianCentroidAggregation{}

	return r
}

type CartesianCentroidAggregationVariant interface {
	CartesianCentroidAggregationCaster() *CartesianCentroidAggregation
}

func (s *CartesianCentroidAggregation) CartesianCentroidAggregationCaster() *CartesianCentroidAggregation {
	return s
}
