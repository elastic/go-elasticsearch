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

// DistanceFeatureQueryBaseDateMathDuration type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/specialized.ts#L40-L44
type DistanceFeatureQueryBaseDateMathDuration struct {
	Boost      *float32 `json:"boost,omitempty"`
	Field      Field    `json:"field"`
	Origin     DateMath `json:"origin"`
	Pivot      Duration `json:"pivot"`
	QueryName_ *string  `json:"_name,omitempty"`
}

// DistanceFeatureQueryBaseDateMathDurationBuilder holds DistanceFeatureQueryBaseDateMathDuration struct and provides a builder API.
type DistanceFeatureQueryBaseDateMathDurationBuilder struct {
	v *DistanceFeatureQueryBaseDateMathDuration
}

// NewDistanceFeatureQueryBaseDateMathDuration provides a builder for the DistanceFeatureQueryBaseDateMathDuration struct.
func NewDistanceFeatureQueryBaseDateMathDurationBuilder() *DistanceFeatureQueryBaseDateMathDurationBuilder {
	r := DistanceFeatureQueryBaseDateMathDurationBuilder{
		&DistanceFeatureQueryBaseDateMathDuration{},
	}

	return &r
}

// Build finalize the chain and returns the DistanceFeatureQueryBaseDateMathDuration struct
func (rb *DistanceFeatureQueryBaseDateMathDurationBuilder) Build() DistanceFeatureQueryBaseDateMathDuration {
	return *rb.v
}

func (rb *DistanceFeatureQueryBaseDateMathDurationBuilder) Boost(boost float32) *DistanceFeatureQueryBaseDateMathDurationBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *DistanceFeatureQueryBaseDateMathDurationBuilder) Field(field Field) *DistanceFeatureQueryBaseDateMathDurationBuilder {
	rb.v.Field = field
	return rb
}

func (rb *DistanceFeatureQueryBaseDateMathDurationBuilder) Origin(origin DateMath) *DistanceFeatureQueryBaseDateMathDurationBuilder {
	rb.v.Origin = origin
	return rb
}

func (rb *DistanceFeatureQueryBaseDateMathDurationBuilder) Pivot(pivot *DurationBuilder) *DistanceFeatureQueryBaseDateMathDurationBuilder {
	v := pivot.Build()
	rb.v.Pivot = v
	return rb
}

func (rb *DistanceFeatureQueryBaseDateMathDurationBuilder) QueryName_(queryname_ string) *DistanceFeatureQueryBaseDateMathDurationBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
