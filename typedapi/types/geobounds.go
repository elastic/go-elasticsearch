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

// GeoBounds holds the union for the following types:
//
//	CoordsGeoBounds
//	TopLeftBottomRightGeoBounds
//	TopRightBottomLeftGeoBounds
//	WktGeoBounds
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Geo.ts#L116-L129
type GeoBounds interface{}

// GeoBoundsBuilder holds GeoBounds struct and provides a builder API.
type GeoBoundsBuilder struct {
	v GeoBounds
}

// NewGeoBounds provides a builder for the GeoBounds struct.
func NewGeoBoundsBuilder() *GeoBoundsBuilder {
	return &GeoBoundsBuilder{}
}

// Build finalize the chain and returns the GeoBounds struct
func (u *GeoBoundsBuilder) Build() GeoBounds {
	return u.v
}

func (u *GeoBoundsBuilder) CoordsGeoBounds(coordsgeobounds *CoordsGeoBoundsBuilder) *GeoBoundsBuilder {
	v := coordsgeobounds.Build()
	u.v = &v
	return u
}

func (u *GeoBoundsBuilder) TopLeftBottomRightGeoBounds(topleftbottomrightgeobounds *TopLeftBottomRightGeoBoundsBuilder) *GeoBoundsBuilder {
	v := topleftbottomrightgeobounds.Build()
	u.v = &v
	return u
}

func (u *GeoBoundsBuilder) TopRightBottomLeftGeoBounds(toprightbottomleftgeobounds *TopRightBottomLeftGeoBoundsBuilder) *GeoBoundsBuilder {
	v := toprightbottomleftgeobounds.Build()
	u.v = &v
	return u
}

func (u *GeoBoundsBuilder) WktGeoBounds(wktgeobounds *WktGeoBoundsBuilder) *GeoBoundsBuilder {
	v := wktgeobounds.Build()
	u.v = &v
	return u
}
