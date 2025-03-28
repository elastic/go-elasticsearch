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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _reindexSource struct {
	v *types.ReindexSource
}

func NewReindexSource() *_reindexSource {

	return &_reindexSource{v: types.NewReindexSource()}

}

// The name of the data stream, index, or alias you are copying from.
// It accepts a comma-separated list to reindex from multiple sources.
func (s *_reindexSource) Index(indices ...string) *_reindexSource {

	s.v.Index = indices

	return s
}

// The documents to reindex, which is defined with Query DSL.
func (s *_reindexSource) Query(query types.QueryVariant) *_reindexSource {

	s.v.Query = query.QueryCaster()

	return s
}

// A remote instance of Elasticsearch that you want to index from.
func (s *_reindexSource) Remote(remote types.RemoteSourceVariant) *_reindexSource {

	s.v.Remote = remote.RemoteSourceCaster()

	return s
}

func (s *_reindexSource) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_reindexSource {

	s.v.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return s
}

// The number of documents to index per batch.
// Use it when you are indexing from remote to ensure that the batches fit
// within the on-heap buffer, which defaults to a maximum size of 100 MB.
func (s *_reindexSource) Size(size int) *_reindexSource {

	s.v.Size = &size

	return s
}

// Slice the reindex request manually using the provided slice ID and total
// number of slices.
func (s *_reindexSource) Slice(slice types.SlicedScrollVariant) *_reindexSource {

	s.v.Slice = slice.SlicedScrollCaster()

	return s
}

// A comma-separated list of `<field>:<direction>` pairs to sort by before
// indexing.
// Use it in conjunction with `max_docs` to control what documents are
// reindexed.
//
// WARNING: Sort in reindex is deprecated.
// Sorting in reindex was never guaranteed to index documents in order and
// prevents further development of reindex such as resilience and performance
// improvements.
// If used in combination with `max_docs`, consider using a query filter
// instead.
func (s *_reindexSource) Sort(sorts ...types.SortCombinationsVariant) *_reindexSource {

	for _, v := range sorts {
		s.v.Sort = append(s.v.Sort, *v.SortCombinationsCaster())
	}

	return s
}

// If `true`, reindex all source fields.
// Set it to a list to reindex select fields.
func (s *_reindexSource) SourceFields_(fields ...string) *_reindexSource {

	s.v.SourceFields_ = fields

	return s
}

func (s *_reindexSource) ReindexSourceCaster() *types.ReindexSource {
	return s.v
}
