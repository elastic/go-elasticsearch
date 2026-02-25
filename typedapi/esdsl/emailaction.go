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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/emailpriority"
)

type _emailAction struct {
	v *types.EmailAction
}

func NewEmailAction(subject string) *_emailAction {

	tmp := &_emailAction{v: types.NewEmailAction()}

	tmp.Subject(subject)

	return tmp

}

func (s *_emailAction) Attachments(attachments map[string]types.EmailAttachmentContainer) *_emailAction {

	s.v.Attachments = attachments
	return s
}

func (s *_emailAction) AddAttachment(key string, value types.EmailAttachmentContainerVariant) *_emailAction {

	var tmp map[string]types.EmailAttachmentContainer
	if s.v.Attachments == nil {
		s.v.Attachments = make(map[string]types.EmailAttachmentContainer)
	} else {
		tmp = s.v.Attachments
	}

	tmp[key] = *value.EmailAttachmentContainerCaster()

	s.v.Attachments = tmp
	return s
}

func (s *_emailAction) Bcc(bccs ...string) *_emailAction {

	s.v.Bcc = make([]string, len(bccs))
	s.v.Bcc = bccs

	return s
}

func (s *_emailAction) Body(body types.EmailBodyVariant) *_emailAction {

	s.v.Body = body.EmailBodyCaster()

	return s
}

func (s *_emailAction) Cc(ccs ...string) *_emailAction {

	s.v.Cc = make([]string, len(ccs))
	s.v.Cc = ccs

	return s
}

func (s *_emailAction) From(from string) *_emailAction {

	s.v.From = &from

	return s
}

func (s *_emailAction) Id(id string) *_emailAction {

	s.v.Id = &id

	return s
}

func (s *_emailAction) Priority(priority emailpriority.EmailPriority) *_emailAction {

	s.v.Priority = &priority
	return s
}

func (s *_emailAction) ReplyTo(replytos ...string) *_emailAction {

	s.v.ReplyTo = make([]string, len(replytos))
	s.v.ReplyTo = replytos

	return s
}

func (s *_emailAction) SentDate(datetime types.DateTimeVariant) *_emailAction {

	s.v.SentDate = *datetime.DateTimeCaster()

	return s
}

func (s *_emailAction) Subject(subject string) *_emailAction {

	s.v.Subject = subject

	return s
}

func (s *_emailAction) To(tos ...string) *_emailAction {

	s.v.To = make([]string, len(tos))
	s.v.To = tos

	return s
}

func (s *_emailAction) EmailActionCaster() *types.EmailAction {
	return s.v
}
