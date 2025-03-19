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

type _azureRepositorySettings struct {
	v *types.AzureRepositorySettings
}

func NewAzureRepositorySettings() *_azureRepositorySettings {

	return &_azureRepositorySettings{v: types.NewAzureRepositorySettings()}

}

// The path to the repository data within the container.
// It defaults to the root directory.
//
// NOTE: Don't set `base_path` when configuring a snapshot repository for
// Elastic Cloud Enterprise.
// Elastic Cloud Enterprise automatically generates the `base_path` for each
// deployment so that multiple deployments can share the same bucket.
func (s *_azureRepositorySettings) BasePath(basepath string) *_azureRepositorySettings {

	s.v.BasePath = &basepath

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
func (s *_azureRepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_azureRepositorySettings {

	s.v.ChunkSize = *bytesize.ByteSizeCaster()

	return s
}

// The name of the Azure repository client to use.
func (s *_azureRepositorySettings) Client(client string) *_azureRepositorySettings {

	s.v.Client = &client

	return s
}

// When set to `true`, metadata files are stored in compressed format.
// This setting doesn't affect index files that are already compressed by
// default.
func (s *_azureRepositorySettings) Compress(compress bool) *_azureRepositorySettings {

	s.v.Compress = &compress

	return s
}

// The Azure container.
func (s *_azureRepositorySettings) Container(container string) *_azureRepositorySettings {

	s.v.Container = &container

	return s
}

// The maxmimum batch size, between 1 and 256, used for `BlobBatch` requests.
// Defaults to 256 which is the maximum number supported by the Azure blob batch
// API.
func (s *_azureRepositorySettings) DeleteObjectsMaxSize(deleteobjectsmaxsize int) *_azureRepositorySettings {

	s.v.DeleteObjectsMaxSize = &deleteobjectsmaxsize

	return s
}

// Either `primary_only` or `secondary_only`.
// Note that if you set it to `secondary_only`, it will force `readonly` to
// `true`.
func (s *_azureRepositorySettings) LocationMode(locationmode string) *_azureRepositorySettings {

	s.v.LocationMode = &locationmode

	return s
}

// The maximum number of concurrent batch delete requests that will be submitted
// for any individual bulk delete with `BlobBatch`.
// Note that the effective number of concurrent deletes is further limited by
// the Azure client connection and event loop thread limits.
// Defaults to 10, minimum is 1, maximum is 100.
func (s *_azureRepositorySettings) MaxConcurrentBatchDeletes(maxconcurrentbatchdeletes int) *_azureRepositorySettings {

	s.v.MaxConcurrentBatchDeletes = &maxconcurrentbatchdeletes

	return s
}

// The maximum snapshot restore rate per node.
// It defaults to unlimited.
// Note that restores are also throttled through recovery settings.
func (s *_azureRepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_azureRepositorySettings {

	s.v.MaxRestoreBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

// The maximum snapshot creation rate per node.
// It defaults to 40mb per second.
// Note that if the recovery settings for managed services are set, then it
// defaults to unlimited, and the rate is additionally throttled through
// recovery settings.
func (s *_azureRepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_azureRepositorySettings {

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
// If `false`, the cluster can write to the repository and create snapshots in
// it.
//
// IMPORTANT: If you register the same snapshot repository with multiple
// clusters, only one cluster should have write access to the repository.
// Having multiple clusters write to the repository at the same time risks
// corrupting the contents of the repository.
func (s *_azureRepositorySettings) Readonly(readonly bool) *_azureRepositorySettings {

	s.v.Readonly = &readonly

	return s
}

func (s *_azureRepositorySettings) AzureRepositorySettingsCaster() *types.AzureRepositorySettings {
	return s.v
}
