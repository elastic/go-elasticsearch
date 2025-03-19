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

type _applicationPrivileges struct {
	v *types.ApplicationPrivileges
}

func NewApplicationPrivileges(application string) *_applicationPrivileges {

	tmp := &_applicationPrivileges{v: types.NewApplicationPrivileges()}

	tmp.Application(application)

	return tmp

}

// The name of the application to which this entry applies.
func (s *_applicationPrivileges) Application(application string) *_applicationPrivileges {

	s.v.Application = application

	return s
}

// A list of strings, where each element is the name of an application privilege
// or action.
func (s *_applicationPrivileges) Privileges(privileges ...string) *_applicationPrivileges {

	for _, v := range privileges {

		s.v.Privileges = append(s.v.Privileges, v)

	}
	return s
}

// A list resources to which the privileges are applied.
func (s *_applicationPrivileges) Resources(resources ...string) *_applicationPrivileges {

	for _, v := range resources {

		s.v.Resources = append(s.v.Resources, v)

	}
	return s
}

func (s *_applicationPrivileges) ApplicationPrivilegesCaster() *types.ApplicationPrivileges {
	return s.v
}
