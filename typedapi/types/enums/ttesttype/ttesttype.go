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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Package ttesttype
package ttesttype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/aggregations/metric.ts#L165-L169
type TTestType struct {
	Name string
}

var (
	Paired = TTestType{"paired"}

	Homoscedastic = TTestType{"homoscedastic"}

	Heteroscedastic = TTestType{"heteroscedastic"}
)

func (t TTestType) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TTestType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "paired":
		*t = Paired
	case "homoscedastic":
		*t = Homoscedastic
	case "heteroscedastic":
		*t = Heteroscedastic
	default:
		*t = TTestType{string(text)}
	}

	return nil
}

func (t TTestType) String() string {
	return t.Name
}
