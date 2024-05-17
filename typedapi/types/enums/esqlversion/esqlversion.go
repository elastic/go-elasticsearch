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
// https://github.com/elastic/elasticsearch-specification/tree/9a0362eb2579c6604966a8fb307caee92de04270

// Package esqlversion
package esqlversion

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/9a0362eb2579c6604966a8fb307caee92de04270/specification/esql/_types/EsqlVersion.ts#L20-L29
type EsqlVersion struct {
	Name string
}

var (
	V20240401 = EsqlVersion{"2024.04.01"}
)

func (e EsqlVersion) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *EsqlVersion) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "2024.04.01":
		*e = V20240401
	default:
		*e = EsqlVersion{string(text)}
	}

	return nil
}

func (e EsqlVersion) String() string {
	return e.Name
}
