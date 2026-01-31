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

type _sharedFileSystemRepository struct {
	v *types.SharedFileSystemRepository
}

func NewSharedFileSystemRepository(settings types.SharedFileSystemRepositorySettingsVariant) *_sharedFileSystemRepository {

	tmp := &_sharedFileSystemRepository{v: types.NewSharedFileSystemRepository()}

	tmp.Settings(settings)

	return tmp

}

func (s *_sharedFileSystemRepository) Settings(settings types.SharedFileSystemRepositorySettingsVariant) *_sharedFileSystemRepository {

	s.v.Settings = *settings.SharedFileSystemRepositorySettingsCaster()

	return s
}

func (s *_sharedFileSystemRepository) Uuid(uuid string) *_sharedFileSystemRepository {

	s.v.Uuid = &uuid

	return s
}

func (s *_sharedFileSystemRepository) SharedFileSystemRepositoryCaster() *types.SharedFileSystemRepository {
	return s.v
}
