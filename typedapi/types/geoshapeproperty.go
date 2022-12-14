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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoorientation"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geostrategy"
)

// GeoShapeProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_types/mapping/geo.ts#L37-L50
type GeoShapeProperty struct {
	Coerce          *bool                          `json:"coerce,omitempty"`
	CopyTo          []string                       `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	IgnoreZValue    *bool                          `json:"ignore_z_value,omitempty"`
	// Meta Metadata about the field.
	Meta        map[string]string              `json:"meta,omitempty"`
	Orientation *geoorientation.GeoOrientation `json:"orientation,omitempty"`
	Properties  map[string]Property            `json:"properties,omitempty"`
	Similarity  *string                        `json:"similarity,omitempty"`
	Store       *bool                          `json:"store,omitempty"`
	Strategy    *geostrategy.GeoStrategy       `json:"strategy,omitempty"`
	Type        string                         `json:"type,omitempty"`
}

// NewGeoShapeProperty returns a GeoShapeProperty.
func NewGeoShapeProperty() *GeoShapeProperty {
	r := &GeoShapeProperty{
		Fields:     make(map[string]Property, 0),
		Meta:       make(map[string]string, 0),
		Properties: make(map[string]Property, 0),
	}

	r.Type = "geo_shape"

	return r
}
