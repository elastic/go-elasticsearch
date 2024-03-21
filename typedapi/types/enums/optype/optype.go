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
// https://github.com/elastic/elasticsearch-specification/tree/accc26662ab4c58f4f6fb0fc1d9fc5249d0de339

// Package optype
package optype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/accc26662ab4c58f4f6fb0fc1d9fc5249d0de339/specification/_types/common.ts#L255-L264
type OpType struct {
	Name string
}

var (
	Index = OpType{"index"}

	Create = OpType{"create"}
)

func (o OpType) MarshalText() (text []byte, err error) {
	return []byte(o.String()), nil
}

func (o *OpType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "index":
		*o = Index
	case "create":
		*o = Create
	default:
		*o = OpType{string(text)}
	}

	return nil
}

func (o OpType) String() string {
	return o.Name
}
