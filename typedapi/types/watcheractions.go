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

// WatcherActions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L385-L387
type WatcherActions struct {
	Actions map[Name]WatcherActionTotals `json:"actions"`
}

// WatcherActionsBuilder holds WatcherActions struct and provides a builder API.
type WatcherActionsBuilder struct {
	v *WatcherActions
}

// NewWatcherActions provides a builder for the WatcherActions struct.
func NewWatcherActionsBuilder() *WatcherActionsBuilder {
	r := WatcherActionsBuilder{
		&WatcherActions{
			Actions: make(map[Name]WatcherActionTotals, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the WatcherActions struct
func (rb *WatcherActionsBuilder) Build() WatcherActions {
	return *rb.v
}

func (rb *WatcherActionsBuilder) Actions(values map[Name]*WatcherActionTotalsBuilder) *WatcherActionsBuilder {
	tmp := make(map[Name]WatcherActionTotals, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Actions = tmp
	return rb
}
