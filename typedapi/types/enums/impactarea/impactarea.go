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
// https://github.com/elastic/elasticsearch-specification/tree/f6a370d0fba975752c644fc730f7c45610e28f36

// Package impactarea
package impactarea

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/f6a370d0fba975752c644fc730f7c45610e28f36/specification/_global/health_report/types.ts#L73-L78
type ImpactArea struct {
	Name string
}

var (
	Search = ImpactArea{"search"}

	Ingest = ImpactArea{"ingest"}

	Backup = ImpactArea{"backup"}

	Deploymentmanagement = ImpactArea{"deployment_management"}
)

func (i ImpactArea) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *ImpactArea) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "search":
		*i = Search
	case "ingest":
		*i = Ingest
	case "backup":
		*i = Backup
	case "deployment_management":
		*i = Deploymentmanagement
	default:
		*i = ImpactArea{string(text)}
	}

	return nil
}

func (i ImpactArea) String() string {
	return i.Name
}
