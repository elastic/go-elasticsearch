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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

type _privilegesCheck struct {
	v *types.PrivilegesCheck
}

func NewPrivilegesCheck() *_privilegesCheck {

	return &_privilegesCheck{v: types.NewPrivilegesCheck()}

}

func (s *_privilegesCheck) Application(applications ...types.ApplicationPrivilegesCheckVariant) *_privilegesCheck {

	for _, v := range applications {

		s.v.Application = append(s.v.Application, *v.ApplicationPrivilegesCheckCaster())

	}
	return s
}

func (s *_privilegesCheck) Cluster(clusters ...clusterprivilege.ClusterPrivilege) *_privilegesCheck {

	for _, v := range clusters {

		s.v.Cluster = append(s.v.Cluster, v)

	}
	return s
}

func (s *_privilegesCheck) Index(indices ...types.IndexPrivilegesCheckVariant) *_privilegesCheck {

	for _, v := range indices {

		s.v.Index = append(s.v.Index, *v.IndexPrivilegesCheckCaster())

	}
	return s
}

func (s *_privilegesCheck) PrivilegesCheckCaster() *types.PrivilegesCheck {
	return s.v
}
