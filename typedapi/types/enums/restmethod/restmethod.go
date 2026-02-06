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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

// Package restmethod
package restmethod

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/_global/capabilities/CapabilitiesRequest.ts#L71-L77
type RestMethod struct {
	Name string
}

var (
	GET = RestMethod{"GET"}

	HEAD = RestMethod{"HEAD"}

	POST = RestMethod{"POST"}

	PUT = RestMethod{"PUT"}

	DELETE = RestMethod{"DELETE"}
)

func (r RestMethod) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RestMethod) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "get":
		*r = GET
	case "head":
		*r = HEAD
	case "post":
		*r = POST
	case "put":
		*r = PUT
	case "delete":
		*r = DELETE
	default:
		*r = RestMethod{string(text)}
	}

	return nil
}

func (r RestMethod) String() string {
	return r.Name
}
