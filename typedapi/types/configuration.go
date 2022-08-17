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

// Configuration type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/slm/_types/SnapshotLifecycle.ts#L99-L129
type Configuration struct {
	// FeatureStates A list of feature states to be included in this snapshot. A list of features
	// available for inclusion in the snapshot and their descriptions be can be
	// retrieved using the get features API.
	// Each feature state includes one or more system indices containing data
	// necessary for the function of that feature. Providing an empty array will
	// include no feature states in the snapshot, regardless of the value of
	// include_global_state. By default, all available feature states will be
	// included in the snapshot if include_global_state is true, or no feature
	// states if include_global_state is false.
	FeatureStates []string `json:"feature_states,omitempty"`
	// IgnoreUnavailable If false, the snapshot fails if any data stream or index in indices is
	// missing or closed. If true, the snapshot ignores missing or closed data
	// streams and indices.
	IgnoreUnavailable *bool `json:"ignore_unavailable,omitempty"`
	// IncludeGlobalState If true, the current global state is included in the snapshot.
	IncludeGlobalState *bool `json:"include_global_state,omitempty"`
	// Indices A comma-separated list of data streams and indices to include in the
	// snapshot. Multi-index syntax is supported.
	// By default, a snapshot includes all data streams and indices in the cluster.
	// If this argument is provided, the snapshot only includes the specified data
	// streams and clusters.
	Indices Indices `json:"indices"`
	// Metadata Attaches arbitrary metadata to the snapshot, such as a record of who took the
	// snapshot, why it was taken, or any other useful data. Metadata must be less
	// than 1024 bytes.
	Metadata *Metadata `json:"metadata,omitempty"`
	// Partial If false, the entire snapshot will fail if one or more indices included in
	// the snapshot do not have all primary shards available.
	Partial *bool `json:"partial,omitempty"`
}

// ConfigurationBuilder holds Configuration struct and provides a builder API.
type ConfigurationBuilder struct {
	v *Configuration
}

// NewConfiguration provides a builder for the Configuration struct.
func NewConfigurationBuilder() *ConfigurationBuilder {
	r := ConfigurationBuilder{
		&Configuration{},
	}

	return &r
}

// Build finalize the chain and returns the Configuration struct
func (rb *ConfigurationBuilder) Build() Configuration {
	return *rb.v
}

// FeatureStates A list of feature states to be included in this snapshot. A list of features
// available for inclusion in the snapshot and their descriptions be can be
// retrieved using the get features API.
// Each feature state includes one or more system indices containing data
// necessary for the function of that feature. Providing an empty array will
// include no feature states in the snapshot, regardless of the value of
// include_global_state. By default, all available feature states will be
// included in the snapshot if include_global_state is true, or no feature
// states if include_global_state is false.

func (rb *ConfigurationBuilder) FeatureStates(feature_states ...string) *ConfigurationBuilder {
	rb.v.FeatureStates = feature_states
	return rb
}

// IgnoreUnavailable If false, the snapshot fails if any data stream or index in indices is
// missing or closed. If true, the snapshot ignores missing or closed data
// streams and indices.

func (rb *ConfigurationBuilder) IgnoreUnavailable(ignoreunavailable bool) *ConfigurationBuilder {
	rb.v.IgnoreUnavailable = &ignoreunavailable
	return rb
}

// IncludeGlobalState If true, the current global state is included in the snapshot.

func (rb *ConfigurationBuilder) IncludeGlobalState(includeglobalstate bool) *ConfigurationBuilder {
	rb.v.IncludeGlobalState = &includeglobalstate
	return rb
}

// Indices A comma-separated list of data streams and indices to include in the
// snapshot. Multi-index syntax is supported.
// By default, a snapshot includes all data streams and indices in the cluster.
// If this argument is provided, the snapshot only includes the specified data
// streams and clusters.

func (rb *ConfigurationBuilder) Indices(indices *IndicesBuilder) *ConfigurationBuilder {
	v := indices.Build()
	rb.v.Indices = v
	return rb
}

// Metadata Attaches arbitrary metadata to the snapshot, such as a record of who took the
// snapshot, why it was taken, or any other useful data. Metadata must be less
// than 1024 bytes.

func (rb *ConfigurationBuilder) Metadata(metadata *MetadataBuilder) *ConfigurationBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

// Partial If false, the entire snapshot will fail if one or more indices included in
// the snapshot do not have all primary shards available.

func (rb *ConfigurationBuilder) Partial(partial bool) *ConfigurationBuilder {
	rb.v.Partial = &partial
	return rb
}
