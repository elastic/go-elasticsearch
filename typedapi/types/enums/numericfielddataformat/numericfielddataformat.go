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

// Package numericfielddataformat
package numericfielddataformat

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/indices/_types/NumericFielddataFormat.ts#L20-L23
type NumericFielddataFormat struct {
	Name string
}

var (
	Array = NumericFielddataFormat{"array"}

	Disabled = NumericFielddataFormat{"disabled"}
)

func (n NumericFielddataFormat) MarshalText() (text []byte, err error) {
	return []byte(n.String()), nil
}

func (n *NumericFielddataFormat) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "array":
		*n = Array
	case "disabled":
		*n = Disabled
	default:
		*n = NumericFielddataFormat{string(text)}
	}

	return nil
}

func (n NumericFielddataFormat) String() string {
	return n.Name
}
