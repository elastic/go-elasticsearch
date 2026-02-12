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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/versiontype"
)

type _reindexDestination struct {
	v *types.ReindexDestination
}

func NewReindexDestination() *_reindexDestination {

	return &_reindexDestination{v: types.NewReindexDestination()}

}

func (s *_reindexDestination) Index(indexname string) *_reindexDestination {

	s.v.Index = indexname

	return s
}

func (s *_reindexDestination) OpType(optype optype.OpType) *_reindexDestination {

	s.v.OpType = &optype
	return s
}

func (s *_reindexDestination) Pipeline(pipeline string) *_reindexDestination {

	s.v.Pipeline = &pipeline

	return s
}

func (s *_reindexDestination) Routing(routings ...string) *_reindexDestination {

	s.v.Routing = routings

	return s
}

func (s *_reindexDestination) VersionType(versiontype versiontype.VersionType) *_reindexDestination {

	s.v.VersionType = &versiontype
	return s
}

func (s *_reindexDestination) ReindexDestinationCaster() *types.ReindexDestination {
	return s.v
}
