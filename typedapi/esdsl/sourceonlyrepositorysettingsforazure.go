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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sourceOnlyRepositorySettingsForAzure struct {
	v *types.SourceOnlyRepositorySettingsForAzure
}

func NewSourceOnlyRepositorySettingsForAzure() *_sourceOnlyRepositorySettingsForAzure {

	return &_sourceOnlyRepositorySettingsForAzure{v: types.NewSourceOnlyRepositorySettingsForAzure()}

}

func (s *_sourceOnlyRepositorySettingsForAzure) BasePath(basepath string) *_sourceOnlyRepositorySettingsForAzure {

	s.v.BasePath = &basepath

	return s
}

func (s *_sourceOnlyRepositorySettingsForAzure) Client(client string) *_sourceOnlyRepositorySettingsForAzure {

	s.v.Client = &client

	return s
}

func (s *_sourceOnlyRepositorySettingsForAzure) Container(container string) *_sourceOnlyRepositorySettingsForAzure {

	s.v.Container = &container

	return s
}

func (s *_sourceOnlyRepositorySettingsForAzure) DeleteObjectsMaxSize(deleteobjectsmaxsize int) *_sourceOnlyRepositorySettingsForAzure {

	s.v.DeleteObjectsMaxSize = &deleteobjectsmaxsize

	return s
}

func (s *_sourceOnlyRepositorySettingsForAzure) LocationMode(locationmode string) *_sourceOnlyRepositorySettingsForAzure {

	s.v.LocationMode = &locationmode

	return s
}

func (s *_sourceOnlyRepositorySettingsForAzure) MaxConcurrentBatchDeletes(maxconcurrentbatchdeletes int) *_sourceOnlyRepositorySettingsForAzure {

	s.v.MaxConcurrentBatchDeletes = &maxconcurrentbatchdeletes

	return s
}

func (s *_sourceOnlyRepositorySettingsForAzure) Readonly(readonly bool) *_sourceOnlyRepositorySettingsForAzure {

	s.v.Readonly = &readonly

	return s
}

func (s *_sourceOnlyRepositorySettingsForAzure) SourceOnlyRepositorySettingsForAzureCaster() *types.SourceOnlyRepositorySettingsForAzure {
	return s.v
}
