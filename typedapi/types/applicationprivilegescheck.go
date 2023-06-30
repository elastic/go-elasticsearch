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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

// ApplicationPrivilegesCheck type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/security/has_privileges/types.ts#L24-L31
type ApplicationPrivilegesCheck struct {
	// Application The name of the application.
	Application string `json:"application"`
	// Privileges A list of the privileges that you want to check for the specified resources.
	// May be either application privilege names, or the names of actions that are
	// granted by those privileges
	Privileges []string `json:"privileges"`
	// Resources A list of resource names against which the privileges should be checked
	Resources []string `json:"resources"`
}

// NewApplicationPrivilegesCheck returns a ApplicationPrivilegesCheck.
func NewApplicationPrivilegesCheck() *ApplicationPrivilegesCheck {
	r := &ApplicationPrivilegesCheck{}

	return r
}
