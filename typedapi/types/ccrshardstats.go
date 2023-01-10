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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// CcrShardStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ccr/_types/FollowIndexStats.ts#L35-L69
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
	TimeSinceLastRead             *Duration       `json:"time_since_last_read,omitempty"`
	TimeSinceLastReadMillis       int64           `json:"time_since_last_read_millis"`
	TotalReadRemoteExecTime       *Duration       `json:"total_read_remote_exec_time,omitempty"`
	TotalReadRemoteExecTimeMillis int64           `json:"total_read_remote_exec_time_millis"`
	TotalReadTime                 *Duration       `json:"total_read_time,omitempty"`
	TotalReadTimeMillis           int64           `json:"total_read_time_millis"`
	TotalWriteTime                *Duration       `json:"total_write_time,omitempty"`
	TotalWriteTimeMillis          int64           `json:"total_write_time_millis"`
	WriteBufferOperationCount     int64           `json:"write_buffer_operation_count"`
	WriteBufferSizeInBytes        ByteSize        `json:"write_buffer_size_in_bytes"`
}

// NewCcrShardStats returns a CcrShardStats.
func NewCcrShardStats() *CcrShardStats {
	r := &CcrShardStats{}

	return r
}
