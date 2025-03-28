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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _dataStreamLifecycle struct {
	v *types.DataStreamLifecycle
}

func NewDataStreamLifecycle() *_dataStreamLifecycle {

	return &_dataStreamLifecycle{v: types.NewDataStreamLifecycle()}

}

// If defined, every document added to this data stream will be stored at least
// for this time frame.
// Any time after this duration the document could be deleted.
// When empty, every document in this data stream will be stored indefinitely.
func (s *_dataStreamLifecycle) DataRetention(duration types.DurationVariant) *_dataStreamLifecycle {

	s.v.DataRetention = *duration.DurationCaster()

	return s
}

// The downsampling configuration to execute for the managed backing index after
// rollover.
func (s *_dataStreamLifecycle) Downsampling(downsampling types.DataStreamLifecycleDownsamplingVariant) *_dataStreamLifecycle {

	s.v.Downsampling = downsampling.DataStreamLifecycleDownsamplingCaster()

	return s
}

// If defined, it turns data stream lifecycle on/off (`true`/`false`) for this
// data stream. A data stream lifecycle
// that's disabled (enabled: `false`) will have no effect on the data stream.
func (s *_dataStreamLifecycle) Enabled(enabled bool) *_dataStreamLifecycle {

	s.v.Enabled = &enabled

	return s
}

func (s *_dataStreamLifecycle) DataStreamLifecycleCaster() *types.DataStreamLifecycle {
	return s.v
}
