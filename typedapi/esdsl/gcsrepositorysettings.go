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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _gcsRepositorySettings struct {
	v *types.GcsRepositorySettings
}

func NewGcsRepositorySettings(bucket string) *_gcsRepositorySettings {

	tmp := &_gcsRepositorySettings{v: types.NewGcsRepositorySettings()}

	tmp.Bucket(bucket)

	return tmp

}

// The name used by the client when it uses the Google Cloud Storage service.
func (s *_gcsRepositorySettings) ApplicationName(applicationname string) *_gcsRepositorySettings {

	s.v.ApplicationName = &applicationname

	return s
}

// The path to the repository data within the bucket.
// It defaults to the root of the bucket.
//
// NOTE: Don't set `base_path` when configuring a snapshot repository for
// Elastic Cloud Enterprise.
// Elastic Cloud Enterprise automatically generates the `base_path` for each
// deployment so that multiple deployments can share the same bucket.
func (s *_gcsRepositorySettings) BasePath(basepath string) *_gcsRepositorySettings {

	s.v.BasePath = &basepath

	return s
}

// The name of the bucket to be used for snapshots.
func (s *_gcsRepositorySettings) Bucket(bucket string) *_gcsRepositorySettings {

	s.v.Bucket = bucket

	return s
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
func (s *_gcsRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_gcsRepositorySettings {

	s.v.ChunkSize = *bytesize.ByteSizeCaster()

	return s
}

// The name of the client to use to connect to Google Cloud Storage.
func (s *_gcsRepositorySettings) Client(client string) *_gcsRepositorySettings {

	s.v.Client = &client

	return s
}

// When set to `true`, metadata files are stored in compressed format.
// This setting doesn't affect index files that are already compressed by
// default.
func (s *_gcsRepositorySettings) Compress(compress bool) *_gcsRepositorySettings {

	s.v.Compress = &compress

	return s
}

// The maximum snapshot restore rate per node.
// It defaults to unlimited.
// Note that restores are also throttled through recovery settings.
func (s *_gcsRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_gcsRepositorySettings {

	s.v.MaxRestoreBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

// The maximum snapshot creation rate per node.
// It defaults to 40mb per second.
// Note that if the recovery settings for managed services are set, then it
// defaults to unlimited, and the rate is additionally throttled through
// recovery settings.
func (s *_gcsRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_gcsRepositorySettings {

	s.v.MaxSnapshotBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

// If `true`, the repository is read-only.
// The cluster can retrieve and restore snapshots from the repository but not
// write to the repository or create snapshots in it.
//
// Only a cluster with write access can create snapshots in the repository.
// All other clusters connected to the repository should have the `readonly`
// parameter set to `true`.
//
// If `false`, the cluster can write to the repository and create snapshots in
// it.
//
// IMPORTANT: If you register the same snapshot repository with multiple
// clusters, only one cluster should have write access to the repository.
// Having multiple clusters write to the repository at the same time risks
// corrupting the contents of the repository.
func (s *_gcsRepositorySettings) Readonly(readonly bool) *_gcsRepositorySettings {

	s.v.Readonly = &readonly

	return s
}

func (s *_gcsRepositorySettings) GcsRepositorySettingsCaster() *types.GcsRepositorySettings {
	return s.v
}
