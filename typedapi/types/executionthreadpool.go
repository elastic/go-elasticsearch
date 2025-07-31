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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ExecutionThreadPool type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/watcher/_types/Execution.ts#L94-L103
type ExecutionThreadPool struct {
	// MaxSize The largest size of the execution thread pool, which indicates the largest
	// number of concurrent running watches.
	MaxSize int64 `json:"max_size"`
	// QueueSize The number of watches that were triggered and are currently queued.
	QueueSize int64 `json:"queue_size"`
}

func (s *ExecutionThreadPool) UnmarshalJSON(data []byte) error {

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

		case "max_size":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxSize", err)
				}
				s.MaxSize = value
			case float64:
				f := int64(v)
				s.MaxSize = f
			}

		case "queue_size":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "QueueSize", err)
				}
				s.QueueSize = value
			case float64:
				f := int64(v)
				s.QueueSize = f
			}

		}
	}
	return nil
}

// NewExecutionThreadPool returns a ExecutionThreadPool.
func NewExecutionThreadPool() *ExecutionThreadPool {
	r := &ExecutionThreadPool{}

	return r
}
