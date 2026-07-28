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
// https://github.com/elastic/elasticsearch-specification/tree/7560c979602e6941815872bdaec801200bc7ec4e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _esqlQuerySettings struct {
	v *types.EsqlQuerySettings
}

func NewEsqlQuerySettings() *_esqlQuerySettings {

	return &_esqlQuerySettings{v: types.NewEsqlQuerySettings()}

}

func (s *_esqlQuerySettings) Approximation(esqlapproximation types.EsqlApproximationVariant) *_esqlQuerySettings {

	s.v.Approximation = *esqlapproximation.EsqlApproximationCaster()

	return s
}

func (s *_esqlQuerySettings) ColumnMetadata(stringifiedboolean types.StringifiedbooleanVariant) *_esqlQuerySettings {

	s.v.ColumnMetadata = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_esqlQuerySettings) ProjectRouting(projectrouting string) *_esqlQuerySettings {

	s.v.ProjectRouting = &projectrouting

	return s
}

func (s *_esqlQuerySettings) TimeZone(timezone string) *_esqlQuerySettings {

	s.v.TimeZone = &timezone

	return s
}

func (s *_esqlQuerySettings) EsqlQuerySettingsCaster() *types.EsqlQuerySettings {
	return s.v
}
