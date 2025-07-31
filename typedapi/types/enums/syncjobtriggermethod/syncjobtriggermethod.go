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

// Package syncjobtriggermethod
package syncjobtriggermethod

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/connector/_types/SyncJob.ts#L48-L51
type SyncJobTriggerMethod struct {
	Name string
}

var (
	Ondemand = SyncJobTriggerMethod{"on_demand"}

	Scheduled = SyncJobTriggerMethod{"scheduled"}
)

func (s SyncJobTriggerMethod) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SyncJobTriggerMethod) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "on_demand":
		*s = Ondemand
	case "scheduled":
		*s = Scheduled
	default:
		*s = SyncJobTriggerMethod{string(text)}
	}

	return nil
}

func (s SyncJobTriggerMethod) String() string {
	return s.Name
}
