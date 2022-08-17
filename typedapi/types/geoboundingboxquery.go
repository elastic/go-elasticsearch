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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoexecution"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geovalidationmethod"
)

// GeoBoundingBoxQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/geo.ts#L32-L41
type GeoBoundingBoxQuery struct {
	Boost               *float32                                 `json:"boost,omitempty"`
	GeoBoundingBoxQuery map[Field]GeoBounds                      `json:"-"`
	IgnoreUnmapped      *bool                                    `json:"ignore_unmapped,omitempty"`
	QueryName_          *string                                  `json:"_name,omitempty"`
	Type                *geoexecution.GeoExecution               `json:"type,omitempty"`
	ValidationMethod    *geovalidationmethod.GeoValidationMethod `json:"validation_method,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s GeoBoundingBoxQuery) MarshalJSON() ([]byte, error) {
	type opt GeoBoundingBoxQuery
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
	for key, value := range s.GeoBoundingBoxQuery {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GeoBoundingBoxQueryBuilder holds GeoBoundingBoxQuery struct and provides a builder API.
type GeoBoundingBoxQueryBuilder struct {
	v *GeoBoundingBoxQuery
}

// NewGeoBoundingBoxQuery provides a builder for the GeoBoundingBoxQuery struct.
func NewGeoBoundingBoxQueryBuilder() *GeoBoundingBoxQueryBuilder {
	r := GeoBoundingBoxQueryBuilder{
		&GeoBoundingBoxQuery{
			GeoBoundingBoxQuery: make(map[Field]GeoBounds, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GeoBoundingBoxQuery struct
func (rb *GeoBoundingBoxQueryBuilder) Build() GeoBoundingBoxQuery {
	return *rb.v
}

func (rb *GeoBoundingBoxQueryBuilder) Boost(boost float32) *GeoBoundingBoxQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *GeoBoundingBoxQueryBuilder) GeoBoundingBoxQuery(values map[Field]*GeoBoundsBuilder) *GeoBoundingBoxQueryBuilder {
	tmp := make(map[Field]GeoBounds, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.GeoBoundingBoxQuery = tmp
	return rb
}

func (rb *GeoBoundingBoxQueryBuilder) IgnoreUnmapped(ignoreunmapped bool) *GeoBoundingBoxQueryBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

func (rb *GeoBoundingBoxQueryBuilder) QueryName_(queryname_ string) *GeoBoundingBoxQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *GeoBoundingBoxQueryBuilder) Type_(type_ geoexecution.GeoExecution) *GeoBoundingBoxQueryBuilder {
	rb.v.Type = &type_
	return rb
}

func (rb *GeoBoundingBoxQueryBuilder) ValidationMethod(validationmethod geovalidationmethod.GeoValidationMethod) *GeoBoundingBoxQueryBuilder {
	rb.v.ValidationMethod = &validationmethod
	return rb
}
