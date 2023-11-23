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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Package decision
package decision

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/allocation_explain/types.ts#L86-L95
type Decision struct {
	Name string
}

var (
	Yes = Decision{"yes"}

	No = Decision{"no"}

	Worsebalance = Decision{"worse_balance"}

	Throttled = Decision{"throttled"}

	Awaitinginfo = Decision{"awaiting_info"}

	Allocationdelayed = Decision{"allocation_delayed"}

	Novalidshardcopy = Decision{"no_valid_shard_copy"}

	Noattempt = Decision{"no_attempt"}
)

func (d Decision) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *Decision) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "yes":
		*d = Yes
	case "no":
		*d = No
	case "worse_balance":
		*d = Worsebalance
	case "throttled":
		*d = Throttled
	case "awaiting_info":
		*d = Awaitinginfo
	case "allocation_delayed":
		*d = Allocationdelayed
	case "no_valid_shard_copy":
		*d = Novalidshardcopy
	case "no_attempt":
		*d = Noattempt
	default:
		*d = Decision{string(text)}
	}

	return nil
}

func (d Decision) String() string {
	return d.Name
}
