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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type _geoDecayFunction struct {
	v *types.GeoDecayFunction
}

// Function that scores a document with a linear decay, depending on the
// distance of a numeric field value of the document from an origin.
func NewGeoDecayFunction() *_geoDecayFunction {

	return &_geoDecayFunction{v: types.NewGeoDecayFunction()}

}

func (s *_geoDecayFunction) DecayFunctionBaseGeoLocationDistance(decayfunctionbasegeolocationdistance map[string]types.DecayPlacementGeoLocationDistance) *_geoDecayFunction {

	s.v.DecayFunctionBaseGeoLocationDistance = decayfunctionbasegeolocationdistance
	return s
}

func (s *_geoDecayFunction) AddDecayFunctionBaseGeoLocationDistance(key string, value types.DecayPlacementGeoLocationDistanceVariant) *_geoDecayFunction {

	var tmp map[string]types.DecayPlacementGeoLocationDistance
	if s.v.DecayFunctionBaseGeoLocationDistance == nil {
		s.v.DecayFunctionBaseGeoLocationDistance = make(map[string]types.DecayPlacementGeoLocationDistance)
	} else {
		tmp = s.v.DecayFunctionBaseGeoLocationDistance
	}

	tmp[key] = *value.DecayPlacementGeoLocationDistanceCaster()

	s.v.DecayFunctionBaseGeoLocationDistance = tmp
	return s
}

func (s *_geoDecayFunction) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *_geoDecayFunction {

	s.v.MultiValueMode = &multivaluemode
	return s
}

func (s *_geoDecayFunction) FunctionScoreCaster() *types.FunctionScore {
	container := types.NewFunctionScore()

	container.Linear = s.v

	return container
}

func (s *_geoDecayFunction) GeoDecayFunctionCaster() *types.GeoDecayFunction {
	return s.v
}
