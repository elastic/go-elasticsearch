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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _slackAttachment struct {
	v *types.SlackAttachment
}

func NewSlackAttachment(authorname string, title string) *_slackAttachment {

	tmp := &_slackAttachment{v: types.NewSlackAttachment()}

	tmp.AuthorName(authorname)

	tmp.Title(title)

	return tmp

}

func (s *_slackAttachment) AuthorIcon(authoricon string) *_slackAttachment {

	s.v.AuthorIcon = &authoricon

	return s
}

func (s *_slackAttachment) AuthorLink(authorlink string) *_slackAttachment {

	s.v.AuthorLink = &authorlink

	return s
}

func (s *_slackAttachment) AuthorName(authorname string) *_slackAttachment {

	s.v.AuthorName = authorname

	return s
}

func (s *_slackAttachment) Color(color string) *_slackAttachment {

	s.v.Color = &color

	return s
}

func (s *_slackAttachment) Fallback(fallback string) *_slackAttachment {

	s.v.Fallback = &fallback

	return s
}

func (s *_slackAttachment) Fields(fields ...types.SlackAttachmentFieldVariant) *_slackAttachment {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, *v.SlackAttachmentFieldCaster())

	}
	return s
}

func (s *_slackAttachment) Footer(footer string) *_slackAttachment {

	s.v.Footer = &footer

	return s
}

func (s *_slackAttachment) FooterIcon(footericon string) *_slackAttachment {

	s.v.FooterIcon = &footericon

	return s
}

func (s *_slackAttachment) ImageUrl(imageurl string) *_slackAttachment {

	s.v.ImageUrl = &imageurl

	return s
}

func (s *_slackAttachment) Pretext(pretext string) *_slackAttachment {

	s.v.Pretext = &pretext

	return s
}

func (s *_slackAttachment) Text(text string) *_slackAttachment {

	s.v.Text = &text

	return s
}

func (s *_slackAttachment) ThumbUrl(thumburl string) *_slackAttachment {

	s.v.ThumbUrl = &thumburl

	return s
}

func (s *_slackAttachment) Title(title string) *_slackAttachment {

	s.v.Title = title

	return s
}

func (s *_slackAttachment) TitleLink(titlelink string) *_slackAttachment {

	s.v.TitleLink = &titlelink

	return s
}

func (s *_slackAttachment) Ts(epochtimeunitseconds int64) *_slackAttachment {

	s.v.Ts = &epochtimeunitseconds

	return s
}

func (s *_slackAttachment) SlackAttachmentCaster() *types.SlackAttachment {
	return s.v
}
