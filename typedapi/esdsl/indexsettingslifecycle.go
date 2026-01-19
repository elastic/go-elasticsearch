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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexSettingsLifecycle struct {
	v *types.IndexSettingsLifecycle
}

func NewIndexSettingsLifecycle() *_indexSettingsLifecycle {

	return &_indexSettingsLifecycle{v: types.NewIndexSettingsLifecycle()}

}

func (s *_indexSettingsLifecycle) IndexingComplete(stringifiedboolean types.StringifiedbooleanVariant) *_indexSettingsLifecycle {

	s.v.IndexingComplete = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_indexSettingsLifecycle) Name(name string) *_indexSettingsLifecycle {

	s.v.Name = &name

	return s
}

func (s *_indexSettingsLifecycle) OriginationDate(originationdate int64) *_indexSettingsLifecycle {

	s.v.OriginationDate = &originationdate

	return s
}

func (s *_indexSettingsLifecycle) ParseOriginationDate(parseoriginationdate bool) *_indexSettingsLifecycle {

	s.v.ParseOriginationDate = &parseoriginationdate

	return s
}

func (s *_indexSettingsLifecycle) PreferIlm(preferilm string) *_indexSettingsLifecycle {

	s.v.PreferIlm = &preferilm

	return s
}

func (s *_indexSettingsLifecycle) RolloverAlias(rolloveralias string) *_indexSettingsLifecycle {

	s.v.RolloverAlias = &rolloveralias

	return s
}

func (s *_indexSettingsLifecycle) Step(step types.IndexSettingsLifecycleStepVariant) *_indexSettingsLifecycle {

	s.v.Step = step.IndexSettingsLifecycleStepCaster()

	return s
}

func (s *_indexSettingsLifecycle) IndexSettingsLifecycleCaster() *types.IndexSettingsLifecycle {
	return s.v
}
