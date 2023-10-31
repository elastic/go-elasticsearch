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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ThreadPoolRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cat/thread_pool/types.ts#L22-L124
type ThreadPoolRecord struct {
	// Active The number of active threads in the current thread pool.
	Active *string `json:"active,omitempty"`
	// Completed The number of completed tasks.
	Completed *string `json:"completed,omitempty"`
	// Core The core number of active threads allowed in a scaling thread pool.
	Core string `json:"core,omitempty"`
	// EphemeralNodeId The ephemeral node identifier.
	EphemeralNodeId *string `json:"ephemeral_node_id,omitempty"`
	// Host The host name for the current node.
	Host *string `json:"host,omitempty"`
	// Ip The IP address for the current node.
	Ip *string `json:"ip,omitempty"`
	// KeepAlive The thread keep alive time.
	KeepAlive string `json:"keep_alive,omitempty"`
	// Largest The highest number of active threads in the current thread pool.
	Largest *string `json:"largest,omitempty"`
	// Max The maximum number of active threads allowed in a scaling thread pool.
	Max string `json:"max,omitempty"`
	// Name The thread pool name.
	Name *string `json:"name,omitempty"`
	// NodeId The persistent node identifier.
	NodeId *string `json:"node_id,omitempty"`
	// NodeName The node name.
	NodeName *string `json:"node_name,omitempty"`
	// Pid The process identifier.
	Pid *string `json:"pid,omitempty"`
	// PoolSize The number of threads in the current thread pool.
	PoolSize *string `json:"pool_size,omitempty"`
	// Port The bound transport port for the current node.
	Port *string `json:"port,omitempty"`
	// Queue The number of tasks currently in queue.
	Queue *string `json:"queue,omitempty"`
	// QueueSize The maximum number of tasks permitted in the queue.
	QueueSize *string `json:"queue_size,omitempty"`
	// Rejected The number of rejected tasks.
	Rejected *string `json:"rejected,omitempty"`
	// Size The number of active threads allowed in a fixed thread pool.
	Size string `json:"size,omitempty"`
	// Type The thread pool type.
	// Returned values include `fixed`, `fixed_auto_queue_size`, `direct`, and
	// `scaling`.
	Type *string `json:"type,omitempty"`
}

func (s *ThreadPoolRecord) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "active", "a":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Active = &o

		case "completed", "c":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Completed = &o

		case "core", "cr":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Core = o

		case "ephemeral_node_id", "eid":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.EphemeralNodeId = &o

		case "host", "h":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Host = &o

		case "ip", "i":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Ip = &o

		case "keep_alive", "ka":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.KeepAlive = o

		case "largest", "l":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Largest = &o

		case "max", "mx":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Max = o

		case "name", "n":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = &o

		case "node_id", "id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return err
			}

		case "node_name", "nn":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.NodeName = &o

		case "pid", "p":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pid = &o

		case "pool_size", "psz":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PoolSize = &o

		case "port", "po":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Port = &o

		case "queue", "q":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Queue = &o

		case "queue_size", "qs":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.QueueSize = &o

		case "rejected", "r":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Rejected = &o

		case "size", "sz":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Size = o

		case "type", "t":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Type = &o

		}
	}
	return nil
}

// NewThreadPoolRecord returns a ThreadPoolRecord.
func NewThreadPoolRecord() *ThreadPoolRecord {
	r := &ThreadPoolRecord{}

	return r
}
