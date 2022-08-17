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

// SlackMessage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L130-L137
type SlackMessage struct {
	Attachments        []SlackAttachment       `json:"attachments"`
	DynamicAttachments *SlackDynamicAttachment `json:"dynamic_attachments,omitempty"`
	From               string                  `json:"from"`
	Icon               *string                 `json:"icon,omitempty"`
	Text               string                  `json:"text"`
	To                 []string                `json:"to"`
}

// SlackMessageBuilder holds SlackMessage struct and provides a builder API.
type SlackMessageBuilder struct {
	v *SlackMessage
}

// NewSlackMessage provides a builder for the SlackMessage struct.
func NewSlackMessageBuilder() *SlackMessageBuilder {
	r := SlackMessageBuilder{
		&SlackMessage{},
	}

	return &r
}

// Build finalize the chain and returns the SlackMessage struct
func (rb *SlackMessageBuilder) Build() SlackMessage {
	return *rb.v
}

func (rb *SlackMessageBuilder) Attachments(attachments []SlackAttachmentBuilder) *SlackMessageBuilder {
	tmp := make([]SlackAttachment, len(attachments))
	for _, value := range attachments {
		tmp = append(tmp, value.Build())
	}
	rb.v.Attachments = tmp
	return rb
}

func (rb *SlackMessageBuilder) DynamicAttachments(dynamicattachments *SlackDynamicAttachmentBuilder) *SlackMessageBuilder {
	v := dynamicattachments.Build()
	rb.v.DynamicAttachments = &v
	return rb
}

func (rb *SlackMessageBuilder) From(from string) *SlackMessageBuilder {
	rb.v.From = from
	return rb
}

func (rb *SlackMessageBuilder) Icon(icon string) *SlackMessageBuilder {
	rb.v.Icon = &icon
	return rb
}

func (rb *SlackMessageBuilder) Text(text string) *SlackMessageBuilder {
	rb.v.Text = text
	return rb
}

func (rb *SlackMessageBuilder) To(to ...string) *SlackMessageBuilder {
	rb.v.To = to
	return rb
}
