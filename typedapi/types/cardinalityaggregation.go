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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/cardinalityexecutionmode"
)

// CardinalityAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/aggregations/metric.ts#L62-L66
type CardinalityAggregation struct {
	ExecutionHint      *cardinalityexecutionmode.CardinalityExecutionMode `json:"execution_hint,omitempty"`
	Field              *string                                            `json:"field,omitempty"`
	Missing            Missing                                            `json:"missing,omitempty"`
	PrecisionThreshold *int                                               `json:"precision_threshold,omitempty"`
	Rehash             *bool                                              `json:"rehash,omitempty"`
	Script             Script                                             `json:"script,omitempty"`
}

func (s *CardinalityAggregation) UnmarshalJSON(data []byte) error {

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

		case "execution_hint":
			if err := dec.Decode(&s.ExecutionHint); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return err
			}

		case "precision_threshold":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PrecisionThreshold = &value
			case float64:
				f := int(v)
				s.PrecisionThreshold = &f
			}

		case "rehash":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Rehash = &value
			case bool:
				s.Rehash = &v
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewCardinalityAggregation returns a CardinalityAggregation.
func NewCardinalityAggregation() *CardinalityAggregation {
	r := &CardinalityAggregation{}

	return r
}
