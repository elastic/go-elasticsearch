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

// NodeInfoHttp type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/nodes/info/types.ts#L298-L303
type NodeInfoHttp struct {
	BoundAddress            []string `json:"bound_address"`
	MaxContentLength        ByteSize `json:"max_content_length,omitempty"`
	MaxContentLengthInBytes int64    `json:"max_content_length_in_bytes"`
	PublishAddress          string   `json:"publish_address"`
}

func (s *NodeInfoHttp) UnmarshalJSON(data []byte) error {

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

		case "bound_address":
			if err := dec.Decode(&s.BoundAddress); err != nil {
				return err
			}

		case "max_content_length":
			if err := dec.Decode(&s.MaxContentLength); err != nil {
				return err
			}

		case "max_content_length_in_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MaxContentLengthInBytes = value
			case float64:
				f := int64(v)
				s.MaxContentLengthInBytes = f
			}

		case "publish_address":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PublishAddress = o

		}
	}
	return nil
}

// NewNodeInfoHttp returns a NodeInfoHttp.
func NewNodeInfoHttp() *NodeInfoHttp {
	r := &NodeInfoHttp{}

	return r
}
