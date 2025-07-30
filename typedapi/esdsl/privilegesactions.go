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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _privilegesActions struct {
	v *types.PrivilegesActions
}

func NewPrivilegesActions() *_privilegesActions {

	return &_privilegesActions{v: types.NewPrivilegesActions()}

}

func (s *_privilegesActions) Actions(actions ...string) *_privilegesActions {

	for _, v := range actions {

		s.v.Actions = append(s.v.Actions, v)

	}
	return s
}

func (s *_privilegesActions) Application(application string) *_privilegesActions {

	s.v.Application = &application

	return s
}

func (s *_privilegesActions) Metadata(metadata types.MetadataVariant) *_privilegesActions {

	s.v.Metadata = *metadata.MetadataCaster()

	return s
}

func (s *_privilegesActions) Name(name string) *_privilegesActions {

	s.v.Name = &name

	return s
}

func (s *_privilegesActions) PrivilegesActionsCaster() *types.PrivilegesActions {
	return s.v
}
