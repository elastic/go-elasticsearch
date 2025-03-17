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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/simplequerystringflag"
)

// This is provide all the types that are part of the union.
type _pipeSeparatedFlagsSimpleQueryStringFlag struct {
	v types.PipeSeparatedFlagsSimpleQueryStringFlag
}

func NewPipeSeparatedFlagsSimpleQueryStringFlag() *_pipeSeparatedFlagsSimpleQueryStringFlag {
	return &_pipeSeparatedFlagsSimpleQueryStringFlag{v: nil}
}

func (u *_pipeSeparatedFlagsSimpleQueryStringFlag) SimpleQueryStringFlag(simplequerystringflag simplequerystringflag.SimpleQueryStringFlag) *_pipeSeparatedFlagsSimpleQueryStringFlag {

	u.v = &simplequerystringflag
	return u
}

func (u *_pipeSeparatedFlagsSimpleQueryStringFlag) String(string string) *_pipeSeparatedFlagsSimpleQueryStringFlag {

	u.v = &string

	return u
}

func (u *_pipeSeparatedFlagsSimpleQueryStringFlag) PipeSeparatedFlagsSimpleQueryStringFlagCaster() *types.PipeSeparatedFlagsSimpleQueryStringFlag {
	return &u.v
}
