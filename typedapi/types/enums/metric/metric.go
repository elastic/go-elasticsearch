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

// Package metric
package metric

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/rollup/_types/Metric.ts#L22-L28
type Metric struct {
	Name string
}

var (
	Min = Metric{"min"}

	Max = Metric{"max"}

	Sum = Metric{"sum"}

	Avg = Metric{"avg"}

	Valuecount = Metric{"value_count"}
)

func (m Metric) MarshalText() (text []byte, err error) {
	return []byte(m.String()), nil
}

func (m *Metric) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "min":
		*m = Min
	case "max":
		*m = Max
	case "sum":
		*m = Sum
	case "avg":
		*m = Avg
	case "value_count":
		*m = Valuecount
	default:
		*m = Metric{string(text)}
	}

	return nil
}

func (m Metric) String() string {
	return m.Name
}
