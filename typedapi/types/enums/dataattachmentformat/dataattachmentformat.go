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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Package dataattachmentformat
package dataattachmentformat

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/watcher/_types/Actions.ts#L187-L190
type DataAttachmentFormat struct {
	Name string
}

var (
	Json = DataAttachmentFormat{"json"}

	Yaml = DataAttachmentFormat{"yaml"}
)

func (d DataAttachmentFormat) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DataAttachmentFormat) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "json":
		*d = Json
	case "yaml":
		*d = Yaml
	default:
		*d = DataAttachmentFormat{string(text)}
	}

	return nil
}

func (d DataAttachmentFormat) String() string {
	return d.Name
}
