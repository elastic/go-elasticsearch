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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _indexOperation struct {
	v *types.IndexOperation
}

// Index the specified document.
// If the document exists, it replaces the document and increments the version.
// The following line must contain the source data to be indexed.
func NewIndexOperation() *_indexOperation {

	return &_indexOperation{v: types.NewIndexOperation()}

}

func (s *_indexOperation) DynamicTemplates(dynamictemplates map[string]string) *_indexOperation {

	s.v.DynamicTemplates = dynamictemplates
	return s
}

func (s *_indexOperation) AddDynamicTemplate(key string, value string) *_indexOperation {

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

func (s *_indexOperation) Id_(id string) *_indexOperation {

	s.v.Id_ = &id

	return s
}

func (s *_indexOperation) IfPrimaryTerm(ifprimaryterm int64) *_indexOperation {

	s.v.IfPrimaryTerm = &ifprimaryterm

	return s
}

func (s *_indexOperation) IfSeqNo(sequencenumber int64) *_indexOperation {

	s.v.IfSeqNo = &sequencenumber

	return s
}

func (s *_indexOperation) Index_(indexname string) *_indexOperation {

	s.v.Index_ = &indexname

	return s
}

func (s *_indexOperation) Pipeline(pipeline string) *_indexOperation {

	s.v.Pipeline = &pipeline

	return s
}

func (s *_indexOperation) RequireAlias(requirealias bool) *_indexOperation {

	s.v.RequireAlias = &requirealias

	return s
}

func (s *_indexOperation) Routing(routing string) *_indexOperation {

	s.v.Routing = &routing

	return s
}

func (s *_indexOperation) Version(versionnumber int64) *_indexOperation {

	s.v.Version = &versionnumber

	return s
}

func (s *_indexOperation) VersionType(versiontype versiontype.VersionType) *_indexOperation {

	s.v.VersionType = &versiontype
	return s
}

func (s *_indexOperation) OperationContainerCaster() *types.OperationContainer {
	container := types.NewOperationContainer()

	container.Index = s.v

	return container
}

func (s *_indexOperation) IndexOperationCaster() *types.IndexOperation {
	return s.v
}
