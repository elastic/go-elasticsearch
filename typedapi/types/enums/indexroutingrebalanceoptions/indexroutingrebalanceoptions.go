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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Package indexroutingrebalanceoptions
package indexroutingrebalanceoptions

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/_types/IndexRouting.ts#L45-L50
type IndexRoutingRebalanceOptions struct {
	Name string
}

var (
	All = IndexRoutingRebalanceOptions{"all"}

	Primaries = IndexRoutingRebalanceOptions{"primaries"}

	Replicas = IndexRoutingRebalanceOptions{"replicas"}

	None = IndexRoutingRebalanceOptions{"none"}
)

func (i IndexRoutingRebalanceOptions) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *IndexRoutingRebalanceOptions) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "all":
		*i = All
	case "primaries":
		*i = Primaries
	case "replicas":
		*i = Replicas
	case "none":
		*i = None
	default:
		*i = IndexRoutingRebalanceOptions{string(text)}
	}

	return nil
}

func (i IndexRoutingRebalanceOptions) String() string {
	return i.Name
}
