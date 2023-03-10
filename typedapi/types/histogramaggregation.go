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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// HistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/aggregations/bucket.ts#L235-L247
type HistogramAggregation struct {
	ExtendedBounds *ExtendedBoundsdouble      `json:"extended_bounds,omitempty"`
	Field          *string                    `json:"field,omitempty"`
	Format         *string                    `json:"format,omitempty"`
	HardBounds     *ExtendedBoundsdouble      `json:"hard_bounds,omitempty"`
	Interval       *Float64                   `json:"interval,omitempty"`
	Keyed          *bool                      `json:"keyed,omitempty"`
	Meta           map[string]json.RawMessage `json:"meta,omitempty"`
	MinDocCount    *int                       `json:"min_doc_count,omitempty"`
	Missing        *Float64                   `json:"missing,omitempty"`
	Name           *string                    `json:"name,omitempty"`
	Offset         *Float64                   `json:"offset,omitempty"`
	Order          AggregateOrder             `json:"order,omitempty"`
	Script         Script                     `json:"script,omitempty"`
}

func (s *HistogramAggregation) UnmarshalJSON(data []byte) error {
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

		case "extended_bounds":
			if err := dec.Decode(&s.ExtendedBounds); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "format":
			if err := dec.Decode(&s.Format); err != nil {
				return err
			}

		case "hard_bounds":
			if err := dec.Decode(&s.HardBounds); err != nil {
				return err
			}

		case "interval":
			if err := dec.Decode(&s.Interval); err != nil {
				return err
			}

		case "keyed":
			if err := dec.Decode(&s.Keyed); err != nil {
				return err
			}

		case "meta":
			if err := dec.Decode(&s.Meta); err != nil {
				return err
			}

		case "min_doc_count":
			if err := dec.Decode(&s.MinDocCount); err != nil {
				return err
			}

		case "missing":
			if err := dec.Decode(&s.Missing); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "offset":
			if err := dec.Decode(&s.Offset); err != nil {
				return err
			}

		case "order":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {

			case '{':
				o := make(map[string]sortorder.SortOrder, 0)
				localDec.Decode(&o)
				s.Order = o

			case '[':
				o := make([]map[string]sortorder.SortOrder, 0)
				localDec.Decode(&o)
				s.Order = o
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewHistogramAggregation returns a HistogramAggregation.
func NewHistogramAggregation() *HistogramAggregation {
	r := &HistogramAggregation{}

	return r
}
