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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _autoscalingPolicy struct {
	v *types.AutoscalingPolicy
}

func NewAutoscalingPolicy() *_autoscalingPolicy {

	return &_autoscalingPolicy{v: types.NewAutoscalingPolicy()}

}

// Decider settings.
func (s *_autoscalingPolicy) Deciders(deciders map[string]json.RawMessage) *_autoscalingPolicy {

	s.v.Deciders = deciders
	return s
}

func (s *_autoscalingPolicy) AddDecider(key string, value json.RawMessage) *_autoscalingPolicy {

	var tmp map[string]json.RawMessage
	if s.v.Deciders == nil {
		s.v.Deciders = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Deciders
	}

	tmp[key] = value

	s.v.Deciders = tmp
	return s
}

func (s *_autoscalingPolicy) Roles(roles ...string) *_autoscalingPolicy {

	for _, v := range roles {

		s.v.Roles = append(s.v.Roles, v)

	}
	return s
}

func (s *_autoscalingPolicy) AutoscalingPolicyCaster() *types.AutoscalingPolicy {
	return s.v
}
