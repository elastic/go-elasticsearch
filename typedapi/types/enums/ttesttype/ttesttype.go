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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

// Package ttesttype
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/metric.ts#L156-L160
package ttesttype

import "strings"

type TTestType struct {
	name string
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
	switch strings.ToLower(string(text)) {

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
	return t.name
}
