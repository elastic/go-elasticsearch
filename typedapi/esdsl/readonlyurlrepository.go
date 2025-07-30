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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _readOnlyUrlRepository struct {
	v *types.ReadOnlyUrlRepository
}

func NewReadOnlyUrlRepository(settings types.ReadOnlyUrlRepositorySettingsVariant) *_readOnlyUrlRepository {

	tmp := &_readOnlyUrlRepository{v: types.NewReadOnlyUrlRepository()}

	tmp.Settings(settings)

	return tmp

}

func (s *_readOnlyUrlRepository) Settings(settings types.ReadOnlyUrlRepositorySettingsVariant) *_readOnlyUrlRepository {

	s.v.Settings = *settings.ReadOnlyUrlRepositorySettingsCaster()

	return s
}

func (s *_readOnlyUrlRepository) Uuid(uuid string) *_readOnlyUrlRepository {

	s.v.Uuid = &uuid

	return s
}

func (s *_readOnlyUrlRepository) ReadOnlyUrlRepositoryCaster() *types.ReadOnlyUrlRepository {
	return s.v
}
