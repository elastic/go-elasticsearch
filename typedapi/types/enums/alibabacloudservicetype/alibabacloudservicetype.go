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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

// Package alibabacloudservicetype
package alibabacloudservicetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/cbfcc73d01310bed2a480ec35aaef98138b598e5/specification/inference/_types/CommonTypes.ts#L289-L291
type AlibabaCloudServiceType struct {
	Name string
}

var (
	AlibabacloudAiSearch = AlibabaCloudServiceType{"alibabacloud-ai-search"}
)

func (a AlibabaCloudServiceType) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *AlibabaCloudServiceType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "alibabacloud-ai-search":
		*a = AlibabacloudAiSearch
	default:
		*a = AlibabaCloudServiceType{string(text)}
	}

	return nil
}

func (a AlibabaCloudServiceType) String() string {
	return a.Name
}
