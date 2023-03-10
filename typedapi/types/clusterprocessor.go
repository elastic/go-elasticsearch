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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// ClusterProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/cluster/stats/types.ts#L269-L275
type ClusterProcessor struct {
	Count        int64    `json:"count"`
	Current      int64    `json:"current"`
	Failed       int64    `json:"failed"`
	Time         Duration `json:"time,omitempty"`
	TimeInMillis int64    `json:"time_in_millis"`
}

// NewClusterProcessor returns a ClusterProcessor.
func NewClusterProcessor() *ClusterProcessor {
	r := &ClusterProcessor{}

	return r
}
