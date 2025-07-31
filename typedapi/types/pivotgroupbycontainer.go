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
	"encoding/json"
	"fmt"
)

// PivotGroupByContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/transform/_types/Transform.ts#L70-L78
type PivotGroupByContainer struct {
	AdditionalPivotGroupByContainerProperty map[string]json.RawMessage `json:"-"`
	DateHistogram                           *DateHistogramAggregation  `json:"date_histogram,omitempty"`
	GeotileGrid                             *GeoTileGridAggregation    `json:"geotile_grid,omitempty"`
	Histogram                               *HistogramAggregation      `json:"histogram,omitempty"`
	Terms                                   *TermsAggregation          `json:"terms,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s PivotGroupByContainer) MarshalJSON() ([]byte, error) {
	type opt PivotGroupByContainer
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]any, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.AdditionalPivotGroupByContainerProperty {
		tmp[fmt.Sprintf("%s", key)] = value
	}
	delete(tmp, "AdditionalPivotGroupByContainerProperty")

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NewPivotGroupByContainer returns a PivotGroupByContainer.
func NewPivotGroupByContainer() *PivotGroupByContainer {
	r := &PivotGroupByContainer{
		AdditionalPivotGroupByContainerProperty: make(map[string]json.RawMessage),
	}

	return r
}
