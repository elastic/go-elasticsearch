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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

// Package openshiftaisimilaritytype
package openshiftaisimilaritytype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27/specification/inference/_types/CommonTypes.ts#L2016-L2020
type OpenShiftAiSimilarityType struct {
	Name string
}

var (
	Cosine = OpenShiftAiSimilarityType{"cosine"}

	Dotproduct = OpenShiftAiSimilarityType{"dot_product"}

	L2norm = OpenShiftAiSimilarityType{"l2_norm"}
)

func (o OpenShiftAiSimilarityType) MarshalText() (text []byte, err error) {
	return []byte(o.String()), nil
}

func (o *OpenShiftAiSimilarityType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "cosine":
		*o = Cosine
	case "dot_product":
		*o = Dotproduct
	case "l2_norm":
		*o = L2norm
	default:
		*o = OpenShiftAiSimilarityType{string(text)}
	}

	return nil
}

func (o OpenShiftAiSimilarityType) String() string {
	return o.Name
}
