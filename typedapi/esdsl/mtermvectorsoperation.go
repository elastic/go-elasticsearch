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
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

type _mTermVectorsOperation struct {
	v *types.MTermVectorsOperation
}

func NewMTermVectorsOperation() *_mTermVectorsOperation {

	return &_mTermVectorsOperation{v: types.NewMTermVectorsOperation()}

}

// An artificial document (a document not present in the index) for which you
// want to retrieve term vectors.
func (s *_mTermVectorsOperation) Doc(doc json.RawMessage) *_mTermVectorsOperation {

	s.v.Doc = doc

	return s
}

// If `true`, the response includes the document count, sum of document
// frequencies, and sum of total term frequencies.
func (s *_mTermVectorsOperation) FieldStatistics(fieldstatistics bool) *_mTermVectorsOperation {

	s.v.FieldStatistics = &fieldstatistics

	return s
}

// Comma-separated list or wildcard expressions of fields to include in the
// statistics.
// Used as the default list unless a specific field list is provided in the
// `completion_fields` or `fielddata_fields` parameters.
func (s *_mTermVectorsOperation) Fields(fields ...string) *_mTermVectorsOperation {

	s.v.Fields = fields

	return s
}

// Filter terms based on their tf-idf scores.
func (s *_mTermVectorsOperation) Filter(filter types.TermVectorsFilterVariant) *_mTermVectorsOperation {

	s.v.Filter = filter.TermVectorsFilterCaster()

	return s
}

// The ID of the document.
func (s *_mTermVectorsOperation) Id_(id string) *_mTermVectorsOperation {

	s.v.Id_ = &id

	return s
}

// The index of the document.
func (s *_mTermVectorsOperation) Index_(indexname string) *_mTermVectorsOperation {

	s.v.Index_ = &indexname

	return s
}

// If `true`, the response includes term offsets.
func (s *_mTermVectorsOperation) Offsets(offsets bool) *_mTermVectorsOperation {

	s.v.Offsets = &offsets

	return s
}

// If `true`, the response includes term payloads.
func (s *_mTermVectorsOperation) Payloads(payloads bool) *_mTermVectorsOperation {

	s.v.Payloads = &payloads

	return s
}

// If `true`, the response includes term positions.
func (s *_mTermVectorsOperation) Positions(positions bool) *_mTermVectorsOperation {

	s.v.Positions = &positions

	return s
}

// Custom value used to route operations to a specific shard.
func (s *_mTermVectorsOperation) Routing(routing string) *_mTermVectorsOperation {

	s.v.Routing = &routing

	return s
}

// If true, the response includes term frequency and document frequency.
func (s *_mTermVectorsOperation) TermStatistics(termstatistics bool) *_mTermVectorsOperation {

	s.v.TermStatistics = &termstatistics

	return s
}

// If `true`, returns the document version as part of a hit.
func (s *_mTermVectorsOperation) Version(versionnumber int64) *_mTermVectorsOperation {

	s.v.Version = &versionnumber

	return s
}

// Specific version type.
func (s *_mTermVectorsOperation) VersionType(versiontype versiontype.VersionType) *_mTermVectorsOperation {

	s.v.VersionType = &versiontype
	return s
}

func (s *_mTermVectorsOperation) MTermVectorsOperationCaster() *types.MTermVectorsOperation {
	return s.v
}
