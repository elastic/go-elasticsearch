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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// SlackAttachment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/watcher/_types/Actions.ts#L101-L117
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
	Ts         *int64                 `json:"ts,omitempty"`
}

// NewSlackAttachment returns a SlackAttachment.
func NewSlackAttachment() *SlackAttachment {
	r := &SlackAttachment{}

	return r
}
