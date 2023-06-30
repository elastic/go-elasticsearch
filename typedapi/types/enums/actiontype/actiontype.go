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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Package actiontype
package actiontype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/watcher/_types/Action.ts#L61-L68
type ActionType struct {
	Name string
}

var (
	Email = ActionType{"email"}

	Webhook = ActionType{"webhook"}

	Index = ActionType{"index"}

	Logging = ActionType{"logging"}

	Slack = ActionType{"slack"}

	Pagerduty = ActionType{"pagerduty"}
)

func (a ActionType) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *ActionType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "email":
		*a = Email
	case "webhook":
		*a = Webhook
	case "index":
		*a = Index
	case "logging":
		*a = Logging
	case "slack":
		*a = Slack
	case "pagerduty":
		*a = Pagerduty
	default:
		*a = ActionType{string(text)}
	}

	return nil
}

func (a ActionType) String() string {
	return a.Name
}
