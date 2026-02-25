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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _readOnlyUrlRepositorySettings struct {
	v *types.ReadOnlyUrlRepositorySettings
}

func NewReadOnlyUrlRepositorySettings(url string) *_readOnlyUrlRepositorySettings {

	tmp := &_readOnlyUrlRepositorySettings{v: types.NewReadOnlyUrlRepositorySettings()}

	tmp.Url(url)

	return tmp

}

func (s *_readOnlyUrlRepositorySettings) HttpMaxRetries(httpmaxretries int) *_readOnlyUrlRepositorySettings {

	s.v.HttpMaxRetries = &httpmaxretries

	return s
}

func (s *_readOnlyUrlRepositorySettings) HttpSocketTimeout(duration types.DurationVariant) *_readOnlyUrlRepositorySettings {

	s.v.HttpSocketTimeout = *duration.DurationCaster()

	return s
}

func (s *_readOnlyUrlRepositorySettings) MaxNumberOfSnapshots(maxnumberofsnapshots int) *_readOnlyUrlRepositorySettings {

	s.v.MaxNumberOfSnapshots = &maxnumberofsnapshots

	return s
}

func (s *_readOnlyUrlRepositorySettings) Url(url string) *_readOnlyUrlRepositorySettings {

	s.v.Url = url

	return s
}

func (s *_readOnlyUrlRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_readOnlyUrlRepositorySettings {

	s.v.ChunkSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_readOnlyUrlRepositorySettings) Compress(compress bool) *_readOnlyUrlRepositorySettings {

	s.v.Compress = &compress

	return s
}

func (s *_readOnlyUrlRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_readOnlyUrlRepositorySettings {

	s.v.MaxRestoreBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_readOnlyUrlRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_readOnlyUrlRepositorySettings {

	s.v.MaxSnapshotBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_readOnlyUrlRepositorySettings) ReadOnlyUrlRepositorySettingsCaster() *types.ReadOnlyUrlRepositorySettings {
	return s.v
}
