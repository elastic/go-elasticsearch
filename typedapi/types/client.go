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

// Client type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/nodes/_types/Stats.ts#L649-L696
type Client struct {
	// Agent Reported agent for the HTTP client.
	// If unavailable, this property is not included in the response.
	Agent *string `json:"agent,omitempty"`
	// ClosedTimeMillis Time at which the client closed the connection if the connection is closed.
	ClosedTimeMillis *int64 `json:"closed_time_millis,omitempty"`
	// Id Unique ID for the HTTP client.
	Id *int64 `json:"id,omitempty"`
	// LastRequestTimeMillis Time of the most recent request from this client.
	LastRequestTimeMillis *int64 `json:"last_request_time_millis,omitempty"`
	// LastUri The URI of the client’s most recent request.
	LastUri *string `json:"last_uri,omitempty"`
	// LocalAddress Local address for the HTTP connection.
	LocalAddress *string `json:"local_address,omitempty"`
	// OpenedTimeMillis Time at which the client opened the connection.
	OpenedTimeMillis *int64 `json:"opened_time_millis,omitempty"`
	// RemoteAddress Remote address for the HTTP connection.
	RemoteAddress *string `json:"remote_address,omitempty"`
	// RequestCount Number of requests from this client.
	RequestCount *int64 `json:"request_count,omitempty"`
	// RequestSizeBytes Cumulative size in bytes of all requests from this client.
	RequestSizeBytes *int64 `json:"request_size_bytes,omitempty"`
	// XOpaqueId Value from the client’s `x-opaque-id` HTTP header.
	// If unavailable, this property is not included in the response.
	XOpaqueId *string `json:"x_opaque_id,omitempty"`
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
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
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
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LastUri = &o

		case "local_address":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
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
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
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
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
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
