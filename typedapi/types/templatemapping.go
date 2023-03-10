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

import (
	"encoding/json"
)

// TemplateMapping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/indices/_types/TemplateMapping.ts#L27-L34
type TemplateMapping struct {
	Aliases       map[string]Alias           `json:"aliases"`
	IndexPatterns []string                   `json:"index_patterns"`
	Mappings      TypeMapping                `json:"mappings"`
	Order         int                        `json:"order"`
	Settings      map[string]json.RawMessage `json:"settings"`
	Version       *int64                     `json:"version,omitempty"`
}

// NewTemplateMapping returns a TemplateMapping.
func NewTemplateMapping() *TemplateMapping {
	r := &TemplateMapping{
		Aliases:  make(map[string]Alias, 0),
		Settings: make(map[string]json.RawMessage, 0),
	}

	return r
}
