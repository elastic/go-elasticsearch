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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

// Package retentionsource
package retentionsource

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/indices/_types/DataStreamLifecycle.ts#L67-L72
type RetentionSource struct {
	Name string
}

var (
	Datastreamconfiguration = RetentionSource{"data_stream_configuration"}

	Defaultglobalretention = RetentionSource{"default_global_retention"}

	Maxglobalretention = RetentionSource{"max_global_retention"}

	Defaultfailuresretention = RetentionSource{"default_failures_retention"}
)

func (r RetentionSource) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RetentionSource) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "data_stream_configuration":
		*r = Datastreamconfiguration
	case "default_global_retention":
		*r = Defaultglobalretention
	case "max_global_retention":
		*r = Maxglobalretention
	case "default_failures_retention":
		*r = Defaultfailuresretention
	default:
		*r = RetentionSource{string(text)}
	}

	return nil
}

func (r RetentionSource) String() string {
	return r.Name
}
