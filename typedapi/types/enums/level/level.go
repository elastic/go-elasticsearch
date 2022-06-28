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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


// Package level
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/common.ts#L222-L226
package level

import "strings"

type Level struct {
	name string
}

var (
	Cluster = Level{"cluster"}

	Indices = Level{"indices"}

	Shards = Level{"shards"}
)

func (l Level) MarshalText() (text []byte, err error) {
	return []byte(l.String()), nil
}

func (l *Level) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "cluster":
		*l = Cluster
	case "indices":
		*l = Indices
	case "shards":
		*l = Shards
	default:
		*l = Level{string(text)}
	}

	return nil
}

func (l Level) String() string {
	return l.name
}
