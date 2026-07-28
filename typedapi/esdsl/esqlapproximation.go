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
// https://github.com/elastic/elasticsearch-specification/tree/7fd0bd13eaf28bd179fc906f57da09e852eb818e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _esqlApproximation struct {
	v types.EsqlApproximation
}

func NewEsqlApproximation() *_esqlApproximation {
	return &_esqlApproximation{v: nil}
}

func (u *_esqlApproximation) Bool(bool bool) *_esqlApproximation {

	u.v = &bool

	return u
}

func (u *_esqlApproximation) EsqlApproximationSettings(esqlapproximationsettings types.EsqlApproximationSettingsVariant) *_esqlApproximation {

	u.v = esqlapproximationsettings.EsqlApproximationSettingsCaster()

	return u
}

// Interface implementation for EsqlApproximationSettings in EsqlApproximation union
func (u *_esqlApproximationSettings) EsqlApproximationCaster() *types.EsqlApproximation {
	t := types.EsqlApproximation(u.v)
	return &t
}

func (u *_esqlApproximation) EsqlApproximationCaster() *types.EsqlApproximation {
	return &u.v
}
