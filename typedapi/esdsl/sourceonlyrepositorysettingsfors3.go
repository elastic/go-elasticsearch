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

type _sourceOnlyRepositorySettingsForS3 struct {
	v *types.SourceOnlyRepositorySettingsForS3
}

func NewSourceOnlyRepositorySettingsForS3(bucket string) *_sourceOnlyRepositorySettingsForS3 {

	tmp := &_sourceOnlyRepositorySettingsForS3{v: types.NewSourceOnlyRepositorySettingsForS3()}

	tmp.Bucket(bucket)

	return tmp

}

func (s *_sourceOnlyRepositorySettingsForS3) BasePath(basepath string) *_sourceOnlyRepositorySettingsForS3 {

	s.v.BasePath = &basepath

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) Bucket(bucket string) *_sourceOnlyRepositorySettingsForS3 {

	s.v.Bucket = bucket

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) BufferSize(bytesize types.ByteSizeVariant) *_sourceOnlyRepositorySettingsForS3 {

	s.v.BufferSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) CannedAcl(cannedacl string) *_sourceOnlyRepositorySettingsForS3 {

	s.v.CannedAcl = &cannedacl

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) Client(client string) *_sourceOnlyRepositorySettingsForS3 {

	s.v.Client = &client

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) DeleteObjectsMaxSize(deleteobjectsmaxsize int) *_sourceOnlyRepositorySettingsForS3 {

	s.v.DeleteObjectsMaxSize = &deleteobjectsmaxsize

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) GetRegisterRetryDelay(duration types.DurationVariant) *_sourceOnlyRepositorySettingsForS3 {

	s.v.GetRegisterRetryDelay = *duration.DurationCaster()

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) MaxMultipartParts(maxmultipartparts int) *_sourceOnlyRepositorySettingsForS3 {

	s.v.MaxMultipartParts = &maxmultipartparts

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) MaxMultipartUploadCleanupSize(maxmultipartuploadcleanupsize int) *_sourceOnlyRepositorySettingsForS3 {

	s.v.MaxMultipartUploadCleanupSize = &maxmultipartuploadcleanupsize

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) Readonly(readonly bool) *_sourceOnlyRepositorySettingsForS3 {

	s.v.Readonly = &readonly

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) ServerSideEncryption(serversideencryption bool) *_sourceOnlyRepositorySettingsForS3 {

	s.v.ServerSideEncryption = &serversideencryption

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) StorageClass(storageclass string) *_sourceOnlyRepositorySettingsForS3 {

	s.v.StorageClass = &storageclass

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) ThrottledDeleteRetryDelayIncrement(duration types.DurationVariant) *_sourceOnlyRepositorySettingsForS3 {

	s.v.ThrottledDeleteRetryDelayIncrement = *duration.DurationCaster()

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) ThrottledDeleteRetryMaximumDelay(duration types.DurationVariant) *_sourceOnlyRepositorySettingsForS3 {

	s.v.ThrottledDeleteRetryMaximumDelay = *duration.DurationCaster()

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) ThrottledDeleteRetryMaximumNumberOfRetries(throttleddeleteretrymaximumnumberofretries int) *_sourceOnlyRepositorySettingsForS3 {

	s.v.ThrottledDeleteRetryMaximumNumberOfRetries = &throttleddeleteretrymaximumnumberofretries

	return s
}

func (s *_sourceOnlyRepositorySettingsForS3) SourceOnlyRepositorySettingsForS3Caster() *types.SourceOnlyRepositorySettingsForS3 {
	return s.v
}
