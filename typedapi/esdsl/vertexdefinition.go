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

type _vertexDefinition struct {
	v *types.VertexDefinition
}

func NewVertexDefinition() *_vertexDefinition {

	return &_vertexDefinition{v: types.NewVertexDefinition()}

}

// Prevents the specified terms from being included in the results.
func (s *_vertexDefinition) Exclude(excludes ...string) *_vertexDefinition {

	for _, v := range excludes {

		s.v.Exclude = append(s.v.Exclude, v)

	}
	return s
}

// Identifies a field in the documents of interest.
func (s *_vertexDefinition) Field(field string) *_vertexDefinition {

	s.v.Field = field

	return s
}

// Identifies the terms of interest that form the starting points from which you
// want to spider out.
func (s *_vertexDefinition) Include(includes ...types.VertexIncludeVariant) *_vertexDefinition {

	for _, v := range includes {

		s.v.Include = append(s.v.Include, *v.VertexIncludeCaster())

	}
	return s
}

// Specifies how many documents must contain a pair of terms before it is
// considered to be a useful connection.
// This setting acts as a certainty threshold.
func (s *_vertexDefinition) MinDocCount(mindoccount int64) *_vertexDefinition {

	s.v.MinDocCount = &mindoccount

	return s
}

// Controls how many documents on a particular shard have to contain a pair of
// terms before the connection is returned for global consideration.
func (s *_vertexDefinition) ShardMinDocCount(shardmindoccount int64) *_vertexDefinition {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

// Specifies the maximum number of vertex terms returned for each field.
func (s *_vertexDefinition) Size(size int) *_vertexDefinition {

	s.v.Size = &size

	return s
}

func (s *_vertexDefinition) VertexDefinitionCaster() *types.VertexDefinition {
	return s.v
}
