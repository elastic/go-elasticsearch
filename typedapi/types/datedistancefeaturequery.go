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

// DateDistanceFeatureQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/specialized.ts#L51-L54
type DateDistanceFeatureQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	Field      Field    `json:"field"`
	Origin     DateMath `json:"origin"`
	Pivot      Duration `json:"pivot"`
	QueryName_ *string  `json:"_name,omitempty"`
}

// DateDistanceFeatureQueryBuilder holds DateDistanceFeatureQuery struct and provides a builder API.
type DateDistanceFeatureQueryBuilder struct {
	v *DateDistanceFeatureQuery
}

// NewDateDistanceFeatureQuery provides a builder for the DateDistanceFeatureQuery struct.
func NewDateDistanceFeatureQueryBuilder() *DateDistanceFeatureQueryBuilder {
	r := DateDistanceFeatureQueryBuilder{
		&DateDistanceFeatureQuery{},
	}

	return &r
}

// Build finalize the chain and returns the DateDistanceFeatureQuery struct
func (rb *DateDistanceFeatureQueryBuilder) Build() DateDistanceFeatureQuery {
	return *rb.v
}

func (rb *DateDistanceFeatureQueryBuilder) Boost(boost float32) *DateDistanceFeatureQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *DateDistanceFeatureQueryBuilder) Field(field Field) *DateDistanceFeatureQueryBuilder {
	rb.v.Field = field
	return rb
}

func (rb *DateDistanceFeatureQueryBuilder) Origin(origin DateMath) *DateDistanceFeatureQueryBuilder {
	rb.v.Origin = origin
	return rb
}

func (rb *DateDistanceFeatureQueryBuilder) Pivot(pivot *DurationBuilder) *DateDistanceFeatureQueryBuilder {
	v := pivot.Build()
	rb.v.Pivot = v
	return rb
}

func (rb *DateDistanceFeatureQueryBuilder) QueryName_(queryname_ string) *DateDistanceFeatureQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
