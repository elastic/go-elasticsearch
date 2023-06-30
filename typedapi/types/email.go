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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/emailpriority"
)

// Email type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/watcher/_types/Actions.ts#L238-L250
type Email struct {
	Attachments map[string]EmailAttachmentContainer `json:"attachments,omitempty"`
	Bcc         []string                            `json:"bcc,omitempty"`
	Body        *EmailBody                          `json:"body,omitempty"`
	Cc          []string                            `json:"cc,omitempty"`
	From        *string                             `json:"from,omitempty"`
	Id          *string                             `json:"id,omitempty"`
	Priority    *emailpriority.EmailPriority        `json:"priority,omitempty"`
	ReplyTo     []string                            `json:"reply_to,omitempty"`
	SentDate    DateTime                            `json:"sent_date,omitempty"`
	Subject     string                              `json:"subject"`
	To          []string                            `json:"to"`
}

func (s *Email) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "attachments":
			if s.Attachments == nil {
				s.Attachments = make(map[string]EmailAttachmentContainer, 0)
			}
			if err := dec.Decode(&s.Attachments); err != nil {
				return err
			}

		case "bcc":
			if err := dec.Decode(&s.Bcc); err != nil {
				return err
			}

		case "body":
			if err := dec.Decode(&s.Body); err != nil {
				return err
			}

		case "cc":
			if err := dec.Decode(&s.Cc); err != nil {
				return err
			}

		case "from":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.From = &o

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "priority":
			if err := dec.Decode(&s.Priority); err != nil {
				return err
			}

		case "reply_to":
			if err := dec.Decode(&s.ReplyTo); err != nil {
				return err
			}

		case "sent_date":
			if err := dec.Decode(&s.SentDate); err != nil {
				return err
			}

		case "subject":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Subject = o

		case "to":
			if err := dec.Decode(&s.To); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewEmail returns a Email.
func NewEmail() *Email {
	r := &Email{
		Attachments: make(map[string]EmailAttachmentContainer, 0),
	}

	return r
}
