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

package types

// RuntimeFieldsType type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/xpack/usage/types.ts#L273-L288
type RuntimeFieldsType struct {
	CharsMax        int64    `json:"chars_max"`
	CharsTotal      int64    `json:"chars_total"`
	Count           int64    `json:"count"`
	DocMax          int64    `json:"doc_max"`
	DocTotal        int64    `json:"doc_total"`
	IndexCount      int64    `json:"index_count"`
	Lang            []string `json:"lang"`
	LinesMax        int64    `json:"lines_max"`
	LinesTotal      int64    `json:"lines_total"`
	Name            string   `json:"name"`
	ScriptlessCount int64    `json:"scriptless_count"`
	ShadowedCount   int64    `json:"shadowed_count"`
	SourceMax       int64    `json:"source_max"`
	SourceTotal     int64    `json:"source_total"`
}

// NewRuntimeFieldsType returns a RuntimeFieldsType.
func NewRuntimeFieldsType() *RuntimeFieldsType {
	r := &RuntimeFieldsType{}

	return r
}
