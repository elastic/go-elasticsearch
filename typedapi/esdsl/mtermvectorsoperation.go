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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _mTermVectorsOperation struct {
	v *types.MTermVectorsOperation
}

func NewMTermVectorsOperation() *_mTermVectorsOperation {

	return &_mTermVectorsOperation{v: types.NewMTermVectorsOperation()}

}

func (s *_mTermVectorsOperation) Doc(doc json.RawMessage) *_mTermVectorsOperation {

	s.v.Doc = doc

	return s
}

func (s *_mTermVectorsOperation) FieldStatistics(fieldstatistics bool) *_mTermVectorsOperation {

	s.v.FieldStatistics = &fieldstatistics

	return s
}

func (s *_mTermVectorsOperation) Fields(fields ...string) *_mTermVectorsOperation {

	s.v.Fields = fields

	return s
}

func (s *_mTermVectorsOperation) Filter(filter types.TermVectorsFilterVariant) *_mTermVectorsOperation {

	s.v.Filter = filter.TermVectorsFilterCaster()

	return s
}

func (s *_mTermVectorsOperation) Id_(id string) *_mTermVectorsOperation {

	s.v.Id_ = &id

	return s
}

func (s *_mTermVectorsOperation) Index_(indexname string) *_mTermVectorsOperation {

	s.v.Index_ = &indexname

	return s
}

func (s *_mTermVectorsOperation) Offsets(offsets bool) *_mTermVectorsOperation {

	s.v.Offsets = &offsets

	return s
}

func (s *_mTermVectorsOperation) Payloads(payloads bool) *_mTermVectorsOperation {

	s.v.Payloads = &payloads

	return s
}

func (s *_mTermVectorsOperation) Positions(positions bool) *_mTermVectorsOperation {

	s.v.Positions = &positions

	return s
}

func (s *_mTermVectorsOperation) Routing(routings ...string) *_mTermVectorsOperation {

	s.v.Routing = routings

	return s
}

func (s *_mTermVectorsOperation) TermStatistics(termstatistics bool) *_mTermVectorsOperation {

	s.v.TermStatistics = &termstatistics

	return s
}

func (s *_mTermVectorsOperation) Version(versionnumber int64) *_mTermVectorsOperation {

	s.v.Version = &versionnumber

	return s
}

func (s *_mTermVectorsOperation) VersionType(versiontype versiontype.VersionType) *_mTermVectorsOperation {

	s.v.VersionType = &versiontype
	return s
}

func (s *_mTermVectorsOperation) MTermVectorsOperationCaster() *types.MTermVectorsOperation {
	return s.v
}
