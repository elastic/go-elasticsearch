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
)

// ShapeQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/specialized.ts#L176-L181
type ShapeQuery struct {
	Boost          *float32                  `json:"boost,omitempty"`
	IgnoreUnmapped *bool                     `json:"ignore_unmapped,omitempty"`
	QueryName_     *string                   `json:"_name,omitempty"`
	ShapeQuery     map[Field]ShapeFieldQuery `json:"-"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s ShapeQuery) MarshalJSON() ([]byte, error) {
	type opt ShapeQuery
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
	for key, value := range s.ShapeQuery {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ShapeQueryBuilder holds ShapeQuery struct and provides a builder API.
type ShapeQueryBuilder struct {
	v *ShapeQuery
}

// NewShapeQuery provides a builder for the ShapeQuery struct.
func NewShapeQueryBuilder() *ShapeQueryBuilder {
	r := ShapeQueryBuilder{
		&ShapeQuery{
			ShapeQuery: make(map[Field]ShapeFieldQuery, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ShapeQuery struct
func (rb *ShapeQueryBuilder) Build() ShapeQuery {
	return *rb.v
}

func (rb *ShapeQueryBuilder) Boost(boost float32) *ShapeQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *ShapeQueryBuilder) IgnoreUnmapped(ignoreunmapped bool) *ShapeQueryBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

func (rb *ShapeQueryBuilder) QueryName_(queryname_ string) *ShapeQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *ShapeQueryBuilder) ShapeQuery(values map[Field]*ShapeFieldQueryBuilder) *ShapeQueryBuilder {
	tmp := make(map[Field]ShapeFieldQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.ShapeQuery = tmp
	return rb
}
