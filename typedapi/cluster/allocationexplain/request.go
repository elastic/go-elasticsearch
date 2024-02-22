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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package allocationexplain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Request holds the request body struct for the package allocationexplain
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/cluster/allocation_explain/ClusterAllocationExplainRequest.ts#L24-L61
type Request struct {

	// CurrentNode Specifies the node ID or the name of the node to only explain a shard that is
	// currently located on the specified node.
	CurrentNode *string `json:"current_node,omitempty"`
	// Index Specifies the name of the index that you would like an explanation for.
	Index *string `json:"index,omitempty"`
	// Primary If true, returns explanation for the primary shard for the given shard ID.
	Primary *bool `json:"primary,omitempty"`
	// Shard Specifies the ID of the shard that you would like an explanation for.
	Shard *int `json:"shard,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Allocationexplain request: %w", err)
	}

	return &req, nil
}

func (s *Request) UnmarshalJSON(data []byte) error {
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

		case "current_node":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CurrentNode = &o

		case "index":
			if err := dec.Decode(&s.Index); err != nil {
				return err
			}

		case "primary":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Primary = &value
			case bool:
				s.Primary = &v
			}

		case "shard":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Shard = &value
			case float64:
				f := int(v)
				s.Shard = &f
			}

		}
	}
	return nil
}
