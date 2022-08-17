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

// Monitoring type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L370-L373
type Monitoring struct {
	Available         bool             `json:"available"`
	CollectionEnabled bool             `json:"collection_enabled"`
	Enabled           bool             `json:"enabled"`
	EnabledExporters  map[string]int64 `json:"enabled_exporters"`
}

// MonitoringBuilder holds Monitoring struct and provides a builder API.
type MonitoringBuilder struct {
	v *Monitoring
}

// NewMonitoring provides a builder for the Monitoring struct.
func NewMonitoringBuilder() *MonitoringBuilder {
	r := MonitoringBuilder{
		&Monitoring{
			EnabledExporters: make(map[string]int64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Monitoring struct
func (rb *MonitoringBuilder) Build() Monitoring {
	return *rb.v
}

func (rb *MonitoringBuilder) Available(available bool) *MonitoringBuilder {
	rb.v.Available = available
	return rb
}

func (rb *MonitoringBuilder) CollectionEnabled(collectionenabled bool) *MonitoringBuilder {
	rb.v.CollectionEnabled = collectionenabled
	return rb
}

func (rb *MonitoringBuilder) Enabled(enabled bool) *MonitoringBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *MonitoringBuilder) EnabledExporters(value map[string]int64) *MonitoringBuilder {
	rb.v.EnabledExporters = value
	return rb
}
