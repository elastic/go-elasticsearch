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

// CcrShardStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ccr/_types/FollowIndexStats.ts#L35-L69
type CcrShardStats struct {
	BytesRead                     int64           `json:"bytes_read"`
	FailedReadRequests            int64           `json:"failed_read_requests"`
	FailedWriteRequests           int64           `json:"failed_write_requests"`
	FatalException                *ErrorCause     `json:"fatal_exception,omitempty"`
	FollowerAliasesVersion        int64           `json:"follower_aliases_version"`
	FollowerGlobalCheckpoint      int64           `json:"follower_global_checkpoint"`
	FollowerIndex                 string          `json:"follower_index"`
	FollowerMappingVersion        int64           `json:"follower_mapping_version"`
	FollowerMaxSeqNo              int64           `json:"follower_max_seq_no"`
	FollowerSettingsVersion       int64           `json:"follower_settings_version"`
	LastRequestedSeqNo            int64           `json:"last_requested_seq_no"`
	LeaderGlobalCheckpoint        int64           `json:"leader_global_checkpoint"`
	LeaderIndex                   string          `json:"leader_index"`
	LeaderMaxSeqNo                int64           `json:"leader_max_seq_no"`
	OperationsRead                int64           `json:"operations_read"`
	OperationsWritten             int64           `json:"operations_written"`
	OutstandingReadRequests       int             `json:"outstanding_read_requests"`
	OutstandingWriteRequests      int             `json:"outstanding_write_requests"`
	ReadExceptions                []ReadException `json:"read_exceptions"`
	RemoteCluster                 string          `json:"remote_cluster"`
	ShardId                       int             `json:"shard_id"`
	SuccessfulReadRequests        int64           `json:"successful_read_requests"`
	SuccessfulWriteRequests       int64           `json:"successful_write_requests"`
	TimeSinceLastRead             Duration        `json:"time_since_last_read,omitempty"`
	TimeSinceLastReadMillis       int64           `json:"time_since_last_read_millis"`
	TotalReadRemoteExecTime       Duration        `json:"total_read_remote_exec_time,omitempty"`
	TotalReadRemoteExecTimeMillis int64           `json:"total_read_remote_exec_time_millis"`
	TotalReadTime                 Duration        `json:"total_read_time,omitempty"`
	TotalReadTimeMillis           int64           `json:"total_read_time_millis"`
	TotalWriteTime                Duration        `json:"total_write_time,omitempty"`
	TotalWriteTimeMillis          int64           `json:"total_write_time_millis"`
	WriteBufferOperationCount     int64           `json:"write_buffer_operation_count"`
	WriteBufferSizeInBytes        ByteSize        `json:"write_buffer_size_in_bytes"`
}

func (s *CcrShardStats) UnmarshalJSON(data []byte) error {

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

		case "bytes_read":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.BytesRead = value
			case float64:
				f := int64(v)
				s.BytesRead = f
			}

		case "failed_read_requests":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.FailedReadRequests = value
			case float64:
				f := int64(v)
				s.FailedReadRequests = f
			}

		case "failed_write_requests":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.FailedWriteRequests = value
			case float64:
				f := int64(v)
				s.FailedWriteRequests = f
			}

		case "fatal_exception":
			if err := dec.Decode(&s.FatalException); err != nil {
				return err
			}

		case "follower_aliases_version":
			if err := dec.Decode(&s.FollowerAliasesVersion); err != nil {
				return err
			}

		case "follower_global_checkpoint":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.FollowerGlobalCheckpoint = value
			case float64:
				f := int64(v)
				s.FollowerGlobalCheckpoint = f
			}

		case "follower_index":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FollowerIndex = o

		case "follower_mapping_version":
			if err := dec.Decode(&s.FollowerMappingVersion); err != nil {
				return err
			}

		case "follower_max_seq_no":
			if err := dec.Decode(&s.FollowerMaxSeqNo); err != nil {
				return err
			}

		case "follower_settings_version":
			if err := dec.Decode(&s.FollowerSettingsVersion); err != nil {
				return err
			}

		case "last_requested_seq_no":
			if err := dec.Decode(&s.LastRequestedSeqNo); err != nil {
				return err
			}

		case "leader_global_checkpoint":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LeaderGlobalCheckpoint = value
			case float64:
				f := int64(v)
				s.LeaderGlobalCheckpoint = f
			}

		case "leader_index":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LeaderIndex = o

		case "leader_max_seq_no":
			if err := dec.Decode(&s.LeaderMaxSeqNo); err != nil {
				return err
			}

		case "operations_read":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.OperationsRead = value
			case float64:
				f := int64(v)
				s.OperationsRead = f
			}

		case "operations_written":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.OperationsWritten = value
			case float64:
				f := int64(v)
				s.OperationsWritten = f
			}

		case "outstanding_read_requests":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.OutstandingReadRequests = value
			case float64:
				f := int(v)
				s.OutstandingReadRequests = f
			}

		case "outstanding_write_requests":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.OutstandingWriteRequests = value
			case float64:
				f := int(v)
				s.OutstandingWriteRequests = f
			}

		case "read_exceptions":
			if err := dec.Decode(&s.ReadExceptions); err != nil {
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
			s.RemoteCluster = o

		case "shard_id":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ShardId = value
			case float64:
				f := int(v)
				s.ShardId = f
			}

		case "successful_read_requests":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SuccessfulReadRequests = value
			case float64:
				f := int64(v)
				s.SuccessfulReadRequests = f
			}

		case "successful_write_requests":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.SuccessfulWriteRequests = value
			case float64:
				f := int64(v)
				s.SuccessfulWriteRequests = f
			}

		case "time_since_last_read":
			if err := dec.Decode(&s.TimeSinceLastRead); err != nil {
				return err
			}

		case "time_since_last_read_millis":
			if err := dec.Decode(&s.TimeSinceLastReadMillis); err != nil {
				return err
			}

		case "total_read_remote_exec_time":
			if err := dec.Decode(&s.TotalReadRemoteExecTime); err != nil {
				return err
			}

		case "total_read_remote_exec_time_millis":
			if err := dec.Decode(&s.TotalReadRemoteExecTimeMillis); err != nil {
				return err
			}

		case "total_read_time":
			if err := dec.Decode(&s.TotalReadTime); err != nil {
				return err
			}

		case "total_read_time_millis":
			if err := dec.Decode(&s.TotalReadTimeMillis); err != nil {
				return err
			}

		case "total_write_time":
			if err := dec.Decode(&s.TotalWriteTime); err != nil {
				return err
			}

		case "total_write_time_millis":
			if err := dec.Decode(&s.TotalWriteTimeMillis); err != nil {
				return err
			}

		case "write_buffer_operation_count":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.WriteBufferOperationCount = value
			case float64:
				f := int64(v)
				s.WriteBufferOperationCount = f
			}

		case "write_buffer_size_in_bytes":
			if err := dec.Decode(&s.WriteBufferSizeInBytes); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewCcrShardStats returns a CcrShardStats.
func NewCcrShardStats() *CcrShardStats {
	r := &CcrShardStats{}

	return r
}
