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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/pagerdutyeventtype"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// PagerDutyEvent type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/watcher/_types/Actions.ts#L40-L52
type PagerDutyEvent struct {
	Account       *string                                `json:"account,omitempty"`
	AttachPayload bool                                   `json:"attach_payload"`
	Client        *string                                `json:"client,omitempty"`
	ClientUrl     *string                                `json:"client_url,omitempty"`
	Contexts      []PagerDutyContext                     `json:"contexts,omitempty"`
	Description   string                                 `json:"description"`
	EventType     *pagerdutyeventtype.PagerDutyEventType `json:"event_type,omitempty"`
	IncidentKey   string                                 `json:"incident_key"`
	Proxy         *PagerDutyEventProxy                   `json:"proxy,omitempty"`
}

func (s *PagerDutyEvent) UnmarshalJSON(data []byte) error {

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

		case "account":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Account = &o

		case "attach_payload":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AttachPayload = value
			case bool:
				s.AttachPayload = v
			}

		case "client":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Client = &o

		case "client_url":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.ClientUrl = &o

		case "contexts", "context":
			if err := dec.Decode(&s.Contexts); err != nil {
				return err
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Description = o

		case "event_type":
			if err := dec.Decode(&s.EventType); err != nil {
				return err
			}

		case "incident_key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.IncidentKey = o

		case "proxy":
			if err := dec.Decode(&s.Proxy); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewPagerDutyEvent returns a PagerDutyEvent.
func NewPagerDutyEvent() *PagerDutyEvent {
	r := &PagerDutyEvent{}

	return r
}
