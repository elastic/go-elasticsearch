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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package types

import (
	"encoding/json"
)

// WatcherInput type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b1811e10a0722431d79d1c234dd412ff47d8656f/specification/watcher/_types/Input.ts#L87-L95
type WatcherInput struct {
	Chain  *ChainInput                `json:"chain,omitempty"`
	Http   *HttpInput                 `json:"http,omitempty"`
	Search *SearchInput               `json:"search,omitempty"`
	Simple map[string]json.RawMessage `json:"simple,omitempty"`
}

// NewWatcherInput returns a WatcherInput.
func NewWatcherInput() *WatcherInput {
	r := &WatcherInput{
		Simple: make(map[string]json.RawMessage),
	}

	return r
}

type WatcherInputVariant interface {
	WatcherInputCaster() *WatcherInput
}

func (s *WatcherInput) WatcherInputCaster() *WatcherInput {
	return s
}
