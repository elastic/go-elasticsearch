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

// SlackAttachment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L101-L117
type SlackAttachment struct {
	AuthorIcon *string                `json:"author_icon,omitempty"`
	AuthorLink *string                `json:"author_link,omitempty"`
	AuthorName string                 `json:"author_name"`
	Color      *string                `json:"color,omitempty"`
	Fallback   *string                `json:"fallback,omitempty"`
	Fields     []SlackAttachmentField `json:"fields,omitempty"`
	Footer     *string                `json:"footer,omitempty"`
	FooterIcon *string                `json:"footer_icon,omitempty"`
	ImageUrl   *string                `json:"image_url,omitempty"`
	Pretext    *string                `json:"pretext,omitempty"`
	Text       *string                `json:"text,omitempty"`
	ThumbUrl   *string                `json:"thumb_url,omitempty"`
	Title      string                 `json:"title"`
	TitleLink  *string                `json:"title_link,omitempty"`
	Ts         *EpochTimeUnitSeconds  `json:"ts,omitempty"`
}

// SlackAttachmentBuilder holds SlackAttachment struct and provides a builder API.
type SlackAttachmentBuilder struct {
	v *SlackAttachment
}

// NewSlackAttachment provides a builder for the SlackAttachment struct.
func NewSlackAttachmentBuilder() *SlackAttachmentBuilder {
	r := SlackAttachmentBuilder{
		&SlackAttachment{},
	}

	return &r
}

// Build finalize the chain and returns the SlackAttachment struct
func (rb *SlackAttachmentBuilder) Build() SlackAttachment {
	return *rb.v
}

func (rb *SlackAttachmentBuilder) AuthorIcon(authoricon string) *SlackAttachmentBuilder {
	rb.v.AuthorIcon = &authoricon
	return rb
}

func (rb *SlackAttachmentBuilder) AuthorLink(authorlink string) *SlackAttachmentBuilder {
	rb.v.AuthorLink = &authorlink
	return rb
}

func (rb *SlackAttachmentBuilder) AuthorName(authorname string) *SlackAttachmentBuilder {
	rb.v.AuthorName = authorname
	return rb
}

func (rb *SlackAttachmentBuilder) Color(color string) *SlackAttachmentBuilder {
	rb.v.Color = &color
	return rb
}

func (rb *SlackAttachmentBuilder) Fallback(fallback string) *SlackAttachmentBuilder {
	rb.v.Fallback = &fallback
	return rb
}

func (rb *SlackAttachmentBuilder) Fields(fields []SlackAttachmentFieldBuilder) *SlackAttachmentBuilder {
	tmp := make([]SlackAttachmentField, len(fields))
	for _, value := range fields {
		tmp = append(tmp, value.Build())
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *SlackAttachmentBuilder) Footer(footer string) *SlackAttachmentBuilder {
	rb.v.Footer = &footer
	return rb
}

func (rb *SlackAttachmentBuilder) FooterIcon(footericon string) *SlackAttachmentBuilder {
	rb.v.FooterIcon = &footericon
	return rb
}

func (rb *SlackAttachmentBuilder) ImageUrl(imageurl string) *SlackAttachmentBuilder {
	rb.v.ImageUrl = &imageurl
	return rb
}

func (rb *SlackAttachmentBuilder) Pretext(pretext string) *SlackAttachmentBuilder {
	rb.v.Pretext = &pretext
	return rb
}

func (rb *SlackAttachmentBuilder) Text(text string) *SlackAttachmentBuilder {
	rb.v.Text = &text
	return rb
}

func (rb *SlackAttachmentBuilder) ThumbUrl(thumburl string) *SlackAttachmentBuilder {
	rb.v.ThumbUrl = &thumburl
	return rb
}

func (rb *SlackAttachmentBuilder) Title(title string) *SlackAttachmentBuilder {
	rb.v.Title = title
	return rb
}

func (rb *SlackAttachmentBuilder) TitleLink(titlelink string) *SlackAttachmentBuilder {
	rb.v.TitleLink = &titlelink
	return rb
}

func (rb *SlackAttachmentBuilder) Ts(ts *EpochTimeUnitSecondsBuilder) *SlackAttachmentBuilder {
	v := ts.Build()
	rb.v.Ts = &v
	return rb
}
