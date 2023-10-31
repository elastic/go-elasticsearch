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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// GeoLineAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/metric.ts#L121-L146
type GeoLineAggregation struct {
	// IncludeSort When `true`, returns an additional array of the sort values in the feature
	// properties.
	IncludeSort *bool `json:"include_sort,omitempty"`
	// Point The name of the geo_point field.
	Point GeoLinePoint `json:"point"`
	// Size The maximum length of the line represented in the aggregation.
	// Valid sizes are between 1 and 10000.
	Size *int `json:"size,omitempty"`
	// Sort The name of the numeric field to use as the sort key for ordering the points.
	// When the `geo_line` aggregation is nested inside a `time_series` aggregation,
	// this field defaults to `@timestamp`, and any other value will result in
	// error.
	Sort GeoLineSort `json:"sort"`
	// SortOrder The order in which the line is sorted (ascending or descending).
	SortOrder *sortorder.SortOrder `json:"sort_order,omitempty"`
}

func (s *GeoLineAggregation) UnmarshalJSON(data []byte) error {

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

		case "include_sort":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IncludeSort = &value
			case bool:
				s.IncludeSort = &v
			}

		case "point":
			if err := dec.Decode(&s.Point); err != nil {
				return err
			}

		case "size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		case "sort":
			if err := dec.Decode(&s.Sort); err != nil {
				return err
			}

		case "sort_order":
			if err := dec.Decode(&s.SortOrder); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewGeoLineAggregation returns a GeoLineAggregation.
func NewGeoLineAggregation() *GeoLineAggregation {
	r := &GeoLineAggregation{}

	return r
}
