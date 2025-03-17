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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

type _updateOperation struct {
	v *types.UpdateOperation
}

// Perform a partial document update.
// The following line must contain the partial document and update options.
func NewUpdateOperation() *_updateOperation {

	return &_updateOperation{v: types.NewUpdateOperation()}

}

// The document ID.
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

// The name of the index or index alias to perform the action on.
func (s *_updateOperation) Index_(indexname string) *_updateOperation {

	s.v.Index_ = &indexname

	return s
}

// If `true`, the request's actions must target an index alias.
func (s *_updateOperation) RequireAlias(requirealias bool) *_updateOperation {

	s.v.RequireAlias = &requirealias

	return s
}

// The number of times an update should be retried in the case of a version
// conflict.
func (s *_updateOperation) RetryOnConflict(retryonconflict int) *_updateOperation {

	s.v.RetryOnConflict = &retryonconflict

	return s
}

// A custom value used to route operations to a specific shard.
func (s *_updateOperation) Routing(routing string) *_updateOperation {

	s.v.Routing = &routing

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
