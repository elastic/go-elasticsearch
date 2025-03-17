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

type _s3Repository struct {
	v *types.S3Repository
}

func NewS3Repository(settings types.S3RepositorySettingsVariant) *_s3Repository {

	tmp := &_s3Repository{v: types.NewS3Repository()}

	tmp.Settings(settings)

	return tmp

}

// The repository settings.
//
// NOTE: In addition to the specified settings, you can also use all non-secure
// client settings in the repository settings.
// In this case, the client settings found in the repository settings will be
// merged with those of the named client used by the repository.
// Conflicts between client and repository settings are resolved by the
// repository settings taking precedence over client settings.
func (s *_s3Repository) Settings(settings types.S3RepositorySettingsVariant) *_s3Repository {

	s.v.Settings = *settings.S3RepositorySettingsCaster()

	return s
}

func (s *_s3Repository) Uuid(uuid string) *_s3Repository {

	s.v.Uuid = &uuid

	return s
}

func (s *_s3Repository) S3RepositoryCaster() *types.S3Repository {
	return s.v
}
