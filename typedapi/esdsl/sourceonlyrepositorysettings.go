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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _sourceOnlyRepositorySettings struct {
	v *types.SourceOnlyRepositorySettings
}

func NewSourceOnlyRepositorySettings() *_sourceOnlyRepositorySettings {

	return &_sourceOnlyRepositorySettings{v: types.NewSourceOnlyRepositorySettings()}

}

func (s *_sourceOnlyRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_sourceOnlyRepositorySettings {

	s.v.ChunkSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_sourceOnlyRepositorySettings) Compress(compress bool) *_sourceOnlyRepositorySettings {

	s.v.Compress = &compress

	return s
}

func (s *_sourceOnlyRepositorySettings) DelegateType(delegatetype string) *_sourceOnlyRepositorySettings {

	s.v.DelegateType = &delegatetype

	return s
}

func (s *_sourceOnlyRepositorySettings) MaxNumberOfSnapshots(maxnumberofsnapshots int) *_sourceOnlyRepositorySettings {

	s.v.MaxNumberOfSnapshots = &maxnumberofsnapshots

	return s
}

func (s *_sourceOnlyRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_sourceOnlyRepositorySettings {

	s.v.MaxRestoreBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_sourceOnlyRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_sourceOnlyRepositorySettings {

	s.v.MaxSnapshotBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_sourceOnlyRepositorySettings) ReadOnly(readonly bool) *_sourceOnlyRepositorySettings {

	s.v.ReadOnly = &readonly

	return s
}

func (s *_sourceOnlyRepositorySettings) SourceOnlyRepositorySettingsCaster() *types.SourceOnlyRepositorySettings {
	return s.v
}
