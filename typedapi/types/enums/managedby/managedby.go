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
// https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64

// Package managedby
package managedby

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64/specification/indices/_types/DataStream.ts#L32-L37
type ManagedBy struct {
	Name string
}

var (
	Ilm = ManagedBy{"Index Lifecycle Management"}

	Datastream = ManagedBy{"Data stream lifecycle"}

	Unmanaged = ManagedBy{"Unmanaged"}
)

func (m ManagedBy) MarshalText() (text []byte, err error) {
	return []byte(m.String()), nil
}

func (m *ManagedBy) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "Index Lifecycle Management":
		*m = Ilm
	case "Data stream lifecycle":
		*m = Datastream
	case "Unmanaged":
		*m = Unmanaged
	default:
		*m = ManagedBy{string(text)}
	}

	return nil
}

func (m ManagedBy) String() string {
	return m.Name
}
