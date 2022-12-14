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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

// ElisionTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_types/analysis/token_filters.ts#L186-L191
type ElisionTokenFilter struct {
	Articles     []string `json:"articles,omitempty"`
	ArticlesCase *bool    `json:"articles_case,omitempty"`
	ArticlesPath *string  `json:"articles_path,omitempty"`
	Type         string   `json:"type,omitempty"`
	Version      *string  `json:"version,omitempty"`
}

// NewElisionTokenFilter returns a ElisionTokenFilter.
func NewElisionTokenFilter() *ElisionTokenFilter {
	r := &ElisionTokenFilter{}

	r.Type = "elision"

	return r
}
