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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _updateAction struct {
	v *types.UpdateAction
}

func NewUpdateAction() *_updateAction {

	return &_updateAction{v: types.NewUpdateAction()}

}

func (s *_updateAction) DetectNoop(detectnoop bool) *_updateAction {

	s.v.DetectNoop = &detectnoop

	return s
}

func (s *_updateAction) Doc(doc json.RawMessage) *_updateAction {

	s.v.Doc = doc

	return s
}

func (s *_updateAction) DocAsUpsert(docasupsert bool) *_updateAction {

	s.v.DocAsUpsert = &docasupsert

	return s
}

func (s *_updateAction) Script(script types.ScriptVariant) *_updateAction {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_updateAction) ScriptedUpsert(scriptedupsert bool) *_updateAction {

	s.v.ScriptedUpsert = &scriptedupsert

	return s
}

func (s *_updateAction) Source_(sourceconfig types.SourceConfigVariant) *_updateAction {

	s.v.Source_ = *sourceconfig.SourceConfigCaster()

	return s
}

func (s *_updateAction) Upsert(upsert json.RawMessage) *_updateAction {

	s.v.Upsert = upsert

	return s
}

func (s *_updateAction) UpdateActionCaster() *types.UpdateAction {
	return s.v
}
