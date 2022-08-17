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

// PagerDutyAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Actions.ts#L54-L54
type PagerDutyAction struct {
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

// PagerDutyActionBuilder holds PagerDutyAction struct and provides a builder API.
type PagerDutyActionBuilder struct {
	v *PagerDutyAction
}

// NewPagerDutyAction provides a builder for the PagerDutyAction struct.
func NewPagerDutyActionBuilder() *PagerDutyActionBuilder {
	r := PagerDutyActionBuilder{
		&PagerDutyAction{},
	}

	return &r
}

// Build finalize the chain and returns the PagerDutyAction struct
func (rb *PagerDutyActionBuilder) Build() PagerDutyAction {
	return *rb.v
}

func (rb *PagerDutyActionBuilder) Account(account string) *PagerDutyActionBuilder {
	rb.v.Account = &account
	return rb
}

func (rb *PagerDutyActionBuilder) AttachPayload(attachpayload bool) *PagerDutyActionBuilder {
	rb.v.AttachPayload = attachpayload
	return rb
}

func (rb *PagerDutyActionBuilder) Client(client string) *PagerDutyActionBuilder {
	rb.v.Client = &client
	return rb
}

func (rb *PagerDutyActionBuilder) ClientUrl(clienturl string) *PagerDutyActionBuilder {
	rb.v.ClientUrl = &clienturl
	return rb
}

func (rb *PagerDutyActionBuilder) Contexts(contexts []PagerDutyContextBuilder) *PagerDutyActionBuilder {
	tmp := make([]PagerDutyContext, len(contexts))
	for _, value := range contexts {
		tmp = append(tmp, value.Build())
	}
	rb.v.Contexts = tmp
	return rb
}

func (rb *PagerDutyActionBuilder) Description(description string) *PagerDutyActionBuilder {
	rb.v.Description = description
	return rb
}

func (rb *PagerDutyActionBuilder) EventType(eventtype pagerdutyeventtype.PagerDutyEventType) *PagerDutyActionBuilder {
	rb.v.EventType = &eventtype
	return rb
}

func (rb *PagerDutyActionBuilder) IncidentKey(incidentkey string) *PagerDutyActionBuilder {
	rb.v.IncidentKey = incidentkey
	return rb
}

func (rb *PagerDutyActionBuilder) Proxy(proxy *PagerDutyEventProxyBuilder) *PagerDutyActionBuilder {
	v := proxy.Build()
	rb.v.Proxy = &v
	return rb
}
