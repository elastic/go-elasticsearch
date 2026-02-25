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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package catthreadpoolcolumn
package catthreadpoolcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/cat/_types/CatBase.ts#L3210-L3310
type CatThreadPoolColumn struct {
	Name string
}

var (

	// Active Number of active threads in the current thread pool.
	Active = CatThreadPoolColumn{"active"}

	// Completed Number of tasks completed by the thread pool executor.
	Completed = CatThreadPoolColumn{"completed"}

	// Core Configured core number of active threads allowed in the current thread pool.
	Core = CatThreadPoolColumn{"core"}

	// Ephemeralid Ephemeral node ID.
	Ephemeralid = CatThreadPoolColumn{"ephemeral_id"}

	// Host Hostname for the current node.
	Host = CatThreadPoolColumn{"host"}

	// Ip IP address for the current node.
	Ip = CatThreadPoolColumn{"ip"}

	// Keepalive Configured keep alive time for threads.
	Keepalive = CatThreadPoolColumn{"keep_alive"}

	// Largest Highest number of active threads in the current thread pool.
	Largest = CatThreadPoolColumn{"largest"}

	// Max Configured maximum number of active threads allowed in the current thread
	// pool.
	Max = CatThreadPoolColumn{"max"}

	// Name Name of the thread pool, such as `analyze` or `generic`.
	Name = CatThreadPoolColumn{"name"}

	// Nodeid ID of the node, such as `k0zy`.
	Nodeid = CatThreadPoolColumn{"node_id"}

	// Nodename Node name, such as `I8hydUG`.
	Nodename = CatThreadPoolColumn{"node_name"}

	// Pid Process ID of the running node.
	Pid = CatThreadPoolColumn{"pid"}

	// Poolsize Number of threads in the current thread pool.
	Poolsize = CatThreadPoolColumn{"pool_size"}

	// Port Bound transport port for the current node.
	Port = CatThreadPoolColumn{"port"}

	// Queue Number of tasks in the queue for the current thread pool.
	Queue = CatThreadPoolColumn{"queue"}

	// Queuesize Maximum number of tasks permitted in the queue for the current thread pool.
	Queuesize = CatThreadPoolColumn{"queue_size"}

	// Rejected Number of tasks rejected by the thread pool executor.
	Rejected = CatThreadPoolColumn{"rejected"}

	// Size Configured fixed number of active threads allowed in the current thread pool.
	Size = CatThreadPoolColumn{"size"}

	// Type Type of thread pool. Returned values are `fixed`, `fixed_auto_queue_size`,
	// `direct`, or `scaling`.
	Type = CatThreadPoolColumn{"type"}
)

func (c CatThreadPoolColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatThreadPoolColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "active":
		*c = Active
	case "completed":
		*c = Completed
	case "core":
		*c = Core
	case "ephemeral_id":
		*c = Ephemeralid
	case "host":
		*c = Host
	case "ip":
		*c = Ip
	case "keep_alive":
		*c = Keepalive
	case "largest":
		*c = Largest
	case "max":
		*c = Max
	case "name":
		*c = Name
	case "node_id":
		*c = Nodeid
	case "node_name":
		*c = Nodename
	case "pid":
		*c = Pid
	case "pool_size":
		*c = Poolsize
	case "port":
		*c = Port
	case "queue":
		*c = Queue
	case "queue_size":
		*c = Queuesize
	case "rejected":
		*c = Rejected
	case "size":
		*c = Size
	case "type":
		*c = Type
	default:
		*c = CatThreadPoolColumn{string(text)}
	}

	return nil
}

func (c CatThreadPoolColumn) String() string {
	return c.Name
}
