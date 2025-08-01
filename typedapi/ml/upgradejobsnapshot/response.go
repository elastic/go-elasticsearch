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

package upgradejobsnapshot

// Response holds the response body struct for the package upgradejobsnapshot
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/upgrade_job_snapshot/MlUpgradeJobSnapshotResponse.ts#L22-L31
type Response struct {

	// Completed When true, this means the task is complete. When false, it is still running.
	Completed bool `json:"completed"`
	// Node The ID of the node that the upgrade task was started on if it is still
	// running. In serverless this will be the "serverless".
	Node string `json:"node"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
