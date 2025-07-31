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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _gcsRepositorySettings struct {
	v *types.GcsRepositorySettings
}

func NewGcsRepositorySettings(bucket string) *_gcsRepositorySettings {

	tmp := &_gcsRepositorySettings{v: types.NewGcsRepositorySettings()}

	tmp.Bucket(bucket)

	return tmp

}

func (s *_gcsRepositorySettings) ApplicationName(applicationname string) *_gcsRepositorySettings {

	s.v.ApplicationName = &applicationname

	return s
}

func (s *_gcsRepositorySettings) BasePath(basepath string) *_gcsRepositorySettings {

	s.v.BasePath = &basepath

	return s
}

func (s *_gcsRepositorySettings) Bucket(bucket string) *_gcsRepositorySettings {

	s.v.Bucket = bucket

	return s
}

func (s *_gcsRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_gcsRepositorySettings {

	s.v.ChunkSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_gcsRepositorySettings) Client(client string) *_gcsRepositorySettings {

	s.v.Client = &client

	return s
}

func (s *_gcsRepositorySettings) Compress(compress bool) *_gcsRepositorySettings {

	s.v.Compress = &compress

	return s
}

func (s *_gcsRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_gcsRepositorySettings {

	s.v.MaxRestoreBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_gcsRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_gcsRepositorySettings {

	s.v.MaxSnapshotBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_gcsRepositorySettings) Readonly(readonly bool) *_gcsRepositorySettings {

	s.v.Readonly = &readonly

	return s
}

func (s *_gcsRepositorySettings) GcsRepositorySettingsCaster() *types.GcsRepositorySettings {
	return s.v
}
