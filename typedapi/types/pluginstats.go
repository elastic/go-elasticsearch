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

// PluginStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L137-L148
type PluginStats struct {
	Classname            string        `json:"classname"`
	Description          string        `json:"description"`
	ElasticsearchVersion VersionString `json:"elasticsearch_version"`
	ExtendedPlugins      []string      `json:"extended_plugins"`
	HasNativeController  bool          `json:"has_native_controller"`
	JavaVersion          VersionString `json:"java_version"`
	Licensed             bool          `json:"licensed"`
	Name                 Name          `json:"name"`
	Type                 string        `json:"type"`
	Version              VersionString `json:"version"`
}

// PluginStatsBuilder holds PluginStats struct and provides a builder API.
type PluginStatsBuilder struct {
	v *PluginStats
}

// NewPluginStats provides a builder for the PluginStats struct.
func NewPluginStatsBuilder() *PluginStatsBuilder {
	r := PluginStatsBuilder{
		&PluginStats{},
	}

	return &r
}

// Build finalize the chain and returns the PluginStats struct
func (rb *PluginStatsBuilder) Build() PluginStats {
	return *rb.v
}

func (rb *PluginStatsBuilder) Classname(classname string) *PluginStatsBuilder {
	rb.v.Classname = classname
	return rb
}

func (rb *PluginStatsBuilder) Description(description string) *PluginStatsBuilder {
	rb.v.Description = description
	return rb
}

func (rb *PluginStatsBuilder) ElasticsearchVersion(elasticsearchversion VersionString) *PluginStatsBuilder {
	rb.v.ElasticsearchVersion = elasticsearchversion
	return rb
}

func (rb *PluginStatsBuilder) ExtendedPlugins(extended_plugins ...string) *PluginStatsBuilder {
	rb.v.ExtendedPlugins = extended_plugins
	return rb
}

func (rb *PluginStatsBuilder) HasNativeController(hasnativecontroller bool) *PluginStatsBuilder {
	rb.v.HasNativeController = hasnativecontroller
	return rb
}

func (rb *PluginStatsBuilder) JavaVersion(javaversion VersionString) *PluginStatsBuilder {
	rb.v.JavaVersion = javaversion
	return rb
}

func (rb *PluginStatsBuilder) Licensed(licensed bool) *PluginStatsBuilder {
	rb.v.Licensed = licensed
	return rb
}

func (rb *PluginStatsBuilder) Name(name Name) *PluginStatsBuilder {
	rb.v.Name = name
	return rb
}

func (rb *PluginStatsBuilder) Type_(type_ string) *PluginStatsBuilder {
	rb.v.Type = type_
	return rb
}

func (rb *PluginStatsBuilder) Version(version VersionString) *PluginStatsBuilder {
	rb.v.Version = version
	return rb
}
