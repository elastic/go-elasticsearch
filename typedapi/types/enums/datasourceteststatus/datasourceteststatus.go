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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

// Package datasourceteststatus
package datasourceteststatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/esql/test_data_source_connection/TestDataSourceConnectionResponse.ts#L20-L35
type DataSourceTestStatus struct {
	Name string
}

var (

	// Success The probe ran and the backend is reachable from every eligible data node.
	Success = DataSourceTestStatus{"success"}

	// Failure The probe ran but the backend was unreachable or rejected the connection. The
	// `error` field carries a human-readable reason.
	Failure = DataSourceTestStatus{"failure"}

	// Untestable The type is valid and registered, but the probe cannot give a meaningful
	// result — for example, no connectivity target, anonymous credentials, or
	// bucket-scoped credentials. The optional `message` field may carry
	// user-visible guidance.
	Untestable = DataSourceTestStatus{"untestable"}
)

func (d DataSourceTestStatus) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DataSourceTestStatus) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "success":
		*d = Success
	case "failure":
		*d = Failure
	case "untestable":
		*d = Untestable
	default:
		*d = DataSourceTestStatus{string(text)}
	}

	return nil
}

func (d DataSourceTestStatus) String() string {
	return d.Name
}
