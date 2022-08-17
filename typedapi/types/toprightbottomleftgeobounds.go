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

// TopRightBottomLeftGeoBounds type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Geo.ts#L147-L150
type TopRightBottomLeftGeoBounds struct {
	BottomLeft GeoLocation `json:"bottom_left"`
	TopRight   GeoLocation `json:"top_right"`
}

// TopRightBottomLeftGeoBoundsBuilder holds TopRightBottomLeftGeoBounds struct and provides a builder API.
type TopRightBottomLeftGeoBoundsBuilder struct {
	v *TopRightBottomLeftGeoBounds
}

// NewTopRightBottomLeftGeoBounds provides a builder for the TopRightBottomLeftGeoBounds struct.
func NewTopRightBottomLeftGeoBoundsBuilder() *TopRightBottomLeftGeoBoundsBuilder {
	r := TopRightBottomLeftGeoBoundsBuilder{
		&TopRightBottomLeftGeoBounds{},
	}

	return &r
}

// Build finalize the chain and returns the TopRightBottomLeftGeoBounds struct
func (rb *TopRightBottomLeftGeoBoundsBuilder) Build() TopRightBottomLeftGeoBounds {
	return *rb.v
}

func (rb *TopRightBottomLeftGeoBoundsBuilder) BottomLeft(bottomleft *GeoLocationBuilder) *TopRightBottomLeftGeoBoundsBuilder {
	v := bottomleft.Build()
	rb.v.BottomLeft = v
	return rb
}

func (rb *TopRightBottomLeftGeoBoundsBuilder) TopRight(topright *GeoLocationBuilder) *TopRightBottomLeftGeoBoundsBuilder {
	v := topright.Build()
	rb.v.TopRight = v
	return rb
}
