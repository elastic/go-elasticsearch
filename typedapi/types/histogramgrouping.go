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

// HistogramGrouping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/rollup/_types/Groupings.ts#L84-L97
type HistogramGrouping struct {
	// Fields The set of fields that you wish to build histograms for.
	// All fields specified must be some kind of numeric.
	// Order does not matter.
	Fields []string `json:"fields"`
	// Interval The interval of histogram buckets to be generated when rolling up.
	// For example, a value of `5` creates buckets that are five units wide (`0-5`,
	// `5-10`, etc).
	// Note that only one interval can be specified in the histogram group, meaning
	// that all fields being grouped via the histogram must share the same interval.
	Interval int64 `json:"interval"`
}

func (s *HistogramGrouping) UnmarshalJSON(data []byte) error {

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

		case "fields":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Fields", err)
				}

				s.Fields = append(s.Fields, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Fields); err != nil {
					return fmt.Errorf("%s | %w", "Fields", err)
				}
			}

		case "interval":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "Interval", err)
				}
				s.Interval = value
			case float64:
				f := int64(v)
				s.Interval = f
			}

		}
	}
	return nil
}

// NewHistogramGrouping returns a HistogramGrouping.
func NewHistogramGrouping() *HistogramGrouping {
	r := &HistogramGrouping{}

	return r
}

type HistogramGroupingVariant interface {
	HistogramGroupingCaster() *HistogramGrouping
}

func (s *HistogramGrouping) HistogramGroupingCaster() *HistogramGrouping {
	return s
}
