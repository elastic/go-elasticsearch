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

// GeoIpNodeDatabases type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/geo_ip_stats/types.ts#L37-L43
type GeoIpNodeDatabases struct {
	// Databases Downloaded databases for the node.
	Databases []GeoIpNodeDatabaseName `json:"databases"`
	// FilesInTemp Downloaded database files, including related license files. Elasticsearch
	// stores these files in the node’s temporary directory:
	// $ES_TMPDIR/geoip-databases/<node_id>.
	FilesInTemp []string `json:"files_in_temp"`
}

// GeoIpNodeDatabasesBuilder holds GeoIpNodeDatabases struct and provides a builder API.
type GeoIpNodeDatabasesBuilder struct {
	v *GeoIpNodeDatabases
}

// NewGeoIpNodeDatabases provides a builder for the GeoIpNodeDatabases struct.
func NewGeoIpNodeDatabasesBuilder() *GeoIpNodeDatabasesBuilder {
	r := GeoIpNodeDatabasesBuilder{
		&GeoIpNodeDatabases{},
	}

	return &r
}

// Build finalize the chain and returns the GeoIpNodeDatabases struct
func (rb *GeoIpNodeDatabasesBuilder) Build() GeoIpNodeDatabases {
	return *rb.v
}

// Databases Downloaded databases for the node.

func (rb *GeoIpNodeDatabasesBuilder) Databases(databases []GeoIpNodeDatabaseNameBuilder) *GeoIpNodeDatabasesBuilder {
	tmp := make([]GeoIpNodeDatabaseName, len(databases))
	for _, value := range databases {
		tmp = append(tmp, value.Build())
	}
	rb.v.Databases = tmp
	return rb
}

// FilesInTemp Downloaded database files, including related license files. Elasticsearch
// stores these files in the node’s temporary directory:
// $ES_TMPDIR/geoip-databases/<node_id>.

func (rb *GeoIpNodeDatabasesBuilder) FilesInTemp(files_in_temp ...string) *GeoIpNodeDatabasesBuilder {
	rb.v.FilesInTemp = files_in_temp
	return rb
}
