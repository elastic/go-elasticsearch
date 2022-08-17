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

// RecoveryOrigin type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/recovery/types.ts#L76-L89
type RecoveryOrigin struct {
	BootstrapNewHistoryUuid *bool             `json:"bootstrap_new_history_uuid,omitempty"`
	Host                    *Host             `json:"host,omitempty"`
	Hostname                *string           `json:"hostname,omitempty"`
	Id                      *Id               `json:"id,omitempty"`
	Index                   *IndexName        `json:"index,omitempty"`
	Ip                      *Ip               `json:"ip,omitempty"`
	Name                    *Name             `json:"name,omitempty"`
	Repository              *Name             `json:"repository,omitempty"`
	RestoreUUID             *Uuid             `json:"restoreUUID,omitempty"`
	Snapshot                *Name             `json:"snapshot,omitempty"`
	TransportAddress        *TransportAddress `json:"transport_address,omitempty"`
	Version                 *VersionString    `json:"version,omitempty"`
}

// RecoveryOriginBuilder holds RecoveryOrigin struct and provides a builder API.
type RecoveryOriginBuilder struct {
	v *RecoveryOrigin
}

// NewRecoveryOrigin provides a builder for the RecoveryOrigin struct.
func NewRecoveryOriginBuilder() *RecoveryOriginBuilder {
	r := RecoveryOriginBuilder{
		&RecoveryOrigin{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryOrigin struct
func (rb *RecoveryOriginBuilder) Build() RecoveryOrigin {
	return *rb.v
}

func (rb *RecoveryOriginBuilder) BootstrapNewHistoryUuid(bootstrapnewhistoryuuid bool) *RecoveryOriginBuilder {
	rb.v.BootstrapNewHistoryUuid = &bootstrapnewhistoryuuid
	return rb
}

func (rb *RecoveryOriginBuilder) Host(host Host) *RecoveryOriginBuilder {
	rb.v.Host = &host
	return rb
}

func (rb *RecoveryOriginBuilder) Hostname(hostname string) *RecoveryOriginBuilder {
	rb.v.Hostname = &hostname
	return rb
}

func (rb *RecoveryOriginBuilder) Id(id Id) *RecoveryOriginBuilder {
	rb.v.Id = &id
	return rb
}

func (rb *RecoveryOriginBuilder) Index(index IndexName) *RecoveryOriginBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *RecoveryOriginBuilder) Ip(ip Ip) *RecoveryOriginBuilder {
	rb.v.Ip = &ip
	return rb
}

func (rb *RecoveryOriginBuilder) Name(name Name) *RecoveryOriginBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *RecoveryOriginBuilder) Repository(repository Name) *RecoveryOriginBuilder {
	rb.v.Repository = &repository
	return rb
}

func (rb *RecoveryOriginBuilder) RestoreUUID(restoreuuid Uuid) *RecoveryOriginBuilder {
	rb.v.RestoreUUID = &restoreuuid
	return rb
}

func (rb *RecoveryOriginBuilder) Snapshot(snapshot Name) *RecoveryOriginBuilder {
	rb.v.Snapshot = &snapshot
	return rb
}

func (rb *RecoveryOriginBuilder) TransportAddress(transportaddress TransportAddress) *RecoveryOriginBuilder {
	rb.v.TransportAddress = &transportaddress
	return rb
}

func (rb *RecoveryOriginBuilder) Version(version VersionString) *RecoveryOriginBuilder {
	rb.v.Version = &version
	return rb
}
