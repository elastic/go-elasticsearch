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

package resumefollow

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package resumefollow
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ccr/resume_follow/ResumeFollowIndexRequest.ts#L25-L65
type Request struct {
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
		return nil, fmt.Errorf("could not deserialise json into Resumefollow request: %w", err)
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

		case "max_outstanding_read_requests":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxOutstandingReadRequests", err)
				}
				s.MaxOutstandingReadRequests = &value
			case float64:
				f := int64(v)
				s.MaxOutstandingReadRequests = &f
			}

		case "max_outstanding_write_requests":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxOutstandingWriteRequests", err)
				}
				s.MaxOutstandingWriteRequests = &value
			case float64:
				f := int64(v)
				s.MaxOutstandingWriteRequests = &f
			}

		case "max_read_request_operation_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxReadRequestOperationCount", err)
				}
				s.MaxReadRequestOperationCount = &value
			case float64:
				f := int64(v)
				s.MaxReadRequestOperationCount = &f
			}

		case "max_read_request_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MaxReadRequestSize", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxReadRequestSize = &o

		case "max_retry_delay":
			if err := dec.Decode(&s.MaxRetryDelay); err != nil {
				return fmt.Errorf("%s | %w", "MaxRetryDelay", err)
			}

		case "max_write_buffer_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxWriteBufferCount", err)
				}
				s.MaxWriteBufferCount = &value
			case float64:
				f := int64(v)
				s.MaxWriteBufferCount = &f
			}

		case "max_write_buffer_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MaxWriteBufferSize", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxWriteBufferSize = &o

		case "max_write_request_operation_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxWriteRequestOperationCount", err)
				}
				s.MaxWriteRequestOperationCount = &value
			case float64:
				f := int64(v)
				s.MaxWriteRequestOperationCount = &f
			}

		case "max_write_request_size":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "MaxWriteRequestSize", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.MaxWriteRequestSize = &o

		case "read_poll_timeout":
			if err := dec.Decode(&s.ReadPollTimeout); err != nil {
				return fmt.Errorf("%s | %w", "ReadPollTimeout", err)
			}

		}
	}
	return nil
}
