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
// https://github.com/elastic/elasticsearch-specification/tree/ec3159eb31c62611202a4fb157ea88fa6ff78e1a


package types

// DateDistanceFeatureQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/_types/query_dsl/specialized.ts#L51-L54
type DateDistanceFeatureQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	Field      string   `json:"field"`
	Origin     string   `json:"origin"`
	Pivot      Duration `json:"pivot"`
	QueryName_ *string  `json:"_name,omitempty"`
}

// NewDateDistanceFeatureQuery returns a DateDistanceFeatureQuery.
func NewDateDistanceFeatureQuery() *DateDistanceFeatureQuery {
	r := &DateDistanceFeatureQuery{}

	return r
}
