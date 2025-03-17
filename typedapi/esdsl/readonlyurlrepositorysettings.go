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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _readOnlyUrlRepositorySettings struct {
	v *types.ReadOnlyUrlRepositorySettings
}

func NewReadOnlyUrlRepositorySettings(url string) *_readOnlyUrlRepositorySettings {

	tmp := &_readOnlyUrlRepositorySettings{v: types.NewReadOnlyUrlRepositorySettings()}

	tmp.Url(url)

	return tmp

}

// Big files can be broken down into multiple smaller blobs in the blob store
// during snapshotting.
// It is not recommended to change this value from its default unless there is
// an explicit reason for limiting the size of blobs in the repository.
// Setting a value lower than the default can result in an increased number of
// API calls to the blob store during snapshot create and restore operations
// compared to using the default value and thus make both operations slower and
// more costly.
// Specify the chunk size as a byte unit, for example: `10MB`, `5KB`, 500B.
// The default varies by repository type.
func (s *_readOnlyUrlRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_readOnlyUrlRepositorySettings {

	s.v.ChunkSize = *bytesize.ByteSizeCaster()

	return s
}

// When set to `true`, metadata files are stored in compressed format.
// This setting doesn't affect index files that are already compressed by
// default.
func (s *_readOnlyUrlRepositorySettings) Compress(compress bool) *_readOnlyUrlRepositorySettings {

	s.v.Compress = &compress

	return s
}

// The maximum number of retries for HTTP and HTTPS URLs.
func (s *_readOnlyUrlRepositorySettings) HttpMaxRetries(httpmaxretries int) *_readOnlyUrlRepositorySettings {

	s.v.HttpMaxRetries = &httpmaxretries

	return s
}

// The maximum wait time for data transfers over a connection.
func (s *_readOnlyUrlRepositorySettings) HttpSocketTimeout(duration types.DurationVariant) *_readOnlyUrlRepositorySettings {

	s.v.HttpSocketTimeout = *duration.DurationCaster()

	return s
}

// The maximum number of snapshots the repository can contain.
// The default is `Integer.MAX_VALUE`, which is 2^31-1 or `2147483647`.
func (s *_readOnlyUrlRepositorySettings) MaxNumberOfSnapshots(maxnumberofsnapshots int) *_readOnlyUrlRepositorySettings {

	s.v.MaxNumberOfSnapshots = &maxnumberofsnapshots

	return s
}

// The maximum snapshot restore rate per node.
// It defaults to unlimited.
// Note that restores are also throttled through recovery settings.
func (s *_readOnlyUrlRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_readOnlyUrlRepositorySettings {

	s.v.MaxRestoreBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

// The maximum snapshot creation rate per node.
// It defaults to 40mb per second.
// Note that if the recovery settings for managed services are set, then it
// defaults to unlimited, and the rate is additionally throttled through
// recovery settings.
func (s *_readOnlyUrlRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_readOnlyUrlRepositorySettings {

	s.v.MaxSnapshotBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

// The URL location of the root of the shared filesystem repository.
// The following protocols are supported:
//
// * `file`
// * `ftp`
// * `http`
// * `https`
// * `jar`
//
// URLs using the HTTP, HTTPS, or FTP protocols must be explicitly allowed with
// the `repositories.url.allowed_urls` cluster setting.
// This setting supports wildcards in the place of a host, path, query, or
// fragment in the URL.
//
// URLs using the file protocol must point to the location of a shared
// filesystem accessible to all master and data nodes in the cluster.
// This location must be registered in the `path.repo` setting.
// You don't need to register URLs using the FTP, HTTP, HTTPS, or JAR protocols
// in the `path.repo` setting.
func (s *_readOnlyUrlRepositorySettings) Url(url string) *_readOnlyUrlRepositorySettings {

	s.v.Url = url

	return s
}

func (s *_readOnlyUrlRepositorySettings) ReadOnlyUrlRepositorySettingsCaster() *types.ReadOnlyUrlRepositorySettings {
	return s.v
}
