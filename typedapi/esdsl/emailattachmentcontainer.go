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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _emailAttachmentContainer struct {
	v *types.EmailAttachmentContainer
}

func NewEmailAttachmentContainer() *_emailAttachmentContainer {
	return &_emailAttachmentContainer{v: types.NewEmailAttachmentContainer()}
}

// AdditionalEmailAttachmentContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_emailAttachmentContainer) AdditionalEmailAttachmentContainerProperty(key string, value json.RawMessage) *_emailAttachmentContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalEmailAttachmentContainerProperty = tmp
	return s
}

func (s *_emailAttachmentContainer) Data(data types.DataEmailAttachmentVariant) *_emailAttachmentContainer {

	s.v.Data = data.DataEmailAttachmentCaster()

	return s
}

func (s *_emailAttachmentContainer) Http(http types.HttpEmailAttachmentVariant) *_emailAttachmentContainer {

	s.v.Http = http.HttpEmailAttachmentCaster()

	return s
}

func (s *_emailAttachmentContainer) Reporting(reporting types.ReportingEmailAttachmentVariant) *_emailAttachmentContainer {

	s.v.Reporting = reporting.ReportingEmailAttachmentCaster()

	return s
}

func (s *_emailAttachmentContainer) EmailAttachmentContainerCaster() *types.EmailAttachmentContainer {
	return s.v
}
