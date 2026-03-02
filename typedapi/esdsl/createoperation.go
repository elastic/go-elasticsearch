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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _createOperation struct {
	v *types.CreateOperation
}

// Index the specified document if it does not already exist.
// The following line must contain the source data to be indexed.
func NewCreateOperation() *_createOperation {

	return &_createOperation{v: types.NewCreateOperation()}

}

func (s *_createOperation) DynamicTemplates(dynamictemplates map[string]string) *_createOperation {

	s.v.DynamicTemplates = dynamictemplates
	return s
}

func (s *_createOperation) AddDynamicTemplate(key string, value string) *_createOperation {

	var tmp map[string]string
	if s.v.DynamicTemplates == nil {
		s.v.DynamicTemplates = make(map[string]string)
	} else {
		tmp = s.v.DynamicTemplates
	}

	tmp[key] = value

	s.v.DynamicTemplates = tmp
	return s
}

func (s *_createOperation) Id_(id string) *_createOperation {

	s.v.Id_ = &id

	return s
}

func (s *_createOperation) IfPrimaryTerm(ifprimaryterm int64) *_createOperation {

	s.v.IfPrimaryTerm = &ifprimaryterm

	return s
}

func (s *_createOperation) IfSeqNo(sequencenumber int64) *_createOperation {

	s.v.IfSeqNo = &sequencenumber

	return s
}

func (s *_createOperation) Index_(indexname string) *_createOperation {

	s.v.Index_ = &indexname

	return s
}

func (s *_createOperation) Pipeline(pipeline string) *_createOperation {

	s.v.Pipeline = &pipeline

	return s
}

func (s *_createOperation) RequireAlias(requirealias bool) *_createOperation {

	s.v.RequireAlias = &requirealias

	return s
}

func (s *_createOperation) Routing(routing string) *_createOperation {

	s.v.Routing = &routing

	return s
}

func (s *_createOperation) Version(versionnumber int64) *_createOperation {

	s.v.Version = &versionnumber

	return s
}

func (s *_createOperation) VersionType(versiontype versiontype.VersionType) *_createOperation {

	s.v.VersionType = &versiontype
	return s
}

func (s *_createOperation) OperationContainerCaster() *types.OperationContainer {
	container := types.NewOperationContainer()

	container.Create = s.v

	return container
}

func (s *_createOperation) CreateOperationCaster() *types.CreateOperation {
	return s.v
}
