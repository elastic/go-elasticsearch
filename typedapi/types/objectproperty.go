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
// https://github.com/elastic/elasticsearch-specification/tree/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// ObjectProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e/specification/_types/mapping/complex.ts#L46-L49
type ObjectProperty struct {
	CopyTo      []string                       `json:"copy_to,omitempty"`
	Dynamic     *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Enabled     *bool                          `json:"enabled,omitempty"`
	Fields      map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove *int                           `json:"ignore_above,omitempty"`
	// Meta Metadata about the field.
	Meta       map[string]string   `json:"meta,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
	Similarity *string             `json:"similarity,omitempty"`
	Store      *bool               `json:"store,omitempty"`
	Type       string              `json:"type,omitempty"`
}

// NewObjectProperty returns a ObjectProperty.
func NewObjectProperty() *ObjectProperty {
	r := &ObjectProperty{
		Fields:     make(map[string]Property, 0),
		Meta:       make(map[string]string, 0),
		Properties: make(map[string]Property, 0),
	}

	r.Type = "object"

	return r
}
