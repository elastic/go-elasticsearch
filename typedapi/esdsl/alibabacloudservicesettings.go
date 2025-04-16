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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _alibabaCloudServiceSettings struct {
	v *types.AlibabaCloudServiceSettings
}

func NewAlibabaCloudServiceSettings(apikey string, host string, serviceid string, workspace string) *_alibabaCloudServiceSettings {

	tmp := &_alibabaCloudServiceSettings{v: types.NewAlibabaCloudServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.Host(host)

	tmp.ServiceId(serviceid)

	tmp.Workspace(workspace)

	return tmp

}

func (s *_alibabaCloudServiceSettings) ApiKey(apikey string) *_alibabaCloudServiceSettings {

	s.v.ApiKey = apikey

	return s
}

func (s *_alibabaCloudServiceSettings) Host(host string) *_alibabaCloudServiceSettings {

	s.v.Host = host

	return s
}

func (s *_alibabaCloudServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_alibabaCloudServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_alibabaCloudServiceSettings) ServiceId(serviceid string) *_alibabaCloudServiceSettings {

	s.v.ServiceId = serviceid

	return s
}

func (s *_alibabaCloudServiceSettings) Workspace(workspace string) *_alibabaCloudServiceSettings {

	s.v.Workspace = workspace

	return s
}

func (s *_alibabaCloudServiceSettings) AlibabaCloudServiceSettingsCaster() *types.AlibabaCloudServiceSettings {
	return s.v
}
