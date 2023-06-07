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
// https://github.com/elastic/elasticsearch-specification/tree/0a58ae2e52dd1bc6227f65da9cbbcea5b61dde96

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indicatorhealthstatus"
)

// DiskIndicator type.
//
// https://github.com/elastic/elasticsearch-specification/blob/0a58ae2e52dd1bc6227f65da9cbbcea5b61dde96/specification/_global/health_report/types.ts#L120-L124
type DiskIndicator struct {
	Details   *DiskIndicatorDetails                       `json:"details,omitempty"`
	Diagnosis []Diagnosis                                 `json:"diagnosis,omitempty"`
	Impacts   []Impact                                    `json:"impacts,omitempty"`
	Status    indicatorhealthstatus.IndicatorHealthStatus `json:"status"`
	Symptom   string                                      `json:"symptom"`
}

// NewDiskIndicator returns a DiskIndicator.
func NewDiskIndicator() *DiskIndicator {
	r := &DiskIndicator{}

	return r
}
