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

// Package jobstate
package jobstate

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/Job.ts#L36-L52
type JobState struct {
	Name string
}

var (
	Closing = JobState{"closing"}

	Closed = JobState{"closed"}

	Opened = JobState{"opened"}

	Failed = JobState{"failed"}

	Opening = JobState{"opening"}
)

func (j JobState) MarshalText() (text []byte, err error) {
	return []byte(j.String()), nil
}

func (j *JobState) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "closing":
		*j = Closing
	case "closed":
		*j = Closed
	case "opened":
		*j = Opened
	case "failed":
		*j = Failed
	case "opening":
		*j = Opening
	default:
		*j = JobState{string(text)}
	}

	return nil
}

func (j JobState) String() string {
	return j.Name
}
