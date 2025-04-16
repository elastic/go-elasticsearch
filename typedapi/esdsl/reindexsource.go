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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _reindexSource struct {
	v *types.ReindexSource
}

func NewReindexSource() *_reindexSource {

	return &_reindexSource{v: types.NewReindexSource()}

}

func (s *_reindexSource) Index(indices ...string) *_reindexSource {

	s.v.Index = indices

	return s
}

func (s *_reindexSource) Query(query types.QueryVariant) *_reindexSource {

	s.v.Query = query.QueryCaster()

	return s
}

func (s *_reindexSource) Remote(remote types.RemoteSourceVariant) *_reindexSource {

	s.v.Remote = remote.RemoteSourceCaster()

	return s
}

func (s *_reindexSource) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_reindexSource {

	s.v.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return s
}

func (s *_reindexSource) Size(size int) *_reindexSource {

	s.v.Size = &size

	return s
}

func (s *_reindexSource) Slice(slice types.SlicedScrollVariant) *_reindexSource {

	s.v.Slice = slice.SlicedScrollCaster()

	return s
}

func (s *_reindexSource) Sort(sorts ...types.SortCombinationsVariant) *_reindexSource {

	for _, v := range sorts {
		s.v.Sort = append(s.v.Sort, *v.SortCombinationsCaster())
	}

	return s
}

func (s *_reindexSource) SourceFields_(fields ...string) *_reindexSource {

	s.v.SourceFields_ = fields

	return s
}

func (s *_reindexSource) ReindexSourceCaster() *types.ReindexSource {
	return s.v
}
