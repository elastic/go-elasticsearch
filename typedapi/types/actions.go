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

// Actions type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Action.ts#L59-L59
type Actions map[ActionStatus]IndexName

// ActionsBuilder holds Actions struct and provides a builder API.
type ActionsBuilder struct {
	v Actions
}

// NewActions provides a builder for the Actions struct.
func NewActionsBuilder() *ActionsBuilder {
	return &ActionsBuilder{}
}

// Build finalize the chain and returns the Actions struct
func (b *ActionsBuilder) Build() Actions {
	return b.v
}

func (b *ActionsBuilder) Actions(value Actions) *ActionsBuilder {
	b.v = value
	return b
}
