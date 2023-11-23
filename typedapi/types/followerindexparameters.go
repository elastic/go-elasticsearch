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

// FollowerIndexParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ccr/follow_info/types.ts#L38-L49
type FollowerIndexParameters struct {
	MaxOutstandingReadRequests    int      `json:"max_outstanding_read_requests"`
	MaxOutstandingWriteRequests   int      `json:"max_outstanding_write_requests"`
	MaxReadRequestOperationCount  int      `json:"max_read_request_operation_count"`
	MaxReadRequestSize            string   `json:"max_read_request_size"`
	MaxRetryDelay                 Duration `json:"max_retry_delay"`
	MaxWriteBufferCount           int      `json:"max_write_buffer_count"`
	MaxWriteBufferSize            string   `json:"max_write_buffer_size"`
	MaxWriteRequestOperationCount int      `json:"max_write_request_operation_count"`
	MaxWriteRequestSize           string   `json:"max_write_request_size"`
	ReadPollTimeout               Duration `json:"read_poll_timeout"`
}

func (s *FollowerIndexParameters) UnmarshalJSON(data []byte) error {

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

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxOutstandingReadRequests = value
			case float64:
				f := int(v)
				s.MaxOutstandingReadRequests = f
			}

		case "max_outstanding_write_requests":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxOutstandingWriteRequests = value
			case float64:
				f := int(v)
				s.MaxOutstandingWriteRequests = f
			}

		case "max_read_request_operation_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxReadRequestOperationCount = value
			case float64:
				f := int(v)
				s.MaxReadRequestOperationCount = f
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
			s.MaxReadRequestSize = o

		case "max_retry_delay":
			if err := dec.Decode(&s.MaxRetryDelay); err != nil {
				return err
			}

		case "max_write_buffer_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxWriteBufferCount = value
			case float64:
				f := int(v)
				s.MaxWriteBufferCount = f
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
			s.MaxWriteBufferSize = o

		case "max_write_request_operation_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxWriteRequestOperationCount = value
			case float64:
				f := int(v)
				s.MaxWriteRequestOperationCount = f
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
			s.MaxWriteRequestSize = o

		case "read_poll_timeout":
			if err := dec.Decode(&s.ReadPollTimeout); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewFollowerIndexParameters returns a FollowerIndexParameters.
func NewFollowerIndexParameters() *FollowerIndexParameters {
	r := &FollowerIndexParameters{}

	return r
}
