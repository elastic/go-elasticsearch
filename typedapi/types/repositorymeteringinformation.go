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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// RepositoryMeteringInformation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/RepositoryMeteringInformation.ts#L24-L66
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
	ClusterVersion *VersionNumber `json:"cluster_version,omitempty"`
	// RepositoryEphemeralId An identifier that changes every time the repository is updated.
	RepositoryEphemeralId Id `json:"repository_ephemeral_id"`
	// RepositoryLocation Represents an unique location within the repository.
	RepositoryLocation RepositoryLocation `json:"repository_location"`
	// RepositoryName Repository name.
	RepositoryName Name `json:"repository_name"`
	// RepositoryStartedAt Time the repository was created or updated. Recorded in milliseconds since
	// the Unix Epoch.
	RepositoryStartedAt EpochTimeUnitMillis `json:"repository_started_at"`
	// RepositoryStoppedAt Time the repository was deleted or updated. Recorded in milliseconds since
	// the Unix Epoch.
	RepositoryStoppedAt *EpochTimeUnitMillis `json:"repository_stopped_at,omitempty"`
	// RepositoryType Repository type.
	RepositoryType string `json:"repository_type"`
	// RequestCounts An object with the number of request performed against the repository grouped
	// by request type.
	RequestCounts RequestCounts `json:"request_counts"`
}

// RepositoryMeteringInformationBuilder holds RepositoryMeteringInformation struct and provides a builder API.
type RepositoryMeteringInformationBuilder struct {
	v *RepositoryMeteringInformation
}

// NewRepositoryMeteringInformation provides a builder for the RepositoryMeteringInformation struct.
func NewRepositoryMeteringInformationBuilder() *RepositoryMeteringInformationBuilder {
	r := RepositoryMeteringInformationBuilder{
		&RepositoryMeteringInformation{},
	}

	return &r
}

// Build finalize the chain and returns the RepositoryMeteringInformation struct
func (rb *RepositoryMeteringInformationBuilder) Build() RepositoryMeteringInformation {
	return *rb.v
}

// Archived A flag that tells whether or not this object has been archived. When a
// repository is closed or updated the
// repository metering information is archived and kept for a certain period of
// time. This allows retrieving the
// repository metering information of previous repository instantiations.

func (rb *RepositoryMeteringInformationBuilder) Archived(archived bool) *RepositoryMeteringInformationBuilder {
	rb.v.Archived = archived
	return rb
}

// ClusterVersion The cluster state version when this object was archived, this field can be
// used as a logical timestamp to delete
// all the archived metrics up to an observed version. This field is only
// present for archived repository metering
// information objects. The main purpose of this field is to avoid possible race
// conditions during repository metering
// information deletions, i.e. deleting archived repositories metering
// information that we haven’t observed yet.

func (rb *RepositoryMeteringInformationBuilder) ClusterVersion(clusterversion VersionNumber) *RepositoryMeteringInformationBuilder {
	rb.v.ClusterVersion = &clusterversion
	return rb
}

// RepositoryEphemeralId An identifier that changes every time the repository is updated.

func (rb *RepositoryMeteringInformationBuilder) RepositoryEphemeralId(repositoryephemeralid Id) *RepositoryMeteringInformationBuilder {
	rb.v.RepositoryEphemeralId = repositoryephemeralid
	return rb
}

// RepositoryLocation Represents an unique location within the repository.

func (rb *RepositoryMeteringInformationBuilder) RepositoryLocation(repositorylocation *RepositoryLocationBuilder) *RepositoryMeteringInformationBuilder {
	v := repositorylocation.Build()
	rb.v.RepositoryLocation = v
	return rb
}

// RepositoryName Repository name.

func (rb *RepositoryMeteringInformationBuilder) RepositoryName(repositoryname Name) *RepositoryMeteringInformationBuilder {
	rb.v.RepositoryName = repositoryname
	return rb
}

// RepositoryStartedAt Time the repository was created or updated. Recorded in milliseconds since
// the Unix Epoch.

func (rb *RepositoryMeteringInformationBuilder) RepositoryStartedAt(repositorystartedat *EpochTimeUnitMillisBuilder) *RepositoryMeteringInformationBuilder {
	v := repositorystartedat.Build()
	rb.v.RepositoryStartedAt = v
	return rb
}

// RepositoryStoppedAt Time the repository was deleted or updated. Recorded in milliseconds since
// the Unix Epoch.

func (rb *RepositoryMeteringInformationBuilder) RepositoryStoppedAt(repositorystoppedat *EpochTimeUnitMillisBuilder) *RepositoryMeteringInformationBuilder {
	v := repositorystoppedat.Build()
	rb.v.RepositoryStoppedAt = &v
	return rb
}

// RepositoryType Repository type.

func (rb *RepositoryMeteringInformationBuilder) RepositoryType(repositorytype string) *RepositoryMeteringInformationBuilder {
	rb.v.RepositoryType = repositorytype
	return rb
}

// RequestCounts An object with the number of request performed against the repository grouped
// by request type.

func (rb *RepositoryMeteringInformationBuilder) RequestCounts(requestcounts *RequestCountsBuilder) *RepositoryMeteringInformationBuilder {
	v := requestcounts.Build()
	rb.v.RequestCounts = v
	return rb
}
