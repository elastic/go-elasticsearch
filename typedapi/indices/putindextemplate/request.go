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
// https://github.com/elastic/elasticsearch-specification/tree/5c8fed5fe577b0d5e9fde34fb13795c5a66fe9fe

package putindextemplate

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putindextemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/5c8fed5fe577b0d5e9fde34fb13795c5a66fe9fe/specification/indices/put_index_template/IndicesPutIndexTemplateRequest.ts#L36-L95
type Request struct {

	// ComposedOf An ordered list of component template names.
	// Component templates are merged in the order specified, meaning that the last
	// component template specified has the highest precedence.
	ComposedOf []string `json:"composed_of,omitempty"`
	// DataStream If this object is included, the template is used to create data streams and
	// their backing indices.
	// Supports an empty object.
	// Data streams require a matching index template with a `data_stream` object.
	DataStream *types.DataStreamVisibility `json:"data_stream,omitempty"`
	// IndexPatterns Name of the index template to create.
	IndexPatterns []string `json:"index_patterns,omitempty"`
	// Meta_ Optional user metadata about the index template.
	// May have any contents.
	// This map is not automatically generated by Elasticsearch.
	Meta_ types.Metadata `json:"_meta,omitempty"`
	// Priority Priority to determine index template precedence when a new data stream or
	// index is created.
	// The index template with the highest priority is chosen.
	// If no priority is specified the template is treated as though it is of
	// priority 0 (lowest priority).
	// This number is not automatically generated by Elasticsearch.
	Priority *int `json:"priority,omitempty"`
	// Template Template to be applied.
	// It may optionally include an `aliases`, `mappings`, or `settings`
	// configuration.
	Template *types.IndexTemplateMapping `json:"template,omitempty"`
	// Version Version number used to manage index templates externally.
	// This number is not automatically generated by Elasticsearch.
	Version *int64 `json:"version,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putindextemplate request: %w", err)
	}

	return &req, nil
}
