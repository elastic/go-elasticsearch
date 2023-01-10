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

// RepositoryMeteringInformation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/nodes/_types/RepositoryMeteringInformation.ts#L24-L66
type RepositoryMeteringInformation struct {
	// Archived A flag that tells whether or not this object has been archived. When a
	// repository is closed or updated the
	// repository metering information is archived and kept for a certain period of
	// time. This allows retrieving the
	// repository metering information of previous repository instantiations.
	Archived bool `json:"archived"`
	// ClusterVersion The cluster state version when this object was archived, this field can be
	// used as a logical timestamp to delete
	// all the archived metrics up to an observed version. This field is only
	// present for archived repository metering
	// information objects. The main purpose of this field is to avoid possible race
	// conditions during repository metering
	// information deletions, i.e. deleting archived repositories metering
	// information that we haven’t observed yet.
	ClusterVersion *int64 `json:"cluster_version,omitempty"`
	// RepositoryEphemeralId An identifier that changes every time the repository is updated.
	RepositoryEphemeralId string `json:"repository_ephemeral_id"`
	// RepositoryLocation Represents an unique location within the repository.
	RepositoryLocation RepositoryLocation `json:"repository_location"`
	// RepositoryName Repository name.
	RepositoryName string `json:"repository_name"`
	// RepositoryStartedAt Time the repository was created or updated. Recorded in milliseconds since
	// the Unix Epoch.
	RepositoryStartedAt int64 `json:"repository_started_at"`
	// RepositoryStoppedAt Time the repository was deleted or updated. Recorded in milliseconds since
	// the Unix Epoch.
	RepositoryStoppedAt *int64 `json:"repository_stopped_at,omitempty"`
	// RepositoryType Repository type.
	RepositoryType string `json:"repository_type"`
	// RequestCounts An object with the number of request performed against the repository grouped
	// by request type.
	RequestCounts RequestCounts `json:"request_counts"`
}

// NewRepositoryMeteringInformation returns a RepositoryMeteringInformation.
func NewRepositoryMeteringInformation() *RepositoryMeteringInformation {
	r := &RepositoryMeteringInformation{}

	return r
}
