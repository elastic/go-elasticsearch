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

type _applicationPrivilegesCheck struct {
	v *types.ApplicationPrivilegesCheck
}

func NewApplicationPrivilegesCheck(application string) *_applicationPrivilegesCheck {

	tmp := &_applicationPrivilegesCheck{v: types.NewApplicationPrivilegesCheck()}

	tmp.Application(application)

	return tmp

}

// The name of the application.
func (s *_applicationPrivilegesCheck) Application(application string) *_applicationPrivilegesCheck {

	s.v.Application = application

	return s
}

// A list of the privileges that you want to check for the specified resources.
// It may be either application privilege names or the names of actions that are
// granted by those privileges
func (s *_applicationPrivilegesCheck) Privileges(privileges ...string) *_applicationPrivilegesCheck {

	for _, v := range privileges {

		s.v.Privileges = append(s.v.Privileges, v)

	}
	return s
}

// A list of resource names against which the privileges should be checked.
func (s *_applicationPrivilegesCheck) Resources(resources ...string) *_applicationPrivilegesCheck {

	for _, v := range resources {

		s.v.Resources = append(s.v.Resources, v)

	}
	return s
}

func (s *_applicationPrivilegesCheck) ApplicationPrivilegesCheckCaster() *types.ApplicationPrivilegesCheck {
	return s.v
}
