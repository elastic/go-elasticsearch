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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _slackMessage struct {
	v *types.SlackMessage
}

func NewSlackMessage(from string, text string) *_slackMessage {

	tmp := &_slackMessage{v: types.NewSlackMessage()}

	tmp.From(from)

	tmp.Text(text)

	return tmp

}

func (s *_slackMessage) Attachments(attachments ...types.SlackAttachmentVariant) *_slackMessage {

	for _, v := range attachments {

		s.v.Attachments = append(s.v.Attachments, *v.SlackAttachmentCaster())

	}
	return s
}

func (s *_slackMessage) DynamicAttachments(dynamicattachments types.SlackDynamicAttachmentVariant) *_slackMessage {

	s.v.DynamicAttachments = dynamicattachments.SlackDynamicAttachmentCaster()

	return s
}

func (s *_slackMessage) From(from string) *_slackMessage {

	s.v.From = from

	return s
}

func (s *_slackMessage) Icon(icon string) *_slackMessage {

	s.v.Icon = &icon

	return s
}

func (s *_slackMessage) Text(text string) *_slackMessage {

	s.v.Text = text

	return s
}

func (s *_slackMessage) To(tos ...string) *_slackMessage {

	for _, v := range tos {

		s.v.To = append(s.v.To, v)

	}
	return s
}

func (s *_slackMessage) SlackMessageCaster() *types.SlackMessage {
	return s.v
}
