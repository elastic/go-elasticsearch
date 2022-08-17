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

// WatcherWatch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L389-L394
type WatcherWatch struct {
	Action    map[Name]Counter    `json:"action,omitempty"`
	Condition map[Name]Counter    `json:"condition,omitempty"`
	Input     map[Name]Counter    `json:"input"`
	Trigger   WatcherWatchTrigger `json:"trigger"`
}

// WatcherWatchBuilder holds WatcherWatch struct and provides a builder API.
type WatcherWatchBuilder struct {
	v *WatcherWatch
}

// NewWatcherWatch provides a builder for the WatcherWatch struct.
func NewWatcherWatchBuilder() *WatcherWatchBuilder {
	r := WatcherWatchBuilder{
		&WatcherWatch{
			Action:    make(map[Name]Counter, 0),
			Condition: make(map[Name]Counter, 0),
			Input:     make(map[Name]Counter, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the WatcherWatch struct
func (rb *WatcherWatchBuilder) Build() WatcherWatch {
	return *rb.v
}

func (rb *WatcherWatchBuilder) Action(values map[Name]*CounterBuilder) *WatcherWatchBuilder {
	tmp := make(map[Name]Counter, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Action = tmp
	return rb
}

func (rb *WatcherWatchBuilder) Condition(values map[Name]*CounterBuilder) *WatcherWatchBuilder {
	tmp := make(map[Name]Counter, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Condition = tmp
	return rb
}

func (rb *WatcherWatchBuilder) Input(values map[Name]*CounterBuilder) *WatcherWatchBuilder {
	tmp := make(map[Name]Counter, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Input = tmp
	return rb
}

func (rb *WatcherWatchBuilder) Trigger(trigger *WatcherWatchTriggerBuilder) *WatcherWatchBuilder {
	v := trigger.Build()
	rb.v.Trigger = v
	return rb
}
