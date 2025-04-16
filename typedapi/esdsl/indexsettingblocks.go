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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexSettingBlocks struct {
	v *types.IndexSettingBlocks
}

func NewIndexSettingBlocks() *_indexSettingBlocks {

	return &_indexSettingBlocks{v: types.NewIndexSettingBlocks()}

}

func (s *_indexSettingBlocks) Metadata(stringifiedboolean types.StringifiedbooleanVariant) *_indexSettingBlocks {

	s.v.Metadata = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_indexSettingBlocks) Read(stringifiedboolean types.StringifiedbooleanVariant) *_indexSettingBlocks {

	s.v.Read = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_indexSettingBlocks) ReadOnly(stringifiedboolean types.StringifiedbooleanVariant) *_indexSettingBlocks {

	s.v.ReadOnly = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_indexSettingBlocks) ReadOnlyAllowDelete(stringifiedboolean types.StringifiedbooleanVariant) *_indexSettingBlocks {

	s.v.ReadOnlyAllowDelete = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_indexSettingBlocks) Write(stringifiedboolean types.StringifiedbooleanVariant) *_indexSettingBlocks {

	s.v.Write = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_indexSettingBlocks) IndexSettingBlocksCaster() *types.IndexSettingBlocks {
	return s.v
}
