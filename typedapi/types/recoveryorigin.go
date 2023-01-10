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

// RecoveryOrigin type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/indices/recovery/types.ts#L76-L89
type RecoveryOrigin struct {
	BootstrapNewHistoryUuid *bool   `json:"bootstrap_new_history_uuid,omitempty"`
	Host                    *string `json:"host,omitempty"`
	Hostname                *string `json:"hostname,omitempty"`
	Id                      *string `json:"id,omitempty"`
	Index                   *string `json:"index,omitempty"`
	Ip                      *string `json:"ip,omitempty"`
	Name                    *string `json:"name,omitempty"`
	Repository              *string `json:"repository,omitempty"`
	RestoreUUID             *string `json:"restoreUUID,omitempty"`
	Snapshot                *string `json:"snapshot,omitempty"`
	TransportAddress        *string `json:"transport_address,omitempty"`
	Version                 *string `json:"version,omitempty"`
}

// NewRecoveryOrigin returns a RecoveryOrigin.
func NewRecoveryOrigin() *RecoveryOrigin {
	r := &RecoveryOrigin{}

	return r
}
