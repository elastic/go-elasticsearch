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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _mgetOperation struct {
	v *types.MgetOperation
}

func NewMgetOperation() *_mgetOperation {

	return &_mgetOperation{v: types.NewMgetOperation()}

}

func (s *_mgetOperation) Id_(id string) *_mgetOperation {

	s.v.Id_ = id

	return s
}

func (s *_mgetOperation) Index_(indexname string) *_mgetOperation {

	s.v.Index_ = &indexname

	return s
}

func (s *_mgetOperation) Routing(routing string) *_mgetOperation {

	s.v.Routing = &routing

	return s
}

func (s *_mgetOperation) Source_(sourceconfig types.SourceConfigVariant) *_mgetOperation {

	s.v.Source_ = *sourceconfig.SourceConfigCaster()

	return s
}

func (s *_mgetOperation) StoredFields(fields ...string) *_mgetOperation {

	s.v.StoredFields = fields

	return s
}

func (s *_mgetOperation) Version(versionnumber int64) *_mgetOperation {

	s.v.Version = &versionnumber

	return s
}

func (s *_mgetOperation) VersionType(versiontype versiontype.VersionType) *_mgetOperation {

	s.v.VersionType = &versiontype
	return s
}

func (s *_mgetOperation) MgetOperationCaster() *types.MgetOperation {
	return s.v
}
