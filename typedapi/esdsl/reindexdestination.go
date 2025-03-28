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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

type _reindexDestination struct {
	v *types.ReindexDestination
}

func NewReindexDestination() *_reindexDestination {

	return &_reindexDestination{v: types.NewReindexDestination()}

}

// The name of the data stream, index, or index alias you are copying to.
func (s *_reindexDestination) Index(indexname string) *_reindexDestination {

	s.v.Index = indexname

	return s
}

// If it is `create`, the operation will only index documents that do not
// already exist (also known as "put if absent").
//
// IMPORTANT: To reindex to a data stream destination, this argument must be
// `create`.
func (s *_reindexDestination) OpType(optype optype.OpType) *_reindexDestination {

	s.v.OpType = &optype
	return s
}

// The name of the pipeline to use.
func (s *_reindexDestination) Pipeline(pipeline string) *_reindexDestination {

	s.v.Pipeline = &pipeline

	return s
}

// By default, a document's routing is preserved unless it's changed by the
// script.
// If it is `keep`, the routing on the bulk request sent for each match is set
// to the routing on the match.
// If it is `discard`, the routing on the bulk request sent for each match is
// set to `null`.
// If it is `=value`, the routing on the bulk request sent for each match is set
// to all value specified after the equals sign (`=`).
func (s *_reindexDestination) Routing(routing string) *_reindexDestination {

	s.v.Routing = &routing

	return s
}

// The versioning to use for the indexing operation.
func (s *_reindexDestination) VersionType(versiontype versiontype.VersionType) *_reindexDestination {

	s.v.VersionType = &versiontype
	return s
}

func (s *_reindexDestination) ReindexDestinationCaster() *types.ReindexDestination {
	return s.v
}
