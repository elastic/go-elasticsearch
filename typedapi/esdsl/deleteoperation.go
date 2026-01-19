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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _deleteOperation struct {
	v *types.DeleteOperation
}

// Remove the specified document from the index.
func NewDeleteOperation() *_deleteOperation {

	return &_deleteOperation{v: types.NewDeleteOperation()}

}

func (s *_deleteOperation) Id_(id string) *_deleteOperation {

	s.v.Id_ = &id

	return s
}

func (s *_deleteOperation) IfPrimaryTerm(ifprimaryterm int64) *_deleteOperation {

	s.v.IfPrimaryTerm = &ifprimaryterm

	return s
}

func (s *_deleteOperation) IfSeqNo(sequencenumber int64) *_deleteOperation {

	s.v.IfSeqNo = &sequencenumber

	return s
}

func (s *_deleteOperation) Index_(indexname string) *_deleteOperation {

	s.v.Index_ = &indexname

	return s
}

func (s *_deleteOperation) Routing(routings ...string) *_deleteOperation {

	s.v.Routing = routings

	return s
}

func (s *_deleteOperation) Version(versionnumber int64) *_deleteOperation {

	s.v.Version = &versionnumber

	return s
}

func (s *_deleteOperation) VersionType(versiontype versiontype.VersionType) *_deleteOperation {

	s.v.VersionType = &versiontype
	return s
}

func (s *_deleteOperation) OperationContainerCaster() *types.OperationContainer {
	container := types.NewOperationContainer()

	container.Delete = s.v

	return container
}

func (s *_deleteOperation) DeleteOperationCaster() *types.DeleteOperation {
	return s.v
}
