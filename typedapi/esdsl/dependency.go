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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dependency struct {
	v *types.Dependency
}

func NewDependency(field string) *_dependency {

	tmp := &_dependency{v: types.NewDependency()}

	tmp.Field(field)

	return tmp

}

func (s *_dependency) Field(field string) *_dependency {

	s.v.Field = field

	return s
}

func (s *_dependency) Value(scalarvalue types.ScalarValueVariant) *_dependency {

	s.v.Value = *scalarvalue.ScalarValueCaster()

	return s
}

func (s *_dependency) DependencyCaster() *types.Dependency {
	return s.v
}
