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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

// Package streamtype
package streamtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/eb2e22fb2ac404e676d19bcc7bb089647f029026/specification/streams/_types/StreamType.ts#L20-L36
type StreamType struct {
	Name string
}

var (

	// Logs The logs stream type.
	Logs = StreamType{"logs"}

	// Logsotel The logs.otel stream type, meant for OTel-formatted data.
	Logsotel = StreamType{"logs.otel"}

	// Logsecs The logs.ecs stream type, meant for ECS-formatted data.
	Logsecs = StreamType{"logs.ecs"}
)

func (s StreamType) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *StreamType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "logs":
		*s = Logs
	case "logs.otel":
		*s = Logsotel
	case "logs.ecs":
		*s = Logsecs
	default:
		*s = StreamType{string(text)}
	}

	return nil
}

func (s StreamType) String() string {
	return s.Name
}
