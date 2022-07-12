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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

// Package excludefrequent
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/ml/_types/Detector.ts#L82-L87
package excludefrequent

import "strings"

type ExcludeFrequent struct {
	name string
}

var (
	All = ExcludeFrequent{"all"}

	None = ExcludeFrequent{"none"}

	By = ExcludeFrequent{"by"}

	Over = ExcludeFrequent{"over"}
)

func (e ExcludeFrequent) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *ExcludeFrequent) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "all":
		*e = All
	case "none":
		*e = None
	case "by":
		*e = By
	case "over":
		*e = Over
	default:
		*e = ExcludeFrequent{string(text)}
	}

	return nil
}

func (e ExcludeFrequent) String() string {
	return e.name
}
