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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package types

// SearchApplicationTemplate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b1811e10a0722431d79d1c234dd412ff47d8656f/specification/search_application/_types/SearchApplicationTemplate.ts#L22-L27
type SearchApplicationTemplate struct {
	// Script The associated mustache template.
	Script Script `json:"script"`
}

// NewSearchApplicationTemplate returns a SearchApplicationTemplate.
func NewSearchApplicationTemplate() *SearchApplicationTemplate {
	r := &SearchApplicationTemplate{}

	return r
}

type SearchApplicationTemplateVariant interface {
	SearchApplicationTemplateCaster() *SearchApplicationTemplate
}

func (s *SearchApplicationTemplate) SearchApplicationTemplateCaster() *SearchApplicationTemplate {
	return s
}
