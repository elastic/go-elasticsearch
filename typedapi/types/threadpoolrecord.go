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

// ThreadPoolRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/thread_pool/types.ts#L22-L123
type ThreadPoolRecord struct {
	// Active number of active threads
	Active *string `json:"active,omitempty"`
	// Completed number of completed tasks
	Completed *string `json:"completed,omitempty"`
	// Core core number of threads in a scaling thread pool
	Core string `json:"core,omitempty"`
	// EphemeralNodeId ephemeral node id
	EphemeralNodeId *string `json:"ephemeral_node_id,omitempty"`
	// Host host name
	Host *string `json:"host,omitempty"`
	// Ip ip address
	Ip *string `json:"ip,omitempty"`
	// KeepAlive thread keep alive time
	KeepAlive string `json:"keep_alive,omitempty"`
	// Largest highest number of seen active threads
	Largest *string `json:"largest,omitempty"`
	// Max maximum number of threads in a scaling thread pool
	Max string `json:"max,omitempty"`
	// Name thread pool name
	Name *string `json:"name,omitempty"`
	// NodeId persistent node id
	NodeId *NodeId `json:"node_id,omitempty"`
	// NodeName node name
	NodeName *string `json:"node_name,omitempty"`
	// Pid process id
	Pid *string `json:"pid,omitempty"`
	// PoolSize number of threads
	PoolSize *string `json:"pool_size,omitempty"`
	// Port bound transport port
	Port *string `json:"port,omitempty"`
	// Queue number of tasks currently in queue
	Queue *string `json:"queue,omitempty"`
	// QueueSize maximum number of tasks permitted in queue
	QueueSize *string `json:"queue_size,omitempty"`
	// Rejected number of rejected tasks
	Rejected *string `json:"rejected,omitempty"`
	// Size number of threads in a fixed thread pool
	Size string `json:"size,omitempty"`
	// Type thread pool type
	Type *string `json:"type,omitempty"`
}

// ThreadPoolRecordBuilder holds ThreadPoolRecord struct and provides a builder API.
type ThreadPoolRecordBuilder struct {
	v *ThreadPoolRecord
}

// NewThreadPoolRecord provides a builder for the ThreadPoolRecord struct.
func NewThreadPoolRecordBuilder() *ThreadPoolRecordBuilder {
	r := ThreadPoolRecordBuilder{
		&ThreadPoolRecord{},
	}

	return &r
}

// Build finalize the chain and returns the ThreadPoolRecord struct
func (rb *ThreadPoolRecordBuilder) Build() ThreadPoolRecord {
	return *rb.v
}

// Active number of active threads

func (rb *ThreadPoolRecordBuilder) Active(active string) *ThreadPoolRecordBuilder {
	rb.v.Active = &active
	return rb
}

// Completed number of completed tasks

func (rb *ThreadPoolRecordBuilder) Completed(completed string) *ThreadPoolRecordBuilder {
	rb.v.Completed = &completed
	return rb
}

// Core core number of threads in a scaling thread pool

func (rb *ThreadPoolRecordBuilder) Core(core string) *ThreadPoolRecordBuilder {
	rb.v.Core = core
	return rb
}

// EphemeralNodeId ephemeral node id

func (rb *ThreadPoolRecordBuilder) EphemeralNodeId(ephemeralnodeid string) *ThreadPoolRecordBuilder {
	rb.v.EphemeralNodeId = &ephemeralnodeid
	return rb
}

// Host host name

func (rb *ThreadPoolRecordBuilder) Host(host string) *ThreadPoolRecordBuilder {
	rb.v.Host = &host
	return rb
}

// Ip ip address

func (rb *ThreadPoolRecordBuilder) Ip(ip string) *ThreadPoolRecordBuilder {
	rb.v.Ip = &ip
	return rb
}

// KeepAlive thread keep alive time

func (rb *ThreadPoolRecordBuilder) KeepAlive(keepalive string) *ThreadPoolRecordBuilder {
	rb.v.KeepAlive = keepalive
	return rb
}

// Largest highest number of seen active threads

func (rb *ThreadPoolRecordBuilder) Largest(largest string) *ThreadPoolRecordBuilder {
	rb.v.Largest = &largest
	return rb
}

// Max maximum number of threads in a scaling thread pool

func (rb *ThreadPoolRecordBuilder) Max(max string) *ThreadPoolRecordBuilder {
	rb.v.Max = max
	return rb
}

// Name thread pool name

func (rb *ThreadPoolRecordBuilder) Name(name string) *ThreadPoolRecordBuilder {
	rb.v.Name = &name
	return rb
}

// NodeId persistent node id

func (rb *ThreadPoolRecordBuilder) NodeId(nodeid NodeId) *ThreadPoolRecordBuilder {
	rb.v.NodeId = &nodeid
	return rb
}

// NodeName node name

func (rb *ThreadPoolRecordBuilder) NodeName(nodename string) *ThreadPoolRecordBuilder {
	rb.v.NodeName = &nodename
	return rb
}

// Pid process id

func (rb *ThreadPoolRecordBuilder) Pid(pid string) *ThreadPoolRecordBuilder {
	rb.v.Pid = &pid
	return rb
}

// PoolSize number of threads

func (rb *ThreadPoolRecordBuilder) PoolSize(poolsize string) *ThreadPoolRecordBuilder {
	rb.v.PoolSize = &poolsize
	return rb
}

// Port bound transport port

func (rb *ThreadPoolRecordBuilder) Port(port string) *ThreadPoolRecordBuilder {
	rb.v.Port = &port
	return rb
}

// Queue number of tasks currently in queue

func (rb *ThreadPoolRecordBuilder) Queue(queue string) *ThreadPoolRecordBuilder {
	rb.v.Queue = &queue
	return rb
}

// QueueSize maximum number of tasks permitted in queue

func (rb *ThreadPoolRecordBuilder) QueueSize(queuesize string) *ThreadPoolRecordBuilder {
	rb.v.QueueSize = &queuesize
	return rb
}

// Rejected number of rejected tasks

func (rb *ThreadPoolRecordBuilder) Rejected(rejected string) *ThreadPoolRecordBuilder {
	rb.v.Rejected = &rejected
	return rb
}

// Size number of threads in a fixed thread pool

func (rb *ThreadPoolRecordBuilder) Size(size string) *ThreadPoolRecordBuilder {
	rb.v.Size = size
	return rb
}

// Type thread pool type

func (rb *ThreadPoolRecordBuilder) Type_(type_ string) *ThreadPoolRecordBuilder {
	rb.v.Type = &type_
	return rb
}
