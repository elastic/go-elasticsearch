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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

// RollupCapabilities type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/rollup/get_rollup_caps/types.ts#L24-L29
type RollupCapabilities struct {
	// RollupJobs There can be multiple, independent jobs configured for a single index or
	// index pattern. Each of these jobs may have different configurations, so the
	// API returns a list of all the various configurations available.
	RollupJobs []RollupCapabilitySummary `json:"rollup_jobs"`
}

// NewRollupCapabilities returns a RollupCapabilities.
func NewRollupCapabilities() *RollupCapabilities {
	r := &RollupCapabilities{}

	return r
}
