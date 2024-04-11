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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// HotThread type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/nodes/hot_threads/types.ts#L23-L28
type HotThread struct {
	Hosts    []string `json:"hosts"`
	NodeId   string   `json:"node_id"`
	NodeName string   `json:"node_name"`
	Threads  []string `json:"threads"`
}

func (s *HotThread) UnmarshalJSON(data []byte) error {

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

		case "hosts":
			if err := dec.Decode(&s.Hosts); err != nil {
				return fmt.Errorf("%s | %w", "Hosts", err)
			}

		case "node_id":
			if err := dec.Decode(&s.NodeId); err != nil {
				return fmt.Errorf("%s | %w", "NodeId", err)
			}

		case "node_name":
			if err := dec.Decode(&s.NodeName); err != nil {
				return fmt.Errorf("%s | %w", "NodeName", err)
			}

		case "threads":
			if err := dec.Decode(&s.Threads); err != nil {
				return fmt.Errorf("%s | %w", "Threads", err)
			}

		}
	}
	return nil
}

// NewHotThread returns a HotThread.
func NewHotThread() *HotThread {
	r := &HotThread{}

	return r
}
