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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _downsamplingRound struct {
	v *types.DownsamplingRound
}

func NewDownsamplingRound(config types.DownsampleConfigVariant) *_downsamplingRound {

	tmp := &_downsamplingRound{v: types.NewDownsamplingRound()}

	tmp.Config(config)

	return tmp

}

// The duration since rollover when this downsampling round should execute
func (s *_downsamplingRound) After(duration types.DurationVariant) *_downsamplingRound {

	s.v.After = *duration.DurationCaster()

	return s
}

// The downsample configuration to execute.
func (s *_downsamplingRound) Config(config types.DownsampleConfigVariant) *_downsamplingRound {

	s.v.Config = *config.DownsampleConfigCaster()

	return s
}

func (s *_downsamplingRound) DownsamplingRoundCaster() *types.DownsamplingRound {
	return s.v
}
