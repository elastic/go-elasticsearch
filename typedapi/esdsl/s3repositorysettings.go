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

type _s3RepositorySettings struct {
	v *types.S3RepositorySettings
}

func NewS3RepositorySettings(bucket string) *_s3RepositorySettings {

	tmp := &_s3RepositorySettings{v: types.NewS3RepositorySettings()}

	tmp.Bucket(bucket)

	return tmp

}

func (s *_s3RepositorySettings) BasePath(basepath string) *_s3RepositorySettings {

	s.v.BasePath = &basepath

	return s
}

func (s *_s3RepositorySettings) Bucket(bucket string) *_s3RepositorySettings {

	s.v.Bucket = bucket

	return s
}

func (s *_s3RepositorySettings) BufferSize(bytesize types.ByteSizeVariant) *_s3RepositorySettings {

	s.v.BufferSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_s3RepositorySettings) CannedAcl(cannedacl string) *_s3RepositorySettings {

	s.v.CannedAcl = &cannedacl

	return s
}

func (s *_s3RepositorySettings) Client(client string) *_s3RepositorySettings {

	s.v.Client = &client

	return s
}

func (s *_s3RepositorySettings) DeleteObjectsMaxSize(deleteobjectsmaxsize int) *_s3RepositorySettings {

	s.v.DeleteObjectsMaxSize = &deleteobjectsmaxsize

	return s
}

func (s *_s3RepositorySettings) GetRegisterRetryDelay(duration types.DurationVariant) *_s3RepositorySettings {

	s.v.GetRegisterRetryDelay = *duration.DurationCaster()

	return s
}

func (s *_s3RepositorySettings) MaxMultipartParts(maxmultipartparts int) *_s3RepositorySettings {

	s.v.MaxMultipartParts = &maxmultipartparts

	return s
}

func (s *_s3RepositorySettings) MaxMultipartUploadCleanupSize(maxmultipartuploadcleanupsize int) *_s3RepositorySettings {

	s.v.MaxMultipartUploadCleanupSize = &maxmultipartuploadcleanupsize

	return s
}

func (s *_s3RepositorySettings) Readonly(readonly bool) *_s3RepositorySettings {

	s.v.Readonly = &readonly

	return s
}

func (s *_s3RepositorySettings) ServerSideEncryption(serversideencryption bool) *_s3RepositorySettings {

	s.v.ServerSideEncryption = &serversideencryption

	return s
}

func (s *_s3RepositorySettings) StorageClass(storageclass string) *_s3RepositorySettings {

	s.v.StorageClass = &storageclass

	return s
}

func (s *_s3RepositorySettings) ThrottledDeleteRetryDelayIncrement(duration types.DurationVariant) *_s3RepositorySettings {

	s.v.ThrottledDeleteRetryDelayIncrement = *duration.DurationCaster()

	return s
}

func (s *_s3RepositorySettings) ThrottledDeleteRetryMaximumDelay(duration types.DurationVariant) *_s3RepositorySettings {

	s.v.ThrottledDeleteRetryMaximumDelay = *duration.DurationCaster()

	return s
}

func (s *_s3RepositorySettings) ThrottledDeleteRetryMaximumNumberOfRetries(throttleddeleteretrymaximumnumberofretries int) *_s3RepositorySettings {

	s.v.ThrottledDeleteRetryMaximumNumberOfRetries = &throttleddeleteretrymaximumnumberofretries

	return s
}

func (s *_s3RepositorySettings) ChunkSize(bytesize types.ByteSizeVariant) *_s3RepositorySettings {

	s.v.ChunkSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_s3RepositorySettings) Compress(compress bool) *_s3RepositorySettings {

	s.v.Compress = &compress

	return s
}

func (s *_s3RepositorySettings) MaxRestoreBytesPerSec(bytesize types.ByteSizeVariant) *_s3RepositorySettings {

	s.v.MaxRestoreBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_s3RepositorySettings) MaxSnapshotBytesPerSec(bytesize types.ByteSizeVariant) *_s3RepositorySettings {

	s.v.MaxSnapshotBytesPerSec = *bytesize.ByteSizeCaster()

	return s
}

func (s *_s3RepositorySettings) S3RepositorySettingsCaster() *types.S3RepositorySettings {
	return s.v
}
