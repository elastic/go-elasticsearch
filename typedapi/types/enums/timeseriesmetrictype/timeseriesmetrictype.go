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

// Package timeseriesmetrictype
package timeseriesmetrictype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/mapping/TimeSeriesMetricType.ts#L20-L26
type TimeSeriesMetricType struct {
	Name string
}

var (
	Gauge = TimeSeriesMetricType{"gauge"}

	Counter = TimeSeriesMetricType{"counter"}

	Summary = TimeSeriesMetricType{"summary"}

	Histogram = TimeSeriesMetricType{"histogram"}

	Position = TimeSeriesMetricType{"position"}
)

func (t TimeSeriesMetricType) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TimeSeriesMetricType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "gauge":
		*t = Gauge
	case "counter":
		*t = Counter
	case "summary":
		*t = Summary
	case "histogram":
		*t = Histogram
	case "position":
		*t = Position
	default:
		*t = TimeSeriesMetricType{string(text)}
	}

	return nil
}

func (t TimeSeriesMetricType) String() string {
	return t.Name
}
