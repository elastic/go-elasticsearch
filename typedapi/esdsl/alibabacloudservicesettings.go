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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

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

// A valid API key for the AlibabaCloud AI Search API.
func (s *_alibabaCloudServiceSettings) ApiKey(apikey string) *_alibabaCloudServiceSettings {

	s.v.ApiKey = apikey

	return s
}

// The name of the host address used for the inference task.
// You can find the host address in the API keys section of the documentation.
func (s *_alibabaCloudServiceSettings) Host(host string) *_alibabaCloudServiceSettings {

	s.v.Host = host

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// AlibabaCloud AI Search.
// By default, the `alibabacloud-ai-search` service sets the number of requests
// allowed per minute to `1000`.
func (s *_alibabaCloudServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_alibabaCloudServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

// The name of the model service to use for the inference task.
// The following service IDs are available for the `completion` task:
//
// * `ops-qwen-turbo`
// * `qwen-turbo`
// * `qwen-plus`
// * `qwen-max ÷ qwen-max-longcontext`
//
// The following service ID is available for the `rerank` task:
//
// * `ops-bge-reranker-larger`
//
// The following service ID is available for the `sparse_embedding` task:
//
// * `ops-text-sparse-embedding-001`
//
// The following service IDs are available for the `text_embedding` task:
//
// `ops-text-embedding-001`
// `ops-text-embedding-zh-001`
// `ops-text-embedding-en-001`
// `ops-text-embedding-002`
func (s *_alibabaCloudServiceSettings) ServiceId(serviceid string) *_alibabaCloudServiceSettings {

	s.v.ServiceId = serviceid

	return s
}

// The name of the workspace used for the inference task.
func (s *_alibabaCloudServiceSettings) Workspace(workspace string) *_alibabaCloudServiceSettings {

	s.v.Workspace = workspace

	return s
}

func (s *_alibabaCloudServiceSettings) AlibabaCloudServiceSettingsCaster() *types.AlibabaCloudServiceSettings {
	return s.v
}
