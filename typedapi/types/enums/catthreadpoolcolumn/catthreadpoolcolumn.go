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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package catthreadpoolcolumn
package catthreadpoolcolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/cat/_types/CatBase.ts#L1952-L2052
type CatThreadPoolColumn struct {
	Name string
}

var (
	Active = CatThreadPoolColumn{"active"}

	Completed = CatThreadPoolColumn{"completed"}

	Core = CatThreadPoolColumn{"core"}

	Ephemeralid = CatThreadPoolColumn{"ephemeral_id"}

	Host = CatThreadPoolColumn{"host"}

	Ip = CatThreadPoolColumn{"ip"}

	Keepalive = CatThreadPoolColumn{"keep_alive"}

	Largest = CatThreadPoolColumn{"largest"}

	Max = CatThreadPoolColumn{"max"}

	Name = CatThreadPoolColumn{"name"}

	Nodeid = CatThreadPoolColumn{"node_id"}

	Nodename = CatThreadPoolColumn{"node_name"}

	Pid = CatThreadPoolColumn{"pid"}

	Poolsize = CatThreadPoolColumn{"pool_size"}

	Port = CatThreadPoolColumn{"port"}

	Queue = CatThreadPoolColumn{"queue"}

	Queuesize = CatThreadPoolColumn{"queue_size"}

	Rejected = CatThreadPoolColumn{"rejected"}

	Size = CatThreadPoolColumn{"size"}

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
