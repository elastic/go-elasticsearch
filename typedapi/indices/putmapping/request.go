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


package putmapping

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// Request holds the request body struct for the package putmapping
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/indices/put_mapping/IndicesPutMappingRequest.ts#L42-L116
type Request struct {

	// DateDetection Controls whether dynamic date detection is enabled.
	DateDetection *bool `json:"date_detection,omitempty"`
	// Dynamic Controls whether new fields are added dynamically.
	Dynamic *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	// DynamicDateFormats If date detection is enabled then new string fields are checked
	// against 'dynamic_date_formats' and if the value matches then
	// a new date field is added instead of string.
	DynamicDateFormats []string `json:"dynamic_date_formats,omitempty"`
	// DynamicTemplates Specify dynamic templates for the mapping.
	DynamicTemplates []map[string]types.DynamicTemplate `json:"dynamic_templates,omitempty"`
	// FieldNames_ Control whether field names are enabled for the index.
	FieldNames_ *types.FieldNamesField `json:"_field_names,omitempty"`
	// Meta_ A mapping type can have custom meta data associated with it. These are
	// not used at all by Elasticsearch, but can be used to store
	// application-specific metadata.
	Meta_ map[string]interface{} `json:"_meta,omitempty"`
	// NumericDetection Automatically map strings into numeric data types for all fields.
	NumericDetection *bool `json:"numeric_detection,omitempty"`
	// Properties Mapping for a field. For new fields, this mapping can include:
	//
	// - Field name
	// - Field data type
	// - Mapping parameters
	Properties map[string]types.Property `json:"properties,omitempty"`
	// Routing_ Enable making a routing value required on indexed documents.
	Routing_ *types.RoutingField `json:"_routing,omitempty"`
	// Runtime Mapping of runtime fields for the index.
	Runtime map[string]types.RuntimeField `json:"runtime,omitempty"`
	// Source_ Control whether the _source field is enabled on the index.
	Source_ *types.SourceField `json:"_source,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Properties: make(map[string]types.Property, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putmapping request: %w", err)
	}

	return &req, nil
}
