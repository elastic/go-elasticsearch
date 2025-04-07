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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _reportingEmailAttachment struct {
	v *types.ReportingEmailAttachment
}

func NewReportingEmailAttachment(url string) *_reportingEmailAttachment {

	tmp := &_reportingEmailAttachment{v: types.NewReportingEmailAttachment()}

	tmp.Url(url)

	return tmp

}

func (s *_reportingEmailAttachment) Inline(inline bool) *_reportingEmailAttachment {

	s.v.Inline = &inline

	return s
}

func (s *_reportingEmailAttachment) Interval(duration types.DurationVariant) *_reportingEmailAttachment {

	s.v.Interval = *duration.DurationCaster()

	return s
}

func (s *_reportingEmailAttachment) Request(request types.HttpInputRequestDefinitionVariant) *_reportingEmailAttachment {

	s.v.Request = request.HttpInputRequestDefinitionCaster()

	return s
}

func (s *_reportingEmailAttachment) Retries(retries int) *_reportingEmailAttachment {

	s.v.Retries = &retries

	return s
}

func (s *_reportingEmailAttachment) Url(url string) *_reportingEmailAttachment {

	s.v.Url = url

	return s
}

func (s *_reportingEmailAttachment) EmailAttachmentContainerCaster() *types.EmailAttachmentContainer {
	container := types.NewEmailAttachmentContainer()

	container.Reporting = s.v

	return container
}

func (s *_reportingEmailAttachment) ReportingEmailAttachmentCaster() *types.ReportingEmailAttachment {
	return s.v
}
