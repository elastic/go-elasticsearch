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

// Package normalization
package normalization

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/Similarity.ts#L52-L58
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
