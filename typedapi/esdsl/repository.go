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

// This is provide all the types that are part of the union.
type _repository struct {
	v types.Repository
}

func NewRepository() *_repository {
	return &_repository{v: nil}
}

func (u *_repository) AzureRepository(azurerepository types.AzureRepositoryVariant) *_repository {

	u.v = &azurerepository

	return u
}

// Interface implementation for AzureRepository in Repository union
func (u *_azureRepository) RepositoryCaster() *types.Repository {
	t := types.Repository(u.v)
	return &t
}

func (u *_repository) GcsRepository(gcsrepository types.GcsRepositoryVariant) *_repository {

	u.v = &gcsrepository

	return u
}

// Interface implementation for GcsRepository in Repository union
func (u *_gcsRepository) RepositoryCaster() *types.Repository {
	t := types.Repository(u.v)
	return &t
}

func (u *_repository) S3Repository(s3repository types.S3RepositoryVariant) *_repository {

	u.v = &s3repository

	return u
}

// Interface implementation for S3Repository in Repository union
func (u *_s3Repository) RepositoryCaster() *types.Repository {
	t := types.Repository(u.v)
	return &t
}

func (u *_repository) SharedFileSystemRepository(sharedfilesystemrepository types.SharedFileSystemRepositoryVariant) *_repository {

	u.v = &sharedfilesystemrepository

	return u
}

// Interface implementation for SharedFileSystemRepository in Repository union
func (u *_sharedFileSystemRepository) RepositoryCaster() *types.Repository {
	t := types.Repository(u.v)
	return &t
}

func (u *_repository) ReadOnlyUrlRepository(readonlyurlrepository types.ReadOnlyUrlRepositoryVariant) *_repository {

	u.v = &readonlyurlrepository

	return u
}

// Interface implementation for ReadOnlyUrlRepository in Repository union
func (u *_readOnlyUrlRepository) RepositoryCaster() *types.Repository {
	t := types.Repository(u.v)
	return &t
}

func (u *_repository) SourceOnlyRepository(sourceonlyrepository types.SourceOnlyRepositoryVariant) *_repository {

	u.v = &sourceonlyrepository

	return u
}

// Interface implementation for SourceOnlyRepository in Repository union
func (u *_sourceOnlyRepository) RepositoryCaster() *types.Repository {
	t := types.Repository(u.v)
	return &t
}

func (u *_repository) RepositoryCaster() *types.Repository {
	return &u.v
}
