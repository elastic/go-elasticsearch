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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// Like holds the union for the following types:
//
//	LikeDocument
//	string
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/specialized.ts#L103-L108
type Like interface{}

// LikeBuilder holds Like struct and provides a builder API.
type LikeBuilder struct {
	v Like
}

// NewLike provides a builder for the Like struct.
func NewLikeBuilder() *LikeBuilder {
	return &LikeBuilder{}
}

// Build finalize the chain and returns the Like struct
func (u *LikeBuilder) Build() Like {
	return u.v
}

func (u *LikeBuilder) LikeDocument(likedocument *LikeDocumentBuilder) *LikeBuilder {
	v := likedocument.Build()
	u.v = &v
	return u
}

func (u *LikeBuilder) String(string string) *LikeBuilder {
	u.v = &string
	return u
}
