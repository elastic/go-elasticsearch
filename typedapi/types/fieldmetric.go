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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/metric"
)

// FieldMetric type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/rollup/_types/Metric.ts#L30-L35
type FieldMetric struct {
	// Field The field to collect metrics for. This must be a numeric of some kind.
	Field string `json:"field"`
	// Metrics An array of metrics to collect for the field. At least one metric must be
	// configured.
	Metrics []metric.Metric `json:"metrics"`
}

func (s *FieldMetric) UnmarshalJSON(data []byte) error {

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
			if err := dec.Decode(&s.Metrics); err != nil {
				return fmt.Errorf("%s | %w", "Metrics", err)
			}

		}
	}
	return nil
}

// NewFieldMetric returns a FieldMetric.
func NewFieldMetric() *FieldMetric {
	r := &FieldMetric{}

	return r
}
