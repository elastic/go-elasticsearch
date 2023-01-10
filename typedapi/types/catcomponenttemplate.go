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

// CatComponentTemplate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/component_templates/types.ts#L20-L28
type CatComponentTemplate struct {
	AliasCount    string `json:"alias_count"`
	IncludedIn    string `json:"included_in"`
	MappingCount  string `json:"mapping_count"`
	MetadataCount string `json:"metadata_count"`
	Name          string `json:"name"`
	SettingsCount string `json:"settings_count"`
	Version       string `json:"version"`
}

// NewCatComponentTemplate returns a CatComponentTemplate.
func NewCatComponentTemplate() *CatComponentTemplate {
	r := &CatComponentTemplate{}

	return r
}
