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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _updateOperation struct {
	v *types.UpdateOperation
}

// Perform a partial document update.
// The following line must contain the partial document and update options.
func NewUpdateOperation() *_updateOperation {

	return &_updateOperation{v: types.NewUpdateOperation()}

}

func (s *_updateOperation) Id_(id string) *_updateOperation {

	s.v.Id_ = &id

	return s
}

func (s *_updateOperation) IfPrimaryTerm(ifprimaryterm int64) *_updateOperation {

	s.v.IfPrimaryTerm = &ifprimaryterm

	return s
}

func (s *_updateOperation) IfSeqNo(sequencenumber int64) *_updateOperation {

	s.v.IfSeqNo = &sequencenumber

	return s
}

func (s *_updateOperation) Index_(indexname string) *_updateOperation {

	s.v.Index_ = &indexname

	return s
}

func (s *_updateOperation) RequireAlias(requirealias bool) *_updateOperation {

	s.v.RequireAlias = &requirealias

	return s
}

func (s *_updateOperation) RetryOnConflict(retryonconflict int) *_updateOperation {

	s.v.RetryOnConflict = &retryonconflict

	return s
}

func (s *_updateOperation) Routing(routings ...string) *_updateOperation {

	s.v.Routing = routings

	return s
}

func (s *_updateOperation) Version(versionnumber int64) *_updateOperation {

	s.v.Version = &versionnumber

	return s
}

func (s *_updateOperation) VersionType(versiontype versiontype.VersionType) *_updateOperation {

	s.v.VersionType = &versiontype
	return s
}

func (s *_updateOperation) OperationContainerCaster() *types.OperationContainer {
	container := types.NewOperationContainer()

	container.Update = s.v

	return container
}

func (s *_updateOperation) UpdateOperationCaster() *types.UpdateOperation {
	return s.v
}
