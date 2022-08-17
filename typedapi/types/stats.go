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

// Stats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L30-L53
type Stats struct {
	AdaptiveSelection map[string]AdaptiveSelection `json:"adaptive_selection,omitempty"`
	Attributes        map[Field]string             `json:"attributes,omitempty"`
	Breakers          map[string]Breaker           `json:"breakers,omitempty"`
	Discovery         *Discovery                   `json:"discovery,omitempty"`
	Fs                *FileSystem                  `json:"fs,omitempty"`
	Host              *Host                        `json:"host,omitempty"`
	Http              *Http                        `json:"http,omitempty"`
	IndexingPressure  *IndexingPressure            `json:"indexing_pressure,omitempty"`
	Indices           *ShardStats                  `json:"indices,omitempty"`
	Ingest            *Ingest                      `json:"ingest,omitempty"`
	Ip                []Ip                         `json:"ip,omitempty"`
	Jvm               *Jvm                         `json:"jvm,omitempty"`
	Name              *Name                        `json:"name,omitempty"`
	Os                *OperatingSystem             `json:"os,omitempty"`
	Process           *Process                     `json:"process,omitempty"`
	Roles             *NodeRoles                   `json:"roles,omitempty"`
	Script            *Scripting                   `json:"script,omitempty"`
	ScriptCache       map[string][]ScriptCache     `json:"script_cache,omitempty"`
	ThreadPool        map[string]ThreadCount       `json:"thread_pool,omitempty"`
	Timestamp         *int64                       `json:"timestamp,omitempty"`
	Transport         *Transport                   `json:"transport,omitempty"`
	TransportAddress  *TransportAddress            `json:"transport_address,omitempty"`
}

// StatsBuilder holds Stats struct and provides a builder API.
type StatsBuilder struct {
	v *Stats
}

// NewStats provides a builder for the Stats struct.
func NewStatsBuilder() *StatsBuilder {
	r := StatsBuilder{
		&Stats{
			AdaptiveSelection: make(map[string]AdaptiveSelection, 0),
			Attributes:        make(map[Field]string, 0),
			Breakers:          make(map[string]Breaker, 0),
			ScriptCache:       make(map[string][]ScriptCache, 0),
			ThreadPool:        make(map[string]ThreadCount, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Stats struct
func (rb *StatsBuilder) Build() Stats {
	return *rb.v
}

func (rb *StatsBuilder) AdaptiveSelection(values map[string]*AdaptiveSelectionBuilder) *StatsBuilder {
	tmp := make(map[string]AdaptiveSelection, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.AdaptiveSelection = tmp
	return rb
}

func (rb *StatsBuilder) Attributes(value map[Field]string) *StatsBuilder {
	rb.v.Attributes = value
	return rb
}

func (rb *StatsBuilder) Breakers(values map[string]*BreakerBuilder) *StatsBuilder {
	tmp := make(map[string]Breaker, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Breakers = tmp
	return rb
}

func (rb *StatsBuilder) Discovery(discovery *DiscoveryBuilder) *StatsBuilder {
	v := discovery.Build()
	rb.v.Discovery = &v
	return rb
}

func (rb *StatsBuilder) Fs(fs *FileSystemBuilder) *StatsBuilder {
	v := fs.Build()
	rb.v.Fs = &v
	return rb
}

func (rb *StatsBuilder) Host(host Host) *StatsBuilder {
	rb.v.Host = &host
	return rb
}

func (rb *StatsBuilder) Http(http *HttpBuilder) *StatsBuilder {
	v := http.Build()
	rb.v.Http = &v
	return rb
}

func (rb *StatsBuilder) IndexingPressure(indexingpressure *IndexingPressureBuilder) *StatsBuilder {
	v := indexingpressure.Build()
	rb.v.IndexingPressure = &v
	return rb
}

func (rb *StatsBuilder) Indices(indices *ShardStatsBuilder) *StatsBuilder {
	v := indices.Build()
	rb.v.Indices = &v
	return rb
}

func (rb *StatsBuilder) Ingest(ingest *IngestBuilder) *StatsBuilder {
	v := ingest.Build()
	rb.v.Ingest = &v
	return rb
}

func (rb *StatsBuilder) Ip(arg []Ip) *StatsBuilder {
	rb.v.Ip = arg
	return rb
}

func (rb *StatsBuilder) Jvm(jvm *JvmBuilder) *StatsBuilder {
	v := jvm.Build()
	rb.v.Jvm = &v
	return rb
}

func (rb *StatsBuilder) Name(name Name) *StatsBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *StatsBuilder) Os(os *OperatingSystemBuilder) *StatsBuilder {
	v := os.Build()
	rb.v.Os = &v
	return rb
}

func (rb *StatsBuilder) Process(process *ProcessBuilder) *StatsBuilder {
	v := process.Build()
	rb.v.Process = &v
	return rb
}

func (rb *StatsBuilder) Roles(roles *NodeRolesBuilder) *StatsBuilder {
	v := roles.Build()
	rb.v.Roles = &v
	return rb
}

func (rb *StatsBuilder) Script(script *ScriptingBuilder) *StatsBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *StatsBuilder) ScriptCache(value map[string][]ScriptCache) *StatsBuilder {
	rb.v.ScriptCache = value
	return rb
}

func (rb *StatsBuilder) ThreadPool(values map[string]*ThreadCountBuilder) *StatsBuilder {
	tmp := make(map[string]ThreadCount, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.ThreadPool = tmp
	return rb
}

func (rb *StatsBuilder) Timestamp(timestamp int64) *StatsBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}

func (rb *StatsBuilder) Transport(transport *TransportBuilder) *StatsBuilder {
	v := transport.Build()
	rb.v.Transport = &v
	return rb
}

func (rb *StatsBuilder) TransportAddress(transportaddress TransportAddress) *StatsBuilder {
	rb.v.TransportAddress = &transportaddress
	return rb
}
