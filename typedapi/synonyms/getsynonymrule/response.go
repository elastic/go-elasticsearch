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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package getsynonymrule

// Response holds the response body struct for the package getsynonymrule
//
// https://github.com/elastic/elasticsearch-specification/blob/52c473efb1fb5320a5bac12572d0b285882862fb/specification/synonyms/get_synonym_rule/SynonymRuleGetResponse.ts#L22-L25
type Response struct {

	// Id Synonym Rule identifier
	Id string `json:"id"`
	// Synonyms Synonyms, in Solr format, that conform the synonym rule.
	Synonyms string `json:"synonyms"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
