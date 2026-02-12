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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package nvidiasimilaritytype
package nvidiasimilaritytype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/inference/_types/CommonTypes.ts#L2021-L2025
type NvidiaSimilarityType struct {
	Name string
}

var (
	Cosine = NvidiaSimilarityType{"cosine"}

	Dotproduct = NvidiaSimilarityType{"dot_product"}

	L2norm = NvidiaSimilarityType{"l2_norm"}
)

func (n NvidiaSimilarityType) MarshalText() (text []byte, err error) {
	return []byte(n.String()), nil
}

func (n *NvidiaSimilarityType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "cosine":
		*n = Cosine
	case "dot_product":
		*n = Dotproduct
	case "l2_norm":
		*n = L2norm
	default:
		*n = NvidiaSimilarityType{string(text)}
	}

	return nil
}

func (n NvidiaSimilarityType) String() string {
	return n.Name
}
