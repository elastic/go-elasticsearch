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

package putautofollowpattern

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putautofollowpattern
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ccr/put_auto_follow_pattern/PutAutoFollowPatternRequest.ts#L27-L133
type Request struct {

	// FollowIndexPattern The name of follower index. The template {{leader_index}} can be used to
	// derive the name of the follower index from the name of the leader index. When
	// following a data stream, use {{leader_index}}; CCR does not support changes
	// to the names of a follower data stream’s backing indices.
	FollowIndexPattern *string `json:"follow_index_pattern,omitempty"`
	// LeaderIndexExclusionPatterns An array of simple index patterns that can be used to exclude indices from
	// being auto-followed. Indices in the remote cluster whose names are matching
	// one or more leader_index_patterns and one or more
	// leader_index_exclusion_patterns won’t be followed.
	LeaderIndexExclusionPatterns []string `json:"leader_index_exclusion_patterns,omitempty"`
	// LeaderIndexPatterns An array of simple index patterns to match against indices in the remote
	// cluster specified by the remote_cluster field.
	LeaderIndexPatterns []string `json:"leader_index_patterns,omitempty"`
	// MaxOutstandingReadRequests The maximum number of outstanding reads requests from the remote cluster.
	MaxOutstandingReadRequests *int `json:"max_outstanding_read_requests,omitempty"`
	// MaxOutstandingWriteRequests The maximum number of outstanding reads requests from the remote cluster.
	MaxOutstandingWriteRequests *int `json:"max_outstanding_write_requests,omitempty"`
	// MaxReadRequestOperationCount The maximum number of operations to pull per read from the remote cluster.
	MaxReadRequestOperationCount *int `json:"max_read_request_operation_count,omitempty"`
	// MaxReadRequestSize The maximum size in bytes of per read of a batch of operations pulled from
	// the remote cluster.
	MaxReadRequestSize types.ByteSize `json:"max_read_request_size,omitempty"`
	// MaxRetryDelay The maximum time to wait before retrying an operation that failed
	// exceptionally. An exponential backoff strategy is employed when retrying.
	MaxRetryDelay types.Duration `json:"max_retry_delay,omitempty"`
	// MaxWriteBufferCount The maximum number of operations that can be queued for writing. When this
	// limit is reached, reads from the remote cluster will be deferred until the
	// number of queued operations goes below the limit.
	MaxWriteBufferCount *int `json:"max_write_buffer_count,omitempty"`
	// MaxWriteBufferSize The maximum total bytes of operations that can be queued for writing. When
	// this limit is reached, reads from the remote cluster will be deferred until
	// the total bytes of queued operations goes below the limit.
	MaxWriteBufferSize types.ByteSize `json:"max_write_buffer_size,omitempty"`
	// MaxWriteRequestOperationCount The maximum number of operations per bulk write request executed on the
	// follower.
	MaxWriteRequestOperationCount *int `json:"max_write_request_operation_count,omitempty"`
	// MaxWriteRequestSize The maximum total bytes of operations per bulk write request executed on the
	// follower.
	MaxWriteRequestSize types.ByteSize `json:"max_write_request_size,omitempty"`
	// ReadPollTimeout The maximum time to wait for new operations on the remote cluster when the
	// follower index is synchronized with the leader index. When the timeout has
	// elapsed, the poll for operations will return to the follower so that it can
	// update some statistics. Then the follower will immediately attempt to read
	// from the leader again.
	ReadPollTimeout types.Duration `json:"read_poll_timeout,omitempty"`
	// RemoteCluster The remote cluster containing the leader indices to match against.
	RemoteCluster string `json:"remote_cluster"`
	// Settings Settings to override from the leader index. Note that certain settings can
	// not be overrode (e.g., index.number_of_shards).
	Settings map[string]json.RawMessage `json:"settings,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Settings: make(map[string]json.RawMessage, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putautofollowpattern request: %w", err)
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

		case "follow_index_pattern":
			if err := dec.Decode(&s.FollowIndexPattern); err != nil {
				return fmt.Errorf("%s | %w", "FollowIndexPattern", err)
			}

		case "leader_index_exclusion_patterns":
			if err := dec.Decode(&s.LeaderIndexExclusionPatterns); err != nil {
				return fmt.Errorf("%s | %w", "LeaderIndexExclusionPatterns", err)
			}

		case "leader_index_patterns":
			if err := dec.Decode(&s.LeaderIndexPatterns); err != nil {
				return fmt.Errorf("%s | %w", "LeaderIndexPatterns", err)
			}

		case "max_outstanding_read_requests":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxOutstandingReadRequests", err)
				}
				s.MaxOutstandingReadRequests = &value
			case float64:
				f := int(v)
				s.MaxOutstandingReadRequests = &f
			}

		case "max_outstanding_write_requests":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxOutstandingWriteRequests", err)
				}
				s.MaxOutstandingWriteRequests = &value
			case float64:
				f := int(v)
				s.MaxOutstandingWriteRequests = &f
			}

		case "max_read_request_operation_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxReadRequestOperationCount", err)
				}
				s.MaxReadRequestOperationCount = &value
			case float64:
				f := int(v)
				s.MaxReadRequestOperationCount = &f
			}

		case "max_read_request_size":
			if err := dec.Decode(&s.MaxReadRequestSize); err != nil {
				return fmt.Errorf("%s | %w", "MaxReadRequestSize", err)
			}

		case "max_retry_delay":
			if err := dec.Decode(&s.MaxRetryDelay); err != nil {
				return fmt.Errorf("%s | %w", "MaxRetryDelay", err)
			}

		case "max_write_buffer_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxWriteBufferCount", err)
				}
				s.MaxWriteBufferCount = &value
			case float64:
				f := int(v)
				s.MaxWriteBufferCount = &f
			}

		case "max_write_buffer_size":
			if err := dec.Decode(&s.MaxWriteBufferSize); err != nil {
				return fmt.Errorf("%s | %w", "MaxWriteBufferSize", err)
			}

		case "max_write_request_operation_count":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxWriteRequestOperationCount", err)
				}
				s.MaxWriteRequestOperationCount = &value
			case float64:
				f := int(v)
				s.MaxWriteRequestOperationCount = &f
			}

		case "max_write_request_size":
			if err := dec.Decode(&s.MaxWriteRequestSize); err != nil {
				return fmt.Errorf("%s | %w", "MaxWriteRequestSize", err)
			}

		case "read_poll_timeout":
			if err := dec.Decode(&s.ReadPollTimeout); err != nil {
				return fmt.Errorf("%s | %w", "ReadPollTimeout", err)
			}

		case "remote_cluster":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RemoteCluster", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RemoteCluster = o

		case "settings":
			if s.Settings == nil {
				s.Settings = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Settings); err != nil {
				return fmt.Errorf("%s | %w", "Settings", err)
			}

		}
	}
	return nil
}
