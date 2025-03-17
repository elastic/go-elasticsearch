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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/translogdurability"
)

type _translog struct {
	v *types.Translog
}

func NewTranslog() *_translog {

	return &_translog{v: types.NewTranslog()}

}

// Whether or not to `fsync` and commit the translog after every index, delete,
// update, or bulk request.
func (s *_translog) Durability(durability translogdurability.TranslogDurability) *_translog {

	s.v.Durability = &durability
	return s
}

// The translog stores all operations that are not yet safely persisted in
// Lucene (i.e., are not
// part of a Lucene commit point). Although these operations are available for
// reads, they will need
// to be replayed if the shard was stopped and had to be recovered. This setting
// controls the
// maximum total size of these operations, to prevent recoveries from taking too
// long. Once the
// maximum size has been reached a flush will happen, generating a new Lucene
// commit point.
func (s *_translog) FlushThresholdSize(bytesize types.ByteSizeVariant) *_translog {

	s.v.FlushThresholdSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_translog) Retention(retention types.TranslogRetentionVariant) *_translog {

	s.v.Retention = retention.TranslogRetentionCaster()

	return s
}

// How often the translog is fsynced to disk and committed, regardless of write
// operations.
// Values less than 100ms are not allowed.
func (s *_translog) SyncInterval(duration types.DurationVariant) *_translog {

	s.v.SyncInterval = *duration.DurationCaster()

	return s
}

func (s *_translog) TranslogCaster() *types.Translog {
	return s.v
}
