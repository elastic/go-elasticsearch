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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/emailpriority"
)

// EmailAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L252-L252
type EmailAction struct {
	Attachments map[string]EmailAttachmentContainer `json:"attachments,omitempty"`
	Bcc         []string                            `json:"bcc,omitempty"`
	Body        *EmailBody                          `json:"body,omitempty"`
	Cc          []string                            `json:"cc,omitempty"`
	From        *string                             `json:"from,omitempty"`
	Id          *Id                                 `json:"id,omitempty"`
	Priority    *emailpriority.EmailPriority        `json:"priority,omitempty"`
	ReplyTo     []string                            `json:"reply_to,omitempty"`
	SentDate    *DateTime                           `json:"sent_date,omitempty"`
	Subject     string                              `json:"subject"`
	To          []string                            `json:"to"`
}

// EmailActionBuilder holds EmailAction struct and provides a builder API.
type EmailActionBuilder struct {
	v *EmailAction
}

// NewEmailAction provides a builder for the EmailAction struct.
func NewEmailActionBuilder() *EmailActionBuilder {
	r := EmailActionBuilder{
		&EmailAction{
			Attachments: make(map[string]EmailAttachmentContainer, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the EmailAction struct
func (rb *EmailActionBuilder) Build() EmailAction {
	return *rb.v
}

func (rb *EmailActionBuilder) Attachments(values map[string]*EmailAttachmentContainerBuilder) *EmailActionBuilder {
	tmp := make(map[string]EmailAttachmentContainer, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Attachments = tmp
	return rb
}

func (rb *EmailActionBuilder) Bcc(bcc ...string) *EmailActionBuilder {
	rb.v.Bcc = bcc
	return rb
}

func (rb *EmailActionBuilder) Body(body *EmailBodyBuilder) *EmailActionBuilder {
	v := body.Build()
	rb.v.Body = &v
	return rb
}

func (rb *EmailActionBuilder) Cc(cc ...string) *EmailActionBuilder {
	rb.v.Cc = cc
	return rb
}

func (rb *EmailActionBuilder) From(from string) *EmailActionBuilder {
	rb.v.From = &from
	return rb
}

func (rb *EmailActionBuilder) Id(id Id) *EmailActionBuilder {
	rb.v.Id = &id
	return rb
}

func (rb *EmailActionBuilder) Priority(priority emailpriority.EmailPriority) *EmailActionBuilder {
	rb.v.Priority = &priority
	return rb
}

func (rb *EmailActionBuilder) ReplyTo(reply_to ...string) *EmailActionBuilder {
	rb.v.ReplyTo = reply_to
	return rb
}

func (rb *EmailActionBuilder) SentDate(sentdate *DateTimeBuilder) *EmailActionBuilder {
	v := sentdate.Build()
	rb.v.SentDate = &v
	return rb
}

func (rb *EmailActionBuilder) Subject(subject string) *EmailActionBuilder {
	rb.v.Subject = subject
	return rb
}

func (rb *EmailActionBuilder) To(to ...string) *EmailActionBuilder {
	rb.v.To = to
	return rb
}
