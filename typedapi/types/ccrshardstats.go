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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CcrShardStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ccr/_types/FollowIndexStats.ts#L37-L109
type CcrShardStats struct {
	// BytesRead The total of transferred bytes read from the leader.
	// This is only an estimate and does not account for compression if enabled.
	BytesRead int64 `json:"bytes_read"`
	// FailedReadRequests The number of failed reads.
	FailedReadRequests int64 `json:"failed_read_requests"`
	// FailedWriteRequests The number of failed bulk write requests on the follower.
	FailedWriteRequests int64       `json:"failed_write_requests"`
	FatalException      *ErrorCause `json:"fatal_exception,omitempty"`
	// FollowerAliasesVersion The index aliases version the follower is synced up to.
	FollowerAliasesVersion int64 `json:"follower_aliases_version"`
	// FollowerGlobalCheckpoint The current global checkpoint on the follower.
	// The difference between the `leader_global_checkpoint` and the
	// `follower_global_checkpoint` is an indication of how much the follower is
	// lagging the leader.
	FollowerGlobalCheckpoint int64 `json:"follower_global_checkpoint"`
	// FollowerIndex The name of the follower index.
	FollowerIndex string `json:"follower_index"`
	// FollowerMappingVersion The mapping version the follower is synced up to.
	FollowerMappingVersion int64 `json:"follower_mapping_version"`
	// FollowerMaxSeqNo The current maximum sequence number on the follower.
	FollowerMaxSeqNo int64 `json:"follower_max_seq_no"`
	// FollowerSettingsVersion The index settings version the follower is synced up to.
	FollowerSettingsVersion int64 `json:"follower_settings_version"`
	// LastRequestedSeqNo The starting sequence number of the last batch of operations requested from
	// the leader.
	LastRequestedSeqNo int64 `json:"last_requested_seq_no"`
	// LeaderGlobalCheckpoint The current global checkpoint on the leader known to the follower task.
	LeaderGlobalCheckpoint int64 `json:"leader_global_checkpoint"`
	// LeaderIndex The name of the index in the leader cluster being followed.
	LeaderIndex string `json:"leader_index"`
	// LeaderMaxSeqNo The current maximum sequence number on the leader known to the follower task.
	LeaderMaxSeqNo int64 `json:"leader_max_seq_no"`
	// OperationsRead The total number of operations read from the leader.
	OperationsRead int64 `json:"operations_read"`
	// OperationsWritten The number of operations written on the follower.
	OperationsWritten int64 `json:"operations_written"`
	// OutstandingReadRequests The number of active read requests from the follower.
	OutstandingReadRequests int `json:"outstanding_read_requests"`
	// OutstandingWriteRequests The number of active bulk write requests on the follower.
	OutstandingWriteRequests int `json:"outstanding_write_requests"`
	// ReadExceptions An array of objects representing failed reads.
	ReadExceptions []ReadException `json:"read_exceptions"`
	// RemoteCluster The remote cluster containing the leader index.
	RemoteCluster string `json:"remote_cluster"`
	// ShardId The numerical shard ID, with values from 0 to one less than the number of
	// replicas.
	ShardId int `json:"shard_id"`
	// SuccessfulReadRequests The number of successful fetches.
	SuccessfulReadRequests int64 `json:"successful_read_requests"`
	// SuccessfulWriteRequests The number of bulk write requests run on the follower.
	SuccessfulWriteRequests int64    `json:"successful_write_requests"`
	TimeSinceLastRead       Duration `json:"time_since_last_read,omitempty"`
	// TimeSinceLastReadMillis The number of milliseconds since a read request was sent to the leader.
	// When the follower is caught up to the leader, this number will increase up to
	// the configured `read_poll_timeout` at which point another read request will
	// be sent to the leader.
	TimeSinceLastReadMillis int64    `json:"time_since_last_read_millis"`
	TotalReadRemoteExecTime Duration `json:"total_read_remote_exec_time,omitempty"`
	// TotalReadRemoteExecTimeMillis The total time reads spent running on the remote cluster.
	TotalReadRemoteExecTimeMillis int64    `json:"total_read_remote_exec_time_millis"`
	TotalReadTime                 Duration `json:"total_read_time,omitempty"`
	// TotalReadTimeMillis The total time reads were outstanding, measured from the time a read was sent
	// to the leader to the time a reply was returned to the follower.
	TotalReadTimeMillis int64    `json:"total_read_time_millis"`
	TotalWriteTime      Duration `json:"total_write_time,omitempty"`
	// TotalWriteTimeMillis The total time spent writing on the follower.
	TotalWriteTimeMillis int64 `json:"total_write_time_millis"`
	// WriteBufferOperationCount The number of write operations queued on the follower.
	WriteBufferOperationCount int64 `json:"write_buffer_operation_count"`
	// WriteBufferSizeInBytes The total number of bytes of operations currently queued for writing.
	WriteBufferSizeInBytes ByteSize `json:"write_buffer_size_in_bytes"`
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
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "BytesRead", err)
				}
				s.BytesRead = value
			case float64:
				f := int64(v)
				s.BytesRead = f
			}

		case "failed_read_requests":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "FailedReadRequests", err)
				}
				s.FailedReadRequests = value
			case float64:
				f := int64(v)
				s.FailedReadRequests = f
			}

		case "failed_write_requests":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "FailedWriteRequests", err)
				}
				s.FailedWriteRequests = value
			case float64:
				f := int64(v)
				s.FailedWriteRequests = f
			}

		case "fatal_exception":
			if err := dec.Decode(&s.FatalException); err != nil {
				return fmt.Errorf("%s | %w", "FatalException", err)
			}

		case "follower_aliases_version":
			if err := dec.Decode(&s.FollowerAliasesVersion); err != nil {
				return fmt.Errorf("%s | %w", "FollowerAliasesVersion", err)
			}

		case "follower_global_checkpoint":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "FollowerGlobalCheckpoint", err)
				}
				s.FollowerGlobalCheckpoint = value
			case float64:
				f := int64(v)
				s.FollowerGlobalCheckpoint = f
			}

		case "follower_index":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "FollowerIndex", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FollowerIndex = o

		case "follower_mapping_version":
			if err := dec.Decode(&s.FollowerMappingVersion); err != nil {
				return fmt.Errorf("%s | %w", "FollowerMappingVersion", err)
			}

		case "follower_max_seq_no":
			if err := dec.Decode(&s.FollowerMaxSeqNo); err != nil {
				return fmt.Errorf("%s | %w", "FollowerMaxSeqNo", err)
			}

		case "follower_settings_version":
			if err := dec.Decode(&s.FollowerSettingsVersion); err != nil {
				return fmt.Errorf("%s | %w", "FollowerSettingsVersion", err)
			}

		case "last_requested_seq_no":
			if err := dec.Decode(&s.LastRequestedSeqNo); err != nil {
				return fmt.Errorf("%s | %w", "LastRequestedSeqNo", err)
			}

		case "leader_global_checkpoint":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "LeaderGlobalCheckpoint", err)
				}
				s.LeaderGlobalCheckpoint = value
			case float64:
				f := int64(v)
				s.LeaderGlobalCheckpoint = f
			}

		case "leader_index":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "LeaderIndex", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LeaderIndex = o

		case "leader_max_seq_no":
			if err := dec.Decode(&s.LeaderMaxSeqNo); err != nil {
				return fmt.Errorf("%s | %w", "LeaderMaxSeqNo", err)
			}

		case "operations_read":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "OperationsRead", err)
				}
				s.OperationsRead = value
			case float64:
				f := int64(v)
				s.OperationsRead = f
			}

		case "operations_written":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "OperationsWritten", err)
				}
				s.OperationsWritten = value
			case float64:
				f := int64(v)
				s.OperationsWritten = f
			}

		case "outstanding_read_requests":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OutstandingReadRequests", err)
				}
				s.OutstandingReadRequests = value
			case float64:
				f := int(v)
				s.OutstandingReadRequests = f
			}

		case "outstanding_write_requests":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OutstandingWriteRequests", err)
				}
				s.OutstandingWriteRequests = value
			case float64:
				f := int(v)
				s.OutstandingWriteRequests = f
			}

		case "read_exceptions":
			if err := dec.Decode(&s.ReadExceptions); err != nil {
				return fmt.Errorf("%s | %w", "ReadExceptions", err)
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

		case "shard_id":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ShardId", err)
				}
				s.ShardId = value
			case float64:
				f := int(v)
				s.ShardId = f
			}

		case "successful_read_requests":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SuccessfulReadRequests", err)
				}
				s.SuccessfulReadRequests = value
			case float64:
				f := int64(v)
				s.SuccessfulReadRequests = f
			}

		case "successful_write_requests":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "SuccessfulWriteRequests", err)
				}
				s.SuccessfulWriteRequests = value
			case float64:
				f := int64(v)
				s.SuccessfulWriteRequests = f
			}

		case "time_since_last_read":
			if err := dec.Decode(&s.TimeSinceLastRead); err != nil {
				return fmt.Errorf("%s | %w", "TimeSinceLastRead", err)
			}

		case "time_since_last_read_millis":
			if err := dec.Decode(&s.TimeSinceLastReadMillis); err != nil {
				return fmt.Errorf("%s | %w", "TimeSinceLastReadMillis", err)
			}

		case "total_read_remote_exec_time":
			if err := dec.Decode(&s.TotalReadRemoteExecTime); err != nil {
				return fmt.Errorf("%s | %w", "TotalReadRemoteExecTime", err)
			}

		case "total_read_remote_exec_time_millis":
			if err := dec.Decode(&s.TotalReadRemoteExecTimeMillis); err != nil {
				return fmt.Errorf("%s | %w", "TotalReadRemoteExecTimeMillis", err)
			}

		case "total_read_time":
			if err := dec.Decode(&s.TotalReadTime); err != nil {
				return fmt.Errorf("%s | %w", "TotalReadTime", err)
			}

		case "total_read_time_millis":
			if err := dec.Decode(&s.TotalReadTimeMillis); err != nil {
				return fmt.Errorf("%s | %w", "TotalReadTimeMillis", err)
			}

		case "total_write_time":
			if err := dec.Decode(&s.TotalWriteTime); err != nil {
				return fmt.Errorf("%s | %w", "TotalWriteTime", err)
			}

		case "total_write_time_millis":
			if err := dec.Decode(&s.TotalWriteTimeMillis); err != nil {
				return fmt.Errorf("%s | %w", "TotalWriteTimeMillis", err)
			}

		case "write_buffer_operation_count":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "WriteBufferOperationCount", err)
				}
				s.WriteBufferOperationCount = value
			case float64:
				f := int64(v)
				s.WriteBufferOperationCount = f
			}

		case "write_buffer_size_in_bytes":
			if err := dec.Decode(&s.WriteBufferSizeInBytes); err != nil {
				return fmt.Errorf("%s | %w", "WriteBufferSizeInBytes", err)
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
