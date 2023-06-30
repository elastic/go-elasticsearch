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

// Package normalization
package normalization

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/Similarity.ts#L52-L58
type Normalization struct {
	Name string
}

var (
	No = Normalization{"no"}

	H1 = Normalization{"h1"}

	H2 = Normalization{"h2"}

	H3 = Normalization{"h3"}

	Z = Normalization{"z"}
)

func (n Normalization) MarshalText() (text []byte, err error) {
	return []byte(n.String()), nil
}

func (n *Normalization) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "no":
		*n = No
	case "h1":
		*n = H1
	case "h2":
		*n = H2
	case "h3":
		*n = H3
	case "z":
		*n = Z
	default:
		*n = Normalization{string(text)}
	}

	return nil
}

func (n Normalization) String() string {
	return n.Name
}
