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

// Package clustersearchstatus
package clustersearchstatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/Stats.ts#L37-L43
type ClusterSearchStatus struct {
	Name string
}

var (
	Running = ClusterSearchStatus{"running"}

	Successful = ClusterSearchStatus{"successful"}

	Partial = ClusterSearchStatus{"partial"}

	Skipped = ClusterSearchStatus{"skipped"}

	Failed = ClusterSearchStatus{"failed"}
)

func (c ClusterSearchStatus) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *ClusterSearchStatus) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "running":
		*c = Running
	case "successful":
		*c = Successful
	case "partial":
		*c = Partial
	case "skipped":
		*c = Skipped
	case "failed":
		*c = Failed
	default:
		*c = ClusterSearchStatus{string(text)}
	}

	return nil
}

func (c ClusterSearchStatus) String() string {
	return c.Name
}
