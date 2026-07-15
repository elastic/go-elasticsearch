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

// This is provide all the types that are part of the union.
type _sourceOnlyRepositorySettings struct {
	v types.SourceOnlyRepositorySettings
}

func NewSourceOnlyRepositorySettings() *_sourceOnlyRepositorySettings {
	return &_sourceOnlyRepositorySettings{v: nil}
}

func (u *_sourceOnlyRepositorySettings) SourceOnlyRepositorySettingsForSharedFileSystem(sourceonlyrepositorysettingsforsharedfilesystem types.SourceOnlyRepositorySettingsForSharedFileSystemVariant) *_sourceOnlyRepositorySettings {

	u.v = sourceonlyrepositorysettingsforsharedfilesystem.SourceOnlyRepositorySettingsForSharedFileSystemCaster()

	return u
}

// Interface implementation for SourceOnlyRepositorySettingsForSharedFileSystem in SourceOnlyRepositorySettings union
func (u *_sourceOnlyRepositorySettingsForSharedFileSystem) SourceOnlyRepositorySettingsCaster() *types.SourceOnlyRepositorySettings {
	t := types.SourceOnlyRepositorySettings(u.v)
	return &t
}

func (u *_sourceOnlyRepositorySettings) SourceOnlyRepositorySettingsForReadOnlyUrl(sourceonlyrepositorysettingsforreadonlyurl types.SourceOnlyRepositorySettingsForReadOnlyUrlVariant) *_sourceOnlyRepositorySettings {

	u.v = sourceonlyrepositorysettingsforreadonlyurl.SourceOnlyRepositorySettingsForReadOnlyUrlCaster()

	return u
}

// Interface implementation for SourceOnlyRepositorySettingsForReadOnlyUrl in SourceOnlyRepositorySettings union
func (u *_sourceOnlyRepositorySettingsForReadOnlyUrl) SourceOnlyRepositorySettingsCaster() *types.SourceOnlyRepositorySettings {
	t := types.SourceOnlyRepositorySettings(u.v)
	return &t
}

func (u *_sourceOnlyRepositorySettings) SourceOnlyRepositorySettingsForAzure(sourceonlyrepositorysettingsforazure types.SourceOnlyRepositorySettingsForAzureVariant) *_sourceOnlyRepositorySettings {

	u.v = sourceonlyrepositorysettingsforazure.SourceOnlyRepositorySettingsForAzureCaster()

	return u
}

// Interface implementation for SourceOnlyRepositorySettingsForAzure in SourceOnlyRepositorySettings union
func (u *_sourceOnlyRepositorySettingsForAzure) SourceOnlyRepositorySettingsCaster() *types.SourceOnlyRepositorySettings {
	t := types.SourceOnlyRepositorySettings(u.v)
	return &t
}

func (u *_sourceOnlyRepositorySettings) SourceOnlyRepositorySettingsForGcs(sourceonlyrepositorysettingsforgcs types.SourceOnlyRepositorySettingsForGcsVariant) *_sourceOnlyRepositorySettings {

	u.v = sourceonlyrepositorysettingsforgcs.SourceOnlyRepositorySettingsForGcsCaster()

	return u
}

// Interface implementation for SourceOnlyRepositorySettingsForGcs in SourceOnlyRepositorySettings union
func (u *_sourceOnlyRepositorySettingsForGcs) SourceOnlyRepositorySettingsCaster() *types.SourceOnlyRepositorySettings {
	t := types.SourceOnlyRepositorySettings(u.v)
	return &t
}

func (u *_sourceOnlyRepositorySettings) SourceOnlyRepositorySettingsForS3(sourceonlyrepositorysettingsfors3 types.SourceOnlyRepositorySettingsForS3Variant) *_sourceOnlyRepositorySettings {

	u.v = sourceonlyrepositorysettingsfors3.SourceOnlyRepositorySettingsForS3Caster()

	return u
}

// Interface implementation for SourceOnlyRepositorySettingsForS3 in SourceOnlyRepositorySettings union
func (u *_sourceOnlyRepositorySettingsForS3) SourceOnlyRepositorySettingsCaster() *types.SourceOnlyRepositorySettings {
	t := types.SourceOnlyRepositorySettings(u.v)
	return &t
}

func (u *_sourceOnlyRepositorySettings) SourceOnlyRepositorySettingsCaster() *types.SourceOnlyRepositorySettings {
	return &u.v
}
