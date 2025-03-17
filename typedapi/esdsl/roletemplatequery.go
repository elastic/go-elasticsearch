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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _roleTemplateQuery struct {
	v *types.RoleTemplateQuery
}

func NewRoleTemplateQuery() *_roleTemplateQuery {

	return &_roleTemplateQuery{v: types.NewRoleTemplateQuery()}

}

// When you create a role, you can specify a query that defines the document
// level security permissions. You can optionally
// use Mustache templates in the role query to insert the username of the
// current authenticated user into the role.
// Like other places in Elasticsearch that support templating or scripting, you
// can specify inline, stored, or file-based
// templates and define custom parameters. You access the details for the
// current authenticated user through the _user parameter.
func (s *_roleTemplateQuery) Template(template types.RoleTemplateScriptVariant) *_roleTemplateQuery {

	s.v.Template = template.RoleTemplateScriptCaster()

	return s
}

func (s *_roleTemplateQuery) RoleTemplateQueryCaster() *types.RoleTemplateQuery {
	return s.v
}
