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

type _translogRetention struct {
	v *types.TranslogRetention
}

func NewTranslogRetention() *_translogRetention {

	return &_translogRetention{v: types.NewTranslogRetention()}

}

// This controls the maximum duration for which translog files are kept by each
// shard. Keeping more
// translog files increases the chance of performing an operation based sync
// when recovering replicas. If
// the translog files are not sufficient, replica recovery will fall back to a
// file based sync. This setting
// is ignored, and should not be set, if soft deletes are enabled. Soft deletes
// are enabled by default in
// indices created in Elasticsearch versions 7.0.0 and later.
func (s *_translogRetention) Age(duration types.DurationVariant) *_translogRetention {

	s.v.Age = *duration.DurationCaster()

	return s
}

// This controls the total size of translog files to keep for each shard.
// Keeping more translog files increases
// the chance of performing an operation based sync when recovering a replica.
// If the translog files are not
// sufficient, replica recovery will fall back to a file based sync. This
// setting is ignored, and should not be
// set, if soft deletes are enabled. Soft deletes are enabled by default in
// indices created in Elasticsearch
// versions 7.0.0 and later.
func (s *_translogRetention) Size(bytesize types.ByteSizeVariant) *_translogRetention {

	s.v.Size = *bytesize.ByteSizeCaster()

	return s
}

func (s *_translogRetention) TranslogRetentionCaster() *types.TranslogRetention {
	return s.v
}
