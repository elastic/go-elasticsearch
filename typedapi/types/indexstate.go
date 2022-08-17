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

// IndexState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexState.ts#L26-L33
type IndexState struct {
	Aliases    map[IndexName]Alias `json:"aliases,omitempty"`
	DataStream *DataStreamName     `json:"data_stream,omitempty"`
	// Defaults Default settings, included when the request's `include_default` is `true`.
	Defaults *IndexSettings `json:"defaults,omitempty"`
	Mappings *TypeMapping   `json:"mappings,omitempty"`
	Settings *IndexSettings `json:"settings,omitempty"`
}

// IndexStateBuilder holds IndexState struct and provides a builder API.
type IndexStateBuilder struct {
	v *IndexState
}

// NewIndexState provides a builder for the IndexState struct.
func NewIndexStateBuilder() *IndexStateBuilder {
	r := IndexStateBuilder{
		&IndexState{
			Aliases: make(map[IndexName]Alias, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndexState struct
func (rb *IndexStateBuilder) Build() IndexState {
	return *rb.v
}

func (rb *IndexStateBuilder) Aliases(values map[IndexName]*AliasBuilder) *IndexStateBuilder {
	tmp := make(map[IndexName]Alias, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aliases = tmp
	return rb
}

func (rb *IndexStateBuilder) DataStream(datastream DataStreamName) *IndexStateBuilder {
	rb.v.DataStream = &datastream
	return rb
}

// Defaults Default settings, included when the request's `include_default` is `true`.

func (rb *IndexStateBuilder) Defaults(defaults *IndexSettingsBuilder) *IndexStateBuilder {
	v := defaults.Build()
	rb.v.Defaults = &v
	return rb
}

func (rb *IndexStateBuilder) Mappings(mappings *TypeMappingBuilder) *IndexStateBuilder {
	v := mappings.Build()
	rb.v.Mappings = &v
	return rb
}

func (rb *IndexStateBuilder) Settings(settings *IndexSettingsBuilder) *IndexStateBuilder {
	v := settings.Build()
	rb.v.Settings = &v
	return rb
}
