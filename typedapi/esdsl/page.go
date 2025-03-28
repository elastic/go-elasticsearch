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

type _page struct {
	v *types.Page
}

func NewPage() *_page {

	return &_page{v: types.NewPage()}

}

// Skips the specified number of items.
func (s *_page) From(from int) *_page {

	s.v.From = &from

	return s
}

// Specifies the maximum number of items to obtain.
func (s *_page) Size(size int) *_page {

	s.v.Size = &size

	return s
}

func (s *_page) PageCaster() *types.Page {
	return s.v
}
