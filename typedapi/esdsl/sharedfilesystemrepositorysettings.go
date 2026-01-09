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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sharedFileSystemRepositorySettings struct {
	v *types.SharedFileSystemRepositorySettings
}

func NewSharedFileSystemRepositorySettings(location string) *_sharedFileSystemRepositorySettings {

	tmp := &_sharedFileSystemRepositorySettings{v: types.NewSharedFileSystemRepositorySettings()}

	tmp.Location(location)

	return tmp

}

func (s *_sharedFileSystemRepositorySettings) Location(location string) *_sharedFileSystemRepositorySettings {

	s.v.Location = location

	return s
}

func (s *_sharedFileSystemRepositorySettings) MaxNumberOfSnapshots(maxnumberofsnapshots int) *_sharedFileSystemRepositorySettings {

	s.v.MaxNumberOfSnapshots = &maxnumberofsnapshots

	return s
}

func (s *_sharedFileSystemRepositorySettings) Readonly(readonly bool) *_sharedFileSystemRepositorySettings {

	s.v.Readonly = &readonly

	return s
}

func (s *_sharedFileSystemRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_sharedFileSystemRepositorySettings {

	s.v.ChunkSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_sharedFileSystemRepositorySettings) Compress(compress bool) *_sharedFileSystemRepositorySettings {

	s.v.Compress = &compress

	return s
}

func (s *_sharedFileSystemRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_sharedFileSystemRepositorySettings {

	s.v.MaxRestoreBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_sharedFileSystemRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_sharedFileSystemRepositorySettings {

	s.v.MaxSnapshotBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_sharedFileSystemRepositorySettings) SharedFileSystemRepositorySettingsCaster() *types.SharedFileSystemRepositorySettings {
	return s.v
}
