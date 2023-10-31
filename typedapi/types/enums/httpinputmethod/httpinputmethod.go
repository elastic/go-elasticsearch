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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Package httpinputmethod
package httpinputmethod

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/watcher/_types/Input.ts#L59-L65
type HttpInputMethod struct {
	Name string
}

var (
	Head = HttpInputMethod{"head"}

	Get = HttpInputMethod{"get"}

	Post = HttpInputMethod{"post"}

	Put = HttpInputMethod{"put"}

	Delete = HttpInputMethod{"delete"}
)

func (h HttpInputMethod) MarshalText() (text []byte, err error) {
	return []byte(h.String()), nil
}

func (h *HttpInputMethod) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "head":
		*h = Head
	case "get":
		*h = Get
	case "post":
		*h = Post
	case "put":
		*h = Put
	case "delete":
		*h = Delete
	default:
		*h = HttpInputMethod{string(text)}
	}

	return nil
}

func (h HttpInputMethod) String() string {
	return h.Name
}
