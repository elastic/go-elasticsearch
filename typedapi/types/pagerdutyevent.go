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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/pagerdutyeventtype"
)

// PagerDutyEvent type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L40-L52
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

// PagerDutyEventBuilder holds PagerDutyEvent struct and provides a builder API.
type PagerDutyEventBuilder struct {
	v *PagerDutyEvent
}

// NewPagerDutyEvent provides a builder for the PagerDutyEvent struct.
func NewPagerDutyEventBuilder() *PagerDutyEventBuilder {
	r := PagerDutyEventBuilder{
		&PagerDutyEvent{},
	}

	return &r
}

// Build finalize the chain and returns the PagerDutyEvent struct
func (rb *PagerDutyEventBuilder) Build() PagerDutyEvent {
	return *rb.v
}

func (rb *PagerDutyEventBuilder) Account(account string) *PagerDutyEventBuilder {
	rb.v.Account = &account
	return rb
}

func (rb *PagerDutyEventBuilder) AttachPayload(attachpayload bool) *PagerDutyEventBuilder {
	rb.v.AttachPayload = attachpayload
	return rb
}

func (rb *PagerDutyEventBuilder) Client(client string) *PagerDutyEventBuilder {
	rb.v.Client = &client
	return rb
}

func (rb *PagerDutyEventBuilder) ClientUrl(clienturl string) *PagerDutyEventBuilder {
	rb.v.ClientUrl = &clienturl
	return rb
}

func (rb *PagerDutyEventBuilder) Contexts(contexts []PagerDutyContextBuilder) *PagerDutyEventBuilder {
	tmp := make([]PagerDutyContext, len(contexts))
	for _, value := range contexts {
		tmp = append(tmp, value.Build())
	}
	rb.v.Contexts = tmp
	return rb
}

func (rb *PagerDutyEventBuilder) Description(description string) *PagerDutyEventBuilder {
	rb.v.Description = description
	return rb
}

func (rb *PagerDutyEventBuilder) EventType(eventtype pagerdutyeventtype.PagerDutyEventType) *PagerDutyEventBuilder {
	rb.v.EventType = &eventtype
	return rb
}

func (rb *PagerDutyEventBuilder) IncidentKey(incidentkey string) *PagerDutyEventBuilder {
	rb.v.IncidentKey = incidentkey
	return rb
}

func (rb *PagerDutyEventBuilder) Proxy(proxy *PagerDutyEventProxyBuilder) *PagerDutyEventBuilder {
	v := proxy.Build()
	rb.v.Proxy = &v
	return rb
}
