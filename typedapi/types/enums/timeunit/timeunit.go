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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Package timeunit
package timeunit

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/Time.ts#L69-L77
type TimeUnit struct {
	Name string
}

var (
	Nanoseconds = TimeUnit{"nanos"}

	Microseconds = TimeUnit{"micros"}

	Milliseconds = TimeUnit{"ms"}

	Seconds = TimeUnit{"s"}

	Minutes = TimeUnit{"m"}

	Hours = TimeUnit{"h"}

	Days = TimeUnit{"d"}
)

func (t TimeUnit) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *TimeUnit) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "nanos":
		*t = Nanoseconds
	case "micros":
		*t = Microseconds
	case "ms":
		*t = Milliseconds
	case "s":
		*t = Seconds
	case "m":
		*t = Minutes
	case "h":
		*t = Hours
	case "d":
		*t = Days
	default:
		*t = TimeUnit{string(text)}
	}

	return nil
}

func (t TimeUnit) String() string {
	return t.Name
}
