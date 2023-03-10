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

// Package icutransformdirection
package icutransformdirection

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/analysis/icu-plugin.ts#L73-L76
type IcuTransformDirection struct {
	Name string
}

var (
	Forward = IcuTransformDirection{"forward"}

	Reverse = IcuTransformDirection{"reverse"}
)

func (i IcuTransformDirection) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *IcuTransformDirection) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "forward":
		*i = Forward
	case "reverse":
		*i = Reverse
	default:
		*i = IcuTransformDirection{string(text)}
	}

	return nil
}

func (i IcuTransformDirection) String() string {
	return i.Name
}
