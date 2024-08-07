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
// https://github.com/elastic/elasticsearch-specification/tree/19027dbdd366978ccae41842a040a636730e7c10

// Package filteringrulerule
package filteringrulerule

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/19027dbdd366978ccae41842a040a636730e7c10/specification/connector/_types/Connector.ts#L160-L168
type FilteringRuleRule struct {
	Name string
}

var (
	Contains = FilteringRuleRule{"contains"}

	Endswith = FilteringRuleRule{"ends_with"}

	Equals = FilteringRuleRule{"equals"}

	Regex = FilteringRuleRule{"regex"}

	Startswith = FilteringRuleRule{"starts_with"}

	Greaterthan = FilteringRuleRule{">"}

	Lessthan = FilteringRuleRule{"<"}
)

func (f FilteringRuleRule) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f *FilteringRuleRule) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "contains":
		*f = Contains
	case "ends_with":
		*f = Endswith
	case "equals":
		*f = Equals
	case "regex":
		*f = Regex
	case "starts_with":
		*f = Startswith
	case ">":
		*f = Greaterthan
	case "<":
		*f = Lessthan
	default:
		*f = FilteringRuleRule{string(text)}
	}

	return nil
}

func (f FilteringRuleRule) String() string {
	return f.Name
}
