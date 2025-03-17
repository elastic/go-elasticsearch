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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _hint struct {
	v *types.Hint
}

func NewHint() *_hint {

	return &_hint{v: types.NewHint()}

}

// A single key-value pair to match against the labels section
// of a profile. A profile is considered matching if it matches
// at least one of the strings.
func (s *_hint) Labels(labels map[string][]string) *_hint {

	s.v.Labels = labels
	return s
}

// A list of profile UIDs to match against.
func (s *_hint) Uids(uids ...string) *_hint {

	for _, v := range uids {

		s.v.Uids = append(s.v.Uids, v)

	}
	return s
}

func (s *_hint) HintCaster() *types.Hint {
	return s.v
}
