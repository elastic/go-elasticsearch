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

// Package useragentproperty
package useragentproperty

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ingest/_types/Processors.ts#L260-L271
type UserAgentProperty struct {
	Name string
}

var (
	NAME = UserAgentProperty{"NAME"}

	MAJOR = UserAgentProperty{"MAJOR"}

	MINOR = UserAgentProperty{"MINOR"}

	PATCH = UserAgentProperty{"PATCH"}

	OS = UserAgentProperty{"OS"}

	OSNAME = UserAgentProperty{"OS_NAME"}

	OSMAJOR = UserAgentProperty{"OS_MAJOR"}

	OSMINOR = UserAgentProperty{"OS_MINOR"}

	DEVICE = UserAgentProperty{"DEVICE"}

	BUILD = UserAgentProperty{"BUILD"}
)

func (u UserAgentProperty) MarshalText() (text []byte, err error) {
	return []byte(u.String()), nil
}

func (u *UserAgentProperty) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "NAME":
		*u = NAME
	case "MAJOR":
		*u = MAJOR
	case "MINOR":
		*u = MINOR
	case "PATCH":
		*u = PATCH
	case "OS":
		*u = OS
	case "OS_NAME":
		*u = OSNAME
	case "OS_MAJOR":
		*u = OSMAJOR
	case "OS_MINOR":
		*u = OSMINOR
	case "DEVICE":
		*u = DEVICE
	case "BUILD":
		*u = BUILD
	default:
		*u = UserAgentProperty{string(text)}
	}

	return nil
}

func (u UserAgentProperty) String() string {
	return u.Name
}
