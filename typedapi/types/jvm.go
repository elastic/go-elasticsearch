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

// Jvm type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L318-L327
type Jvm struct {
	BufferPools    map[string]NodeBufferPool `json:"buffer_pools,omitempty"`
	Classes        *JvmClasses               `json:"classes,omitempty"`
	Gc             *GarbageCollector         `json:"gc,omitempty"`
	Mem            *JvmMemoryStats           `json:"mem,omitempty"`
	Threads        *JvmThreads               `json:"threads,omitempty"`
	Timestamp      *int64                    `json:"timestamp,omitempty"`
	Uptime         *string                   `json:"uptime,omitempty"`
	UptimeInMillis *int64                    `json:"uptime_in_millis,omitempty"`
}

// JvmBuilder holds Jvm struct and provides a builder API.
type JvmBuilder struct {
	v *Jvm
}

// NewJvm provides a builder for the Jvm struct.
func NewJvmBuilder() *JvmBuilder {
	r := JvmBuilder{
		&Jvm{
			BufferPools: make(map[string]NodeBufferPool, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Jvm struct
func (rb *JvmBuilder) Build() Jvm {
	return *rb.v
}

func (rb *JvmBuilder) BufferPools(values map[string]*NodeBufferPoolBuilder) *JvmBuilder {
	tmp := make(map[string]NodeBufferPool, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.BufferPools = tmp
	return rb
}

func (rb *JvmBuilder) Classes(classes *JvmClassesBuilder) *JvmBuilder {
	v := classes.Build()
	rb.v.Classes = &v
	return rb
}

func (rb *JvmBuilder) Gc(gc *GarbageCollectorBuilder) *JvmBuilder {
	v := gc.Build()
	rb.v.Gc = &v
	return rb
}

func (rb *JvmBuilder) Mem(mem *JvmMemoryStatsBuilder) *JvmBuilder {
	v := mem.Build()
	rb.v.Mem = &v
	return rb
}

func (rb *JvmBuilder) Threads(threads *JvmThreadsBuilder) *JvmBuilder {
	v := threads.Build()
	rb.v.Threads = &v
	return rb
}

func (rb *JvmBuilder) Timestamp(timestamp int64) *JvmBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}

func (rb *JvmBuilder) Uptime(uptime string) *JvmBuilder {
	rb.v.Uptime = &uptime
	return rb
}

func (rb *JvmBuilder) UptimeInMillis(uptimeinmillis int64) *JvmBuilder {
	rb.v.UptimeInMillis = &uptimeinmillis
	return rb
}
