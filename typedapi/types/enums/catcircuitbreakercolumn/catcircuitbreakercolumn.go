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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

// Package catcircuitbreakercolumn
package catcircuitbreakercolumn

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6/specification/cat/_types/CatBase.ts#L1263-L1310
type CatCircuitBreakerColumn struct {
	Name string
}

var (
	Nodeid = CatCircuitBreakerColumn{"node_id"}

	Nodename = CatCircuitBreakerColumn{"node_name"}

	Breaker = CatCircuitBreakerColumn{"breaker"}

	Limit = CatCircuitBreakerColumn{"limit"}

	Limitbytes = CatCircuitBreakerColumn{"limit_bytes"}

	Estimated = CatCircuitBreakerColumn{"estimated"}

	Estimatedbytes = CatCircuitBreakerColumn{"estimated_bytes"}

	Tripped = CatCircuitBreakerColumn{"tripped"}

	Overhead = CatCircuitBreakerColumn{"overhead"}
)

func (c CatCircuitBreakerColumn) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CatCircuitBreakerColumn) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "node_id":
		*c = Nodeid
	case "node_name":
		*c = Nodename
	case "breaker":
		*c = Breaker
	case "limit":
		*c = Limit
	case "limit_bytes":
		*c = Limitbytes
	case "estimated":
		*c = Estimated
	case "estimated_bytes":
		*c = Estimatedbytes
	case "tripped":
		*c = Tripped
	case "overhead":
		*c = Overhead
	default:
		*c = CatCircuitBreakerColumn{string(text)}
	}

	return nil
}

func (c CatCircuitBreakerColumn) String() string {
	return c.Name
}
