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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Weights type.
//
// https://github.com/elastic/elasticsearch-specification/blob/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6/specification/ml/put_trained_model/types.ts#L109-L111
type Weights struct {
	Weights Float64 `json:"weights"`
}

func (s *Weights) UnmarshalJSON(data []byte) error {

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

		case "weights":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Weights", err)
				}
				f := Float64(value)
				s.Weights = f
			case float64:
				f := Float64(v)
				s.Weights = f
			}

		}
	}
	return nil
}

// NewWeights returns a Weights.
func NewWeights() *Weights {
	r := &Weights{}

	return r
}

type WeightsVariant interface {
	WeightsCaster() *Weights
}

func (s *Weights) WeightsCaster() *Weights {
	return s
}
