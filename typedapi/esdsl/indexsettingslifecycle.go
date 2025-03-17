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

type _indexSettingsLifecycle struct {
	v *types.IndexSettingsLifecycle
}

func NewIndexSettingsLifecycle() *_indexSettingsLifecycle {

	return &_indexSettingsLifecycle{v: types.NewIndexSettingsLifecycle()}

}

// Indicates whether or not the index has been rolled over. Automatically set to
// true when ILM completes the rollover action.
// You can explicitly set it to skip rollover.
func (s *_indexSettingsLifecycle) IndexingComplete(stringifiedboolean types.StringifiedbooleanVariant) *_indexSettingsLifecycle {

	s.v.IndexingComplete = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

// The name of the policy to use to manage the index. For information about how
// Elasticsearch applies policy changes, see Policy updates.
func (s *_indexSettingsLifecycle) Name(name string) *_indexSettingsLifecycle {

	s.v.Name = &name

	return s
}

// If specified, this is the timestamp used to calculate the index age for its
// phase transitions. Use this setting
// if you create a new index that contains old data and want to use the original
// creation date to calculate the index
// age. Specified as a Unix epoch value in milliseconds.
func (s *_indexSettingsLifecycle) OriginationDate(originationdate int64) *_indexSettingsLifecycle {

	s.v.OriginationDate = &originationdate

	return s
}

// Set to true to parse the origination date from the index name. This
// origination date is used to calculate the index age
// for its phase transitions. The index name must match the pattern
// ^.*-{date_format}-\\d+, where the date_format is
// yyyy.MM.dd and the trailing digits are optional. An index that was rolled
// over would normally match the full format,
// for example logs-2016.10.31-000002). If the index name doesn’t match the
// pattern, index creation fails.
func (s *_indexSettingsLifecycle) ParseOriginationDate(parseoriginationdate bool) *_indexSettingsLifecycle {

	s.v.ParseOriginationDate = &parseoriginationdate

	return s
}

// Preference for the system that manages a data stream backing index
// (preferring ILM when both ILM and DLM are
// applicable for an index).
func (s *_indexSettingsLifecycle) PreferIlm(preferilm string) *_indexSettingsLifecycle {

	s.v.PreferIlm = preferilm

	return s
}

// The index alias to update when the index rolls over. Specify when using a
// policy that contains a rollover action.
// When the index rolls over, the alias is updated to reflect that the index is
// no longer the write index. For more
// information about rolling indices, see Rollover.
func (s *_indexSettingsLifecycle) RolloverAlias(rolloveralias string) *_indexSettingsLifecycle {

	s.v.RolloverAlias = &rolloveralias

	return s
}

func (s *_indexSettingsLifecycle) Step(step types.IndexSettingsLifecycleStepVariant) *_indexSettingsLifecycle {

	s.v.Step = step.IndexSettingsLifecycleStepCaster()

	return s
}

func (s *_indexSettingsLifecycle) IndexSettingsLifecycleCaster() *types.IndexSettingsLifecycle {
	return s.v
}
