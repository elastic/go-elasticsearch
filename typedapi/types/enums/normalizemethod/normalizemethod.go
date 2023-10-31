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

// Package normalizemethod
package normalizemethod

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/aggregations/pipeline.ts#L328-L354
type NormalizeMethod struct {
	Name string
}

var (
	Rescale01 = NormalizeMethod{"rescale_0_1"}

	Rescale0100 = NormalizeMethod{"rescale_0_100"}

	Percentofsum = NormalizeMethod{"percent_of_sum"}

	Mean = NormalizeMethod{"mean"}

	ZScore = NormalizeMethod{"z-score"}

	Softmax = NormalizeMethod{"softmax"}
)

func (n NormalizeMethod) MarshalText() (text []byte, err error) {
	return []byte(n.String()), nil
}

func (n *NormalizeMethod) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "rescale_0_1":
		*n = Rescale01
	case "rescale_0_100":
		*n = Rescale0100
	case "percent_of_sum":
		*n = Percentofsum
	case "mean":
		*n = Mean
	case "z-score":
		*n = ZScore
	case "softmax":
		*n = Softmax
	default:
		*n = NormalizeMethod{string(text)}
	}

	return nil
}

func (n NormalizeMethod) String() string {
	return n.Name
}
