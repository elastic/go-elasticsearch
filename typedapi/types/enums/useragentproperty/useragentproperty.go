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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package useragentproperty
package useragentproperty

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ingest/_types/Processors.ts#L547-L553
type UserAgentProperty struct {
	Name string
}

var (
	Name = UserAgentProperty{"name"}

	Os = UserAgentProperty{"os"}

	Device = UserAgentProperty{"device"}

	Original = UserAgentProperty{"original"}

	Version = UserAgentProperty{"version"}
)

func (u UserAgentProperty) MarshalText() (text []byte, err error) {
	return []byte(u.String()), nil
}

func (u *UserAgentProperty) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "name":
		*u = Name
	case "os":
		*u = Os
	case "device":
		*u = Device
	case "original":
		*u = Original
	case "version":
		*u = Version
	default:
		*u = UserAgentProperty{string(text)}
	}

	return nil
}

func (u UserAgentProperty) String() string {
	return u.Name
}
