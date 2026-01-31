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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cohereinputtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/coheretruncatetype"
)

type _cohereTaskSettings struct {
	v *types.CohereTaskSettings
}

func NewCohereTaskSettings(inputtype cohereinputtype.CohereInputType) *_cohereTaskSettings {

	tmp := &_cohereTaskSettings{v: types.NewCohereTaskSettings()}

	tmp.InputType(inputtype)

	return tmp

}

func (s *_cohereTaskSettings) InputType(inputtype cohereinputtype.CohereInputType) *_cohereTaskSettings {

	s.v.InputType = inputtype
	return s
}

func (s *_cohereTaskSettings) ReturnDocuments(returndocuments bool) *_cohereTaskSettings {

	s.v.ReturnDocuments = &returndocuments

	return s
}

func (s *_cohereTaskSettings) TopN(topn int) *_cohereTaskSettings {

	s.v.TopN = &topn

	return s
}

func (s *_cohereTaskSettings) Truncate(truncate coheretruncatetype.CohereTruncateType) *_cohereTaskSettings {

	s.v.Truncate = &truncate
	return s
}

func (s *_cohereTaskSettings) CohereTaskSettingsCaster() *types.CohereTaskSettings {
	return s.v
}
