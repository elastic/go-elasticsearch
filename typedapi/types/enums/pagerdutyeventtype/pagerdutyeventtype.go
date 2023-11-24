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

// Package pagerdutyeventtype
package pagerdutyeventtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/watcher/_types/Actions.ts#L72-L76
type PagerDutyEventType struct {
	Name string
}

var (
	Trigger = PagerDutyEventType{"trigger"}

	Resolve = PagerDutyEventType{"resolve"}

	Acknowledge = PagerDutyEventType{"acknowledge"}
)

func (p PagerDutyEventType) MarshalText() (text []byte, err error) {
	return []byte(p.String()), nil
}

func (p *PagerDutyEventType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "trigger":
		*p = Trigger
	case "resolve":
		*p = Resolve
	case "acknowledge":
		*p = Acknowledge
	default:
		*p = PagerDutyEventType{string(text)}
	}

	return nil
}

func (p PagerDutyEventType) String() string {
	return p.Name
}
