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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// GeoShapeQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/query_dsl/geo.ts#L86-L91
type GeoShapeQuery struct {
	Boost          *float32                     `json:"boost,omitempty"`
	GeoShapeQuery  map[Field]GeoShapeFieldQuery `json:"GeoShapeQuery,omitempty"`
	IgnoreUnmapped *bool                        `json:"ignore_unmapped,omitempty"`
	QueryName_     *string                      `json:"_name,omitempty"`
}

// GeoShapeQueryBuilder holds GeoShapeQuery struct and provides a builder API.
type GeoShapeQueryBuilder struct {
	v *GeoShapeQuery
}

// NewGeoShapeQuery provides a builder for the GeoShapeQuery struct.
func NewGeoShapeQueryBuilder() *GeoShapeQueryBuilder {
	r := GeoShapeQueryBuilder{
		&GeoShapeQuery{
			GeoShapeQuery: make(map[Field]GeoShapeFieldQuery, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GeoShapeQuery struct
func (rb *GeoShapeQueryBuilder) Build() GeoShapeQuery {
	return *rb.v
}

func (rb *GeoShapeQueryBuilder) Boost(boost float32) *GeoShapeQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *GeoShapeQueryBuilder) GeoShapeQuery(values map[Field]*GeoShapeFieldQueryBuilder) *GeoShapeQueryBuilder {
	tmp := make(map[Field]GeoShapeFieldQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.GeoShapeQuery = tmp
	return rb
}

func (rb *GeoShapeQueryBuilder) IgnoreUnmapped(ignoreunmapped bool) *GeoShapeQueryBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

func (rb *GeoShapeQueryBuilder) QueryName_(queryname_ string) *GeoShapeQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
