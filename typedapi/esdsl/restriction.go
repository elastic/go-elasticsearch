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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/restrictionworkflow"
)

type _restriction struct {
	v *types.Restriction
}

func NewRestriction() *_restriction {

	return &_restriction{v: types.NewRestriction()}

}

func (s *_restriction) Workflows(workflows ...restrictionworkflow.RestrictionWorkflow) *_restriction {

	for _, v := range workflows {

		s.v.Workflows = append(s.v.Workflows, v)

	}
	return s
}

func (s *_restriction) RestrictionCaster() *types.Restriction {
	return s.v
}
