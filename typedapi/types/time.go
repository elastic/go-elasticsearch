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


package types

// Time holds the union for the following types:
//     int
//     string
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/Time.ts#L65-L71
type Time interface{}

// TimeBuilder holds Time struct and provides a builder API.
type TimeBuilder struct {
	v Time
}

// NewTime provides a builder for the Time struct.
func NewTimeBuilder() *TimeBuilder {
	return &TimeBuilder{}
}

// Build finalize the chain and returns the Time struct
func (u *TimeBuilder) Build() Time {
	return u.v
}

func (u *TimeBuilder) Int(int int) *TimeBuilder {
	u.v = &int
	return u
}

func (u *TimeBuilder) String(string string) *TimeBuilder {
	u.v = &string
	return u
}
