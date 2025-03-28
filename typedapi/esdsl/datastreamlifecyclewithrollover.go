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

type _dataStreamLifecycleWithRollover struct {
	v *types.DataStreamLifecycleWithRollover
}

func NewDataStreamLifecycleWithRollover() *_dataStreamLifecycleWithRollover {

	return &_dataStreamLifecycleWithRollover{v: types.NewDataStreamLifecycleWithRollover()}

}

// If defined, every document added to this data stream will be stored at least
// for this time frame.
// Any time after this duration the document could be deleted.
// When empty, every document in this data stream will be stored indefinitely.
func (s *_dataStreamLifecycleWithRollover) DataRetention(duration types.DurationVariant) *_dataStreamLifecycleWithRollover {

	s.v.DataRetention = *duration.DurationCaster()

	return s
}

// The downsampling configuration to execute for the managed backing index after
// rollover.
func (s *_dataStreamLifecycleWithRollover) Downsampling(downsampling types.DataStreamLifecycleDownsamplingVariant) *_dataStreamLifecycleWithRollover {

	s.v.Downsampling = downsampling.DataStreamLifecycleDownsamplingCaster()

	return s
}

// If defined, it turns data stream lifecycle on/off (`true`/`false`) for this
// data stream. A data stream lifecycle
// that's disabled (enabled: `false`) will have no effect on the data stream.
func (s *_dataStreamLifecycleWithRollover) Enabled(enabled bool) *_dataStreamLifecycleWithRollover {

	s.v.Enabled = &enabled

	return s
}

// The conditions which will trigger the rollover of a backing index as
// configured by the cluster setting `cluster.lifecycle.default.rollover`.
// This property is an implementation detail and it will only be retrieved when
// the query param `include_defaults` is set to true.
// The contents of this field are subject to change.
func (s *_dataStreamLifecycleWithRollover) Rollover(rollover types.DataStreamLifecycleRolloverConditionsVariant) *_dataStreamLifecycleWithRollover {

	s.v.Rollover = rollover.DataStreamLifecycleRolloverConditionsCaster()

	return s
}

func (s *_dataStreamLifecycleWithRollover) DataStreamLifecycleWithRolloverCaster() *types.DataStreamLifecycleWithRollover {
	return s.v
}
