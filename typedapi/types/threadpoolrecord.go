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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// ThreadPoolRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/thread_pool/types.ts#L22-L123
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
	NodeId *string `json:"node_id,omitempty"`
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

// NewThreadPoolRecord returns a ThreadPoolRecord.
func NewThreadPoolRecord() *ThreadPoolRecord {
	r := &ThreadPoolRecord{}

	return r
}
