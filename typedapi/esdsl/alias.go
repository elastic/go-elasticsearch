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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _alias struct {
	v *types.Alias
}

func NewAlias() *_alias {

	return &_alias{v: types.NewAlias()}

}

func (s *_alias) Filter(filter types.QueryVariant) *_alias {

	s.v.Filter = filter.QueryCaster()

	return s
}

func (s *_alias) IndexRouting(routings ...string) *_alias {

	s.v.IndexRouting = routings

	return s
}

func (s *_alias) IsHidden(ishidden bool) *_alias {

	s.v.IsHidden = &ishidden

	return s
}

func (s *_alias) IsWriteIndex(iswriteindex bool) *_alias {

	s.v.IsWriteIndex = &iswriteindex

	return s
}

func (s *_alias) Routing(routings ...string) *_alias {

	s.v.Routing = routings

	return s
}

func (s *_alias) SearchRouting(routings ...string) *_alias {

	s.v.SearchRouting = routings

	return s
}

func (s *_alias) AliasCaster() *types.Alias {
	return s.v
}
