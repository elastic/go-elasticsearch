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
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/calendarinterval"
)

// DateHistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/_types/aggregations/bucket.ts#L93-L110
type DateHistogramAggregation struct {
	CalendarInterval *calendarinterval.CalendarInterval `json:"calendar_interval,omitempty"`
	ExtendedBounds   *ExtendedBoundsFieldDateMath       `json:"extended_bounds,omitempty"`
	Field            *string                            `json:"field,omitempty"`
	FixedInterval    *Duration                          `json:"fixed_interval,omitempty"`
	Format           *string                            `json:"format,omitempty"`
	HardBounds       *ExtendedBoundsFieldDateMath       `json:"hard_bounds,omitempty"`
	Interval         *Duration                          `json:"interval,omitempty"`
	Keyed            *bool                              `json:"keyed,omitempty"`
	Meta             map[string]interface{}             `json:"meta,omitempty"`
	MinDocCount      *int                               `json:"min_doc_count,omitempty"`
	Missing          *DateTime                          `json:"missing,omitempty"`
	Name             *string                            `json:"name,omitempty"`
	Offset           *Duration                          `json:"offset,omitempty"`
	Order            *AggregateOrder                    `json:"order,omitempty"`
	Params           map[string]interface{}             `json:"params,omitempty"`
	Script           *Script                            `json:"script,omitempty"`
	TimeZone         *string                            `json:"time_zone,omitempty"`
}

// NewDateHistogramAggregation returns a DateHistogramAggregation.
func NewDateHistogramAggregation() *DateHistogramAggregation {
	r := &DateHistogramAggregation{
		Params: make(map[string]interface{}, 0),
	}

	return r
}
