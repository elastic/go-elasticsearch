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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _azureRepositorySettings struct {
	v *types.AzureRepositorySettings
}

func NewAzureRepositorySettings() *_azureRepositorySettings {

	return &_azureRepositorySettings{v: types.NewAzureRepositorySettings()}

}

func (s *_azureRepositorySettings) BasePath(basepath string) *_azureRepositorySettings {

	s.v.BasePath = &basepath

	return s
}

func (s *_azureRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_azureRepositorySettings {

	s.v.ChunkSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_azureRepositorySettings) Client(client string) *_azureRepositorySettings {

	s.v.Client = &client

	return s
}

func (s *_azureRepositorySettings) Compress(compress bool) *_azureRepositorySettings {

	s.v.Compress = &compress

	return s
}

func (s *_azureRepositorySettings) Container(container string) *_azureRepositorySettings {

	s.v.Container = &container

	return s
}

func (s *_azureRepositorySettings) DeleteObjectsMaxSize(deleteobjectsmaxsize int) *_azureRepositorySettings {

	s.v.DeleteObjectsMaxSize = &deleteobjectsmaxsize

	return s
}

func (s *_azureRepositorySettings) LocationMode(locationmode string) *_azureRepositorySettings {

	s.v.LocationMode = &locationmode

	return s
}

func (s *_azureRepositorySettings) MaxConcurrentBatchDeletes(maxconcurrentbatchdeletes int) *_azureRepositorySettings {

	s.v.MaxConcurrentBatchDeletes = &maxconcurrentbatchdeletes

	return s
}

func (s *_azureRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_azureRepositorySettings {

	s.v.MaxRestoreBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_azureRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_azureRepositorySettings {

	s.v.MaxSnapshotBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_azureRepositorySettings) Readonly(readonly bool) *_azureRepositorySettings {

	s.v.Readonly = &readonly

	return s
}

func (s *_azureRepositorySettings) AzureRepositorySettingsCaster() *types.AzureRepositorySettings {
	return s.v
}
