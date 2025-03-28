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

type _connectorScheduling struct {
	v *types.ConnectorScheduling
}

func NewConnectorScheduling(enabled bool, interval string) *_connectorScheduling {

	tmp := &_connectorScheduling{v: types.NewConnectorScheduling()}

	tmp.Enabled(enabled)

	tmp.Interval(interval)

	return tmp

}

func (s *_connectorScheduling) Enabled(enabled bool) *_connectorScheduling {

	s.v.Enabled = enabled

	return s
}

// The interval is expressed using the crontab syntax
func (s *_connectorScheduling) Interval(interval string) *_connectorScheduling {

	s.v.Interval = interval

	return s
}

func (s *_connectorScheduling) ConnectorSchedulingCaster() *types.ConnectorScheduling {
	return s.v
}
