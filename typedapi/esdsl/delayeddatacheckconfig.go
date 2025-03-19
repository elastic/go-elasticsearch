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

type _delayedDataCheckConfig struct {
	v *types.DelayedDataCheckConfig
}

func NewDelayedDataCheckConfig(enabled bool) *_delayedDataCheckConfig {

	tmp := &_delayedDataCheckConfig{v: types.NewDelayedDataCheckConfig()}

	tmp.Enabled(enabled)

	return tmp

}

// The window of time that is searched for late data. This window of time ends
// with the latest finalized bucket.
// It defaults to null, which causes an appropriate `check_window` to be
// calculated when the real-time datafeed runs.
// In particular, the default `check_window` span calculation is based on the
// maximum of `2h` or `8 * bucket_span`.
func (s *_delayedDataCheckConfig) CheckWindow(duration types.DurationVariant) *_delayedDataCheckConfig {

	s.v.CheckWindow = *duration.DurationCaster()

	return s
}

// Specifies whether the datafeed periodically checks for delayed data.
func (s *_delayedDataCheckConfig) Enabled(enabled bool) *_delayedDataCheckConfig {

	s.v.Enabled = enabled

	return s
}

func (s *_delayedDataCheckConfig) DelayedDataCheckConfigCaster() *types.DelayedDataCheckConfig {
	return s.v
}
