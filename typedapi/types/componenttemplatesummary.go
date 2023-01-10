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

// ComponentTemplateSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cluster/_types/ComponentTemplate.ts#L38-L45
type ComponentTemplateSummary struct {
	Aliases  map[string]AliasDefinition `json:"aliases,omitempty"`
	Mappings *TypeMapping               `json:"mappings,omitempty"`
	Meta_    map[string]interface{}     `json:"_meta,omitempty"`
	Settings map[string]IndexSettings   `json:"settings,omitempty"`
	Version  *int64                     `json:"version,omitempty"`
}

// NewComponentTemplateSummary returns a ComponentTemplateSummary.
func NewComponentTemplateSummary() *ComponentTemplateSummary {
	r := &ComponentTemplateSummary{
		Aliases:  make(map[string]AliasDefinition, 0),
		Settings: make(map[string]IndexSettings, 0),
	}

	return r
}
