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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Package waitforactiveshardoptions
package waitforactiveshardoptions

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/common.ts#L304-L308
type WaitForActiveShardOptions struct {
	Name string
}

var (
	All = WaitForActiveShardOptions{"all"}

	IndexSetting = WaitForActiveShardOptions{"index-setting"}
)

func (w WaitForActiveShardOptions) MarshalText() (text []byte, err error) {
	return []byte(w.String()), nil
}

func (w *WaitForActiveShardOptions) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "all":
		*w = All
	case "index-setting":
		*w = IndexSetting
	default:
		*w = WaitForActiveShardOptions{string(text)}
	}

	return nil
}

func (w WaitForActiveShardOptions) String() string {
	return w.Name
}
