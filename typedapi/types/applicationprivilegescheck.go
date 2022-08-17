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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// ApplicationPrivilegesCheck type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/has_privileges/types.ts#L24-L31
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

// ApplicationPrivilegesCheckBuilder holds ApplicationPrivilegesCheck struct and provides a builder API.
type ApplicationPrivilegesCheckBuilder struct {
	v *ApplicationPrivilegesCheck
}

// NewApplicationPrivilegesCheck provides a builder for the ApplicationPrivilegesCheck struct.
func NewApplicationPrivilegesCheckBuilder() *ApplicationPrivilegesCheckBuilder {
	r := ApplicationPrivilegesCheckBuilder{
		&ApplicationPrivilegesCheck{},
	}

	return &r
}

// Build finalize the chain and returns the ApplicationPrivilegesCheck struct
func (rb *ApplicationPrivilegesCheckBuilder) Build() ApplicationPrivilegesCheck {
	return *rb.v
}

// Application The name of the application.

func (rb *ApplicationPrivilegesCheckBuilder) Application(application string) *ApplicationPrivilegesCheckBuilder {
	rb.v.Application = application
	return rb
}

// Privileges A list of the privileges that you want to check for the specified resources.
// May be either application privilege names, or the names of actions that are
// granted by those privileges

func (rb *ApplicationPrivilegesCheckBuilder) Privileges(privileges ...string) *ApplicationPrivilegesCheckBuilder {
	rb.v.Privileges = privileges
	return rb
}

// Resources A list of resource names against which the privileges should be checked

func (rb *ApplicationPrivilegesCheckBuilder) Resources(resources ...string) *ApplicationPrivilegesCheckBuilder {
	rb.v.Resources = resources
	return rb
}
