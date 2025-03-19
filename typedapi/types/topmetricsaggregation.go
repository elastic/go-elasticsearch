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
// https://github.com/elastic/elasticsearch-specification/tree/3ea9ce260df22d3244bff5bace485dd97ff4046d

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// TopMetricsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3ea9ce260df22d3244bff5bace485dd97ff4046d/specification/_types/aggregations/metric.ts#L408-L425
type TopMetricsAggregation struct {
	// Field The field on which to run the aggregation.
	Field *string `json:"field,omitempty"`
	// Metrics The fields of the top document to return.
	Metrics []TopMetricsValue `json:"metrics,omitempty"`
	// Missing The value to apply to documents that do not have a value.
	// By default, documents without a value are ignored.
	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
	// Size The number of top documents from which to return metrics.
	Size *int `json:"size,omitempty"`
	// Sort The sort order of the documents.
	Sort []SortCombinations `json:"sort,omitempty"`
}

func (s *TopMetricsAggregation) UnmarshalJSON(data []byte) error {

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

		case "metrics":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewTopMetricsValue()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Metrics", err)
				}

				s.Metrics = append(s.Metrics, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Metrics); err != nil {
					return fmt.Errorf("%s | %w", "Metrics", err)
				}
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return fmt.Errorf("%s | %w", "Missing", err)
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return fmt.Errorf("%s | %w", "Script", err)
			}

		case "size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Size", err)
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		case "sort":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(SortCombinations)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Sort", err)
				}

				s.Sort = append(s.Sort, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Sort); err != nil {
					return fmt.Errorf("%s | %w", "Sort", err)
				}
			}

		}
	}
	return nil
}

// NewTopMetricsAggregation returns a TopMetricsAggregation.
func NewTopMetricsAggregation() *TopMetricsAggregation {
	r := &TopMetricsAggregation{}

	return r
}

// true

type TopMetricsAggregationVariant interface {
	TopMetricsAggregationCaster() *TopMetricsAggregation
}

func (s *TopMetricsAggregation) TopMetricsAggregationCaster() *TopMetricsAggregation {
	return s
}
