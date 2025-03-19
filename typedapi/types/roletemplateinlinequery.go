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

package types

// RoleTemplateInlineQuery holds the union for the following types:
//
//	string
//	Query
//
// https://github.com/elastic/elasticsearch-specification/blob/c75a0abec670d027d13eb8d6f23374f86621c76b/specification/security/_types/Privileges.ts#L383-L384
type RoleTemplateInlineQuery any

type RoleTemplateInlineQueryVariant interface {
	RoleTemplateInlineQueryCaster() *RoleTemplateInlineQuery
}
