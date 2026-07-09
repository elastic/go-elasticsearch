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
// https://github.com/elastic/elasticsearch-specification/tree/c0021097996e8ff7ae5fe8995f26b148dc329bae

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sourceOnlyRepositorySettingsForGcs struct {
	v *types.SourceOnlyRepositorySettingsForGcs
}

func NewSourceOnlyRepositorySettingsForGcs(bucket string) *_sourceOnlyRepositorySettingsForGcs {

	tmp := &_sourceOnlyRepositorySettingsForGcs{v: types.NewSourceOnlyRepositorySettingsForGcs()}

	tmp.Bucket(bucket)

	return tmp

}

func (s *_sourceOnlyRepositorySettingsForGcs) ApplicationName(applicationname string) *_sourceOnlyRepositorySettingsForGcs {

	s.v.ApplicationName = &applicationname

	return s
}

func (s *_sourceOnlyRepositorySettingsForGcs) BasePath(basepath string) *_sourceOnlyRepositorySettingsForGcs {

	s.v.BasePath = &basepath

	return s
}

func (s *_sourceOnlyRepositorySettingsForGcs) Bucket(bucket string) *_sourceOnlyRepositorySettingsForGcs {

	s.v.Bucket = bucket

	return s
}

func (s *_sourceOnlyRepositorySettingsForGcs) Client(client string) *_sourceOnlyRepositorySettingsForGcs {

	s.v.Client = &client

	return s
}

func (s *_sourceOnlyRepositorySettingsForGcs) Readonly(readonly bool) *_sourceOnlyRepositorySettingsForGcs {

	s.v.Readonly = &readonly

	return s
}

func (s *_sourceOnlyRepositorySettingsForGcs) SourceOnlyRepositorySettingsForGcsCaster() *types.SourceOnlyRepositorySettingsForGcs {
	return s.v
}
