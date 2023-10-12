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
// https://github.com/elastic/elasticsearch-specification/tree/3b09f9d8e90178243f8a340a7bc324aab152c602

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Jvm type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3b09f9d8e90178243f8a340a7bc324aab152c602/specification/nodes/_types/Stats.ts#L324-L333
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

func (s *Jvm) UnmarshalJSON(data []byte) error {

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

		case "buffer_pools":
			if s.BufferPools == nil {
				s.BufferPools = make(map[string]NodeBufferPool, 0)
			}
			if err := dec.Decode(&s.BufferPools); err != nil {
				return err
			}

		case "classes":
			if err := dec.Decode(&s.Classes); err != nil {
				return err
			}

		case "gc":
			if err := dec.Decode(&s.Gc); err != nil {
				return err
			}

		case "mem":
			if err := dec.Decode(&s.Mem); err != nil {
				return err
			}

		case "threads":
			if err := dec.Decode(&s.Threads); err != nil {
				return err
			}

		case "timestamp":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Timestamp = &value
			case float64:
				f := int64(v)
				s.Timestamp = &f
			}

		case "uptime":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Uptime = &o

		case "uptime_in_millis":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.UptimeInMillis = &value
			case float64:
				f := int64(v)
				s.UptimeInMillis = &f
			}

		}
	}
	return nil
}

// NewJvm returns a Jvm.
func NewJvm() *Jvm {
	r := &Jvm{
		BufferPools: make(map[string]NodeBufferPool, 0),
	}

	return r
}
