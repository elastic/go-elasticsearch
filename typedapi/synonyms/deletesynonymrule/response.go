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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package deletesynonymrule

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/result"
)

// Response holds the response body struct for the package deletesynonymrule
//
// https://github.com/elastic/elasticsearch-specification/blob/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6/specification/synonyms/delete_synonym_rule/SynonymRuleDeleteResponse.ts#L22-L25
type Response struct {
	// ReloadAnalyzersDetails Updating synonyms in a synonym set can reload the associated analyzers in
	// case refresh is set to true.
	// This information is the analyzers reloading result.
	ReloadAnalyzersDetails *types.ReloadResult `json:"reload_analyzers_details,omitempty"`
	// Result The update operation result.
	Result result.Result `json:"result"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
