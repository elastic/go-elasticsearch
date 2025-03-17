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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _transformSource struct {
	v *types.TransformSource
}

func NewTransformSource() *_transformSource {

	return &_transformSource{v: types.NewTransformSource()}

}

// The source indices for the transform. It can be a single index, an index
// pattern (for example, `"my-index-*""`), an
// array of indices (for example, `["my-index-000001", "my-index-000002"]`), or
// an array of index patterns (for
// example, `["my-index-*", "my-other-index-*"]`. For remote indices use the
// syntax `"remote_name:index_name"`. If
// any indices are in remote clusters then the master node and at least one
// transform node must have the `remote_cluster_client` node role.
func (s *_transformSource) Index(indices ...string) *_transformSource {

	s.v.Index = indices

	return s
}

// A query clause that retrieves a subset of data from the source index.
func (s *_transformSource) Query(query types.QueryVariant) *_transformSource {

	s.v.Query = query.QueryCaster()

	return s
}

// Definitions of search-time runtime fields that can be used by the transform.
// For search runtime fields all data
// nodes, including remote nodes, must be 7.12 or later.
func (s *_transformSource) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_transformSource {

	s.v.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return s
}

func (s *_transformSource) TransformSourceCaster() *types.TransformSource {
	return s.v
}
