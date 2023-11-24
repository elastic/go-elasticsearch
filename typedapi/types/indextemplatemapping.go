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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

// IndexTemplateMapping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/put_index_template/IndicesPutIndexTemplateRequest.ts#L97-L119
type IndexTemplateMapping struct {
	// Aliases Aliases to add.
	// If the index template includes a `data_stream` object, these are data stream
	// aliases.
	// Otherwise, these are index aliases.
	// Data stream aliases ignore the `index_routing`, `routing`, and
	// `search_routing` options.
	Aliases   map[string]Alias     `json:"aliases,omitempty"`
	Lifecycle *DataStreamLifecycle `json:"lifecycle,omitempty"`
	// Mappings Mapping for fields in the index.
	// If specified, this mapping can include field names, field data types, and
	// mapping parameters.
	Mappings *TypeMapping `json:"mappings,omitempty"`
	// Settings Configuration options for the index.
	Settings *IndexSettings `json:"settings,omitempty"`
}

// NewIndexTemplateMapping returns a IndexTemplateMapping.
func NewIndexTemplateMapping() *IndexTemplateMapping {
	r := &IndexTemplateMapping{
		Aliases: make(map[string]Alias, 0),
	}

	return r
}
