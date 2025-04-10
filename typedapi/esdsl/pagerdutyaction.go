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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/pagerdutyeventtype"
)

type _pagerDutyAction struct {
	v *types.PagerDutyAction
}

func NewPagerDutyAction(attachpayload bool, description string, incidentkey string) *_pagerDutyAction {

	tmp := &_pagerDutyAction{v: types.NewPagerDutyAction()}

	tmp.AttachPayload(attachpayload)

	tmp.Description(description)

	tmp.IncidentKey(incidentkey)

	return tmp

}

func (s *_pagerDutyAction) Account(account string) *_pagerDutyAction {

	s.v.Account = &account

	return s
}

func (s *_pagerDutyAction) AttachPayload(attachpayload bool) *_pagerDutyAction {

	s.v.AttachPayload = attachpayload

	return s
}

func (s *_pagerDutyAction) Client(client string) *_pagerDutyAction {

	s.v.Client = &client

	return s
}

func (s *_pagerDutyAction) ClientUrl(clienturl string) *_pagerDutyAction {

	s.v.ClientUrl = &clienturl

	return s
}

func (s *_pagerDutyAction) Contexts(contexts ...types.PagerDutyContextVariant) *_pagerDutyAction {

	for _, v := range contexts {

		s.v.Contexts = append(s.v.Contexts, *v.PagerDutyContextCaster())

	}
	return s
}

func (s *_pagerDutyAction) Description(description string) *_pagerDutyAction {

	s.v.Description = description

	return s
}

func (s *_pagerDutyAction) EventType(eventtype pagerdutyeventtype.PagerDutyEventType) *_pagerDutyAction {

	s.v.EventType = &eventtype
	return s
}

func (s *_pagerDutyAction) IncidentKey(incidentkey string) *_pagerDutyAction {

	s.v.IncidentKey = incidentkey

	return s
}

func (s *_pagerDutyAction) Proxy(proxy types.PagerDutyEventProxyVariant) *_pagerDutyAction {

	s.v.Proxy = proxy.PagerDutyEventProxyCaster()

	return s
}

func (s *_pagerDutyAction) PagerDutyActionCaster() *types.PagerDutyAction {
	return s.v
}
