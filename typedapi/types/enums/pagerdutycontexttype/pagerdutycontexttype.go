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

// Package pagerdutycontexttype
package pagerdutycontexttype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/watcher/_types/Actions.ts#L67-L70
type PagerDutyContextType struct {
	Name string
}

var (
	Link = PagerDutyContextType{"link"}

	Image = PagerDutyContextType{"image"}
)

func (p PagerDutyContextType) MarshalText() (text []byte, err error) {
	return []byte(p.String()), nil
}

func (p *PagerDutyContextType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "link":
		*p = Link
	case "image":
		*p = Image
	default:
		*p = PagerDutyContextType{string(text)}
	}

	return nil
}

func (p PagerDutyContextType) String() string {
	return p.Name
}
