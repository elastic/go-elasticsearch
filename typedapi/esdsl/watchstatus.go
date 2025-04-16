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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _watchStatus struct {
	v *types.WatchStatus
}

func NewWatchStatus(state types.ActivationStateVariant) *_watchStatus {

	tmp := &_watchStatus{v: types.NewWatchStatus()}

	tmp.State(state)

	return tmp

}

func (s *_watchStatus) Actions(watcherstatusactions types.WatcherStatusActionsVariant) *_watchStatus {

	s.v.Actions = *watcherstatusactions.WatcherStatusActionsCaster()

	return s
}

func (s *_watchStatus) ExecutionState(executionstate string) *_watchStatus {

	s.v.ExecutionState = &executionstate

	return s
}

func (s *_watchStatus) LastChecked(datetime types.DateTimeVariant) *_watchStatus {

	s.v.LastChecked = *datetime.DateTimeCaster()

	return s
}

func (s *_watchStatus) LastMetCondition(datetime types.DateTimeVariant) *_watchStatus {

	s.v.LastMetCondition = *datetime.DateTimeCaster()

	return s
}

func (s *_watchStatus) State(state types.ActivationStateVariant) *_watchStatus {

	s.v.State = *state.ActivationStateCaster()

	return s
}

func (s *_watchStatus) Version(versionnumber int64) *_watchStatus {

	s.v.Version = versionnumber

	return s
}

func (s *_watchStatus) WatchStatusCaster() *types.WatchStatus {
	return s.v
}
