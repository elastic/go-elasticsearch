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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hop struct {
	v *types.Hop
}

func NewHop() *_hop {

	return &_hop{v: types.NewHop()}

}

func (s *_hop) Connections(connections types.HopVariant) *_hop {

	s.v.Connections = connections.HopCaster()

	return s
}

func (s *_hop) Query(query types.QueryVariant) *_hop {

	s.v.Query = query.QueryCaster()

	return s
}

func (s *_hop) Vertices(vertices ...types.VertexDefinitionVariant) *_hop {

	for _, v := range vertices {

		s.v.Vertices = append(s.v.Vertices, *v.VertexDefinitionCaster())

	}
	return s
}

func (s *_hop) HopCaster() *types.Hop {
	return s.v
}
