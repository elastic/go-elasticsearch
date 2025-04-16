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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/translogdurability"
)

type _translog struct {
	v *types.Translog
}

func NewTranslog() *_translog {

	return &_translog{v: types.NewTranslog()}

}

func (s *_translog) Durability(durability translogdurability.TranslogDurability) *_translog {

	s.v.Durability = &durability
	return s
}

func (s *_translog) FlushThresholdSize(bytesize types.ByteSizeVariant) *_translog {

	s.v.FlushThresholdSize = *bytesize.ByteSizeCaster()

	return s
}

func (s *_translog) Retention(retention types.TranslogRetentionVariant) *_translog {

	s.v.Retention = retention.TranslogRetentionCaster()

	return s
}

func (s *_translog) SyncInterval(duration types.DurationVariant) *_translog {

	s.v.SyncInterval = *duration.DurationCaster()

	return s
}

func (s *_translog) TranslogCaster() *types.Translog {
	return s.v
}
