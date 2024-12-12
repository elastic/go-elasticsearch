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
// https://github.com/elastic/elasticsearch-specification/tree/1ed5f4795fc7c4d9875601f883b8d5fb9023c526

package types

// RoleTemplateQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1ed5f4795fc7c4d9875601f883b8d5fb9023c526/specification/security/_types/Privileges.ts#L327-L337
type RoleTemplateQuery struct {
	// Template When you create a role, you can specify a query that defines the document
	// level security permissions. You can optionally
	// use Mustache templates in the role query to insert the username of the
	// current authenticated user into the role.
	// Like other places in Elasticsearch that support templating or scripting, you
	// can specify inline, stored, or file-based
	// templates and define custom parameters. You access the details for the
	// current authenticated user through the _user parameter.
	Template *RoleTemplateScript `json:"template,omitempty"`
}

// NewRoleTemplateQuery returns a RoleTemplateQuery.
func NewRoleTemplateQuery() *RoleTemplateQuery {
	r := &RoleTemplateQuery{}

	return r
}
