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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// SearchTemplateRequestBody type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/watcher/_types/Input.ts#L128-L145
type SearchTemplateRequestBody struct {
	Explain *bool `json:"explain,omitempty"`
	// Id ID of the search template to use. If no source is specified,
	// this parameter is required.
	Id      *string                `json:"id,omitempty"`
	Params  map[string]interface{} `json:"params,omitempty"`
	Profile *bool                  `json:"profile,omitempty"`
	// Source An inline search template. Supports the same parameters as the search API's
	// request body. Also supports Mustache variables. If no id is specified, this
	// parameter is required.
	Source *string `json:"source,omitempty"`
}

// NewSearchTemplateRequestBody returns a SearchTemplateRequestBody.
func NewSearchTemplateRequestBody() *SearchTemplateRequestBody {
	r := &SearchTemplateRequestBody{
		Params: make(map[string]interface{}, 0),
	}

	return r
}
