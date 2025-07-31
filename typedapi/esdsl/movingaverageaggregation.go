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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

// This is provide all the types that are part of the union.
type _movingAverageAggregation struct {
	v types.MovingAverageAggregation
}

func NewMovingAverageAggregation() *_movingAverageAggregation {
	return &_movingAverageAggregation{v: nil}
}

func (u *_movingAverageAggregation) LinearMovingAverageAggregation(linearmovingaverageaggregation types.LinearMovingAverageAggregationVariant) *_movingAverageAggregation {

	u.v = linearmovingaverageaggregation.LinearMovingAverageAggregationCaster()

	return u
}

// Interface implementation for LinearMovingAverageAggregation in MovingAverageAggregation union
func (u *_linearMovingAverageAggregation) MovingAverageAggregationCaster() *types.MovingAverageAggregation {
	t := types.MovingAverageAggregation(u.v)
	return &t
}

func (u *_movingAverageAggregation) SimpleMovingAverageAggregation(simplemovingaverageaggregation types.SimpleMovingAverageAggregationVariant) *_movingAverageAggregation {

	u.v = simplemovingaverageaggregation.SimpleMovingAverageAggregationCaster()

	return u
}

// Interface implementation for SimpleMovingAverageAggregation in MovingAverageAggregation union
func (u *_simpleMovingAverageAggregation) MovingAverageAggregationCaster() *types.MovingAverageAggregation {
	t := types.MovingAverageAggregation(u.v)
	return &t
}

func (u *_movingAverageAggregation) EwmaMovingAverageAggregation(ewmamovingaverageaggregation types.EwmaMovingAverageAggregationVariant) *_movingAverageAggregation {

	u.v = ewmamovingaverageaggregation.EwmaMovingAverageAggregationCaster()

	return u
}

// Interface implementation for EwmaMovingAverageAggregation in MovingAverageAggregation union
func (u *_ewmaMovingAverageAggregation) MovingAverageAggregationCaster() *types.MovingAverageAggregation {
	t := types.MovingAverageAggregation(u.v)
	return &t
}

func (u *_movingAverageAggregation) HoltMovingAverageAggregation(holtmovingaverageaggregation types.HoltMovingAverageAggregationVariant) *_movingAverageAggregation {

	u.v = holtmovingaverageaggregation.HoltMovingAverageAggregationCaster()

	return u
}

// Interface implementation for HoltMovingAverageAggregation in MovingAverageAggregation union
func (u *_holtMovingAverageAggregation) MovingAverageAggregationCaster() *types.MovingAverageAggregation {
	t := types.MovingAverageAggregation(u.v)
	return &t
}

func (u *_movingAverageAggregation) HoltWintersMovingAverageAggregation(holtwintersmovingaverageaggregation types.HoltWintersMovingAverageAggregationVariant) *_movingAverageAggregation {

	u.v = holtwintersmovingaverageaggregation.HoltWintersMovingAverageAggregationCaster()

	return u
}

// Interface implementation for HoltWintersMovingAverageAggregation in MovingAverageAggregation union
func (u *_holtWintersMovingAverageAggregation) MovingAverageAggregationCaster() *types.MovingAverageAggregation {
	t := types.MovingAverageAggregation(u.v)
	return &t
}

func (u *_movingAverageAggregation) MovingAverageAggregationCaster() *types.MovingAverageAggregation {
	return &u.v
}
