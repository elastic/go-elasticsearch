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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _vertexDefinition struct {
	v *types.VertexDefinition
}

func NewVertexDefinition() *_vertexDefinition {

	return &_vertexDefinition{v: types.NewVertexDefinition()}

}

func (s *_vertexDefinition) Exclude(excludes ...string) *_vertexDefinition {

	for _, v := range excludes {

		s.v.Exclude = append(s.v.Exclude, v)

	}
	return s
}

func (s *_vertexDefinition) Field(field string) *_vertexDefinition {

	s.v.Field = field

	return s
}

func (s *_vertexDefinition) Include(includes ...types.VertexIncludeVariant) *_vertexDefinition {

	for _, v := range includes {

		s.v.Include = append(s.v.Include, *v.VertexIncludeCaster())

	}
	return s
}

func (s *_vertexDefinition) MinDocCount(mindoccount int64) *_vertexDefinition {

	s.v.MinDocCount = &mindoccount

	return s
}

func (s *_vertexDefinition) ShardMinDocCount(shardmindoccount int64) *_vertexDefinition {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

func (s *_vertexDefinition) Size(size int) *_vertexDefinition {

	s.v.Size = &size

	return s
}

func (s *_vertexDefinition) VertexDefinitionCaster() *types.VertexDefinition {
	return s.v
}
