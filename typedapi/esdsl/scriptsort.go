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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptsorttype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type _scriptSort struct {
	v *types.ScriptSort
}

func NewScriptSort(script types.ScriptVariant) *_scriptSort {

	tmp := &_scriptSort{v: types.NewScriptSort()}

	tmp.Script(script)

	return tmp

}

func (s *_scriptSort) Mode(mode sortmode.SortMode) *_scriptSort {

	s.v.Mode = &mode
	return s
}

func (s *_scriptSort) Nested(nested types.NestedSortValueVariant) *_scriptSort {

	s.v.Nested = nested.NestedSortValueCaster()

	return s
}

func (s *_scriptSort) Order(order sortorder.SortOrder) *_scriptSort {

	s.v.Order = &order
	return s
}

func (s *_scriptSort) Script(script types.ScriptVariant) *_scriptSort {

	s.v.Script = *script.ScriptCaster()

	return s
}

func (s *_scriptSort) Type(type_ scriptsorttype.ScriptSortType) *_scriptSort {

	s.v.Type = &type_
	return s
}

func (s *_scriptSort) SortOptionsCaster() *types.SortOptions {
	container := types.NewSortOptions()

	container.Script_ = s.v

	return container
}

func (s *_scriptSort) ScriptSortCaster() *types.ScriptSort {
	return s.v
}
