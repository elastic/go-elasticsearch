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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// PercentileRanksAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/aggregations/metric.ts#L105-L110
type PercentileRanksAggregation struct {
	Field   *string    `json:"field,omitempty"`
	Format  *string    `json:"format,omitempty"`
	Hdr     *HdrMethod `json:"hdr,omitempty"`
	Keyed   *bool      `json:"keyed,omitempty"`
	Missing Missing    `json:"missing,omitempty"`
	Script  Script     `json:"script,omitempty"`
	Tdigest *TDigest   `json:"tdigest,omitempty"`
	Values  []Float64  `json:"values,omitempty"`
}

func (s *PercentileRanksAggregation) UnmarshalJSON(data []byte) error {

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
				return err
			}

		case "format":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Format = &o

		case "hdr":
			if err := dec.Decode(&s.Hdr); err != nil {
				return err
			}

		case "keyed":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Keyed = &value
			case bool:
				s.Keyed = &v
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return err
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		case "tdigest":
			if err := dec.Decode(&s.Tdigest); err != nil {
				return err
			}

		case "values":
			if err := dec.Decode(&s.Values); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewPercentileRanksAggregation returns a PercentileRanksAggregation.
func NewPercentileRanksAggregation() *PercentileRanksAggregation {
	r := &PercentileRanksAggregation{}

	return r
}
