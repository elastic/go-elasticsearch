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

// Package appliesto
package appliesto

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/Rule.ts#L67-L72
type AppliesTo struct {
	Name string
}

var (
	Actual = AppliesTo{"actual"}

	Typical = AppliesTo{"typical"}

	Difffromtypical = AppliesTo{"diff_from_typical"}

	Time = AppliesTo{"time"}
)

func (a AppliesTo) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *AppliesTo) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "actual":
		*a = Actual
	case "typical":
		*a = Typical
	case "diff_from_typical":
		*a = Difffromtypical
	case "time":
		*a = Time
	default:
		*a = AppliesTo{string(text)}
	}

	return nil
}

func (a AppliesTo) String() string {
	return a.Name
}
