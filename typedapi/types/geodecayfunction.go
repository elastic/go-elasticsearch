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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/multivaluemode"
)

// GeoDecayFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/query_dsl/compound.ts#L96-L98
type GeoDecayFunction struct {
	GeoDecayFunction map[Field]DecayPlacementGeoLocationDistance `json:"GeoDecayFunction,omitempty"`
	MultiValueMode   *multivaluemode.MultiValueMode              `json:"multi_value_mode,omitempty"`
}

// GeoDecayFunctionBuilder holds GeoDecayFunction struct and provides a builder API.
type GeoDecayFunctionBuilder struct {
	v *GeoDecayFunction
}

// NewGeoDecayFunction provides a builder for the GeoDecayFunction struct.
func NewGeoDecayFunctionBuilder() *GeoDecayFunctionBuilder {
	r := GeoDecayFunctionBuilder{
		&GeoDecayFunction{
			GeoDecayFunction: make(map[Field]DecayPlacementGeoLocationDistance, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GeoDecayFunction struct
func (rb *GeoDecayFunctionBuilder) Build() GeoDecayFunction {
	return *rb.v
}

func (rb *GeoDecayFunctionBuilder) GeoDecayFunction(values map[Field]*DecayPlacementGeoLocationDistanceBuilder) *GeoDecayFunctionBuilder {
	tmp := make(map[Field]DecayPlacementGeoLocationDistance, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.GeoDecayFunction = tmp
	return rb
}

func (rb *GeoDecayFunctionBuilder) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *GeoDecayFunctionBuilder {
	rb.v.MultiValueMode = &multivaluemode
	return rb
}
