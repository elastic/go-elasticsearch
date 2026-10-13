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
// https://github.com/elastic/elasticsearch-specification/tree/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6

// Package dynamic
package dynamic

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/56c1eabdd35f941d1fbb3ad7ad8a9676664223f6/specification/esql/_types/types.ts#L82-L89
type Dynamic struct {
	Name string
}

var (
	True = Dynamic{"true"}

	False = Dynamic{"false"}
)

func (d *Dynamic) UnmarshalJSON(data []byte) error {
	return d.UnmarshalText(data)
}

func (d Dynamic) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *Dynamic) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "true":
		*d = True
	case "false":
		*d = False
	default:
		*d = Dynamic{string(text)}
	}

	return nil
}

func (d Dynamic) String() string {
	return d.Name
}
