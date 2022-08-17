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

// CoordsGeoBounds type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Geo.ts#L135-L140
type CoordsGeoBounds struct {
	Bottom float64 `json:"bottom"`
	Left   float64 `json:"left"`
	Right  float64 `json:"right"`
	Top    float64 `json:"top"`
}

// CoordsGeoBoundsBuilder holds CoordsGeoBounds struct and provides a builder API.
type CoordsGeoBoundsBuilder struct {
	v *CoordsGeoBounds
}

// NewCoordsGeoBounds provides a builder for the CoordsGeoBounds struct.
func NewCoordsGeoBoundsBuilder() *CoordsGeoBoundsBuilder {
	r := CoordsGeoBoundsBuilder{
		&CoordsGeoBounds{},
	}

	return &r
}

// Build finalize the chain and returns the CoordsGeoBounds struct
func (rb *CoordsGeoBoundsBuilder) Build() CoordsGeoBounds {
	return *rb.v
}

func (rb *CoordsGeoBoundsBuilder) Bottom(bottom float64) *CoordsGeoBoundsBuilder {
	rb.v.Bottom = bottom
	return rb
}

func (rb *CoordsGeoBoundsBuilder) Left(left float64) *CoordsGeoBoundsBuilder {
	rb.v.Left = left
	return rb
}

func (rb *CoordsGeoBoundsBuilder) Right(right float64) *CoordsGeoBoundsBuilder {
	rb.v.Right = right
	return rb
}

func (rb *CoordsGeoBoundsBuilder) Top(top float64) *CoordsGeoBoundsBuilder {
	rb.v.Top = top
	return rb
}
