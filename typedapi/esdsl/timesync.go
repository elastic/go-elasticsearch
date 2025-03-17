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

type _timeSync struct {
	v *types.TimeSync
}

// Specifies that the transform uses a time field to synchronize the source and
// destination indices.
func NewTimeSync() *_timeSync {

	return &_timeSync{v: types.NewTimeSync()}

}

// The time delay between the current time and the latest input data time.
func (s *_timeSync) Delay(duration types.DurationVariant) *_timeSync {

	s.v.Delay = *duration.DurationCaster()

	return s
}

// The date field that is used to identify new documents in the source. In
// general, it’s a good idea to use a field
// that contains the ingest timestamp. If you use a different field, you might
// need to set the delay such that it
// accounts for data transmission delays.
func (s *_timeSync) Field(field string) *_timeSync {

	s.v.Field = field

	return s
}

func (s *_timeSync) SyncContainerCaster() *types.SyncContainer {
	container := types.NewSyncContainer()

	container.Time = s.v

	return container
}

func (s *_timeSync) TimeSyncCaster() *types.TimeSync {
	return s.v
}
