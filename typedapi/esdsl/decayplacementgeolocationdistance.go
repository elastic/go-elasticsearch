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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _decayPlacementGeoLocationDistance struct {
	v *types.DecayPlacementGeoLocationDistance
}

func NewDecayPlacementGeoLocationDistance() *_decayPlacementGeoLocationDistance {

	return &_decayPlacementGeoLocationDistance{v: types.NewDecayPlacementGeoLocationDistance()}

}

// Defines how documents are scored at the distance given at scale.
func (s *_decayPlacementGeoLocationDistance) Decay(decay types.Float64) *_decayPlacementGeoLocationDistance {

	s.v.Decay = &decay

	return s
}

// If defined, the decay function will only compute the decay function for
// documents with a distance greater than the defined `offset`.
func (s *_decayPlacementGeoLocationDistance) Offset(distance string) *_decayPlacementGeoLocationDistance {

	s.v.Offset = &distance

	return s
}

// The point of origin used for calculating distance. Must be given as a number
// for numeric field, date for date fields and geo point for geo fields.
func (s *_decayPlacementGeoLocationDistance) Origin(geolocation types.GeoLocationVariant) *_decayPlacementGeoLocationDistance {

	s.v.Origin = *geolocation.GeoLocationCaster()

	return s
}

// Defines the distance from origin + offset at which the computed score will
// equal `decay` parameter.
func (s *_decayPlacementGeoLocationDistance) Scale(distance string) *_decayPlacementGeoLocationDistance {

	s.v.Scale = &distance

	return s
}

func (s *_decayPlacementGeoLocationDistance) DecayPlacementGeoLocationDistanceCaster() *types.DecayPlacementGeoLocationDistance {
	return s.v
}
