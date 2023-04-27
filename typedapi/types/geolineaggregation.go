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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// GeoLineAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/aggregations/metric.ts#L81-L87
type GeoLineAggregation struct {
	IncludeSort *bool                `json:"include_sort,omitempty"`
	Point       GeoLinePoint         `json:"point"`
	Size        *int                 `json:"size,omitempty"`
	Sort        GeoLineSort          `json:"sort"`
	SortOrder   *sortorder.SortOrder `json:"sort_order,omitempty"`
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
