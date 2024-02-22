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

package follow

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package follow
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/ccr/follow/CreateFollowIndexRequest.ts#L25-L51
type Request struct {
	LeaderIndex                   *string        `json:"leader_index,omitempty"`
	MaxOutstandingReadRequests    *int64         `json:"max_outstanding_read_requests,omitempty"`
	MaxOutstandingWriteRequests   *int64         `json:"max_outstanding_write_requests,omitempty"`
	MaxReadRequestOperationCount  *int64         `json:"max_read_request_operation_count,omitempty"`
	MaxReadRequestSize            *string        `json:"max_read_request_size,omitempty"`
	MaxRetryDelay                 types.Duration `json:"max_retry_delay,omitempty"`
	MaxWriteBufferCount           *int64         `json:"max_write_buffer_count,omitempty"`
	MaxWriteBufferSize            *string        `json:"max_write_buffer_size,omitempty"`
	MaxWriteRequestOperationCount *int64         `json:"max_write_request_operation_count,omitempty"`
	MaxWriteRequestSize           *string        `json:"max_write_request_size,omitempty"`
	ReadPollTimeout               types.Duration `json:"read_poll_timeout,omitempty"`
	RemoteCluster                 *string        `json:"remote_cluster,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Follow request: %w", err)
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

		case "leader_index":
			if err := dec.Decode(&s.LeaderIndex); err != nil {
				return err
			}

		case "max_outstanding_read_requests":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MaxOutstandingReadRequests = &value
			case float64:
				f := int64(v)
				s.MaxOutstandingReadRequests = &f
			}

		case "max_outstanding_write_requests":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MaxOutstandingWriteRequests = &value
			case float64:
				f := int64(v)
				s.MaxOutstandingWriteRequests = &f
			}

		case "max_read_request_operation_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MaxReadRequestOperationCount = &value
			case float64:
				f := int64(v)
				s.MaxReadRequestOperationCount = &f
			}

		case "max_read_request_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxReadRequestSize = &o

		case "max_retry_delay":
			if err := dec.Decode(&s.MaxRetryDelay); err != nil {
				return err
			}

		case "max_write_buffer_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MaxWriteBufferCount = &value
			case float64:
				f := int64(v)
				s.MaxWriteBufferCount = &f
			}

		case "max_write_buffer_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxWriteBufferSize = &o

		case "max_write_request_operation_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.MaxWriteRequestOperationCount = &value
			case float64:
				f := int64(v)
				s.MaxWriteRequestOperationCount = &f
			}

		case "max_write_request_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxWriteRequestSize = &o

		case "read_poll_timeout":
			if err := dec.Decode(&s.ReadPollTimeout); err != nil {
				return err
			}

		case "remote_cluster":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RemoteCluster = &o

		}
	}
	return nil
}
