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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geovalidationmethod"
)

// GeoPolygonQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/geo.ts#L63-L71
type GeoPolygonQuery struct {
	Boost            *float32                                 `json:"boost,omitempty"`
	GeoPolygonQuery  map[Field]GeoPolygonPoints               `json:"-"`
	IgnoreUnmapped   *bool                                    `json:"ignore_unmapped,omitempty"`
	QueryName_       *string                                  `json:"_name,omitempty"`
	ValidationMethod *geovalidationmethod.GeoValidationMethod `json:"validation_method,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s GeoPolygonQuery) MarshalJSON() ([]byte, error) {
	type opt GeoPolygonQuery
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.GeoPolygonQuery {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GeoPolygonQueryBuilder holds GeoPolygonQuery struct and provides a builder API.
type GeoPolygonQueryBuilder struct {
	v *GeoPolygonQuery
}

// NewGeoPolygonQuery provides a builder for the GeoPolygonQuery struct.
func NewGeoPolygonQueryBuilder() *GeoPolygonQueryBuilder {
	r := GeoPolygonQueryBuilder{
		&GeoPolygonQuery{
			GeoPolygonQuery: make(map[Field]GeoPolygonPoints, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GeoPolygonQuery struct
func (rb *GeoPolygonQueryBuilder) Build() GeoPolygonQuery {
	return *rb.v
}

func (rb *GeoPolygonQueryBuilder) Boost(boost float32) *GeoPolygonQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *GeoPolygonQueryBuilder) GeoPolygonQuery(values map[Field]*GeoPolygonPointsBuilder) *GeoPolygonQueryBuilder {
	tmp := make(map[Field]GeoPolygonPoints, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.GeoPolygonQuery = tmp
	return rb
}

func (rb *GeoPolygonQueryBuilder) IgnoreUnmapped(ignoreunmapped bool) *GeoPolygonQueryBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

func (rb *GeoPolygonQueryBuilder) QueryName_(queryname_ string) *GeoPolygonQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *GeoPolygonQueryBuilder) ValidationMethod(validationmethod geovalidationmethod.GeoValidationMethod) *GeoPolygonQueryBuilder {
	rb.v.ValidationMethod = &validationmethod
	return rb
}
