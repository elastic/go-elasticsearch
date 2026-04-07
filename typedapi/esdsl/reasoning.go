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
// https://github.com/elastic/elasticsearch-specification/tree/836fca874204ca4173ae5c36fb6b5107d28d2fc0

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/reasoningeffort"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/reasoningsummary"
)

type _reasoning struct {
	v *types.Reasoning
}

func NewReasoning() *_reasoning {

	return &_reasoning{v: types.NewReasoning()}

}

func (s *_reasoning) Effort(effort reasoningeffort.ReasoningEffort) *_reasoning {

	s.v.Effort = &effort
	return s
}

func (s *_reasoning) Enabled(enabled bool) *_reasoning {

	s.v.Enabled = &enabled

	return s
}

func (s *_reasoning) Exclude(exclude bool) *_reasoning {

	s.v.Exclude = &exclude

	return s
}

func (s *_reasoning) Summary(summary reasoningsummary.ReasoningSummary) *_reasoning {

	s.v.Summary = &summary
	return s
}

func (s *_reasoning) ReasoningCaster() *types.Reasoning {
	return s.v
}
