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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// GarbageCollectorTotal type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/nodes/_types/Stats.ts#L930-L943
type GarbageCollectorTotal struct {
	// CollectionCount Total number of JVM garbage collectors that collect objects.
	CollectionCount *int64 `json:"collection_count,omitempty"`
	// CollectionTime Total time spent by JVM collecting objects.
	CollectionTime *string `json:"collection_time,omitempty"`
	// CollectionTimeInMillis Total time, in milliseconds, spent by JVM collecting objects.
	CollectionTimeInMillis *int64 `json:"collection_time_in_millis,omitempty"`
}

func (s *GarbageCollectorTotal) UnmarshalJSON(data []byte) error {

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

		case "collection_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CollectionCount = &value
			case float64:
				f := int64(v)
				s.CollectionCount = &f
			}

		case "collection_time":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CollectionTime = &o

		case "collection_time_in_millis":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.CollectionTimeInMillis = &value
			case float64:
				f := int64(v)
				s.CollectionTimeInMillis = &f
			}

		}
	}
	return nil
}

// NewGarbageCollectorTotal returns a GarbageCollectorTotal.
func NewGarbageCollectorTotal() *GarbageCollectorTotal {
	r := &GarbageCollectorTotal{}

	return r
}
