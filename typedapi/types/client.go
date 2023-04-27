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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// Client type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/nodes/_types/Stats.ts#L272-L284
type Client struct {
	Agent                 *string `json:"agent,omitempty"`
	ClosedTimeMillis      *int64  `json:"closed_time_millis,omitempty"`
	Id                    *int64  `json:"id,omitempty"`
	LastRequestTimeMillis *int64  `json:"last_request_time_millis,omitempty"`
	LastUri               *string `json:"last_uri,omitempty"`
	LocalAddress          *string `json:"local_address,omitempty"`
	OpenedTimeMillis      *int64  `json:"opened_time_millis,omitempty"`
	RemoteAddress         *string `json:"remote_address,omitempty"`
	RequestCount          *int64  `json:"request_count,omitempty"`
	RequestSizeBytes      *int64  `json:"request_size_bytes,omitempty"`
	XOpaqueId             *string `json:"x_opaque_id,omitempty"`
}

func (s *Client) UnmarshalJSON(data []byte) error {

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

		case "agent":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Agent = &o

		case "closed_time_millis":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.ClosedTimeMillis = &value
			case float64:
				f := int64(v)
				s.ClosedTimeMillis = &f
			}

		case "id":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Id = &value
			case float64:
				f := int64(v)
				s.Id = &f
			}

		case "last_request_time_millis":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LastRequestTimeMillis = &value
			case float64:
				f := int64(v)
				s.LastRequestTimeMillis = &f
			}

		case "last_uri":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.LastUri = &o

		case "local_address":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.LocalAddress = &o

		case "opened_time_millis":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.OpenedTimeMillis = &value
			case float64:
				f := int64(v)
				s.OpenedTimeMillis = &f
			}

		case "remote_address":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.RemoteAddress = &o

		case "request_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RequestCount = &value
			case float64:
				f := int64(v)
				s.RequestCount = &f
			}

		case "request_size_bytes":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.RequestSizeBytes = &value
			case float64:
				f := int64(v)
				s.RequestSizeBytes = &f
			}

		case "x_opaque_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.XOpaqueId = &o

		}
	}
	return nil
}

// NewClient returns a Client.
func NewClient() *Client {
	r := &Client{}

	return r
}
