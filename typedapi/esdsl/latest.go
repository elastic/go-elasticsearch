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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _latest struct {
	v *types.Latest
}

func NewLatest() *_latest {

	return &_latest{v: types.NewLatest()}

}

func (s *_latest) Sort(field string) *_latest {

	s.v.Sort = field

	return s
}

func (s *_latest) UniqueKey(uniquekeys ...string) *_latest {

	for _, v := range uniquekeys {

		s.v.UniqueKey = append(s.v.UniqueKey, v)

	}
	return s
}

func (s *_latest) LatestCaster() *types.Latest {
	return s.v
}
