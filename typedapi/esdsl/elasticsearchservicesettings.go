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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _elasticsearchServiceSettings struct {
	v *types.ElasticsearchServiceSettings
}

func NewElasticsearchServiceSettings(modelid string, numthreads int) *_elasticsearchServiceSettings {

	tmp := &_elasticsearchServiceSettings{v: types.NewElasticsearchServiceSettings()}

	tmp.ModelId(modelid)

	tmp.NumThreads(numthreads)

	return tmp

}

func (s *_elasticsearchServiceSettings) AdaptiveAllocations(adaptiveallocations types.AdaptiveAllocationsVariant) *_elasticsearchServiceSettings {

	s.v.AdaptiveAllocations = adaptiveallocations.AdaptiveAllocationsCaster()

	return s
}

func (s *_elasticsearchServiceSettings) DeploymentId(deploymentid string) *_elasticsearchServiceSettings {

	s.v.DeploymentId = &deploymentid

	return s
}

func (s *_elasticsearchServiceSettings) LongDocumentStrategy(longdocumentstrategy string) *_elasticsearchServiceSettings {

	s.v.LongDocumentStrategy = &longdocumentstrategy

	return s
}

func (s *_elasticsearchServiceSettings) MaxChunksPerDoc(maxchunksperdoc int) *_elasticsearchServiceSettings {

	s.v.MaxChunksPerDoc = &maxchunksperdoc

	return s
}

func (s *_elasticsearchServiceSettings) ModelId(modelid string) *_elasticsearchServiceSettings {

	s.v.ModelId = modelid

	return s
}

func (s *_elasticsearchServiceSettings) NumAllocations(numallocations int) *_elasticsearchServiceSettings {

	s.v.NumAllocations = &numallocations

	return s
}

func (s *_elasticsearchServiceSettings) NumThreads(numthreads int) *_elasticsearchServiceSettings {

	s.v.NumThreads = numthreads

	return s
}

func (s *_elasticsearchServiceSettings) ElasticsearchServiceSettingsCaster() *types.ElasticsearchServiceSettings {
	return s.v
}
