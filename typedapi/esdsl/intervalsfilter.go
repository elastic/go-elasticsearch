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

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _intervalsFilter struct {
	v *types.IntervalsFilter
}

func NewIntervalsFilter() *_intervalsFilter {
	return &_intervalsFilter{v: types.NewIntervalsFilter()}
}

// AdditionalIntervalsFilterProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_intervalsFilter) AdditionalIntervalsFilterProperty(key string, value json.RawMessage) *_intervalsFilter {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalIntervalsFilterProperty = tmp
	return s
}

func (s *_intervalsFilter) After(after types.IntervalsVariant) *_intervalsFilter {

	s.v.After = after.IntervalsCaster()

	return s
}

func (s *_intervalsFilter) Before(before types.IntervalsVariant) *_intervalsFilter {

	s.v.Before = before.IntervalsCaster()

	return s
}

func (s *_intervalsFilter) ContainedBy(containedby types.IntervalsVariant) *_intervalsFilter {

	s.v.ContainedBy = containedby.IntervalsCaster()

	return s
}

func (s *_intervalsFilter) Containing(containing types.IntervalsVariant) *_intervalsFilter {

	s.v.Containing = containing.IntervalsCaster()

	return s
}

func (s *_intervalsFilter) NotContainedBy(notcontainedby types.IntervalsVariant) *_intervalsFilter {

	s.v.NotContainedBy = notcontainedby.IntervalsCaster()

	return s
}

func (s *_intervalsFilter) NotContaining(notcontaining types.IntervalsVariant) *_intervalsFilter {

	s.v.NotContaining = notcontaining.IntervalsCaster()

	return s
}

func (s *_intervalsFilter) NotOverlapping(notoverlapping types.IntervalsVariant) *_intervalsFilter {

	s.v.NotOverlapping = notoverlapping.IntervalsCaster()

	return s
}

func (s *_intervalsFilter) Overlapping(overlapping types.IntervalsVariant) *_intervalsFilter {

	s.v.Overlapping = overlapping.IntervalsCaster()

	return s
}

func (s *_intervalsFilter) Script(script types.ScriptVariant) *_intervalsFilter {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_intervalsFilter) IntervalsFilterCaster() *types.IntervalsFilter {
	return s.v
}
